package shared

import (
	"net/url"
	"os/exec"
	"runtime"
)

func IfNilEmpty(s *string) string {
	if s == nil {
		return "n/a"
	} else {
		if len(*s) == 0 {
			return "n/a"
		}
		return *s
	}
}

func IfNilPlaceholder(s *string) string {
	if s == nil {
		return "https://www.froben11.de/wp-content/uploads/2016/10/orionthemes-placeholder-image.png"
	} else {
		if len(*s) < 1 {
			return "https://www.froben11.de/wp-content/uploads/2016/10/orionthemes-placeholder-image.png"
		}
		return *s
	}
}

func GetHostNameFromUrl(u string) string {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return u
	}

	return parsedUrl.Hostname()
}

func OpenBrowser(id int, url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}

	if err != nil {
		OError(id, "Failed opening browser for: "+url)
	}
}
