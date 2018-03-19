package clients

import (
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"

	"context"

	"net/http"

	"net/url"

	"fmt"

	btypes "git.containerum.net/ch/json-types/billing"
	rstypes "git.containerum.net/ch/json-types/resource-service"
	"git.containerum.net/ch/kube-client/pkg/cherry"
	"git.containerum.net/ch/kube-client/pkg/cherry/adaptors/cherrylog"
	"git.containerum.net/ch/utils"
	"gopkg.in/resty.v1"
)

// Billing is an interface to billing service
type Billing interface {
	Subscribe(ctx context.Context, userID string, resource rstypes.Resource, resourceKind rstypes.Kind) error
	Unsubscribe(ctx context.Context, userID string, resource rstypes.Resource) error

	GetNamespaceTariff(ctx context.Context, tariffID string) (btypes.NamespaceTariff, error)
	GetVolumeTariff(ctx context.Context, tariffID string) (btypes.VolumeTariff, error)

	//ActivateNamespaceTariff(ctx context.Context, ...)
	//ActivateVolumeTariff(ctx, ...)
}

// Data for dummy client

type DummyBillingClient struct {
	log *cherrylog.LogrusAdapter
}

var fakeNSData = `
[
  {
    "id": "f3091cc9-6dc3-470e-ac54-84defe011111",
    "created_at": "2017-12-26T13:53:56Z",
    "cpu_limit": 500,
    "memory_limit": 512,
    "traffic": 20,
    "traffic_price": 0.333,
    "external_services": 2,
    "internal_services": 5,
    "is_active": true,
    "is_public": true,
    "price": 0
  },
  {
    "id": "4563e8c1-fb41-416a-9798-e949a2616260",
    "created_at": "2017-12-26T13:57:45Z",
    "cpu_limit": 900,
    "memory_limit": 1024,
    "traffic": 50,
    "traffic_price": 0.5,
    "external_services": 10,
    "internal_services": 20,
    "is_active": true,
    "is_public": true,
    "price": 0
  }
]
`

var fakeVolumeData = `
[
  {
    "id": "15348470-e98f-4da0-8d2e-8c65e15d6eeb",
    "created_at": "2017-12-27T07:55:22Z",
    "storage_limit": 1,
    "replicas_limit": 2,
    "is_persistent": false,
    "is_active": true,
    "is_public": true,
    "price": 0
  },
  {
    "id": "11a35f90-c343-4fc1-a966-381f75568036",
    "created_at": "2017-12-27T07:55:22Z",
    "storage_limit": 2,
    "replicas_limit": 2,
    "is_persistent": false,
    "is_active": true,
    "is_public": true,
    "price": 0
  }
]
`

var (
	fakeNSTariffs     []btypes.NamespaceTariff
	fakeVolumeTariffs []btypes.VolumeTariff
)

func init() {
	var err error
	err = jsoniter.Unmarshal([]byte(fakeNSData), &fakeNSTariffs)
	if err != nil {
		panic(err)
	}
	err = jsoniter.Unmarshal([]byte(fakeVolumeData), &fakeVolumeTariffs)
	if err != nil {
		panic(err)
	}
}

var buildErr = cherry.BuildErr(cherry.Billing) //FIXME: add package "billing" to "cherry"

func nsTariffNotFound() *cherry.Err {
	return buildErr("namespace tariff not found", http.StatusNotFound, 1)
}

func volTarifNotFound() *cherry.Err {
	return buildErr("volume tariff not found", http.StatusNotFound, 2)
}

type BillingHTTP struct {
	client *resty.Client
	log    *cherrylog.LogrusAdapter
}

func NewHTTPBillingClient(u *url.URL) *BillingHTTP {
	log := logrus.WithField("component", "billing_client")
	client := resty.New().
		SetHostURL(u.String()).
		SetLogger(log.WriterLevel(logrus.DebugLevel)).
		SetDebug(true).
		SetError(cherry.Err{}).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	return &BillingHTTP{
		client: client,
		log:    cherrylog.NewLogrusAdapter(log),
	}
}

func (b *BillingHTTP) Subscribe(ctx context.Context, userID string, resource rstypes.Resource, resourceKind rstypes.Kind) error {
	b.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"tariff_id":   resource.TariffID,
		"resource_id": resource.ID,
		"kind":        resourceKind,
	}).Infoln("subscribing")

	//TODO: request when method will be implemented

	return nil
}

func (b *BillingHTTP) Unsubscribe(ctx context.Context, userID string, resource rstypes.Resource) error {
	b.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"resource_id": resource.ID,
	}).Infoln("unsubscribing")

	//TODO: request when method will be implemented

	return nil
}

func (b *BillingHTTP) GetNamespaceTariff(ctx context.Context, tariffID string) (btypes.NamespaceTariff, error) {
	b.log.WithField("tariff_id", tariffID).Infoln("get namespace tariff")

	resp, err := resty.R().
		SetContext(ctx).
		SetHeaders(utils.RequestXHeadersMap(ctx)).
		SetResult(btypes.NamespaceTariff{}).
		Get(fmt.Sprintf("/tariffs/namespace/%s", tariffID))
	if err != nil {
		return btypes.NamespaceTariff{}, err
	}
	if resp.Error() != nil {
		return btypes.NamespaceTariff{}, resp.Error().(*cherry.Err)
	}

	return *resp.Result().(*btypes.NamespaceTariff), nil
}

func (b *BillingHTTP) GetVolumeTariff(ctx context.Context, tariffID string) (btypes.VolumeTariff, error) {
	b.log.WithField("tariff_id", tariffID).Infoln("get volume tariff")

	resp, err := resty.R().
		SetContext(ctx).
		SetHeaders(utils.RequestXHeadersMap(ctx)).
		SetResult(btypes.NamespaceTariff{}).
		Get(fmt.Sprintf("/tariffs/volume/%s", tariffID))
	if err != nil {
		return btypes.VolumeTariff{}, err
	}
	if resp.Error() != nil {
		return btypes.VolumeTariff{}, resp.Error().(*cherry.Err)
	}

	return *resp.Result().(*btypes.VolumeTariff), nil
}

func (b BillingHTTP) String() string {
	return fmt.Sprintf("billing service http client: url=%s", b.client.HostURL)
}

// NewDummyBilling creates a dummy billing service client. It does nothing but logs actions.
func NewDummyBillingClient() DummyBillingClient {
	return DummyBillingClient{
		log: cherrylog.NewLogrusAdapter(logrus.WithField("component", "billing_dummy")),
	}
}

func (b DummyBillingClient) Subscribe(ctx context.Context, userID string, resource rstypes.Resource, resourceKind rstypes.Kind) error {
	b.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"tariff_id":   resource.TariffID,
		"resource_id": resource.ID,
		"kind":        resourceKind,
	}).Infoln("subscribing")
	return nil
}

func (b DummyBillingClient) Unsubscribe(ctx context.Context, userID string, resource rstypes.Resource) error {
	b.log.WithFields(logrus.Fields{
		"user_id":     userID,
		"resource_id": resource.ID,
	}).Infoln("unsubscribing")
	return nil
}

func (b DummyBillingClient) GetNamespaceTariff(ctx context.Context, tariffID string) (btypes.NamespaceTariff, error) {
	b.log.WithField("tariff_id", tariffID).Infoln("get namespace tariff")
	for _, nsTariff := range fakeNSTariffs {
		if nsTariff.ID != "" && nsTariff.ID == tariffID {
			return nsTariff, nil
		}
	}
	return btypes.NamespaceTariff{}, nsTariffNotFound()
}

func (b DummyBillingClient) GetVolumeTariff(ctx context.Context, tariffID string) (btypes.VolumeTariff, error) {
	b.log.WithField("tariff_id", tariffID).Infoln("get volume tariff")
	for _, volumeTariff := range fakeVolumeTariffs {
		if volumeTariff.ID != "" && volumeTariff.ID == tariffID {
			return volumeTariff, nil
		}
	}
	return btypes.VolumeTariff{}, volTarifNotFound()
}

func (b DummyBillingClient) String() string {
	return "billing service dummy"
}
