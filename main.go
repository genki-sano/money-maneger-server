package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/sheets/v4"
)

type itemResponse struct {
	Name     interface{}
	Date     interface{}
	Price    interface{}
	Category interface{}
	Memo     interface{}
}

func getValues(r *sheets.ValueRange) []itemResponse {
	ret := make([]itemResponse, 0)
	if r == nil {
		return ret
	}
	for i, items := range r.Values {
		if i == 0 {
			continue
		}
		ret = append(ret, itemResponse{
			Name:     items[1],
			Date:     items[2],
			Price:    items[3],
			Category: items[4],
			Memo:     items[5],
		})
	}
	return ret
}

func getList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		sheetsService, err := sheets.NewService(ctx)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}

		valueRange, err := sheetsService.Spreadsheets.Values.Get(
			os.Getenv("GOOGLE_SPREDSHEET_ID"),
			os.Getenv("GOOGLE_SPREDSHEET_RANGE"),
		).Do()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"values": getValues(valueRange),
		})
	}
}

func router() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.GET("/list", getList())

	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router().Run(":" + port); err != nil {
		panic(err)
	}
}
