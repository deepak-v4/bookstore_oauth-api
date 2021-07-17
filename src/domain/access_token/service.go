package access_token

import "github.com/deepak-v4/bookstore_oauth-api/src/utils/rest_errors"

type Respository interface {
	GetById(string) (*AccessToken, *rest_errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *rest_errors.RestErr)
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
