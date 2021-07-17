package app

import (
	"github.com/deepak-v4/bookstore_oauth-api/src/Repository/db"
	"github.com/deepak-v4/bookstore_oauth-api/src/domain/access_token"
	"github.com/deepak-v4/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/acess_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")

}
