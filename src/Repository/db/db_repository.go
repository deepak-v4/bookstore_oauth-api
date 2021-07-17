package db

import (
	"github.com/deepak-v4/bookstore_oauth-api/src/domain/access_token"
	"github.com/deepak-v4/bookstore_oauth-api/src/utils/rest_errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestErr) {
	return nil, rest_errors.NewInternalServerError("db connection error")
}
