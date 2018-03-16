package sdk

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/slrem/AliyunSMS"
	"hbsms/models"

	"encoding/json"
)

const (
	ALIKEY     = "aliyun::accessKeyId"   // 配置文件中的阿里云Key Name
	ALISECRET  = "aliyun::accessSecret"  // 配置文件中的阿里云Secret Name
)

type Aliyun struct {
	accessKeyId  string
	accessSecret string
}

func (ali Aliyun) Init() {
	accessKeyId := beego.AppConfig.String(ALIKEY)
	accessSecret := beego.AppConfig.String(ALISECRET)
	ali.accessKeyId = accessKeyId
	ali.accessSecret = accessSecret
}

func (ali Aliyun) Send(sms models.SMS, st models.SMSTemplate) string {
	if sms.Valid() {
		logs.Error("sms object exist empty val")
	}
	params, err := json.Marshal(sms.Data)
	if err != nil {
		logs.Error("sms data to json error")
	}
	bizId, err:= AliyunSMS.SendSms(ali.accessKeyId, ali.accessSecret, st.SignName, sms.Mobile, st.TemplateCode, string(params))
	if err != nil {
		logs.Error("aliyun sms send message error")
	}

	return bizId
}