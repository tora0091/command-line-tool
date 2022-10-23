package auth

import "encoding/xml"

type Header struct {
	Text    string  `xml:",chardata"`
	Session Session `xml:"Session"`
}

type Session struct {
	Text           string `xml:",chardata"`
	SessionId      string `xml:"SessionId"`
	SequenceNumber string `xml:"SequenceNumber"`
	SecurityToken  string `xml:"SecurityToken"`
}

type Body struct {
	Text                      string                    `xml:",chardata"`
	SecurityAuthenticateReply SecurityAuthenticateReply `xml:"Security_AuthenticateReply"`
}

type SecurityAuthenticateReply struct {
	Text          string        `xml:",chardata"`
	Xmlns         string        `xml:"xmlns,attr"`
	ProcessStatus ProcessStatus `xml:"processStatus"`
}

type ProcessStatus struct {
	Text       string `xml:",chardata"`
	StatusCode string `xml:"statusCode"`
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Awss    string   `xml:"awss,attr"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
}
