package cio

import (
	"fmt"

	"github.com/suicidegang/chttp"
)

const CIO_LITE_API = "https://api.context.io/lite/"

type ContextIOLite struct {
	token  string
	secret string
}

func (c ContextIOLite) GET(url string, f ...interface{}) (Req, error) {
	return Request(
		GET(fmt.Sprintf(CIO_LITE_API+url, f...)),
		Header("Content-Type", "application/x-www-form-urlencoded"),
		Header("Accept", "application/json"),
		Header("Accept-Charset", "utf-8"),
		Header("User-Agent", "suicidegang/chttp.context-io"),
		Oauth(c.token, c.secret),
	)
}

func ContextIO(token, secret string) ContextIOLite {
	return ContextIOLite{token, secret}
}
