package authmanager

import (
	"github.com/kaanaktas/openbanking-paymentinitiation/api"
	"github.com/kaanaktas/openbanking-paymentinitiation/pkg/consent"
	"github.com/kaanaktas/openbanking-paymentinitiation/pkg/token"
	"github.com/pkg/errors"
)

type AuthManager interface {
	GetAuthorisedTokenByCid(aspspId, cid string) (string, string, error)
}

type authManager struct {
	consentServiceRead  consent.ServiceRead
	consentServiceWrite consent.ServiceWrite
	tokenService        token.Service
}

func NewAuthManager(consentServiceRead consent.ServiceRead, consentServiceWrite consent.ServiceWrite, tokenService token.Service) AuthManager {
	return &authManager{
		consentServiceRead:  consentServiceRead,
		consentServiceWrite: consentServiceWrite,
		tokenService:        tokenService,
	}
}

func (s authManager) GetAuthorisedTokenByCid(aspspId, cid string) (string, string, error) {
	consentResp, err := s.consentServiceRead.FindConsentByCidAndStatus(cid, api.Authorised)
	if err != nil || consentResp.AspspId != aspspId {
		return "", "", errors.WithMessagef(err, "couldn't retrieve the consentResp. cid: %v aspspId: %v", cid, aspspId)
	}

	authorisedToken := consentResp.Tokens[0]
	payload := consentResp.ObjectState
	return *authorisedToken.ResourceAccessToken, payload, nil
}
