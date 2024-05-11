package doku

import "encoding/json"

type Client struct {
	Config    *Config
	requestId string
}

type Config struct {
	BaseUrl       string `json:"base_url"`
	ClientID      string `json:"client_id"`
	SecretKey     string `json:"secret_key"`
	PublicKeyPath string `json:"doku_public_key_path"`
}

type Order struct {
	InvoiceNumber       string      `json:"invoice_number"`
	Amount              json.Number `json:"amount"`
	MinAmount           json.Number `json:"min_amount,omitempty"`
	MaxAmount           json.Number `json:"max_amount,omitempty"`
	Currency            string      `json:"currency,omitempty"`
	CallbackUrl         string      `json:"callback_url,omitempty"`
	CallbackUrlCancel   string      `json:"callback_url_cancel,omitempty"`
	Language            string      `json:"language,omitempty"`
	AutoRedirect        bool        `json:"auto_redirect,omitempty"`
	DisableRetryPayment bool        `json:"disable_retry_payment,omitempty"`
	SessionId           string      `json:"session_id,omitempty"`
}

type Payment struct {
	PaymentDueDate     int      `json:"payment_due_date"` // in minutes
	PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
	TokenId            string   `json:"token_id,omitempty"`
	Url                string   `json:"url"`
	ExpiredDate        string   `json:"expired_date"`
}

type Transaction struct {
	Status            string `json:"status"`
	OriginalRequestId string `json:"original_request_id"`
}

type Service struct {
	Id string `json:"id"`
}

type Acquirer struct {
	Id string `json:"id"`
}

type Channel struct {
	Id string `json:"id"`
}

type Customer struct {
	Id       string `json:",omitempty"`
	Name     string `json:"name"`
	LastName string `json:"last_name,omitempty"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Address  string `json:"address,omitempty"`
	Postcode string `json:"postcode,omitempty"`
	State    string `json:"state,omitempty"`
	City     string `json:"city,omitempty"`
	Country  string `json:"country,omitempty"`
}

type AdditionalInfo struct {
	AllowTenor          []int  `json:"allow_tenor,omitempty"`
	CloseRedirect       string `json:"close_redirect,omitempty"`
	DokuWalletNotifyUrl string `json:"doku_wallet_notify_url,omitempty"`
	Origin              struct {
		Product            string `json:"product"`
		System             string `json:"system"`
		Source             string `json:"source"`
		StandardApiVersion string `json:"standardApiVersion"`
	} `json:"origin,omitempty"`
}

type VirtualAccountInfo struct {
	BillingType          string `json:"billing_type"`
	VirtualAccountNumber string `json:"virtual_account_number,omitempty"`
	CreatedDate          string `json:"created_date,omitempty"`
	CreatedDateUtc       string `json:"created_date_utc,omitempty"`
	ExpiredDate          string `json:"expired_date,omitempty"`
	ExpiredDateUtc       string `json:"expired_date_utc,omitempty"`
	ExpiredTime          int    `json:"expired_time,omitempty"`
	ReusableStatus       bool   `json:"reusable_status"`
	Info1                string `json:"info1,omitempty"`
	Info2                string `json:"info2,omitempty"`
	Info3                string `json:"info3,omitempty"`
	HowToPayPage         string `json:"how_to_pay_page,omitempty"`
	HowToPayApi          string `json:"how_to_pay_api,omitempty"`
}

type VirtualAccountPayment struct {
	Identifier []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"identifier"`
}

type Headers struct {
	RequestId string `json:"request_id"`
	Signature string `json:"signature"`
	Date      string `json:"date"`
	ClientId  string `json:"client_id"`
}

type Request struct {
	Order              Order               `json:"order"`
	Payment            *Payment            `json:"payment,omitempty"`
	VirtualAccountInfo *VirtualAccountInfo `json:"virtual_account_info,omitempty"`
	Customer           *Customer           `json:"customer,omitempty"`
	AdditionalInfo     *AdditionalInfo     `json:"additional_info,omitempty"`
}

type ErrorResponse struct {
	Order       *Order      `json:"order,omitempty"`
	Code        int         `json:"code"`
	ErrorDetail ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

type CheckoutResponse struct {
	Message  []string `json:"message"`
	Response struct {
		Order          Order          `json:"order"`
		Payment        Payment        `json:"payment,omitempty"`
		AdditionalInfo AdditionalInfo `json:"additional_info,omitempty"`
		Uuid           json.Number    `json:"uuid"`
		Headers        Headers        `json:"headers"`
	} `json:"response"`
}

type CheckStatusResponse struct {
	Order              Order              `json:"order"`
	Transaction        Transaction        `json:"transaction"`
	Service            Service            `json:"service"`
	Acquirer           Acquirer           `json:"acquirer"`
	Channel            Channel            `json:"channel"`
	VirtualAccountInfo VirtualAccountInfo `json:"virtual_account_info"`
	AdditionalInfo     AdditionalInfo     `json:"additional_info"`
}

type CreateVirtualAccountResponse struct {
	Order              Order              `json:"order"`
	VirtualAccountInfo VirtualAccountInfo `json:"virtual_account_info"`
}
