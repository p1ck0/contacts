package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	name string
}

func check(db *gorm.DB) (bool, error) {
	rows, err := db.Raw("select datname from pg_database").Rows()
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		d := database{}
		err := rows.Scan(&d.name)
		if err != nil {
			log.Println(err)
			continue
		}
		if d.name == dbname {
			return true, nil
		}
	}
	return false, nil
}

func CreateDB() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=postgres",
		host, port, user, pass, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	ok, err := check(db)
	if err != nil {
		return err
	}
	if !ok {
		db.Exec("CREATE DATABASE " + dbname)
		db = Connector()
		db.AutoMigrate(&Contact{})
	}
	return nil
}

func Connector() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=%s",
		host, port, user, pass, sslmode, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
