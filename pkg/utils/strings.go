package utils

import "net/url"

func StrPointer(str string) *string {
	return &str
}

func StrToURL(str string) (*url.URL, error) {
	return url.Parse(str)
}
