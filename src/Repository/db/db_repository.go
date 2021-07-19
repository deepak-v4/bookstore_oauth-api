package db

import (
	"fmt"

	"github.com/deepak-v4/bookstore_oauth-api/src/clients/cassandra"
	"github.com/deepak-v4/bookstore_oauth-api/src/domain/access_token"
	"github.com/deepak-v4/bookstore_oauth-api/src/utils/rest_errors"
)

const (
	queryGetAccessToken    = "SELECT access_token,user_id,client_id,expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token,user_id,client_id,expires) VALUES(?,?,?,?);"
	queryUpdateAccessToken = "UPDATE access_tokens SET expires=? where access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	CreateId(access_token.AccessToken) *rest_errors.RestErr
	UpdateId(access_token.AccessToken) *rest_errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestErr) {

	session, err := cassandra.GetSession()
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())

	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		fmt.Println(err.Error())
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) CreateId(at access_token.AccessToken) *rest_errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())

	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateId(at access_token.AccessToken) *rest_errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error())

	}
	defer session.Close()

	if err := session.Query(queryUpdateAccessToken,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	return nil
}
