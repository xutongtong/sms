package models

import "fmt"

const COUNTRYCODE  = "86" // 默认的国家码

type SMS struct {
	CountryCode string    // 国家码
	Mobile string         // 手机号
	Source string         // 来源
	Template string       // 短信模版名称
	Data map[string]interface{} // 短信参数以及数据
}

func (sms *SMS) Valid() bool {
	return sms.CountryCode == "" || sms.Mobile == "" || sms.Template == ""
}

func (sms *SMS) GetMobile() string {
	mobile := sms.Mobile
	//国际手机号添加+
	if sms.CountryCode != COUNTRYCODE {
		mobile = fmt.Sprintf("+%d%d", sms.CountryCode, sms.Mobile)
	}
	return mobile
}