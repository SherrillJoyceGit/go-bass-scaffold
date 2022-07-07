package util

import (
	"errors"
	"regexp"
)

// 识别手机号码
func IsMobile(mobile string) error {
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if result {
		return nil
	} else {
		return errors.New("手机号不正确！")
	}
}

//手机号加密
func EncryptionPhone(phone string) string {
	phoneNo := []byte(phone)
	n := len(phoneNo)
	for i := 0; i < n-4; i++ {
		if i < 3 {
			continue
		}
		phoneNo[i] = '*'
	}
	return string(phoneNo[:])
}
