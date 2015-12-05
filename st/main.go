package main

import (
	"fmt"
	"github.com/shaalx/goutils"
	"github.com/shaalx/jsnm"
	"net/url"
)

func main() {
	// TestJson()
	jm := jsnm.FileNameFmt("site.md")
	arr := jm.Arr()
	for i, it := range arr {
		fmt.Print(i)
		parseURI(it.RawData().String())
	}
	t := 1
	for k, v := range mp {
		fmt.Println(t)
		t++
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println()
	}
}

var mp map[string]string

func init() {
	mp = make(map[string]string)
}

func parseURI(uri string) {
	URL, err := url.Parse(uri)
	if goutils.CheckErr(err) {
		return
	}
	qry := URL.Query()
	targetURL := qry.Get("url")
	// fmt.Printf("target URL: %s\n", targetURL)
	// addgroupid := qry.Get("adgroupid")
	// fmt.Printf("adgroupid URL: %s\n", addgroupid)
	// adid := qry.Get("adid")
	// fmt.Printf("adid URL: %s\n", adid)
	if _, exist := mp[targetURL]; !exist {
		mp[targetURL] = uri
	}
}
