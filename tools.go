package Tools

import (
	"encoding/base64"
)

type Tools struct {
}

func (t *Tools) BasicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
