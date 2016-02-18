// Watch build numbers for Turner sites
// by Matt Fleury <matthew.fleury@turner.com>

package buildchecker

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"regexp"
)

func CheckNumber(site string) (string, string) {
	number := GetBuildNumber(site)
	siteUrl, err := url.Parse(site)
	if err != nil {
		return site, "Unable to parse URL"
	}
	return siteUrl.Host, number
}

func GetBuildNumber(site string)string {
	resp, err := http.Get(site)
	if err != nil {
		return "Unable to reach page"
	}

	if resp.StatusCode != 200 {
		return "Invalid status code"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Unable to read body"
	}
	text := string(body)
	resp.Body.Close()

	tester := regexp.MustCompile(`build.number=([0-9]+)`)
	number := tester.FindStringSubmatch(text)
	return number[1]
}
