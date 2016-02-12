// Watch build numbers for Turner sites
// by Matt Fleury <matthew.fleury@turner.com>

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"regexp"
)

var sites []string = []string {
	"http://dev.ncaa.com/build-number.txt",
	"http://qa.ncaa.com/build-number.txt",
	"http://staging.ncaa.com/build-number.txt",
	"http://admin.ncaa.com/build-number.txt",
	"http://gametool.ncaa.com/build-number.txt",
	"http://carmen-qa.ncaa.com/build.html",
	"http://carmen.ncaa.com/build.html",
	"http://carmen-staging.ncaa.com/build.html",
}

func checkNumber(site string) {
	number := getBuildNumber(site)
	siteUrl, err := url.Parse(site)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(siteUrl.Host, ":", number)
}

func getBuildNumber(site string)string {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	text := string(body)
	resp.Body.Close()

	tester := regexp.MustCompile(`build.number=([0-9]+)`)
	number := tester.FindStringSubmatch(text)
	return number[0]
}

func main () {
	for _, site := range sites {
		checkNumber(site)
	}
}
