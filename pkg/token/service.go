package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/kaanaktas/openbanking-paymentinitiation/api"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/client"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/security"
	cfg "github.com/kaanaktas/openbanking-paymentinitiation/pkg/configmanager"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
)

type Service interface {
	GetAccessToken(aspspId, scopeType string) (string, error)
	RefreshAccessToken(aspspId, scopeType, refreshTokenData string) (*AccessToken, error)
	GetResourceAccessRefreshToken(aspspId, code string) (*AccessToken, error)
}

type service struct {
	cfg cfg.Service
}

func NewService(cfg cfg.Service) Service {
	return &service{cfg}
}

const (
	grantType           = "grant_type"
	clientId            = "client_id"
	scope               = "scope"
	refreshToken        = "refresh_token"
	code                = "code"
	redirectUri         = "redirect_uri"
	clientAssertionType = "client_assertion_type"
	clientAssertion     = "client_assertion"
)

func (s service) GetAccessToken(aspspId, scopeType string) (string, error) {
	var errMessage = "error in GetAccessToken()"
	endpointOauth2, err := s.cfg.FindByConfigName(aspspId, api.EndpointOauth2)
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}

	clientIdValue, err := s.cfg.FindByConfigName(aspspId, api.ClientId)
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}

	financialId, err := s.cfg.FindByConfigName(aspspId, api.FapiFinancialId)
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}
	iss, err := s.cfg.FindByConfigName(aspspId, api.Iss)
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}
	aud, err := s.cfg.FindByConfigName(aspspId, api.Iss)
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}

	cs := ClientAssertion{
		Iss: iss,
		Sub: iss,
		Exp: security.CreateTokenTime(240),
		Iat: security.CreateTokenTime(0),
		Jti: uuid.New().String(),
		Aud: aud,
	}

	csJson, err := json.Marshal(cs)
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}

	signedClientAssertion, err := security.GenerateJwtWithJsonString(bytes.NewBuffer(csJson).String(), jwt.SigningMethodPS256)
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}

	parameters := url.Values{}
	parameters.Set(grantType, "client_credentials")
	parameters.Set(clientId, clientIdValue)
	parameters.Set(scope, scopeType)
	parameters.Set(clientAssertionType, "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	parameters.Set(clientAssertion, signedClientAssertion)

	httpClient, err := client.NewSecureHttpClient(endpointOauth2, s.setHeader(financialId))
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}

	resp, err := httpClient.Post(strings.NewReader(parameters.Encode()))
	if err != nil {
		return "", errors.WithMessage(err, errMessage)
	}

	if (resp.StatusCode == 200 || resp.StatusCode == 201) && resp.Body != "" {
		var accessToken *AccessToken

		err = json.Unmarshal([]byte(resp.Body), &accessToken)
		if err != nil {
			return "", errors.WithMessage(err, errMessage)
		}

		return accessToken.AccessToken, nil
	} else {
		return "", fmt.Errorf("unexpected result from the token service. resp: %v", *resp)
	}
}

func (s service) RefreshAccessToken(aspspId, scopeType, refreshTokenData string) (*AccessToken, error) {
	var errMessage = "error in RefreshAccessToken()"
	endpointOauth2, err := s.cfg.FindByConfigName(aspspId, api.EndpointOauth2)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	clientIdValue, err := s.cfg.FindByConfigName(aspspId, api.ClientId)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	financialId, err := s.cfg.FindByConfigName(aspspId, api.FapiFinancialId)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	parameters := url.Values{}
	parameters.Set(grantType, "refresh_token")
	parameters.Set(refreshToken, refreshTokenData)
	parameters.Set(clientId, clientIdValue)
	parameters.Set(scope, scopeType)

	httpClient, err := client.NewSecureHttpClient(endpointOauth2, s.setHeader(financialId))
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	resp, err := httpClient.Post(strings.NewReader(parameters.Encode()))
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	var accessToken *AccessToken
	if resp.StatusCode == 200 {
		err = json.Unmarshal([]byte(resp.Body), &accessToken)
		if err != nil {
			return nil, errors.WithMessage(err, errMessage)
		}

		return accessToken, nil
	} else {
		return nil, fmt.Errorf("unexpected result from the token service. resp: %v", *resp)
	}
}

func (s service) GetResourceAccessRefreshToken(aspspId, authCode string) (*AccessToken, error) {
	var errMessage = "error in GetResourceAccessRefreshToken()"
	endpointOauth2, err := s.cfg.FindByConfigName(aspspId, api.EndpointOauth2)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	appRedirectUrl, err := s.cfg.FindByConfigName(aspspId, api.RedirectUrl)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	clientIdParam, err := s.cfg.FindByConfigName(aspspId, api.ClientId)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	financialId, err := s.cfg.FindByConfigName(aspspId, api.FapiFinancialId)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}
	iss, err := s.cfg.FindByConfigName(aspspId, api.Iss)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}
	aud, err := s.cfg.FindByConfigName(aspspId, api.Iss)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	cs := ClientAssertion{
		Iss: iss,
		Sub: iss,
		Exp: security.CreateTokenTime(240),
		Iat: security.CreateTokenTime(0),
		Jti: uuid.New().String(),
		Aud: aud,
	}

	csJson, err := json.Marshal(cs)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	signedClientAssertion, err := security.GenerateJwtWithJsonString(bytes.NewBuffer(csJson).String(), jwt.SigningMethodPS256)
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	parameters := url.Values{}
	parameters.Set(grantType, "authorization_code")
	parameters.Set(clientId, clientIdParam)
	parameters.Set(redirectUri, appRedirectUrl)
	parameters.Set(code, authCode)
	parameters.Set(clientAssertionType, "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	parameters.Set(clientAssertion, signedClientAssertion)

	httpClient, err := client.NewSecureHttpClient(endpointOauth2, s.setHeader(financialId))
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	resp, err := httpClient.Post(strings.NewReader(parameters.Encode()))
	if err != nil {
		return nil, errors.WithMessage(err, errMessage)
	}

	if (resp.StatusCode == 200 || resp.StatusCode == 201) && resp.Body != "" {
		var accessToken *AccessToken
		err = json.Unmarshal([]byte(resp.Body), &accessToken)
		if err != nil {
			return accessToken, errors.WithMessage(err, errMessage)
		}

		return accessToken, nil
	} else {
		return nil, fmt.Errorf("unexpected result from the token service. resp: %v", *resp)
	}
}

func (s service) setHeader(financialId string) http.Header {
	header := http.Header{}
	header.Set(api.Accept, api.ApplicationJson)
	header.Set(api.ContentType, api.ApplicationFormUrlencodedValue)
	header.Set(api.CacheControl, "no-cache")
	header.Set(api.XFapiFinancialId, financialId)

	return header
}
