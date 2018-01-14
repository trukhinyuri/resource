package other

import (
	"fmt"
	"net/url"

	mttypes "git.containerum.net/ch/json-types/mail-templater"
	rstypes "git.containerum.net/ch/json-types/resource-service"

	"context"

	"git.containerum.net/ch/json-types/errors"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

type Mailer interface {
	SendNamespaceCreated(ctx context.Context, userID, nsLabel string, t rstypes.NamespaceTariff) error
	SendNamespaceDeleted(ctx context.Context, userID, nsLabel string, t rstypes.NamespaceTariff) error

	SendVolumeCreated(ctx context.Context, userID, nsLabel string, t rstypes.VolumeTariff) error
	SendVolumeDeleted(ctx context.Context, userID, nsLabel string, t rstypes.VolumeTariff) error
}

type mailerHTTP struct {
	client *resty.Client
	log    *logrus.Entry
}

func NewMailerHTTP(u *url.URL) Mailer {
	log := logrus.WithField("component", "mail_client")
	client := resty.New().
		SetHostURL(u.String()).
		SetLogger(log.WriterLevel(logrus.DebugLevel)).
		SetDebug(true).
		SetError(errors.Error{})
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	return mailerHTTP{
		client: client,
		log:    log,
	}
}

func (ml mailerHTTP) sendRequest(ctx context.Context, eventName string, userID string, vars map[string]interface{}) error {
	ml.log.WithFields(logrus.Fields{
		"event":   eventName,
		"user_id": userID,
	}).Infof("sending mail with vars %v", vars)
	resp, err := ml.client.R().SetContext(ctx).SetBody(mttypes.SimpleSendRequest{
		Template:  eventName,
		UserID:    userID,
		Variables: vars,
	}).SetResult(mttypes.SimpleSendResponse{}).Post("/send")
	if err != nil {
		return err
	}
	if resp.Error() != nil {
		return resp.Error().(*errors.Error)
	}
	result := resp.Result().(*mttypes.SimpleSendResponse)
	ml.log.WithField("user_id", result.UserID).Infoln("sent mail")
	return nil
}

func (ml mailerHTTP) SendNamespaceCreated(ctx context.Context, userID, nsLabel string, t rstypes.NamespaceTariff) error {
	err := ml.sendRequest(ctx, "ns_created", userID, map[string]interface{}{
		"NAMESPACE": nsLabel,
		"CPU":       t.CpuLimit,
		"RAM":       t.MemoryLimit,
		"DAILY_PAY": t.Price,
		//"DAILY_PAY_TOTAL": 0, // FIXME
		//"STORAGE": 0, // FIXME
	})
	if err != nil {
		return err
	}
	return nil
}

func (ml mailerHTTP) SendNamespaceDeleted(ctx context.Context, userID, nsLabel string, t rstypes.NamespaceTariff) error {
	err := ml.sendRequest(ctx, "ns_deleted", userID, map[string]interface{}{
		"NAMESPACE": nsLabel,
	})
	if err != nil {
		return err
	}
	return nil
}

func (ml mailerHTTP) SendVolumeCreated(ctx context.Context, userID, volLabel string, t rstypes.VolumeTariff) error {
	err := ml.sendRequest(ctx, "vol_created", userID, map[string]interface{}{
		"VOLUME":    volLabel,
		"STORAGE":   t.StorageLimit,
		"DAILY_PAY": t.Price,
		//"DAILY_PAY_TOTAL": 0, // FIXME
	})
	if err != nil {
		return err
	}
	return nil
}

func (ml mailerHTTP) SendVolumeDeleted(ctx context.Context, userID, volLabel string, t rstypes.VolumeTariff) error {
	err := ml.sendRequest(ctx, "vol_deleted", userID, map[string]interface{}{
		"VOLUME": volLabel,
	})
	if err != nil {
		return err
	}
	return nil
}

func (ml mailerHTTP) String() string {
	return fmt.Sprintf("mail service http client: url=%v", ml.client.HostURL)
}

type mailerStub struct {
	log *logrus.Entry
}

func NewMailerStub() Mailer {
	return mailerStub{log: logrus.WithField("component", "mailer_stub")}
}

func (ml mailerStub) SendNamespaceCreated(ctx context.Context, userID, nsLabel string, t rstypes.NamespaceTariff) error {
	ml.log.WithFields(logrus.Fields{
		"user_id":  userID,
		"ns_label": nsLabel,
	}).Infof("send namespace created with tariff %+v", t)
	return nil
}

func (ml mailerStub) SendNamespaceDeleted(ctx context.Context, userID, nsLabel string, t rstypes.NamespaceTariff) error {
	ml.log.WithFields(logrus.Fields{
		"user_id":  userID,
		"ns_label": nsLabel,
	}).Infof("send namespace deleted with tariff %+v", t)
	return nil
}

func (ml mailerStub) SendVolumeCreated(ctx context.Context, userID, label string, t rstypes.VolumeTariff) error {
	ml.log.WithFields(logrus.Fields{
		"user_id":   userID,
		"vol_label": label,
	}).Infof("send volume created with tariff %+v", t)
	return nil
}

func (ml mailerStub) SendVolumeDeleted(ctx context.Context, userID, label string, t rstypes.VolumeTariff) error {
	ml.log.WithFields(logrus.Fields{
		"user_id":   userID,
		"vol_label": label,
	}).Infof("send volume deleted with tariff %+v", t)
	return nil
}

func (mailerStub) String() string {
	return "mail service dummy"
}
