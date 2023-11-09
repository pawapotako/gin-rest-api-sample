package repository

import (
	"context"

	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

type externalServiceRepositoryDB struct {
	db *gorm.DB
}

type ExternalServiceRepository interface {
	VerifyGoogleIdToken(idToken string) (*idtoken.Payload, error)
}

func NewExternalServiceRepositoryDB(db *gorm.DB) ExternalServiceRepository {
	return externalServiceRepositoryDB{db}
}

func (r externalServiceRepositoryDB) VerifyGoogleIdToken(idToken string) (*idtoken.Payload, error) {

	payload, err := idtoken.Validate(context.Background(), idToken, "")
	if err != nil {
		return nil, err
	}
	return payload, nil
}
