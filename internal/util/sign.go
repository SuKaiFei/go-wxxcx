package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

type LocalTime time.Time

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

func GetSign(data map[string]interface{}, secret string) string {
	var keyList = make([]string, 0, len(data)-1)
	for k := range data {
		if k == "sign" {
			continue
		}
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	var signStr = new(bytes.Buffer)
	for _, key := range keyList {
		value := data[key]
		kind := reflect.TypeOf(value).Kind()
		if kind == reflect.Slice || kind == reflect.Interface {
			continue
		}
		_, _ = signStr.WriteString(fmt.Sprintf("%v%v", key, value))
	}
	sign := strings.ToUpper(HmacSha256(signStr.String(), secret))
	return sign
}

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
