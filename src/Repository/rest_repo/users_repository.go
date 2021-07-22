package rest_repo

import (
	"encoding/json"
	"time"

	"github.com/deepak-v4/bookstore_oauth-api/src/domain/users"
	"github.com/deepak-v4/bookstore_oauth-api/src/utils/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

type userRepository struct {
}

func NewRepository() RestUsersRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := userRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, nil
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal")
	}
	return &user, nil
}
