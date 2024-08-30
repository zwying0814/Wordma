package ip_region

import (
	"strings"
	"wordma/config"
	"wordma/log"
)

func IP2Region(ip string) string {
	if strings.TrimSpace(ip) == "" {
		return ""
	}

	ip = ipScraper(ip)
	region, err := search(ip, config.IPDataPath, true)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "invalid ip address") {
			log.Warn("[IP2Region] ", err)
		}
		return ""
	}

	return regionScraper(region)
}
