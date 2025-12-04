package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	configuration "kujicole/configration"
	domainErr "kujicole/domain/error"
	"kujicole/domain/repository"
	"kujicole/graph"
	"kujicole/infra/authtoken"
	customEchoMiddleware "kujicole/infra/echo/middleware"
	"kujicole/logger"
	"kujicole/usecase"

	// gqlCustomExtension "kujicole/infra/gql/extension"
	"kujicole/infra/mysql"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/stephenafamo/bob"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "5050"

func main() {
	configuration.Load()
	conf := configuration.Get()

	// --------------------------------------------------
	// Setup database connection
	// --------------------------------------------------
	var (
		conn *bob.DB
		err  error
	)

	maxTry := 5
	for i := 0; i < maxTry; i++ {
		// try connection
		conn, err = mysql.Connect(conf.Database.DSN,
			mysql.MaxOpenConnections(conf.Database.MaxOpenConnections),
			mysql.MaxIdleConnections(conf.Database.MaxIdleConnections),
			mysql.ConnectionMaxLifetime(time.Duration(conf.Database.MaxConnectionLifetime)*time.Second),
			mysql.EnableQueryLogging(conf.Database.QueryLogging),
		)
		if err == nil {
			// success to break the retry loop
			break
		}

		logger.Warnf("connect error (%d/%d). error: %+v", i+1, maxTry, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		logger.Fatalf("failed to connect to database. error: %+v", err)
	}
	defer conn.Close()

	// --------------------------------------------------
	// Setup atomic executor
	// --------------------------------------------------
	repository.SetAtomicExecutor(mysql.NewAtomicExecutor(conn))

	// --------------------------------------------------
	// Setup repositories
	// --------------------------------------------------
	repos := repository.NewRepositories(
		mysql.NewUserRepository(conn),
		mysql.NewIntellectualPropertyRepository(conn),
	)

	// --------------------------------------------------
	// Setup query services
	// --------------------------------------------------
	// itemQueryService := mysql.NewItemQueryService(conn)

	// --------------------------------------------------
	// Setup services
	// --------------------------------------------------
	authTokenService, err := authtoken.NewAuthTokenService(conf.JWTAuth.Base64PrivateKey, conf.JWTAuth.Base64PublicKey, conn)
	if err != nil {
		logger.Fatalf("failed to create auth token service. error: %+v", err)
	}

	// fileService, err := s3.NewFileService(s3.S3FileServiceConfig{
	// 	BucketName:   conf.AWS.S3BucketName,
	// 	BucketRegion: conf.AWS.S3BucketRegion,
	// })
	// if err != nil {
	// 	logger.Fatalf("failed to create file service. error: %+v", err)
	// }

	// --------------------------------------------------
	// Setup use cases
	// --------------------------------------------------
	authUseCase := usecase.NewAuthUseCase(repos, authTokenService)
	userUseCase := usecase.NewUserUseCase(repos)
	intellectualPropertyUseCase := usecase.NewIntellectualPropertyUseCase(repos)
	useCases := usecase.NewUseCases(
		authUseCase,
		userUseCase,
		intellectualPropertyUseCase,
	)

	// --------------------------------------------------
	// Setup graphql handler
	// --------------------------------------------------
	gqlHandler := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			UseCases: useCases,
		},
	}))

	gqlHandler.AddTransport(transport.Options{})
	// なぜGETはないのか
	gqlHandler.AddTransport(transport.POST{})

	// error handler
	gqlHandler.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		logger.Errorf("some error occurred. raw err: %+v", err)
		srcErr := err
		var (
			gqlErr *gqlerror.Error
		)
		if errors.As(err, &gqlErr) {
			srcErr = gqlErr.Err
		}
		if srcErr == nil {
			return gqlerror.Errorf("internal system error")
		}
		logger.Errorf("some error occurred: %+v", srcErr)

		unwrappedErr := errors.Cause(srcErr)

		presentationErr := &gqlerror.Error{
			Message: unwrappedErr.Error(),
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]any{
				// debug に便利なので stack trace を含める（パブリックサービスではないので環境にかかわらず含める）
				"stackDetail": fmt.Sprintf("%+v", srcErr),
			},
		}

		var validationErr *domainErr.ValidationError
		if errors.As(srcErr, &validationErr) {
			presentationErr.Extensions["validationErrors"] = map[string]any{
				"code": validationErr.Code,
			}
		}

		var convertCSVErr domainErr.ConvertCSVError
		if errors.As(srcErr, &convertCSVErr) {
			presentationErr.Extensions["convertCSVError"] = map[string]any{
				"relatedProps": convertCSVErr.RelatedProps(),
				"type":         convertCSVErr.Type(),
			}
		}

		return presentationErr
	})

	e := echo.New()

	allowOrigins := conf.CORS.AllowOrigins
	if len(allowOrigins) == 0 {
		allowOrigins = []string{"*"}
	}
	allowAllOrigins := len(allowOrigins) == 1 && allowOrigins[0] == "*"

	corsConfig := echomiddleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}
	if allowAllOrigins {
		corsConfig.AllowOrigins = nil
		corsConfig.AllowOriginFunc = func(origin string) (bool, error) {
			return true, nil
		}
	}

	e.Use(
		echomiddleware.Logger(),
		echomiddleware.CORSWithConfig(corsConfig),
		(&customEchoMiddleware.AddUserIDMiddleware{
			Service: authTokenService,
		}).MiddlewareFunc(),
	)

	// health check
	const healthCheckPath = "/health-check"
	e.GET(healthCheckPath, func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// graphql
	gqlQueryPath := "/query"
	e.POST(gqlQueryPath, func(c echo.Context) error {
		req := c.Request()
		gqlHandler.ServeHTTP(c.Response().Writer, req)
		return nil
	})

	e.GET("/", func(c echo.Context) error {
		playground.Handler("GraphQL playground", gqlQueryPath)(c.Response().Writer, c.Request())
		return nil
	})
	// Useの使い方について
	gqlHandler.Use(extension.Introspection{})

	// --------------------------------------------------
	// Start the server in the background
	// --------------------------------------------------
	go func() {
		bindAddr := fmt.Sprintf("0.0.0.0:%s", conf.Port)
		e.Logger.Fatal(e.Start(bindAddr))
	}()

	// --------------------------------------------------
	// wait signal for graceful stop
	// --------------------------------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("stopping server...")
	e.Shutdown(context.Background())
	logger.Info("server exiting")
}
