package route

import (
	"log"
	"product_manager/interface/api/handler"
	"product_manager/interface/api/middleware"
	"product_manager/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Routes() {
	if utils.GetEnvWithDefault("ENVIRONMENT", "debug") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, _ := jwt.New(middleware.GetJWTMiddleware(&gin.Context{}))
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/migration", handler.Migration)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	auth := r.Group("/api/v1")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/product", handler.CreateProduct)
		auth.PATCH("/product/:id", handler.UpdateProduct)
		auth.GET("/product/:id", handler.FindProductById)
		auth.GET("/products", handler.FindProducts)
		auth.DELETE("/product/:id", handler.DeleteProduct)
	}

	r.Run()
}
