package main

import (
	"fmt"
	"net/url"
)

func main() {
	var URL = "%5B%7B%22ip%22%3A%22119.90.49.95%22%2C%22port%22%3A%228074%22%7D%2C%7B%22ip%22%3A%22119.90.49.110%22%2C%22port%22%3A%228050%22%7D%2C%7B%22ip%22%3A%22119.90.49.92%22%2C%22port%22%3A%228057%22%7D%2C%7B%22ip%22%3A%22119.90.49.104%22%2C%22port%22%3A%228016%22%7D%2C%7B%22ip%22%3A%22119.90.49.104%22%2C%22port%22%3A%228017%22%7D%2C%7B%22ip%22%3A%22119.90.49.102%22%2C%22port%22%3A%228007%22%7D%2C%7B%22ip%22%3A%22119.90.49.107%22%2C%22port%22%3A%228034%22%7D%2C%7B%22ip%22%3A%22119.90.49.110%22%2C%22port%22%3A%228049%22%7D%2C%7B%22ip%22%3A%22119.90.49.91%22%2C%22port%22%3A%228053%22%7D%2C%7B%22ip%22%3A%22119.90.49.94%22%2C%22port%22%3A%228069%22%7D%5D"
	res, err := url.QueryUnescape(URL)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	urlEnCodeStr := url.QueryEscape(res)
	fmt.Println(urlEnCodeStr)
}
