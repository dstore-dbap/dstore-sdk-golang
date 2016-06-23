package credentials

import (
	"golang.org/x/net/context"
)

type DstoreCredentials struct {
	Username string
	Password string
	AccessToken string
}

func (p DstoreCredentials) RequireTransportSecurity() bool {
	return true
}

func (p DstoreCredentials) GetRequestMetadata(ctx context.Context, s ...string) (map[string]string, error) {
	return map[string]string{
		"Username": p.Username,
		"Password": p.Password,
		"AccessToken" : p.AccessToken,
	}, nil
}