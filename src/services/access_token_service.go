package access_token_service

import (
	"strings"

	"github.com/deepak-v4/bookstore_oauth-api/src/Repository/db"
	"github.com/deepak-v4/bookstore_oauth-api/src/Repository/rest_repo"
	"github.com/deepak-v4/bookstore_oauth-api/src/domain/access_token"
	"github.com/deepak-v4/bookstore_oauth-api/src/utils/rest_errors"
)

/*type Respository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	CreateId(*access_token.AccessToken) *rest_errors.RestErr
	UpdateId(*access_token.AccessToken) *rest_errors.RestErr
}*/

type Service interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	CreateId(access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr)
	UpdateId(access_token.AccessToken) *rest_errors.RestErr
}

type service struct {
	restUserRepo rest_repo.RestUsersRepository
	dbRepo       db.DbRepository
}

func NewService(userRepo rest_repo.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUserRepo: userRepo,
		dbRepo:       dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *rest_errors.RestErr) {
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) CreateId(request access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr) {

	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	err = s.dbRepo.CreateId(at)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateId(at access_token.AccessToken) *rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return rest_errors.NewBadRequestError("Invalid access token")
	}
	err := s.dbRepo.UpdateId(at)
	if err != nil {
		return err
	}
	return nil
}
