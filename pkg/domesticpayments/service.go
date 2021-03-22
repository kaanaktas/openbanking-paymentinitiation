package domesticpayments

import (
	"fmt"
	"github.com/kaanaktas/openbanking-paymentinitiation/api"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/client"
	"github.com/kaanaktas/openbanking-paymentinitiation/pkg/authmanager"
	cfg "github.com/kaanaktas/openbanking-paymentinitiation/pkg/configmanager"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type Service interface {
	DomesticPayment(cid, aspspId string) (string, error)
}

type service struct {
	authManager authmanager.AuthManager
	cfg         cfg.Service
}

func NewService(authManager authmanager.AuthManager, cfg cfg.Service) Service {
	return &service{
		authManager: authManager,
		cfg:         cfg,
	}
}

func (s service) DomesticPayment(cid, aspspId string) (string, error) {
	endpointDomesticPayment, err := s.cfg.FindByConfigName(aspspId, api.EndpointDomesticPayment)
	if err != nil {
		return "", errors.WithMessage(err, "errMessage")
	}

	return s.processCall(cid, aspspId, endpointDomesticPayment)
}

func (s service) processCall(cid, aspspId, endpointDomesticPayment string) (string, error) {
	resourceAccessToken, payload, err := s.authManager.GetAuthorisedTokenByCid(aspspId, cid)
	if err != nil {
		return "", err
	}

	fapiFinancialId, err := s.cfg.FindByConfigName(aspspId, api.FapiFinancialId)
	if err != nil {
		return "", err
	}

	httpClient, err := client.NewSecureHttpClient(endpointDomesticPayment, s.setHeader(resourceAccessToken, fapiFinancialId))
	if err != nil {
		return "", errors.WithMessage(err, "error in processCall()")
	}

	resp, err := httpClient.Post(strings.NewReader(payload))
	if err != nil {
		return "", errors.WithMessage(err, "error in processCall()")
	}

	switch resp.StatusCode {
	case 200, 201:
		return resp.Body, nil
	//TODO case 401 is an anomaly and needs to be taken care either by revoking the consent or refreshing the token
	default:
		return "", fmt.Errorf("unexpected result from the payment service. resp: %v", *resp)
	}
}

func (s service) setHeader(resourceAccessToken, fapiFinancialId string) http.Header {
	header := http.Header{}
	header.Set(api.Accept, api.ApplicationJson)
	header.Set(api.Authorization, "Bearer "+resourceAccessToken)
	header.Set(api.ContentType, api.ApplicationJson)
	header.Set(api.CacheControl, "no-cache")
	header.Set(api.XIdempotencyKey, "test")
	header.Set(api.XFapiFinancialId, fapiFinancialId)
	header.Set(api.XJwsSignature, "test")

	return header
}
