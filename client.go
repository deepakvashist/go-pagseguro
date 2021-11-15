package pags

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// Client represents PagSeguro v4 API client.
type Client struct {
	httpClient *resty.Client
}

// Charge is responsible for creating a new charge (credit card/boleto).
func (c Client) Charge(charge *Charge) (*Charge, error) {
	response, err := c.request("/charges", charge)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Capture is responsible for capturing an authorized transaction.
func (c Client) Capture(tid string, charge *Charge) (*Charge, error) {
	resource := fmt.Sprintf("charges/%s/capture", tid)
	response, err := c.request(resource, charge)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Cancel is responsible for canceling a transaction.
func (c Client) Cancel(tid string, charge *Charge) (*Charge, error) {
	resource := fmt.Sprintf("charges/%s/cancel", tid)
	response, err := c.request(resource, charge)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c Client) request(resource string, body *Charge) (*Charge, error) {
	response, err := c.httpClient.R().SetBody(body).Post(resource)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("invalid response status code: %d. %s", response.StatusCode(), response.Body())
	}

	charge := Charge{}
	if err := json.Unmarshal(response.Body(), &charge); err != nil {
		return nil, err
	}

	return &charge, nil
}

// NewClient returns PagSeguro v4 API client.
func NewClient(baseURL, token string, retryCount int, timeout, retryWait, retryMaxWait time.Duration) *Client {
	httpClient := resty.New().
		SetHostURL(baseURL).
		SetHeader("Content-Type", "application/json").
		SetHeader("x-api-version", "4.0").
		SetHeader("Authorization", token)

	return &Client{
		httpClient: httpClient,
	}
}
