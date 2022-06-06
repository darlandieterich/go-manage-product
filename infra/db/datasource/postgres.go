package datasource

import (
	"fmt"
	"product_manager/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DataSourcePostgres() (*gorm.DB, error) {
	user := utils.GetEnvWithDefault("DB_USER", "root")
	password := utils.GetEnvWithDefault("DB_PASSWORD", "passwd")
	host := utils.GetEnvWithDefault("DB_HOST", "localhost")
	port := utils.GetEnvWithDefault("DB_PORT", "5432")
	dbName := utils.GetEnvWithDefault("DB_NAME", "banco")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbName)

	return gorm.Open(postgres.Open(dsn))
}
