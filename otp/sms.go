package otp

import (
	"github.com/craftzbay/go_grc/v2/client"
)

type smsBody struct {
	PhoneNumbers []string `json:"phone_numbers"`
	MessageValue string   `json:"message_value"`
	SmsType      uint     `json:"sms_type"`
}

func SendSms(url, phone, text string) error {
	sms := smsBody{}
	sms.PhoneNumbers = append(sms.PhoneNumbers, phone)
	sms.MessageValue = text
	sms.SmsType = 10
	// data, err := converter.InterfaceToMap(sms)
	// if err != nil {
	// 	return err
	// }
	headers := make(map[string]string)
	headers["message_code"] = "302402"

	config := &client.RequestConfig{
		Url:     url,
		Method:  "POST",
		Headers: &headers,
		Body:    sms,
	}

	go client.MakeHTTPRequest[map[string]interface{}](config)
	return nil
}
