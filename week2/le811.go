package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/subdomain-visit-count/
// 811. Subdomain Visit Count
// 811. 子域名访问计数

func subdomainVisits(cpdomains []string) []string {
	counts := map[string]int{}
	for _, cpdomain := range cpdomains {
		count, domains := parseDomain(cpdomain)
		for _, domain := range domains {
			counts[domain] += count
		}
	}
	out := make([]string, 0, len(counts))
	for domain, count := range counts {
		out = append(out, fmt.Sprintf("%v %s", count, domain))
	}
	return out
}

func parseDomain(cp string) (count int, domains []string) {
	id := strings.IndexByte(cp, ' ')
	count, _ = strconv.Atoi(cp[0:id])
	domains = make([]string, 0, 3)
	domain := cp[id+1:]
	domains = append(domains, domain)
	for idx := strings.IndexByte(domain, '.'); idx >= 0; idx = strings.IndexByte(domain, '.') {
		domain = domain[idx+1:]
		domains = append(domains, domain)
	}
	return
}
