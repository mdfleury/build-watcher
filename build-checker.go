// Watch build numbers for Turner sites
// by Matt Fleury <matthew.fleury@turner.com>

package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "regexp"
    "strings"
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
    resp, err := http.Get(site)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    text := string(body)
    // fmt.Println(text)
    resp.Body.Close()
    tester := regexp.MustCompile(`build.number=([0-9]+)`)
    number := tester.FindStringSubmatch(text)
    urlParts := strings.Split(site, "/");
    host := urlParts[2]
    fmt.Println(host, ":", number[1])
}

func main () {
    for _, site := range sites {
        checkNumber(site)
    }
}
