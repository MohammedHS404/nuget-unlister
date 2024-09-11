package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	cookie := ""
	requestVerificationToken := ""
	packageName := ""
	packageVersions := [...]string{}

	for i := 0; i < len(packageVersions); i++ {
		unlistPackage(packageName, packageVersions[i], requestVerificationToken, cookie)
	}
}

func unlistPackage(packageName string, packageVersion string, requestVerificationToke string, cookie string) {
	url := "https://www.nuget.org/packages/" + packageName + "/" + packageVersion + "/UpdateListed?aria_label=Change%20listing%20of%20this%20package%20version"
	method := "POST"

	payload := strings.NewReader("__RequestVerificationToken=" + requestVerificationToke + "&version=" + packageVersion + "&Listed=false")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("DNT", "1")
	req.Header.Add("Origin", "https://www.nuget.org")
	req.Header.Add("Referer", "https://www.nuget.org/packages/"+packageName+"/"+packageVersion+"/Manage")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
