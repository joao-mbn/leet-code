package main

import (
	"strings"
)

func simplifyPath(path string) string {
	sections := strings.Split(path, "/")
	canonicalSections := []string{}

	for _, section := range sections {
		if section == "" || section == "." {
			continue
		}

		if section == ".." && len(canonicalSections) > 0 {
			canonicalSections = canonicalSections[:len(canonicalSections)-1]
		} else if section == ".." {
			continue
		} else {
			canonicalSections = append(canonicalSections, section)
		}
	}

	return "/" + strings.Join(canonicalSections, "/")
}
