package repository

import (
	"fmt"
	"time"

	"github.com/ibanezv/fly-devs-practice_go/cmd/api/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	_domainServiceName = "%s:%s@tcp(%s)/%s?%s"
)

type Repositories interface {
}

type Db struct {
	config         settings.Database
	openConnection func(connectionConfig string, opts gorm.Option) (db *gorm.DB, err error)
}

func NewDatabase(dbConfig settings.Database) Db {
	return Db{
		config: dbConfig,
		openConnection: func(connectionConfig string, cfg gorm.Option) (db *gorm.DB, err error) {
			return gorm.Open(mysql.Open(connectionConfig), cfg)
		},
	}
}

func (db *Db) GetConnection() (*gorm.DB, error) {
	connection, err := db.openConnection(getDSN(db.config), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = configureDatabaseConnection(connection, db.config)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func getDSN(dbConfig settings.Database) string {
	return fmt.Sprintf(_domainServiceName,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DatabaseName,
		"",
	)
}

func configureDatabaseConnection(db *gorm.DB, dbConfig settings.Database) error {
	dbConfiguration, err := db.DB()

	if err != nil {
		return err
	}

	dbConfiguration.SetMaxOpenConns(dbConfig.MaxOpenConns)
	dbConfiguration.SetMaxIdleConns(dbConfig.MaxIdleConns)
	dbConfiguration.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Millisecond)
	dbConfiguration.SetConnMaxIdleTime(time.Duration(dbConfig.ConnMaxIdleTime) * time.Millisecond)

	return nil
}
