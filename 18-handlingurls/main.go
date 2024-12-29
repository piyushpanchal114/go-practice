package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.trendmicro.com:443/cloudoneconformity/knowledge-base/aws?coursename=reactjs&payment_id=sdaf"

func main() {
	fmt.Println("Handling Url in Go")

	res, _ := url.Parse(myurl)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Port())
	fmt.Println(res.Path)
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("the type of query is: %T\n", qparams)

	for q, p := range qparams {
		fmt.Printf("The query is %v and the params is %v\n", q, p)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "www.trendmicro.com",
		Path:    "cloudoneconformity/knowledge-base/aws/EC2/",
		RawPath: "user=piyush",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println("The Url is:", anotherUrl)
}
