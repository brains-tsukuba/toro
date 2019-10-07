package util

func Fetch(lang string, since string) ([]string, error) {
	url := "https://github.com/trending"
	if lang != "" {
		url += "/" + lang
	}
	if since != "" {
		url += "?since=" + since
	}
	return Parse(url)
}
