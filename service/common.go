package service

import "strings"

const(
	httpsUrl               = `https://`
	httpUrl               = `http://`
)

func AddHttpsString(url string) string {
	if strings.Index(url, "v.media.daum.net/") == -1{
		return httpsUrl + url
	}else{
		return httpUrl + url
	}
}
