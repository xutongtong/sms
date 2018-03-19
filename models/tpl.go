package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/yaml"
	"github.com/astaxie/beego/logs"
	"hbsms/utils"
)

const (
	PATH         = "conf/sms.yaml"         // 短信配置的文件
	TEMPLATEID   = "aliTemplateCode"       // 配置文件中的阿里云的模版ID
	TEMPLATE     = "template"              // 配置文件中的短信内容模版
	PARAMS       = "params"                // 配置文件中的短信内容对应的参数
	SIGN         = "signname::%s"          // 配置文件中的短信签名的名称
	KY           = "kaiyue"                // 短信签名-凯阅
)

// 短信模板
type SMSTemplate struct {
	TemplateCode string   // 短信对应阿里云的模版ID
	SignName string       // 签名
	ContentFormat string  // 短信内容格式
	Content string        // 加上签名后的短信内容
	Params []string  // 短信内容中需要适配的参数
}

var yamlConf config.Configer

func init() {
	cnf, _ := config.NewConfig("yaml", PATH)
	yamlConf = cnf
}

func (st *SMSTemplate) Init(tpl string) {
	if tpl == "" {
		logs.Error("sms template is null")
	}

	smsConf, err := yamlConf.DIY(tpl)
	if err != nil {
		logs.Error(err)
	}

	template := smsConf.(map[string]interface{})

	st.TemplateCode = template[TEMPLATEID].(string)
	st.Params = template[PARAMS].([]string)
	st.ContentFormat = template[TEMPLATE].(string)
}

func (st *SMSTemplate) Valid(sms SMS) bool {
	// Data与Params需要完全匹配
	for key, _ := range sms.Data {
		if !utils.InArray(key, st.Params) {
			return true
		}
	}
	return false
}

func (st *SMSTemplate) SetContent() {
	content := fmt.Sprintf(st.ContentFormat, st.Params)
	signName := st.GetSignName()
	content = signName + content

	st.Content = content
}

func (st *SMSTemplate) GetContent() string {
	return st.Content
}

func (st *SMSTemplate) SetSignName(source string) {
	signName := beego.AppConfig.String(fmt.Sprintf(SIGN, source))

	// 防止数据异常的情况，签名默认为凯阅
	if signName == "" {
		signName = beego.AppConfig.String(fmt.Sprintf(SIGN, KY))
	}

	st.SignName = signName
}

func (st *SMSTemplate) GetSignName() string {
	return st.SignName
}