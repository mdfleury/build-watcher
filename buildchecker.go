/**
 * @file buildchecker.go
 * Watch build numbers for Turner sites
 *
 * @author Matt Fleury <matthew.fleury@turner.com>
 */

package buildchecker

import (
	"fmt"
	"github.com/mirtchovski/gosxnotifier"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var (
	sitesStore      []string
	pollTime        = 10 * time.Second
	siteNumberStore = map[string]string{}
)

func notify(host string, number string) {
	note := gosxnotifier.NewNotification("Check your Apple Stock!")
	err := note.Push()
	if err != nil {
		fmt.Println("Notifcation Error")
	}
}

func CheckNumber(site string) (string, string) {
	number := GetBuildNumber(site)
	siteUrl, err := url.Parse(site)
	if err != nil {
		return site, "Unable to parse URL"
	}
	return siteUrl.Host, number
}

func GetBuildNumber(site string) string {
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

func CheckSites(sites []string) {
	for _, site := range sites {
		host, number := CheckNumber(site)
		fmt.Println(host, ":", number)
	}
}

func CheckSiteChanges(sites []string) {
	for _, site := range sites {
		host, number := CheckNumber(site)
		fmt.Println(host, ":", number)
		if _, ok := siteNumberStore[host]; ok {
			if siteNumberStore[host] != number {
				notify(host, number)
			}
		} else {
			siteNumberStore[host] = number
		}
	}
}

func WatchSites(sites []string) {
	sitesStore = sites
	CheckSiteChanges(sitesStore)
	for range time.Tick(pollTime) {
		CheckSiteChanges(sitesStore)
	}
}
