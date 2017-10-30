package other

import (
	"net/http"
	"net/url"

	"bitbucket.org/exonch/resource-manager/server/model"
)

type Mailer interface {
	SendNamespaceCreated(user model.User, label string, tariff model.Tariff) error
	SendNamespaceDeleted(user model.User, label string) error
}

type mailerHTTP struct {
	c *http.Client
	u *url.URL
}

func (ml mailerHTTP) SendNamespaceCreated(user model.User, label string, tariff model.Tariff) error {
	return nil
}

func (ml mailerHTTP) SendNamespaceDeleted(user model.User, label string) error {
	return nil
}
