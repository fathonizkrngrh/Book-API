package utils

import (
	"net/url"
	"strconv"
)

func ParseInt(s string) int {
	val, _ := strconv.ParseInt(s, 10, 32)
	return int(val)
}

func IsValidURL(urlString string) bool {
	u, err := url.Parse(urlString)
	return err == nil && u.Scheme != "" && u.Host != ""
}
func CheckThickness(total_page int) string {
	var thickness string
	if total_page <= 100 {
		thickness = "tipis"
	} else if total_page <= 200 {
		thickness = "sedang"
	} else {
		thickness = "tebal"
	}

	return thickness
}