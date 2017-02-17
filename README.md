# Is my build out yet? #

Sample 1: Check builds (runs once then stops)

```
#!go
package main

import (
	"bitbucket.org/mfleury_turner/build-watcher"
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

func main () {
	buildchecker.CheckSites(sites)
}
```

Sample 2: Watch for builds (runs continuously, alerting of changes)

```
#!go
package main

import (
	"bitbucket.org/mfleury_turner/build-watcher"
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

func main () {
	buildchecker.WatchSites(sites)
}

```