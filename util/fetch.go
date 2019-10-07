package util

func Fetch(lang string, since string) ([]string, error) {
	url := "https://github.com/trending/" + lang + "?since=" + since
	return Parse(url)
}
