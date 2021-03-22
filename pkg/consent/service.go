package consent

import (
	"github.com/kaanaktas/openbanking-paymentinitiation/api"
)

type ServiceRead interface {
	FindAuthorisedConsentByUserIdAndTppId(userId, tppId string) ([]ActiveConsent, error)
	FindByCid(consentId string) (*Consent, error)
	FindConsentByCidAndStatus(cid, status string) (*Consent, error)
	FindByTrackingId(trackingId string) (*Consent, error)
}
type ServiceWrite interface {
	ChangeConsentStateByCid(cid, status string) error
	InvalidateAuthorisedTokenByConsentTid(tid int64, status string) error
	SaveToken(token *Token) error
	SaveConsent(consent *Consent) error
}

type serviceRead struct {
	repo RepositoryRead
}

type serviceWrite struct {
	repo RepositoryWrite
}

func NewServiceRead(repo RepositoryRead) ServiceRead {
	return &serviceRead{
		repo: repo,
	}
}

func NewServiceWrite(repo RepositoryWrite) ServiceWrite {
	return &serviceWrite{
		repo: repo,
	}
}

func (sw serviceWrite) SaveConsent(consent *Consent) error {
	return sw.repo.saveConsent(consent)
}

func (sw serviceWrite) SaveToken(token *Token) error {
	return sw.repo.saveToken(token)
}

func (sw serviceWrite) InvalidateAuthorisedTokenByConsentTid(tid int64, status string) error {
	return sw.repo.invalidateAuthorisedTokenByConsentTid(tid, status)
}

func (sw serviceWrite) ChangeConsentStateByCid(cid, status string) error {
	return sw.repo.changeConsentStateByCid(cid, status)
}

func (sr serviceRead) FindConsentByCidAndStatus(cid, status string) (*Consent, error) {
	return sr.repo.findConsentByCidAndStatus(cid, status)
}

func (sr serviceRead) FindByCid(cid string) (*Consent, error) {
	return sr.repo.findByCid(cid)
}

func (sr serviceRead) FindByTrackingId(trackingId string) (*Consent, error) {
	return sr.repo.findByTrackingId(trackingId)
}

func (sr serviceRead) FindAuthorisedConsentByUserIdAndTppId(userId, tppId string) ([]ActiveConsent, error) {
	consents, err := sr.repo.findConsentByUserIdAndTppIdAndStatus(userId, tppId, api.Authorised)
	if err != nil {
		return nil, err
	}

	var activeConsents []ActiveConsent
	for _, consent := range consents {
		activeConsents = append(activeConsents, ActiveConsent{consent.Id, consent.AspspId})
	}

	return activeConsents, nil
}
