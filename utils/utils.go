// Copyright 2020 The shop Authors

// Package utils implements utils.
package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Float64ToString float64转string
func Float64ToString(val float64) string {
	return strconv.FormatFloat(val, 'E', -1, 64)
}

// GetUUID 获取uuid
func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	} else {
		return uuid.String()
	}
}

// GetTimeStampStr 获取时间串
func GetTimeStampStr() string {
	return time.Now().Format("2006/1/2 15:04:05")
}

// GetTimestamp the result likes 1423361979
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// FormatTimestamp the result likes 2015-02-08 10:19:39 AM
func FormatTimestamp(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

// FormatTimestampStr 格式化time.Time为string
func FormatTimestampStr(timestamp time.Time) string {
	tm := time.Unix(timestamp.Unix(), 0)
	return tm.Format("2006/1/2 15:04:05")
}

// RandomStr 生成随机字符串
func RandomStr(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}

// ChangeStr 在字符串中加中横线
func ChangeStr(number string) string {
	byteStr := []byte(number)
	result := make([]string, 0)
	for i := 0; i < len(byteStr); i++ {
		if i != 0 && i%5 == 0 {
			result = append(result, "-")
		}

		result = append(result, string(byteStr[i]))
	}
	return strings.Join(result, "")
}

// MysqlEscapeString sql语句特殊字符处理
func MysqlEscapeString(realStr string) string {
	newStr := strings.ReplaceAll(realStr, "_", "\\_")
	newStr = strings.ReplaceAll(newStr, "%", "\\%")
	newStr = strings.ReplaceAll(newStr, "'", "\\'")
	newStr = strings.ReplaceAll(newStr, "\"", "\\\"")
	return newStr
}
