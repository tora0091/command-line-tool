package cryptic

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
	Text                string              `xml:",chardata"`
	CommandCrypticReply CommandCrypticReply `xml:"Command_CrypticReply"`
}

type CommandCrypticReply struct {
	Text           string         `xml:",chardata"`
	Xmlns          string         `xml:"xmlns,attr"`
	LongTextString LongTextString `xml:"longTextString"`
}

type LongTextString struct {
	Text              string `xml:",chardata"`
	TextStringDetails string `xml:"textStringDetails"`
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"SOAP-ENV,attr"`
	Awss    string   `xml:"awss,attr"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
}
