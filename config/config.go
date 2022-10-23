package config

import (
	"encoding/json"
	"os"
)

var conf Config

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		panic(err)
	}
}

func GetAmadeusEndPoint() string {
	return conf.AmadeusEndPoint
}

func GetAPIHeaderHost() string {
	return conf.APIHeader.Host
}

func GetAPIParams() *APIParams {
	return &conf.APIParams
}

func GetSoapAction() *SoapAction {
	return &conf.APIHeader.SoapAction
}

type Config struct {
	AmadeusEndPoint string    `json:"amadeusEndPoint"`
	APIHeader       APIHeader `json:"apiHeader"`
	APIParams       APIParams `json:"apiParams"`
}

type APIParams struct {
	SourceOffice        string `json:"sourceOffice"`
	OriginatorTypeCode  string `json:"originatorTypeCode"`
	Originator          string `json:"originator"`
	ReferenceQualifier  string `json:"referenceQualifier"`
	ReferenceIdentifier string `json:"referenceIdentifier"`
	OrganizationID      string `json:"organizationId"`
	DataLength          string `json:"dataLength"`
	DataType            string `json:"dataType"`
	BinaryData          string `json:"binaryData"`
}

type APIHeader struct {
	Host       string     `json:"host"`
	SoapAction SoapAction `json:"soapAction"`
}

type SoapAction struct {
	Auth    string `json:"auth"`
	Cryptic string `json:"cryptic"`
}
