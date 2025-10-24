package str_utils

import (
	"regexp"
	"strings"
)

func Slugify(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, "_", "-")
	str = strings.ReplaceAll(str, " ", "-")

	// Remove invalid characters (keep only a–z, 0–9, and -)
	reg := regexp.MustCompile(`[^a-z0-9-]+`)
	str = reg.ReplaceAllString(str, "")

	// Collapse multiple dashes into one
	regDash := regexp.MustCompile(`-+`)
	str = regDash.ReplaceAllString(str, "-")

	return strings.Trim(str, "-")
}

