package main

import (
	"fmt"
)

func main()  {
	str := "[[[u'acl number 3000 name internet_biz_icfw01_acl', u' rule 5 permit ip source 42.36.151.0 0.0.0.255', u' rule 10 permit ip source 42.36.97.0 0.0.0.255'], [u'acl number 3001 name internet_biz_icfw02_acl', u' rule 5 permit ip source 42.36.150.0 0.0.0.255'], [u'acl number 3002 name internet_test_icfw01_acl'], [u'acl number 3003 name internet_test_icfw02_acl']], [[u'acl ipv6 number 3000 name internet_biz_icfw01_acl', u' rule 5 permit ipv6 source 2010:1010::/64', u' rule 10 permit ipv6 source 2010:1100:0:10::/64', u' rule 15 permit ipv6 source 2020:1100::1234:0/112'], [u'acl ipv6 number 3001 name internet_biz_icfw02_acl'], [u'acl ipv6 number 3002 name internet_test_icfw01_acl'], [u'acl ipv6 number 3003 name internet_test_icfw02_acl']]]"
	ipv4Arr := getIpv4Rules(str)
	fmt.Println("ipv4Arr: ", ipv4Arr)

	ipv6Arr := getIpv6Rules(str)
	fmt.Println("ipv6Arr: ", ipv6Arr)
}