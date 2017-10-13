package cio

import (
	"fmt"
	"net/url"

	"github.com/suicidegang/chttp"
)

const CIO_LITE_API = "https://api.context.io/lite/"

func ContextIO(token, secret string) ContextIOLite {
	return ContextIOLite{token, secret}
}

type ContextIOLite struct {
	token  string
	secret string
}

func (c ContextIOLite) GET(url string, f ...interface{}) chttp.Req {
	return chttp.RequestOrPanic(
		chttp.GET(fmt.Sprintf(CIO_LITE_API+url, f...)),
		chttp.Header("Content-Type", "application/x-www-form-urlencoded"),
		chttp.Header("Accept", "application/json"),
		chttp.Header("Accept-Charset", "utf-8"),
		chttp.Header("User-Agent", "suicidegang/chttp.context-io"),
		chttp.Oauth(c.token, c.secret),
	)
}

func (c ContextIOLite) User(id string) User {
	return User{c, id}
}

/*req, err := ctx.GET("users/%s/email_accounts/%s/folders/%s/messages/%s/body", webhook.AccountID, s.Label, url.QueryEscape(s.Folder), id)
if err != nil {
	c.StatusCode(iris.StatusBadRequest)
	c.WriteString(err.Error())
	return
}*/

type User struct {
	Client    ContextIOLite
	AccountID string
}

func (u User) EmailAccount(label string) EmailAccount {
	return EmailAccount{u, label}
}

type EmailAccount struct {
	User  User
	Label string
}

func (e EmailAccount) Folder(label string) Folder {
	return Folder{e, label}
}

type Folder struct {
	EmailAccount EmailAccount
	Label        string
}

func (f Folder) Message(id string) Message {
	return Message{f, id}
}

type Message struct {
	Folder Folder
	Id     string
}

func (m Message) Remote() chttp.Req {
	c := m.Folder.EmailAccount.User.Client
	return c.GET(
		"users/%s/email_accounts/%s/folders/%s/messages/%s",
		m.Folder.EmailAccount.User.AccountID,
		m.Folder.EmailAccount.Label,
		url.QueryEscape(m.Folder.Label),
		m.Id,
	)
}

func (m Message) Body() chttp.Req {
	c := m.Folder.EmailAccount.User.Client
	return c.GET(
		"users/%s/email_accounts/%s/folders/%s/messages/%s/body",
		m.Folder.EmailAccount.User.AccountID,
		m.Folder.EmailAccount.Label,
		url.QueryEscape(m.Folder.Label),
		m.Id,
	)
}

func (m Message) Attachments() chttp.Req {
	c := m.Folder.EmailAccount.User.Client
	return c.GET(
		"users/%s/email_accounts/%s/folders/%s/messages/%s/attachments",
		m.Folder.EmailAccount.User.AccountID,
		m.Folder.EmailAccount.Label,
		url.QueryEscape(m.Folder.Label),
		m.Id,
	)
}
