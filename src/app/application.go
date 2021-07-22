package app

import (
	"github.com/deepak-v4/bookstore_oauth-api/src/Repository/db"
	"github.com/deepak-v4/bookstore_oauth-api/src/Repository/rest_repo"
	"github.com/deepak-v4/bookstore_oauth-api/src/http"
	access_token_service "github.com/deepak-v4/bookstore_oauth-api/src/services"
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
	atHandler := http.NewAccessTokenHandler(
		access_token_service.NewService(rest_repo.NewRepository(), db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.CreateId)
	router.Run(":8080")

}
