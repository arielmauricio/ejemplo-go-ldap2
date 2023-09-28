package main

import (
	"fmt"
	"log"

	"gopkg.in/ldap.v2"
)

func main() {
	ldapServer := "sid.umag.cl:389"
	l, err := ldap.Dial("tcp", ldapServer)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	userName := "coloque su usario aqui"
	password := "coloque su clave aqui"

	//bindUsername := "cn=asantana,dc=umag,dc=cl"
	bindUsername := "uid=" + userName + ",ou=usuarios,dc=umag,dc=cl"
	bindPassword := password

	err = l.Bind(bindUsername, bindPassword)
	if err != nil {
		log.Fatal(err)
	}

	userDN := "uid=" + userName + ",ou=usuarios,dc=umag,dc=cl"
	attributes := []string{"gecos", "employeetype", "employeenumber"}

	searchRequest := ldap.NewSearchRequest(
		userDN,
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)",
		attributes,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		fmt.Println("DN:", entry.DN)
		for _, attr := range entry.Attributes {
			fmt.Printf("%s: %v\n", attr.Name, attr.Values)
		}
	}
}
