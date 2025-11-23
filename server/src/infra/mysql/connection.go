package mysql

import (
	"context"
	"time"

	mysqldriver "github.com/go-sql-driver/mysql" //nolint
	"github.com/pkg/errors"

	"kujicole/logger"
	"github.com/stephenafamo/bob"
	sqldblogger "github.com/simukti/sqldb-logger"
)

// ConnectionOptions は、DB が扱うコネクションプールの設定を表す。
type ConnectionOptions struct {
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifetime time.Duration
	EnableQueryLogging    bool
}

func defaultConnectionOptions() *ConnectionOptions {
	return &ConnectionOptions{
		MaxOpenConnections:    5,
		MaxIdleConnections:    5,
		ConnectionMaxLifetime: 5 * time.Minute,
		EnableQueryLogging:    false,
	}
}

// connectionOption は、コネクションプールの設定オプションであることを表すインタフェースの定義。
type connectionOption interface {
	apply(*ConnectionOptions)
}

// MaxOpenConnections は、コネクションの最大オープン数を設定するオプション。
type MaxOpenConnections int

func (n MaxOpenConnections) apply(o *ConnectionOptions) {
	o.MaxOpenConnections = int(n)
}

// MaxIdleConnections は、コネクションの最大アイドル数を設定するオプション。
type MaxIdleConnections int

func (n MaxIdleConnections) apply(o *ConnectionOptions) {
	o.MaxIdleConnections = int(n)
}

// ConnectionMaxLifetime は、コネクションの最大生存時間を設定するオプション。
type ConnectionMaxLifetime time.Duration

func (d ConnectionMaxLifetime) apply(o *ConnectionOptions) {
	o.ConnectionMaxLifetime = time.Duration(d)
}

type EnableQueryLogging bool

func (b EnableQueryLogging) apply(o *ConnectionOptions) {
	o.EnableQueryLogging = bool(b)
}

// Connect は、データベースに接続し、コネクションプールの設定を行う。
func Connect(dsn string, options ...connectionOption) (*bob.DB, error) {
	o := defaultConnectionOptions()
	for _, f := range options {
		f.apply(o)
	}

	var conn bob.DB
	if o.EnableQueryLogging {
		sqlDB := sqldblogger.OpenDriver(dsn, &mysqldriver.MySQLDriver{}, &queryLogger{})	
		conn = bob.NewDB(sqlDB)
	} else {
		var err error
		conn, err = bob.Open("mysql", dsn)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	conn.SetMaxOpenConns(o.MaxOpenConnections)
	conn.SetMaxIdleConns(o.MaxIdleConnections)
	conn.SetConnMaxLifetime(o.ConnectionMaxLifetime)

	err := conn.Ping()
	if err != nil {
		conn.Close()
		return nil, errors.WithStack(err)
	}

	return &conn, nil
}

type queryLogger struct {
}

func (ql *queryLogger) Log(ctx context.Context, level sqldblogger.Level, msg string, data map[string]interface{}) {
	switch level {
	// case sqldblogger.LevelTrace:
	// 	logger.Debugf("%s, %+v", msg, data)

	// case sqldblogger.LevelDebug:
	// 	logger.Debugf("%s, %+v", msg, data)

	case sqldblogger.LevelInfo:
		logger.Infof("%s, query: %+v args: %+v", msg, data["query"], data["args"])

	case sqldblogger.LevelError:
		logger.Errorf("%s, query: %s args: %+v", msg, data["query"], data["args"])
	}
}
