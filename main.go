package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

// Payment 支払い内容
type Payment struct {
	ID       interface{}
	Name     interface{}
	Date     interface{}
	Price    interface{}
	Category interface{}
	Memo     interface{}
}

func makeValues(r *sheets.ValueRange) []Payment {
	ret := make([]Payment, 0)
	if r == nil {
		return ret
	}

	for _, items := range r.Values {
		ret = append(ret, Payment{
			ID:       items[0],
			Name:     items[1],
			Date:     items[2],
			Price:    items[3],
			Category: items[4],
			Memo:     items[5],
		})
	}
	return ret
}

func getClient(email string, key string) *http.Client {
	conf := &jwt.Config{
		Email:      email,
		PrivateKey: []byte(key),
		TokenURL:   google.JWTTokenURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets.readonly",
		},
	}

	return conf.Client(oauth2.NoContext)
}

func getList() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := os.Getenv("GOOGLE_SERVICE_ACCOUNT_EMAIL")
		privateKey := os.Getenv("GOOGLE_SERVICE_ACCOUNT_PLIVATE_KEY")

		client := getClient(email, privateKey)
		Service, err := sheets.New(client)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}

		spreadsheetID := os.Getenv("GOOGLE_SPREDSHEET_ID")
		valueRange := os.Getenv("GOOGLE_SPREDSHEET_RANGE")

		ctx := context.Background()

		resp, err := Service.Spreadsheets.Values.Get(spreadsheetID, valueRange).Context(ctx).Do()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"values": makeValues(resp),
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
