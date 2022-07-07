// Package util  @Description
// @Author  gaoyaya  2021/2/3 3:29 PM
package util

import (
	"errors"
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

const (
	DATE_PATTERN      = "2006-01-02"
	DATE_TIME_PATTERN = "2006-01-02 15:04:05"
)

type BeanUpdateLog struct {
	Changed    bool                `json:"changed"`
	UpdateLogs []BeanUpdateLogItem `json:"updateLogs"`
}

type BeanUpdateLogItem struct {
	PropertyName     string      `json:"propertyName"`
	OldProperyValue  interface{} `json:"oldProperyValue"`
	NewPropertyValue interface{} `json:"newPropertyValue"`
}

func Copy(source interface{}, dest interface{}, ignoreFields ...string) ([]string, error) {
	destType := reflect.TypeOf(dest)
	destValue := reflect.ValueOf(dest)
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)

	if destType.Kind() != reflect.Ptr {
		return make([]string, 0), fmt.Errorf("a must be a struct pointer")
	}

	destValue = reflect.ValueOf(destValue.Interface())

	m := SliceToMap(ignoreFields)
	fileds := make([]string, 0)

	for i := 0; i < sourceValue.NumField(); i++ {
		name := sourceType.Field(i).Name

		if _, ok := m[name]; !ok {
			fileds = append(fileds, name)
		}
	}

	if len(fileds) == 0 {
		return make([]string, 0), nil
	}

	modified := make([]string, 0, len(fileds))

	for i := 0; i < len(fileds); i++ {
		name := fileds[i]
		sourceFieldValue := sourceValue.FieldByName(name)
		destFieldValue := destValue.Elem().FieldByName(name)

		if destFieldValue.IsValid() {
			sourceFieldValueTypeStr := sourceFieldValue.Type().String()
			destFieldValueTypeStr := destFieldValue.Type().String()

			if (sourceFieldValueTypeStr == "time.Time" ||
				sourceFieldValueTypeStr == "*time.Time") &&
				(destFieldValueTypeStr == "string" ||
					destFieldValueTypeStr == "*string") {
				timeStr := ""

				if sourceFieldValueTypeStr == "time.Time" {
					t := (sourceFieldValue.Interface()).(time.Time)

					if !t.IsZero() {
						timeStr = t.Format(DATE_TIME_PATTERN)
					}

				} else if sourceFieldValueTypeStr == "*time.Time" {
					timePtr := (sourceFieldValue.Interface()).(*time.Time)

					if timePtr != nil && !timePtr.IsZero() {
						timeStr = (*timePtr).Format(DATE_TIME_PATTERN)
					}
				}

				if destFieldValueTypeStr == "string" &&
					reflect.DeepEqual(destFieldValue.Interface(), timeStr) == false {

					destFieldValue.Set(reflect.ValueOf(timeStr))

					modified = append(modified, name)
				} else if destFieldValueTypeStr == "*string" {
					strPtr := (destFieldValue.Interface()).(*string)

					if strPtr != nil && reflect.DeepEqual(*strPtr, timeStr) == false {
						destFieldValue.Set(reflect.ValueOf(&timeStr))

						modified = append(modified, name)
					} else {
						destFieldValue.Set(reflect.ValueOf(&timeStr))

						modified = append(modified, name)
					}
				}
			} else if (sourceFieldValueTypeStr == "string" || sourceFieldValueTypeStr == "*string") &&
				(destFieldValueTypeStr == "time.Time" || destFieldValueTypeStr == "*time.Time") {
				timeStr := ""

				if sourceFieldValueTypeStr == "string" {
					timeStr = (sourceFieldValue.Interface()).(string)
				} else if sourceFieldValueTypeStr == "*string" {
					timePtr := (sourceFieldValue.Interface()).(*string)

					if timePtr != nil {
						timeStr = *timePtr
					}
				}

				if t, err := time.ParseInLocation(DATE_TIME_PATTERN, timeStr, time.Local); err == nil {
					if destFieldValueTypeStr == "time.Time" && reflect.DeepEqual(t, destFieldValue.Interface()) == false {
						destFieldValue.Set(reflect.ValueOf(t))

						modified = append(modified, name)
					} else if destFieldValueTypeStr == "*time.Time" {
						timePtr := (destFieldValue.Interface()).(*time.Time)

						if timePtr != nil && reflect.DeepEqual((*timePtr).Format(DATE_TIME_PATTERN), timeStr) == false {
							destFieldValue.Set(reflect.ValueOf(&t))

							modified = append(modified, name)
						} else if timePtr == nil {
							destFieldValue.Set(reflect.ValueOf(&t))

							modified = append(modified, name)
						}
					}

				}
			} else if sourceFieldValueTypeStr[0] == '*' &&
				sourceFieldValueTypeStr[1:] == destFieldValueTypeStr &&
				!sourceFieldValue.IsNil() && !reflect.DeepEqual(sourceFieldValue.Elem().Interface(), destFieldValue.Interface()) {
				destFieldValue.Set(reflect.ValueOf(sourceFieldValue.Elem().Interface()))

				modified = append(modified, name)
			} else if destFieldValueTypeStr[0] == '*' &&
				destFieldValueTypeStr[1:] == sourceFieldValueTypeStr &&
				((!destFieldValue.IsZero() && !reflect.DeepEqual(sourceFieldValue.Interface(), destFieldValue.Elem().Interface())) || (destFieldValue.IsNil())) {
				switch sourceFieldValue.Kind() {
				case reflect.Bool:
					v := sourceFieldValue.Interface().(bool)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int:
					v := sourceFieldValue.Interface().(int)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int8:
					v := sourceFieldValue.Interface().(int8)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int16:
					v := sourceFieldValue.Interface().(int16)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int32:
					v := sourceFieldValue.Interface().(int32)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int64:
					v := sourceFieldValue.Interface().(int64)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint:
					v := sourceFieldValue.Interface().(uint)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint8:
					v := sourceFieldValue.Interface().(uint8)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint16:
					v := sourceFieldValue.Interface().(uint16)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint32:
					v := sourceFieldValue.Interface().(uint32)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint64:
					v := sourceFieldValue.Interface().(uint64)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Float32:
					v := sourceFieldValue.Interface().(float32)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Float64:
					v := sourceFieldValue.Interface().(float64)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.String:
					v := sourceFieldValue.Interface().(string)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				}

			} else if destFieldValue.Kind() == sourceFieldValue.Kind() &&
				reflect.DeepEqual(destFieldValue.Interface(), sourceFieldValue.Interface()) == false {
				destFieldValue.Set(sourceFieldValue)

				modified = append(modified, name)
			}
		}

	}

	return modified, nil
}

func CopyAndLogProperties(source interface{}, dest interface{}, ignoreFields ...string) (BeanUpdateLog, error) {
	beanUpdateLogs := &BeanUpdateLog{UpdateLogs: make([]BeanUpdateLogItem, 0)}
	destType := reflect.TypeOf(dest)
	destValue := reflect.ValueOf(dest)
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)

	if destType.Kind() != reflect.Ptr {
		return *beanUpdateLogs, fmt.Errorf("a must be a struct pointer")
	}

	destValue = reflect.ValueOf(destValue.Interface())

	m := SliceToMap(ignoreFields)
	fileds := make([]string, 0)

	for i := 0; i < sourceValue.NumField(); i++ {
		name := sourceType.Field(i).Name

		if _, ok := m[name]; !ok {
			fileds = append(fileds, name)
		}
	}

	if len(fileds) == 0 {
		return *beanUpdateLogs, nil
	}

	for i := 0; i < len(fileds); i++ {
		name := fileds[i]
		sourceFieldValue := sourceValue.FieldByName(name)

		destFieldValue := destValue.Elem().FieldByName(name)

		if destFieldValue.IsValid() {
			lname := StringFirstLower(name)
			destFieldValueTypeStr := destFieldValue.Type().String()
			sourceFieldValueTypeStr := sourceFieldValue.Type().String()

			if (sourceFieldValueTypeStr == "time.Time" ||
				sourceFieldValueTypeStr == "*time.Time") &&
				(destFieldValueTypeStr == "string" ||
					destFieldValueTypeStr == "*string") {
				timeStr := ""

				if sourceFieldValueTypeStr == "time.Time" {
					t := (sourceFieldValue.Interface()).(time.Time)

					if !t.IsZero() {
						timeStr = t.Format(DATE_TIME_PATTERN)
					}

				} else if sourceFieldValueTypeStr == "*time.Time" {
					timePtr := (sourceFieldValue.Interface()).(*time.Time)

					if timePtr != nil && !timePtr.IsZero() {
						timeStr = (*timePtr).Format(DATE_TIME_PATTERN)
					}
				}

				if destFieldValueTypeStr == "string" &&
					reflect.DeepEqual(destFieldValue.Interface(), timeStr) == false {
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: timeStr}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					destFieldValue.Set(reflect.ValueOf(timeStr))
				} else if destFieldValueTypeStr == "*string" {
					strPtr := (destFieldValue.Interface()).(*string)

					if strPtr != nil && reflect.DeepEqual(*strPtr, timeStr) == false {
						beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: &timeStr}
						beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

						destFieldValue.Set(reflect.ValueOf(&timeStr))
					} else {
						beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: &timeStr}
						beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

						destFieldValue.Set(reflect.ValueOf(&timeStr))
					}
				}
			} else if (sourceFieldValueTypeStr == "string" || sourceFieldValueTypeStr == "*string") &&
				(destFieldValueTypeStr == "time.Time" || destFieldValueTypeStr == "*time.Time") {
				timeStr := ""

				if sourceFieldValueTypeStr == "string" {
					timeStr = (sourceFieldValue.Interface()).(string)
				} else if sourceFieldValueTypeStr == "*string" {
					timePtr := (sourceFieldValue.Interface()).(*string)

					if timePtr != nil {
						timeStr = *timePtr
					}
				}

				if t, err := time.ParseInLocation(DATE_TIME_PATTERN, timeStr, time.Local); err == nil {
					if destFieldValueTypeStr == "time.Time" && reflect.DeepEqual(t, destFieldValue.Interface()) == false {
						beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: t}
						beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

						destFieldValue.Set(reflect.ValueOf(t))
					} else if destFieldValueTypeStr == "*time.Time" {
						timePtr := (destFieldValue.Interface()).(*time.Time)

						if timePtr != nil && reflect.DeepEqual((*timePtr).Format(DATE_TIME_PATTERN), timeStr) == false {
							beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: &t}
							beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

							destFieldValue.Set(reflect.ValueOf(&t))
						} else if timePtr == nil {
							beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: &timeStr}
							beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

							destFieldValue.Set(reflect.ValueOf(&t))
						}
					}

				}
			} else if sourceFieldValueTypeStr[0] == '*' &&
				sourceFieldValueTypeStr[1:] == destFieldValueTypeStr &&
				!sourceFieldValue.IsNil() && !reflect.DeepEqual(sourceFieldValue.Elem().Interface(), destFieldValue.Interface()) {

				beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
				beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

				destFieldValue.Set(reflect.ValueOf(sourceFieldValue.Elem().Interface()))
			} else if destFieldValueTypeStr[0] == '*' &&
				destFieldValueTypeStr[1:] == sourceFieldValueTypeStr &&
				((!destFieldValue.IsZero() && !reflect.DeepEqual(sourceFieldValue.Interface(), destFieldValue.Elem().Interface())) || (destFieldValue.IsNil())) {
				switch sourceFieldValue.Kind() {
				case reflect.Bool:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(bool)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int8:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int8)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int16:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int16)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int32:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int32)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int64:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int64)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Uint:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(uint)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Uint8:
					v := sourceFieldValue.Interface().(uint8)
					destFieldValue.Set(reflect.ValueOf(&v))

					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)
				case reflect.Uint16:
					v := sourceFieldValue.Interface().(uint16)
					destFieldValue.Set(reflect.ValueOf(&v))

					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)
				case reflect.Uint32:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(uint32)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Uint64:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(uint64)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Float32:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(float32)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Float64:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(float64)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.String:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(string)
					destFieldValue.Set(reflect.ValueOf(&v))
				}

			} else if destFieldValue.Kind() == sourceFieldValue.Kind() &&
				reflect.DeepEqual(destFieldValue.Interface(), sourceFieldValue.Interface()) == false {
				beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
				beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

				destFieldValue.Set(sourceFieldValue)
			}
		}

	}

	beanUpdateLogs.Changed = len(beanUpdateLogs.UpdateLogs) > 0

	return *beanUpdateLogs, nil
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	return m
}

func SliceToMap(s []string) map[string]string {
	m := make(map[string]string, 0)

	if s != nil {
		for _, v := range s {
			m[v] = v
		}
	}

	return m
}

func StringFirstLower(str string) string {
	if isLower := 'a' <= str[0] && str[0] <= 'z'; isLower {
		return str
	}

	newstr := make([]byte, 0, len(str))

	for i := 0; i < len(str); i++ {
		c := str[i]

		if i == 0 {
			c += 'a' - 'A'

			newstr = append(newstr, c)
		} else {
			newstr = append(newstr, c)
		}
	}

	return BytesToString(newstr)
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func SimpleCopyProperties(dst, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针，.Elem()类似于*ptr的操作返回指针指向的地址反射类型
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}

	return nil
}
