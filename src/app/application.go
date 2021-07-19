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

	/*session, dbErr := cassandra.GetSession()
	if dbErr != nil {

		panic(dbErr)
	}
	session.Close()*/

	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewAccessTokenHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.CreateId)
	router.Run(":8080")

}
