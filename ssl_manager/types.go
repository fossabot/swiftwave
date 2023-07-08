package sslmanager

import (
	"context"
	"time"

	"github.com/mholt/acmez"
	"github.com/mholt/acmez/acme"
	"gorm.io/gorm"
)

type SSLManager struct {
	ctx      context.Context
	account  acme.Account
	client   acmez.Client
	dbClient gorm.DB
	options  SSLManagerOptions
}

type SSLManagerOptions struct {
	Email                     string
	AccountPrivateKeyFilePath string
	DomainPrivateKeyStorePath string
	DomainFullChainStorePath  string
}

type http01Solver struct {
	dbClient gorm.DB
}

// GORM Models
type KeyAuthorizationToken struct {
	Token              string `gorm:"primaryKey"`
	AuthorizationToken string
}

type DomainSSLDetails struct {
	Domain       string `gorm:"primaryKey"`
	CreationDate time.Time
}
