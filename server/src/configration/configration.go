package configuration

import (
	"github.com/kelseyhightower/envconfig"
)

// Config
type Config struct {
	Port string `envconfig:"PORT" required:"true"`
	Env  string `envconfig:"ENV" required:"true"`

	CORS struct {
		AllowOrigins []string `envconfig:"CORS_ALLOW_ORIGINS" default:"*"`
	}

	Database struct {
		DSN                   string `envconfig:"DATABASE_DSN" required:"true"`
		MaxOpenConnections    int    `envconfig:"DATABASE_MAX_OPEN_CONNECTIONS" default:"5"`
		MaxIdleConnections    int    `envconfig:"DATABASE_MAX_IDLE_CONNECTIONS" default:"5"`
		MaxConnectionLifetime int    `envconfig:"DATABASE_CONNECTION_MAX_LIFETIME_SECONDS" default:"300"`
		QueryLogging          bool   `envconfig:"DATABASE_QUERY_LOGGING" default:"false"`
	}

	JWTAuth struct {
		Base64PrivateKey string `envconfig:"JWT_AUTH_BASE64_PRIVATE_KEY" required:"true"`
		Base64PublicKey  string `envconfig:"JWT_AUTH_BASE64_PUBLIC_KEY" required:"true"`
	}

	AWS struct {
		S3BucketName   string `envconfig:"S3_BUCKET_NAME" required:"true"`
		S3BucketRegion string `envconfig:"S3_BUCKET_REGION" required:"true"`
	}

	Sentry struct {
		DSN string `envconfig:"SENTRY_DSN" required:"true"`
	}
}

var globalConfig Config

// Load は環境変数を globalConfig に読み込む。
func Load() {
	prefix := ""
	envconfig.MustProcess(prefix, &globalConfig)
}

// Get は Config を取得する。
func Get() Config {
	return globalConfig
}
