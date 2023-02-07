package mytoys

import (
	"regexp"
	"strings"
)

func GetPidFromUrl(url string) string {
	re := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	matches := re.FindStringSubmatch(url)
	slug := matches[2]

	parts := strings.Split(slug, "-")

	pid := parts[len(parts)-1]

	return pid
}
