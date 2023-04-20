package dhl

import (
	"encoding/json"
	cdek "github.com/tazhibayda/testProject_Golang/cdek"
)

var Client cdek.Client

func NewDHL(account, password, apiurl string) *cdek.API {
	return cdek.NewCDEK(account, password, apiurl)
}

type CustomerDetails struct {
	ShipperDetails struct {
		PostalCode  string `json:"postalCode"`
		CityName    string `json:"cityName"`
		CountryCode string `json:"countryCode"`
	} `json:"shipperDetails"`
	ReceiverDetails struct {
		PostalCode  string `json:"postalCode"`
		CityName    string `json:"cityName"`
		CountryCode string `json:"countryCode"`
	} `json:"receiverDetails"`
}

type Account struct {
	TypeCode string `json:"typeCode"`
	Number   string `json:"number"`
}

type Charge struct {
	TypeCode     string `json:"typeCode"`
	Amount       int    `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
}
type Dimensions struct {
	Length int `json:"length"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Package struct {
	TypeCode   string     `json:"typeCode"`
	Weight     int        `json:"weight"`
	Dimensions Dimensions `json:"dimensions"`
}

type Item struct {
	Number                   int    `json:"number"`
	Name                     string `json:"name"`
	Description              string `json:"description"`
	ManufacturerCountry      string `json:"manufacturerCountry"`
	PartNumber               string `json:"partNumber"`
	Quantity                 int    `json:"quantity"`
	QuantityType             string `json:"quantityType"`
	UnitPrice                int    `json:"unitPrice"`
	UnitPriceCurrencyCode    string `json:"unitPriceCurrencyCode"`
	CustomsValue             int    `json:"customsValue"`
	CustomsValueCurrencyCode string `json:"customsValueCurrencyCode"`
	CommodityCode            string `json:"commodityCode"`
	Weight                   int    `json:"weight"`
	WeightUnitOfMeasurement  string `json:"weightUnitOfMeasurement"`
	Category                 string `json:"category"`
	Brand                    string `json:"brand"`
	GoodsCharacteristics     []struct {
		TypeCode string `json:"typeCode"`
		Value    string `json:"value"`
	} `json:"goodsCharacteristics"`
	AdditionalQuantityDefinitions []struct {
		TypeCode string `json:"typeCode"`
		Amount   int    `json:"amount"`
	} `json:"additionalQuantityDefinitions"`
	EstimatedTariffRateType string `json:"estimatedTariffRateType"`
}

type Payload struct {
	CustomerDetails             CustomerDetails `json:"customerDetails"`
	Accounts                    []Account       `json:"accounts"`
	ProductCode                 string          `json:"productCode"`
	LocalProductCode            string          `json:"localProductCode"`
	UnitOfMeasurement           string          `json:"unitOfMeasurement"`
	CurrencyCode                string          `json:"currencyCode"`
	IsCustomsDeclarable         bool            `json:"isCustomsDeclarable"`
	IsDTPRequested              bool            `json:"isDTPRequested"`
	IsInsuranceRequested        bool            `json:"isInsuranceRequested"`
	GetCostBreakdown            bool            `json:"getCostBreakdown"`
	Charges                     []Charge        `json:"charges"`
	ShipmentPurpose             string          `json:"shipmentPurpose"`
	TransportationMode          string          `json:"transportationMode"`
	MerchantSelectedCarrierName string          `json:"merchantSelectedCarrierName"`
	Packages                    []Package       `json:"packages"`
	Items                       Item            `json:"items"`
	GetTariffFormula            bool            `json:"getTariffFormula"`
	GetQuotationID              bool            `json:"getQuotationID"`
}

func Calculate(addrFrom string, addrTo string, size cdek.Size) ([]cdek.PriceSending, error) {
	endpoint := "landed-cost"
	var payload Payload
	token := ""
	dataJSON, err := json.Marshal(payload)
	client := cdek.Client{Token: token, Endpoint: endpoint, Method: "POST"}
	_, err = cdek.SendRequest("", &client, dataJSON)

	if err != nil {
		panic(err)
	}

	return nil, nil
}
