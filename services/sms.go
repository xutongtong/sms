package services

import (
	"hbsms/models"
	"hbsms/sdk"
)

func Send(sms models.SMS, st models.SMSTemplate) string {
	obj := sdk.SMSFactory{}

	if sms.CountryCode != models.COUNTRYCODE {
		// 使用云片发送国际短信号
		yp := obj.CreateSMSSDK(sdk.YP)
		return yp.Send(sms, st)
	}

	// 使用阿里云发送国内短信
	ali := obj.CreateSMSSDK(sdk.ALI)
	return ali.Send(sms, st)
}