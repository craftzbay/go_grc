package otp

import (
	"github.com/craftzbay/go_grc/v2/client"
)

type emailBody struct {
	SenderId     uint                   `json:"sender_id"`
	To           []string               `json:"to"`
	Subject      string                 `json:"subject"`
	TemplateId   uint                   `json:"template_id"`
	TemplateData map[string]interface{} `json:"template_data"`
}

func SendEmail(uri, to, subject string, body map[string]interface{}) error {
	email := emailBody{}
	email.SenderId = 17
	email.TemplateId = 9
	email.To = append(email.To, to)
	email.Subject = subject
	email.TemplateData = body
	// data, err := converter.InterfaceToMap(email)
	// if err != nil {
	// 	return err
	// }
	go client.MakeHTTPRequest[map[string]interface{}](uri, "POST", nil, nil, email)
	return nil
}
