package server

import "log"

const (
	defaultLdapPort  = "389"
	defaultLdapsPort = "636"
)

//new password connector

func NewConnector(c *proto.IamConnector, b backend.Backend, l log.Logger) {

}

//Refresh token identity
func (c *Connector) RefreshIdentity()
