package helpers

import (
	"os"
	"strings"
	"go/build"
	"github.com/mitchellh/go-homedir"
)

var GOPATH string
var GOPATHS []string

func init() {
	var (
		err error
		pth string
		ok  bool
	)

	paths := make(map[string]int)

	for _, pth = range strings.Split(os.Getenv("GOPATH"), ":") {
		if pth != "" {
			if _, ok = paths[pth]; !ok {
				GOPATHS = append(GOPATHS, pth)
			}
		}
	}

	pth, err = homedir.Expand("~/go")
	if err != nil {
		panic(err)
	}

	if _, err = os.Stat(pth); err == nil {
		if _, ok = paths[pth]; !ok {
			GOPATHS = append(GOPATHS, pth)
		}
	}

	pth = build.Default.GOPATH
	if _, ok = paths[pth]; !ok {
		GOPATHS = append(GOPATHS, pth)
	}

	GOPATH = GOPATHS[0]
}