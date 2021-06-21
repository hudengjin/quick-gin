package auth

type Auth interface {
	login() (bool, error)
	logout() (bool, error)
	
}