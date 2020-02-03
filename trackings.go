package aftership

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DeliveryType string
type Tag string

const (
	PickupAtStore   DeliveryType = "pickup_at_store"
	PickupAtCourier              = "pickup_at_courier"
	DoorToDoor                   = "door_to_door"

	Pending        Tag = "Pending"
	InfoReceived       = "InfoReceived"
	InTransit          = "InTransit"
	OutForDelivery     = "OutForDelivery"
	AttemptFail        = "AttemptFail"
	Delivered          = "Delivered"
	Exception          = "Exception"
	Expired            = "Expired"
)

type Tracking struct {
	ID                         string                 `json:"id"`
	CreatedAt                  string                 `json:"created_at"`
	UpdatedAt                  string                 `json:"updated_at"`
	LastUpdatedAt              string                 `json:"last_updated_at"`
	Active                     bool                   `json:"active"`
	TrackingNumber             string                 `json:"tracking_number"`
	UniqueToken                string                 `json:"unique_token"`
	Slug                       interface{}            `json:"slug,omitempty"`
	TrackingPostalCode         string                 `json:"tracking_postal_code,omitempty"`
	TrackingShipDate           string                 `json:"tracking_ship_date,omitempty"`
	TrackingAccountNumber      string                 `json:"tracking_account_number,omitempty"`
	TrackingKey                string                 `json:"tracking_key,omitempty"`
	TrackingOriginCountry      string                 `json:"tracking_origin_country,omitempty"`
	TrackingDestinationCountry string                 `json:"tracking_destination_country,omitempty"`
	TrackingState              string                 `json:"tracking_state,omitempty"`
	Android                    interface{}            `json:"android,omitempty"`
	IOS                        interface{}            `json:"ios,omitempty"`
	EMails                     interface{}            `json:"emails,omitempty"`
	SubscribedEMails           []string               `json:"subscribed_emails,omitempty"`
	SMSes                      interface{}            `json:"smses,omitempty"`
	SubscribedSMSes            []string               `json:"subscribed_smses,omitempty"`
	Title                      string                 `json:"title,omitempty"`
	CustomerName               string                 `json:"customer_name,omitempty"`
	DeliveryTime               int                    `json:"delivery_time"`
	ExpectedDelivery           string                 `json:"expected_delivery"`
	OriginCountryISO3          string                 `json:"origin_country_iso3,omitempty" validate:"len=3"`
	DestinationCountryISO3     string                 `json:"destination_country_iso3,omitempty" validate:"len=3"`
	OrderID                    string                 `json:"order_id,omitempty"`
	OrderIDPath                string                 `json:"order_id_path,omitempty"`
	CustomFields               map[string]interface{} `json:"custom_fields,omitempty"`
	Note                       string                 `json:"note,omitempty"`
	Language                   string                 `json:"language,omitempty"`
	OrderPromisedDeliveryDate  string                 `json:"order_promised_delivery_date,omitempty"`
	DeliveryType               DeliveryType           `json:"delivery_type,omitempty"`
	PickupLocation             string                 `json:"pickup_location,omitempty"`
	PickupNote                 string                 `json:"pickup_note,omitempty"`
	ShipmentPackageCount       int                    `json:"shipment_package_count,omitempty"`
	ShipmentType               string                 `json:"shipment_type,omitempty"`
	ShipmentWeight             float64                `json:"shipment_weight,omitempty"`
	ShipmentWeightUnit         string                 `json:"shipment_weight_unit,omitempty"`
	ShipmentPickupDate         string                 `json:"shipment_pickup_date,omitempty"`
	ShipmentDeliveryDate       string                 `json:"shipment_delivery_date,omitempty"`
	SignedBy                   string                 `json:"signed_by,omitempty"`
	Source                     string                 `json:"source,omitempty"`
	Tag                        Tag                    `json:"tag,omitempty"`
	Subtag                     string                 `json:"subtag,omitempty"`
	SubtagMessage              string                 `json:"subtag_message,omitempty"`
	TrackedCount               int                    `json:"tracked_count,omitempty"`
	LastMileTrackingSupported  *bool                  `json:"last_mile_tracking_supported,omitempty"`
	ReturnToSender             bool                   `json:"return_to_sender,omitempty"`
	Checkpoints                []Checkpoint           `json:"checkpoints"`
}

type Checkpoint struct {
	CreatedAt      time.Time `json:"created_at"`
	Slug           string    `json:"slug"`
	CheckpointTime string    `json:"checkpoint_time,omitempty"`
	Location       string    `json:"location,omitempty"`
	City           string    `json:"city,omitempty"`
	State          string    `json:"state,omitempty"`
	CountryISO3    string    `json:"country_iso3,omitempty"`
	CountryName    string    `json:"country_name,omitempty"`
	Zip            string    `json:"zip,omitempty"`
	Tag            Tag       `json:"tag"`
	Message        string    `json:"message"`
	Subtag         string    `json:"subtag"`
	SubtagMessage  string    `json:"subtag_message"`
}

type CreateTracking struct {
	TrackingNumber             string                 `json:"tracking_number" validate:"required"`
	Slug                       interface{}            `json:"slug,omitempty"`
	TrackingPostalCode         string                 `json:"tracking_postal_code,omitempty"`
	TrackingShipDate           string                 `json:"tracking_ship_date,omitempty"`
	TrackingAccountNumber      string                 `json:"tracking_account_number,omitempty"`
	TrackingKey                string                 `json:"tracking_key,omitempty"`
	TrackingOriginCountry      string                 `json:"tracking_origin_country,omitempty"`
	TrackingDestinationCountry string                 `json:"tracking_destination_country,omitempty"`
	TrackingState              string                 `json:"tracking_state,omitempty"`
	Android                    interface{}            `json:"android,omitempty"`
	IOS                        interface{}            `json:"ios,omitempty"`
	EMails                     interface{}            `json:"emails,omitempty"`
	SMSes                      interface{}            `json:"smses,omitempty"`
	Title                      string                 `json:"title,omitempty"`
	CustomerName               string                 `json:"customer_name,omitempty"`
	OriginCountryISO3          string                 `json:"origin_country_iso3,omitempty" validate:"len=3"`
	DestinationCountryISO3     string                 `json:"destination_country_iso3,omitempty" validate:"len=3"`
	OrderID                    string                 `json:"order_id,omitempty"`
	OrderIDPath                string                 `json:"order_id_path,omitempty"`
	CustomFields               map[string]interface{} `json:"custom_fields,omitempty"`
	Note                       string                 `json:"note,omitempty"`
	Language                   string                 `json:"language,omitempty"`
	OrderPromisedDeliveryDate  string                 `json:"order_promised_delivery_date,omitempty"`
	DeliveryType               DeliveryType           `json:"delivery_type,omitempty"`
	PickupLocation             string                 `json:"pickup_location,omitempty"`
	PickupNote                 string                 `json:"pickup_note,omitempty"`
}

func (a *AfterShip) CreateTracking(ctx context.Context, create CreateTracking) (*Tracking, error) {
	res, err := a.prepareAndSend(ctx, http.MethodPost, "/trackings", struct {
		Tracking CreateTracking `json:"tracking"`
	}{
		Tracking: create,
	})
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		return nil, formatError(ErrUnexpectedResponseStatus, res)
	}

	var body struct {
		Data struct {
			Tracking Tracking `json:"tracking"`
		} `json:"data"`
	}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return &body.Data.Tracking, nil
}

func (a *AfterShip) DeleteTracking(ctx context.Context, slug, trackingNumber string) error {
	res, err := a.prepareAndSend(ctx, http.MethodDelete, fmt.Sprintf("/trackings/%s/%s", slug, trackingNumber), nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return formatError(ErrUnexpectedResponseStatus, res)
	}

	return nil
}
