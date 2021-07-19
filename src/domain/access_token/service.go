package access_token

import (
	"strings"

	"github.com/deepak-v4/bookstore_oauth-api/src/utils/rest_errors"
)

type Respository interface {
	GetById(string) (*AccessToken, *rest_errors.RestErr)
	CreateId(AccessToken) *rest_errors.RestErr
	UpdateId(AccessToken) *rest_errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *rest_errors.RestErr)
	CreateId(AccessToken) *rest_errors.RestErr
	UpdateId(AccessToken) *rest_errors.RestErr
}

type service struct {
	respository Respository
}

func NewService(repo Respository) Service {
	return &service{
		respository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *rest_errors.RestErr) {
	accessToken, err := s.respository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) CreateId(at AccessToken) *rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return rest_errors.NewBadRequestError("Invalid access token")
	}
	err := s.respository.CreateId(at)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateId(at AccessToken) *rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return rest_errors.NewBadRequestError("Invalid access token")
	}
	err := s.respository.UpdateId(at)
	if err != nil {
		return err
	}
	return nil
}
