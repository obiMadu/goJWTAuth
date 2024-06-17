package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/obiMadu/goJWTAuth/internals/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var counts int

func NewDB() (*gorm.DB, error) {
	// new db
	db := connectToPostgres()
	log.Println("Connected to DB")

	// migrate models
	err := migrate(db)
	if err != nil {
		return nil, err
	}
	log.Println("Successfully Migrated Models.")

	return db, nil
}

func NewRawDB() *sql.DB {
	db, err := NewDB()
	if err != nil {
		log.Panicf("Unable to create & configure DB %s\n", err.Error())
	}

	rawDB := RawDB(db)

	return rawDB
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}

func RawDB(db *gorm.DB) *sql.DB {
	rawDB, err := db.DB()
	if err != nil {
		log.Panicf("Unable to get raw sql.DB %s\n", err.Error())
	}

	return rawDB
}

func connectToPostgres() *gorm.DB {
	dsn := os.Getenv("POSTGRES_DSN")

	for {
		connection, err := openPostgres(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Fatal(err)
		}

		log.Println("Backing off for three seconds....")
		time.Sleep(3 * time.Second)
		continue
	}
}

func connectToMysql() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")

	for {
		connection, err := openMysql(dsn)
		if err != nil {
			log.Println("MySQL not yet ready ...")
			counts++
		} else {
			log.Println("Connected to MySQL!")
			return connection
		}

		if counts > 10 {
			log.Fatal(err)
		}

		log.Println("Backing off for three seconds....")
		time.Sleep(3 * time.Second)
		continue
	}
}

func openPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// return *sql.DB from db(*gorm.DB) to enable Ping()
	gormDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// ping database
	err = gormDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func openMysql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// return *sql.DB from db(*gorm.DB) to enable Ping()
	gormDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// ping database
	err = gormDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
