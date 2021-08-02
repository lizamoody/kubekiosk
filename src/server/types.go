package server

import (
	"crypto/tls"
	"log"
	"sync"

	"github.com/go-ldap/ldap/v3"
)

type refreshData struct {
	FederatedID string     `json:"federated_id"`
	Username    string     `json:"username`
	Entry       ldap.Entry `json:"entry"`
}

//a host configuration
type hostConfig struct {
	host      string
	tlsConfig *tls.Config
}

//Password Connector
type Connector struct {
	c                *proto.IamConnector
	b                backend.Backend
	l                log.Logger
	config           *Config
	syncMx           sync.Mutex
	hosts            []*hostConfig
	userSearchScope  int
	groupSearchScope int
}
