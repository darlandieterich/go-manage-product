package db

import (
	"errors"
	"log"
	"product_manager/domain/model"
	"product_manager/infra/db/datasource"

	"gorm.io/gorm"
)

const (
	DriverPostgres DriverName = "postgres"
)

type DriverName string

type Connection struct {
	Conn *gorm.DB
}

func NewConnection(d DriverName) (conn *Connection, err error) {
	switch d {
	case DriverPostgres:
		{
			db, err := datasource.DataSourcePostgres()
			if err != nil {
				return nil, err
			}
			conn = &Connection{Conn: db}
		}
	default:
		return nil, errors.New("driver not found")
	}

	return conn, nil
}

func (c *Connection) RunMigration() {
	// c.Conn.Migrator().DropTable(&model.Product{}, &model.Stock{})

	err := c.Conn.AutoMigrate(&model.Product{})
	if err != nil {
		log.Println("Error on migration of product", err)
	}

	err = c.Conn.AutoMigrate(&model.Stock{})
	if err != nil {
		log.Println("Error on migration of stock", err)
	}
}
