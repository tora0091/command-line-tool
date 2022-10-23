package auth

import (
	"encoding/xml"
	"strings"
	"text/template"

	"amadeuscommand/config"
	"amadeuscommand/sender"
)

type Auth struct {
	Response string
}

func NewAuth() *Auth {
	return &Auth{}
}

type templateData struct {
	SourceOffice        string
	OriginatorTypeCode  string
	Originator          string
	ReferenceQualifier  string
	ReferenceIdentifier string
	OrganizationId      string
	DataLength          string
	DataType            string
	BinaryData          string
}

func (a *Auth) GetSession() (*Session, error) {
	req, err := makeTemplate()
	if err != nil {
		return nil, err
	}

	resp, err := sender.NewSendRequest(config.GetSoapAction().Auth, req).CallApi()
	if err != nil {
		return nil, err
	}

	var envelope Envelope
	err = xml.Unmarshal(resp, &envelope)
	if err != nil {
		return nil, err
	}
	return &envelope.Header.Session, nil
}

func makeTemplate() (string, error) {
	data := &templateData{
		SourceOffice:        config.GetAPIParams().SourceOffice,
		OriginatorTypeCode:  config.GetAPIParams().OriginatorTypeCode,
		Originator:          config.GetAPIParams().Originator,
		ReferenceQualifier:  config.GetAPIParams().ReferenceQualifier,
		ReferenceIdentifier: config.GetAPIParams().ReferenceIdentifier,
		OrganizationId:      config.GetAPIParams().OrganizationID,
		DataLength:          config.GetAPIParams().DataLength,
		DataType:            config.GetAPIParams().DataType,
		BinaryData:          config.GetAPIParams().BinaryData,
	}

	tmp, err := template.New("template.txt").Parse(getTemplate())
	if err != nil {
		return "", err
	}

	writer := new(strings.Builder)

	if err := tmp.Execute(writer, data); err != nil {
		return "", err
	}
	return writer.String(), nil
}

func getTemplate() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:ns1="http://xml.amadeus.com/VLSSLQ_06_1_1A">
<SOAP-ENV:Body>
<ns1:Security_Authenticate>
<ns1:userIdentifier>
<ns1:originIdentification>
<ns1:sourceOffice>{{.SourceOffice}}</ns1:sourceOffice>
</ns1:originIdentification>
<ns1:originatorTypeCode>{{.OriginatorTypeCode}}</ns1:originatorTypeCode>
<ns1:originator>{{.Originator}}</ns1:originator>
</ns1:userIdentifier>
<ns1:dutyCode>
<ns1:dutyCodeDetails>
<ns1:referenceQualifier>{{.ReferenceQualifier}}</ns1:referenceQualifier>
<ns1:referenceIdentifier>{{.ReferenceIdentifier}}</ns1:referenceIdentifier>
</ns1:dutyCodeDetails>
</ns1:dutyCode>
<ns1:systemDetails>
<ns1:organizationDetails>
<ns1:organizationId>{{.OrganizationId}}</ns1:organizationId>
</ns1:organizationDetails>
</ns1:systemDetails>
<ns1:passwordInfo>
<ns1:dataLength>{{.DataLength}}</ns1:dataLength>
<ns1:dataType>{{.DataType}}</ns1:dataType>
<ns1:binaryData>{{.BinaryData}}</ns1:binaryData>
</ns1:passwordInfo>
</ns1:Security_Authenticate>
</SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
}
