package services

import (
	"be-test/config"
	"be-test/domain"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type GoogleOauthService interface {
	ExchangeCode(
		context context.Context,
		code string,
	) (
		domain.GoogleOauthServiceExchangeCodeReturn,
		error,
	)
	GetGoogleUserData(accessToken string) (
		domain.GoogleOauthServiceGetUserDataReturn,
		error,
	)
}

type GoogleOauthServiceImpl struct {
	oauth *config.OauthConfig
}

func (g *GoogleOauthServiceImpl) ExchangeCode(
	ctx context.Context,
	code string,
) (
	domain.GoogleOauthServiceExchangeCodeReturn,
	error,
) {
	token, err := g.oauth.Google.Exchange(
		ctx,
		code,
	)
	if err != nil {
		return domain.GoogleOauthServiceExchangeCodeReturn{}, err
	}
	return domain.GoogleOauthServiceExchangeCodeReturn{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (g *GoogleOauthServiceImpl) GetGoogleUserData(accessToken string) (
	domain.GoogleOauthServiceGetUserDataReturn,
	error,
) {
	var userInfo domain.GoogleOauthServiceGetUserDataReturn
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return domain.GoogleOauthServiceGetUserDataReturn{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(
		body,
		&userInfo,
	)
	return userInfo, nil
}

func NewGoogleOauthService(oauthConfig *config.OauthConfig) *GoogleOauthServiceImpl {
	return &GoogleOauthServiceImpl{oauth: oauthConfig}
}
