package cryptic

import (
	"encoding/xml"
	"strconv"
	"strings"
	"text/template"

	"amadeuscommand/auth"
	"amadeuscommand/config"
	"amadeuscommand/sender"
)

type Cryptic struct {
	Session *auth.Session
}

func NewCryptic(session *auth.Session) *Cryptic {
	return &Cryptic{
		Session: session,
	}
}

type templateData struct {
	SessionId         string
	SequenceNumber    string
	SecurityToken     string
	MessageFunction   string
	TextStringDetails string
}

func (c *Cryptic) GetCommandResult(command string) (*string, error) {
	req, err := c.makeTemplate(strings.ToUpper(command))
	if err != nil {
		return nil, err
	}

	resp, err := sender.NewSendRequest(config.GetSoapAction().Cryptic, req).CallApi()
	if err != nil {
		return nil, err
	}

	var envelope Envelope
	if err := xml.Unmarshal(resp, &envelope); err != nil {
		return nil, err
	}

	return &envelope.Body.CommandCrypticReply.LongTextString.TextStringDetails, nil
}

func (c *Cryptic) makeTemplate(command string) (string, error) {
	data := &templateData{
		SessionId:         c.Session.SessionId,
		SequenceNumber:    setSequenceNumber(c.Session.SequenceNumber),
		SecurityToken:     c.Session.SecurityToken,
		MessageFunction:   "M",
		TextStringDetails: command,
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
xmlns:ns1="http://xml.amadeus.com/HSFREQ_07_3_1A"
xmlns:ns2="http://xml.amadeus.com/ws/2009/01/WBS_Session-2.0.xsd">
<SOAP-ENV:Header>
<ns2:Session>
<ns2:SessionId>{{.SessionId}}</ns2:SessionId>
<ns2:SequenceNumber>{{.SequenceNumber}}</ns2:SequenceNumber>
<ns2:SecurityToken>{{.SecurityToken}}</ns2:SecurityToken>
</ns2:Session>
</SOAP-ENV:Header>
<SOAP-ENV:Body>
<ns1:Command_Cryptic>
<ns1:messageAction>
<ns1:messageFunctionDetails>
<ns1:messageFunction>{{.MessageFunction}}</ns1:messageFunction>
</ns1:messageFunctionDetails>
</ns1:messageAction>
<ns1:longTextString>
<ns1:textStringDetails>{{.TextStringDetails}}</ns1:textStringDetails>
</ns1:longTextString>
</ns1:Command_Cryptic>
</SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
}

func setSequenceNumber(number string) string {
	i, _ := strconv.Atoi(number)
	i = i + 1
	return strconv.Itoa(i)
}
