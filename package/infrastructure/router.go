package infrastructure

import (
	"errors"
	"net/http"
	"os"

	"github.com/genki-sano/money-maneger-server/package/infrastructure/di"
	"github.com/genki-sano/money-maneger-server/package/interfaces/handlers"
	"github.com/gin-gonic/gin"
)

// Route ルーティングの設定
func Route() *gin.Engine {
	e := createEngine()

	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"healthCheck": "ok"})
	})

	api := e.Group("/api")
	{
		api.GET("/payments", func(c *gin.Context) { di.InitializePayment().Index(c) })
		api.GET("/payments/:date", func(c *gin.Context) { di.InitializePayment().Show(c) })
	}

	e.NoRoute(func(c *gin.Context) {
		err := errors.New("no route to host")
		c.JSON(http.StatusNotFound, handlers.NewError(err))
	})

	return e
}

func createEngine() *gin.Engine {
	gin.DisableConsoleColor()
	setMode()

	return gin.Default()
}

func setMode() {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
}
