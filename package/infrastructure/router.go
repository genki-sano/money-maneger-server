package infrastructure

import (
	"errors"
	"net/http"
	"os"

	"github.com/genki-sano/money-maneger-server/package/infrastructure/di"
	"github.com/genki-sano/money-maneger-server/package/interfaces/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Route ルーティングの設定
func Route() *gin.Engine {
	e := createEngine()

	e.NoRoute(func(c *gin.Context) {
		err := errors.New("no route to host")
		c.JSON(http.StatusNotFound, handlers.NewError(err))
	})

	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"healthCheck": "ok"})
	})

	api := e.Group("/api")
	{
		payments := api.Group("/payments")
		{
			payments.GET("/:date", func(c *gin.Context) { di.InitializePayment().List(c) })
		}
	}

	return e
}

func createEngine() *gin.Engine {
	gin.DisableConsoleColor()
	setMode()

	r := gin.Default()

	// COR設定の追加
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{
		"http://localhost:3000", // ローカル開発用
		"https://money-maneger-client.herokuapp.com",
	}
	r.Use(cors.New(config))

	return r
}

func setMode() {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
}
