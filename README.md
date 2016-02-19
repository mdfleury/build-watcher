# Is my build out yet? #

Sample 1: Check builds

```
#!go
package main

import (
	"mfleury/buildchecker"
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

Sample 2: Watch for builds

```
#!go
package main

import (
	"mfleury/buildchecker"
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