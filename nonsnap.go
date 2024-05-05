package doku

import (
	"errors"
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
	"time"
)

func (c *Client) SendRequest(method, targetUri string, requestData interface{}, result interface{}) error {
	requestId := c.requestId

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}

	requestTimestamp := time.Now().In(loc).Format("2006-01-02T15:04:05Z")

	signature, err := c.generateRequestSignature(targetUri, requestTimestamp, requestData)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Content-Type":      "application/json",
		"Accept":            "application/json",
		"Client-Id":         c.Config.ClientID,
		"Request-Id":        requestId,
		"Request-Timestamp": requestTimestamp,
		"Signature":         *signature,
	}

	targetUrl := fmt.Sprintf("%s%s", c.Config.BaseUrl, targetUri)

	if _, err = goutil.SendHttpRequest(method, targetUrl, requestData, &headers, result); err != nil {
		return err
	}

	return nil
}

func (c *Client) Checkout(requestData Request) (*CheckoutResponse, error) {
	var result CheckoutResponse

	if err := c.SendRequest(http.MethodPost, TargetPathCheckout, &requestData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CheckStatus(invoiceNumber string) (*CheckStatusResponse, error) {
	var result CheckStatusResponse

	targetUri := fmt.Sprintf("%s/%s", TargetPathCheckStatus, invoiceNumber)

	if err := c.SendRequest(http.MethodGet, targetUri, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateVirtualAccount(vaType string, requestData Request) (*CreateVirtualAccountResponse, error) {
	var result CreateVirtualAccountResponse

	targetUri := ""

	switch vaType {
	default:
		return nil, errors.New("invalid virtual account type")
	case VirtualAccountBCA:
		targetUri = TargetPathBCAVirtualAccount
	case VirtualAccountMandiri:
		targetUri = TargetPathMandiriVirtualAccount
	case VirtualAccountBRI:
		targetUri = TargetPathBRIVirtualAccount
	case VirtualAccountBSI:
		targetUri = TargetPathBSIVirtualAccount
	case VirtualAccountDOKU:
		targetUri = TargetPathDOKUVirtualAccount
	}

	if err := c.SendRequest(http.MethodPost, targetUri, &requestData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
