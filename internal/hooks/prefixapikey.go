package hooks

import (
	"fmt"
	"net/http"
	"strings"
)

type PrefixApiKey struct{}

const prefix = "Token"

var (
	_ beforeRequestHook = (*PrefixApiKey)(nil)
)

func (i *PrefixApiKey) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	token := req.Header.Get("Authorization")

	if strings.HasPrefix(token, prefix) {
		return req, nil
	}

	req.Header.Set("Auhorization", fmt.Sprintf("Token %s", token))

	return req, nil
}
