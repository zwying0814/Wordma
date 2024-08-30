package ip_region

import (
	"strings"
)

func ipScraper(ip string) string {
	ip = strings.TrimSpace(ip)
	if ip == "" {
		return ""
	}

	// 多 IP 仅选第一个
	ipSep := strings.Split(ip, ",")
	if len(ipSep) > 1 {
		ip = strings.TrimSpace(ipSep[0])
	}

	return ip
}

func regionScraper(raw string) string {
	sep := strings.Split(raw, "|")
	if len(sep) < 5 {
		return ""
	}

	var (
		country  = sep[0]
		_        = sep[1] // Area
		province = strings.TrimSuffix(strings.TrimSuffix(sep[2], "省"), "市")
		// isp      = sep[4]
	)

	if country == "0" {
		return ""
	}

	if province == "0" {
		return country
	}
	return province
}
