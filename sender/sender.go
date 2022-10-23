package sender

import (
	"io"
	"net/http"
	"strings"

	"amadeuscommand/config"
)

type SendRequest struct {
	SoapAction string
	RequestXml string
}

func NewSendRequest(action, xml string) *SendRequest {
	return &SendRequest{
		SoapAction: action,
		RequestXml: xml,
	}
}

func (r *SendRequest) CallApi() ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		config.GetAmadeusEndPoint(),
		strings.NewReader(r.RequestXml),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", r.SoapAction)
	req.Header.Set("Host", config.GetAPIHeaderHost())

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
