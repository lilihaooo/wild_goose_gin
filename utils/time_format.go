package utils

import "time"

func TimeFormat(ISO time.Time) string {
	return ISO.Format("2006-01-02 15:04:05")
}
