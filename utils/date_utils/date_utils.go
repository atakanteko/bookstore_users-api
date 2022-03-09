package date_utils

import "time"

const apiDateLayout = "02-01-2006 15:04:05"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
