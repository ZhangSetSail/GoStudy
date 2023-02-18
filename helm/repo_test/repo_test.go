package main

func main() {

}

type repoAddOptions struct {
	name                 string
	url                  string
	username             string
	password             string
	passwordFromStdinOpt bool
	passCredentialsAll   bool
	forceUpdate          bool
	allowDeprecatedRepos bool

	certFile              string
	keyFile               string
	caFile                string
	insecureSkipTLSverify bool

	repoFile  string
	repoCache string

	// Deprecated, but cannot be removed until Helm 4
	deprecatedNoUpdate bool
}