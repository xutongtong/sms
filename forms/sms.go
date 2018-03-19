package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"strings"
	"hbsms/utils"
)

const (
	KEY = "sms"
	TYPE = "types"
	SOURCES = "sources"
	CHAR = ","
	COUNTRYCODE = "86"
	SOURCE = "kaiyue"
)

var smsConf map[string]string

type SMS struct {
	CountryCode string `json:"country_code" valid:"Required"`
	Mobile string `json:"mobile" valid:"Match(/^[1-9]\\d{5,10}$/)"`
	Type string `json:"type" valid:"Required"`
	Source string `json:"source" valid:"Required"`
}

func init() {
	conf, _ := beego.AppConfig.DIY(KEY)
	smsConf = conf.(map[string]string)
}

func (sms *SMS) Valid(v *validation.Validation) {
	if hasError(sms) {
		v.SetError("SMS", "sms obj error")
	}
}

func hasError(sms *SMS) bool {
	fields := make(map[string]string)
	fields[TYPE] = sms.Type
	fields[SOURCES] = sms.Source

	for k, v := range fields {
		checked := strings.Split(smsConf[k], CHAR)
		if len(checked) < 0 {
			return true
		}
		if !utils.InArray(v, checked) {
			return true
		}
	}

	return false
}