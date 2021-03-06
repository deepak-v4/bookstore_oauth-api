package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/deepak-v4/bookstore_oauth-api/src/domain/access_token"
	access_token_service "github.com/deepak-v4/bookstore_oauth-api/src/services"
	"github.com/deepak-v4/bookstore_oauth-api/src/utils/rest_errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	CreateId(*gin.Context)
}

type accessTokenHandler struct {
	service access_token_service.Service
}

func NewAccessTokenHandler(service access_token_service.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {

	log.Println("error 1")
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) CreateId(c *gin.Context) {

	var at access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := rest_errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := h.service.CreateId(at)
	if err != nil {
		c.JSON(err.Status, err)
		return

	}

	c.JSON(http.StatusCreated, accessToken)
}
