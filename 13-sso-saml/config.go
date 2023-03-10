package main

import "fmt"

var (
	samlCertificatePath = "./myservice.cert"
	samlPrivateKeyPath  = "./myservice.key"
	samlIDPMetadata     = "http://samltest.id/saml/idp"

	webServerPort   = 9123
	webServerRootUR = fmt.Sprintf("http://localhost:%d", webServerPort)
)
