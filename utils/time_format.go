package utils

import "time"

// TimeFormat_YMD 格式化时间年月日
func TimeFormat_YMD(ISO time.Time) string {
	return ISO.Format("2006-01-02")
}

// TimeFormat_YMDHMS 格式化时间年月日时分秒
func TimeFormat_YMDHMS(ISO time.Time) string {
	return ISO.Format("2006-01-02 15:04:05")
}
