package sdk

import (
	"github.com/astaxie/beego"
	YunpianSMS "github.com/yunpian/yunpian-go-sdk/sdk"
	"hbsms/models"
)

const YPKEY = "yunpian::apikey" // 配置文件中云片网的Key Name

type Yunpian struct {
	apiKey string
}

func (yp Yunpian) Init() {
	apiKey := beego.AppConfig.String(YPKEY)
	yp.apiKey = apiKey
}

func (yp Yunpian) Send(sms models.SMS, st models.SMSTemplate) string {
	clnt := YunpianSMS.New(yp.apiKey)
	param := YunpianSMS.NewParam(2)

	param[YunpianSMS.MOBILE] = sms.GetMobile()
	st.SetContent()
	param[YunpianSMS.TEXT] = st.GetContent()
	r := clnt.Sms().SingleSend(param)

	return r.String()
}
