package main

import (
	"context"
	"time"

	"kujicole/cmd/testdatafactory/init/csvs"
	"kujicole/configration"
	"kujicole/infra/bob/models"
	"kujicole/infra/mysql"
	"kujicole/logger"
	"github.com/stephenafamo/bob"
)

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

	tx, err := conn.BeginTx(context.Background(), nil)
	if err != nil {
		logger.Fatalf("failed to begin transaction. error: %+v", err)
	}
	defer tx.Rollback(context.Background())

	records, err := csvs.ReadRecords()
	if err != nil {
		logger.Fatalf("failed to read records. error: %+v", err)
	}

	ctx := context.Background()

	// Insert users
	for _, user := range records.Users {
		_, err := models.Users.Insert(user).One(ctx, tx)
		if err != nil {
			logger.Fatalf("failed to insert user. error: %+v", err)
		}
	}

	// Insert intellectual property categories
	for _, category := range records.IntellectualPropertyCategories {
		_, err := models.IntellectualPropertyCategories.Insert(category).One(ctx, tx)
		if err != nil {
			logger.Fatalf("failed to insert intellectual property category. error: %+v", err)
		}
	}

	// Insert intellectual property rank groups
	for _, rankGroup := range records.IntellectualPropertyRankGroups {
		_, err := models.IntellectualPropertyRankGroups.Insert(rankGroup).One(ctx, tx)
		if err != nil {
			logger.Fatalf("failed to insert intellectual property rank group. error: %+v", err)
		}
	}

	// Insert intellectual properties
	for _, property := range records.IntellectualProperties {
		_, err := models.IntellectualProperties.Insert(property).One(ctx, tx)
		if err != nil {
			logger.Fatalf("failed to insert intellectual property. error: %+v", err)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		logger.Fatalf("failed to commit transaction. error: %+v", err)
	}
}
