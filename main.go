package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	var allNews string

	for i := 1; i <= 5; i++ {
		allNews += getNews(i)
	}

	errFile := os.WriteFile("news.html", []byte(allNews), 0644)
	if errFile != nil {
		log.Panic(errFile)
	}

}

func getNews(pageNumber int) string {
	//url := fmt.Sprint("https://www.jcnet.com.br/webparts.php?sesit=5&site=25&file=%252Fapp%252Fview%252Fparts%252Fpart-listao&p=", pageNumber)
	url := fmt.Sprint("https://www.jcnet.com.br/webparts.php?sesit=5&site=25&file=%2Fapp%2Fview%2Fparts%2Fpart-listao&p=", pageNumber)

	fmt.Println(url)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("authority", "www.jcnet.com.br")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
	req.Header.Add("accept", "text/html, */*; q=0.01")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://www.jcnet.com.br/noticias/ultimas_noticias")
	req.Header.Add("accept-language", "pt,pt-BR;q=0.9,en;q=0.8")
	req.Header.Add("cookie", "_ga=GA1.3.1210185455.1639703925; _gid=GA1.3.802578550.1639703925; _fw_plugins=1196662060; _fw_fonts=459217059; _fw_clientIp=179.127.52.121; _fw_validIP=true; evercookie_png=61bbe57944b64; evercookie_etag=61bbe57944b64; evercookie_cache=61bbe57944b64; _everfw4=61bbe57944b64; _everfw4=61bbe57944b64; _fw_views=1; cookieconsent_status=dismiss; paywall=true; _fw_userName=; _fw_signId=; _fw_email=; _fw_plan=; _fw_planId=; _fw_planValid=false; _dc_gtm_UA-24117260-1=1; _gat_gtag_UA_175164381_36=1")

	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	//fmt.Println(res)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	return string(body)
}
