package handler

import (
	"context"
	"product_manager/application"
	"product_manager/infra/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	driver = db.DriverPostgres
)

func CreateProduct(c *gin.Context) {
	ib, err := initBase()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	product := &application.ProductParam{}
	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	productID, err := ib.CreateProduct(context.Background(), product)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":    "product has created",
		"product_id": productID,
	})
}

func UpdateProduct(c *gin.Context) {
	ib, err := initBase()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id := c.Params.ByName("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	product := &application.ProductParam{}
	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = ib.UpdateProduct(context.Background(), uuid, product)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "product has updated",
	})
}

func initBase() (*application.ProductService, error) {
	conn, err := db.NewConnection(driver)

	if err != nil {
		return nil, err
	}

	return application.NewProductService(conn.Conn), nil
}

func Migration(c *gin.Context) {
	conn, _ := db.NewConnection(driver)

	conn.RunMigration()
}
