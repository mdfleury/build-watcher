// Watch build numbers for Turner sites
// by Matt Fleury <matthew.fleury@turner.com>

package buildchecker

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"regexp"
	"log"
)

func CheckNumber(site string) (string, string) {
	number := GetBuildNumber(site)
	siteUrl, err := url.Parse(site)
	if err != nil {
		log.Fatal(err)
	}
	return siteUrl.Host, number
}

func GetBuildNumber(site string)string {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		return "Unable to scrape page"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	text := string(body)
	resp.Body.Close()

	tester := regexp.MustCompile(`build.number=([0-9]+)`)
	number := tester.FindStringSubmatch(text)
	return number[1]
}
