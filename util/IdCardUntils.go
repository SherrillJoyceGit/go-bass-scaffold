package util

import (
	"errors"
	"strconv"
	"time"
)

var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var valid_value = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
var valid_province = map[string]string{
	"11": "北京市",
	"12": "天津市",
	"13": "河北省",
	"14": "山西省",
	"15": "内蒙古自治区",
	"21": "辽宁省",
	"22": "吉林省",
	"23": "黑龙江省",
	"31": "上海市",
	"32": "江苏省",
	"33": "浙江省",
	"34": "安徽省",
	"35": "福建省",
	"36": "山西省",
	"37": "山东省",
	"41": "河南省",
	"42": "湖北省",
	"43": "湖南省",
	"44": "广东省",
	"45": "广西壮族自治区",
	"46": "海南省",
	"50": "重庆市",
	"51": "四川省",
	"52": "贵州省",
	"53": "云南省",
	"54": "西藏自治区",
	"61": "陕西省",
	"62": "甘肃省",
	"63": "青海省",
	"64": "宁夏回族自治区",
	"65": "新疆维吾尔自治区",
	"71": "台湾省",
	"81": "香港特别行政区",
	"91": "澳门特别行政区",
}

// 效验18位身份证
func IsValidCitizenNo18(citizenNo18 *[]byte) bool {
	nLen := len(*citizenNo18)
	if nLen != 18 {
		return false
	}

	nSum := 0
	for i := 0; i < nLen-1; i++ {
		n, _ := strconv.Atoi(string((*citizenNo18)[i]))
		nSum += n * weight[i]
	}
	mod := nSum % 11
	if valid_value[mod] == (*citizenNo18)[17] {
		return true
	}
	return false
}

//是否是闰年
func IsLeapYear(nYear int) bool {
	if nYear <= 0 {
		return false
	}
	if (nYear%4 == 0 && nYear%100 != 0) || nYear%400 == 0 {
		return true
	}
	return false
}

// 生日日期格式效验
func CheckBirthdayValid(nYear, nMonth, nDay int) bool {
	if nYear < 1900 || nMonth <= 0 || nMonth > 12 || nDay <= 0 || nDay > 31 {
		return false
	}

	curYear, curMonth, curDay := time.Now().Date()
	if nYear == curYear {
		if nMonth > int(curMonth) {
			return false
		} else if nMonth == int(curMonth) && nDay > curDay {
			return false
		}
	}

	if 2 == nMonth {
		if IsLeapYear(nYear) && nDay > 29 {
			return false
		} else if nDay > 28 {
			return false
		}
	} else if 4 == nMonth || 6 == nMonth || 9 == nMonth || 11 == nMonth {
		if nDay > 30 {
			return false
		}
	}

	return true
}

// 省份号码效验
func CheckProvinceValid(citizenNo []byte) bool {
	provinceCode := make([]byte, 0)
	provinceCode = append(provinceCode, citizenNo[:2]...)
	provinceStr := string(provinceCode)

	for i := range valid_province {
		if provinceStr == i {
			return true
		}
	}

	return false
}

// 效验有效地身份证号码
func IsValidCitizenNo(citizenNo *[]byte) bool {
	nLen := len(*citizenNo)
	if nLen == 15 {
		*citizenNo = Change(*citizenNo)
	}
	if !IsValidCitizenNo18(citizenNo) {
		return false
	}

	for i, v := range *citizenNo {
		n, _ := strconv.Atoi(string(v))
		if n >= 0 && n <= 9 {
			continue
		}
		if v == 'X' && i == 16 {
			continue
		}
		return false
	}
	if !CheckProvinceValid(*citizenNo) {
		return false
	}
	nYear, _ := strconv.Atoi(string((*citizenNo)[6:10]))
	nMonth, _ := strconv.Atoi(string((*citizenNo)[10:12]))
	nDay, _ := strconv.Atoi(string((*citizenNo)[12:14]))
	if !CheckBirthdayValid(nYear, nMonth, nDay) {
		return false
	}
	return true
}

//转换15位身份证为18位身份证
func Change(citizenNo []byte) []byte {
	//15位转17位 在年份前面添加19
	var citizenNo18 []byte
	for i := 0; i < 6; i++ {
		citizenNo18 = append(citizenNo18, citizenNo[i])
	}
	id2 := citizenNo[6:15]
	citizenNo18 = append(citizenNo18, '1', '9')
	for _, id := range id2 {
		citizenNo18 = append(citizenNo18, id)
	}

	//对应每一位都与自己的的系数相乘的和取模
	var coef = []uint8{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var result uint32 = 0
	for i, id := range citizenNo18 {
		var idnum = id - 48
		result = result + uint32(idnum)*uint32(coef[i])
	}
	result = result % 11

	//根据取模后的结果获得最后一位的校验值
	var Mant = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	citizenNo18 = append(citizenNo18, Mant[result])
	return citizenNo18
}

//校验身份证号
func GetIsValidCitizenNo(idNo string) (err error) {
	citizenNo := []byte(idNo)
	err = nil
	if !IsValidCitizenNo(&citizenNo) {
		return errors.New("不合法的身份证号码。")
	}
	return nil
}

//获取生日
func GetBirthdayCitizenNo(idNo string) time.Time {
	citizenNo := []byte(idNo)
	if len(citizenNo) == 15 {
		citizenNo = Change(citizenNo)
	}
	birthday, _ := time.Parse("2006-01-02", string(citizenNo[6:10])+"-"+string(citizenNo[10:12])+"-"+string(citizenNo[12:14]))
	return birthday
}

//获取性别
func GetSexCitizenNo(idNo string) string {
	citizenNo := []byte(idNo)
	if len(citizenNo) == 15 {
		citizenNo = Change(citizenNo)
	}
	genderMask, _ := strconv.Atoi(string(citizenNo[16]))
	var sex string
	if genderMask%2 == 0 {
		sex = "女"
	} else {
		sex = "男"
	}
	return sex
}

//获取省份地址
func GetAddressCitizenNo(idNo string) string {
	citizenNo := []byte(idNo)
	address := valid_province[string(citizenNo[:2])]
	return address
}

//获取年龄
func GetAgeCitizenNo(idNo string) int64 {
	citizenNo := []byte(idNo)
	if len(citizenNo) == 15 {
		citizenNo = Change(citizenNo)
	}
	birthday := GetBirthdayCitizenNo(idNo)
	age := time.Now().Year() - birthday.Year()
	return int64(age)
}

//身份证加密
func EncryptionCitizenNo(idNo string) string {
	citizenNo := []byte(idNo)
	n := len(citizenNo)
	for i := 0; i < n-2; i++ {
		if i < 6 {
			continue
		}
		citizenNo[i] = '*'
	}
	return string(citizenNo[:])
}
