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

	//bindUsername := "cn=asantana,dc=umag,dc=cl" // Reemplaza con tu nombre de usuario
	bindUsername := "uid=" + userName + ",ou=usuarios,dc=umag,dc=cl" // Reemplaza con tu nombre de usuario
	bindPassword := password                                         // Reemplaza con tu contraseña

	err = l.Bind(bindUsername, bindPassword)
	if err != nil {
		log.Fatal(err)
	}

	userDN := "uid=" + userName + ",ou=usuarios,dc=umag,dc=cl"        // Reemplaza con el DN del usuario
	attributes := []string{"gecos", "employeetype", "employeenumber"} // Reemplaza con los atributos que necesitas

	// Realiza la búsqueda del usuario
	searchRequest := ldap.NewSearchRequest(
		userDN,               // Base DN (Distinguished Name)
		ldap.ScopeBaseObject, // Alcance de búsqueda (base)
		ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)", // Filtro de búsqueda (todos los objetos)
		attributes,        // Atributos a recuperar
		nil,
	)

	// Realiza la búsqueda
	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Imprime los resultados
	for _, entry := range sr.Entries {
		fmt.Println("DN:", entry.DN)
		for _, attr := range entry.Attributes {
			fmt.Printf("%s: %v\n", attr.Name, attr.Values)
		}
	}
}
