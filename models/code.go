package models

import (
	"time"
	"fmt"
	"math/rand"
)

type SmsCode struct {
	CountryCode string  `valid:"Required" json:"country_code"`
	Mobile string `valid:"Match(/\d+{5, 20}/)" json:"mobile"`
	Type string `valid:"Required" json:"type"`
	Code string `"json:code"`
	Source string `json:"source"`
	SendDate string `json:"send_date"`
	BizId string `json:"biz_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 随机生成六位随机数字
func GenerateSmsCode() string{
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	SmsCode := fmt.Sprintf("%06v", random.Int31n(1000000))

	return SmsCode
}

//
//func SendCode(smsCode SmsCode) {
//	db.GetDB().Create(&smsCode)
//}