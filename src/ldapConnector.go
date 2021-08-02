package main

import (
	"log"

	"github.com/go-ldap/ldap/v3" // gopkg.in/ldap.v3
)

func verifyCredentials(username string, pass string) bool {
	verified := false
	ldapURL := "ldap://ldap.forumsys.com:389"
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	err = l.Bind("cn=read-only-admin,dc=example,dc=com", "password")
	if err != nil {
		log.Fatal(err)
	}
	baseDN := "DC=example,DC=com"
	filter := "(objectClass=Person)" //fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(user))
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, 0, 0, 0, false, filter, []string{"sAMAccountName"}, []ldap.Control{})
	result_users, err := l.Search(searchReq)
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range result_users.Entries {
		if w.DN == ("uid=" + username + ",dc=example,dc=com") {
			err = l.Bind(w.DN, pass)
			if err != nil {
				log.Println(err)
			} else {
				log.Print("user pass is valid: ")
				verified = true
			}
			// Rebind as the read only user for any further queries
			err = l.Bind("cn=read-only-admin,dc=example,dc=com", "password")
			if err != nil {
				log.Fatal(err)
			}
			return verified
		}
	}
	return verified
}
