package repositories

import (
	"fmt"
	"github.com/mavricknz/ldap"
	"ldap-auth/config"
	"strconv"
	"strings"
)


func ConnectToLDAP() (*ldap.LDAPConnection, error) {

	host := config.C.Ldap.Host
	portStr := config.C.Ldap.Port
	user := config.C.Ldap.User
	password := config.C.Ldap.Password

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid LDAP port: %v", err)
	}

	ldapConn := ldap.NewLDAPConnection(host, uint16(port))
	ldapConn.Debug = false

	if err := ldapConn.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to LDAP: %v", err)
	}

	if err := ldapConn.Bind(user, password); err != nil {
		ldapConn.Close()
		return nil, fmt.Errorf("LDAP bind failed: %v", err)
	}

	return ldapConn, nil
}

func GetUserDN(username string) (string, error) {
	ldapConn, err := ConnectToLDAP()
	if err != nil {
		return "", err
	}
	defer ldapConn.Close()

	searchReq := ldap.NewSimpleSearchRequest(
		config.C.Ldap.Base,
		ldap.ScopeWholeSubtree,
		fmt.Sprintf("(userPrincipalName=%s)", username),
		[]string{"dn"},
	)

	result, err := ldapConn.Search(searchReq)
	if err != nil || len(result.Entries) == 0 {
		return "", fmt.Errorf("user not found in LDAP")
	}

	return result.Entries[0].DN, nil
}

func GetUserRole(username string) (string, error) {
	ldapConn, err := ConnectToLDAP()
	if err != nil {
		return "", err
	}
	defer ldapConn.Close()

	searchReq := ldap.NewSimpleSearchRequest(
		config.C.Ldap.Base,
		ldap.ScopeWholeSubtree,
		fmt.Sprintf("(userPrincipalName=%s)", username),
		[]string{"memberOf"},
	)

	result, err := ldapConn.Search(searchReq)
	if err != nil || len(result.Entries) == 0 {
		return "", fmt.Errorf("user not found in LDAP")
	}

	for _, attr := range result.Entries[0].Attributes {
		for _, val := range attr.Values {
			if strings.Contains(val, "CN=YOUR-ADMIN-CN") {
				return "admin", nil
			} else if strings.Contains(val, "CN=YOUR-USER-CN") {
				return "user", nil
			}
		}
	}

	return "", fmt.Errorf("role not found")
}
