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
	ID       string
	Name     string
	Date     string
	Price    string
	Category string
	Memo     string
}

func makeValues(r *sheets.ValueRange) []Payment {
	ret := make([]Payment, 0)
	if r == nil {
		return ret
	}

	for _, items := range r.Values {
		ret = append(ret, Payment{
			ID:       items[0].(string),
			Name:     items[1].(string),
			Date:     items[2].(string),
			Price:    items[3].(string),
			Category: items[4].(string),
			Memo:     items[5].(string),
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
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errors": []string{err.Error()}})
			return
		}

		spreadsheetID := os.Getenv("GOOGLE_SPREDSHEET_ID")
		valueRange := os.Getenv("GOOGLE_SPREDSHEET_RANGE")

		ctx := context.Background()

		resp, err := Service.Spreadsheets.Values.Get(spreadsheetID, valueRange).Context(ctx).Do()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"errors": []string{err.Error()}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"body": makeValues(resp)})
	}
}

func router() *gin.Engine {
	e := createEngine()

	api := e.Group("/api")
	{
		api.GET("/payments", getList())
	}

	e.NoRoute(
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotFound, gin.H{"errors": []string{"指定したURLが存在しません。"}})
		},
	)

	return e
}

func createEngine() *gin.Engine {
	gin.DisableConsoleColor()

	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	return gin.Default()
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
