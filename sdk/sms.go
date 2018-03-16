package sdk

import "hbsms/models"

const (
	ALI = "ali"
	YP  = "yp"
)

type SMS interface {
	Send(sms models.SMS, st models.SMSTemplate) string
}

type SMSFactory struct {

}

func (sf SMSFactory) CreateSMSSDK(platform string) SMS {
	switch platform {
	case YP:
		yp := Yunpian{}
		return yp
	case ALI:
		ali := Aliyun{}
		return ali
	default:
		ali := Aliyun{}
		return ali
	}
	return nil
}