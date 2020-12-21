package main

import (
	"regexp"
)

func getIpv4Rules(str string) []string {
	ipv4Regular := `(((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3})|( 0\n)`
	reg := regexp.MustCompile(ipv4Regular)
	ipArr := reg.FindAllString(str, -1)
	return ipArr
}
