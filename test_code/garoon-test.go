package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"net/http"
	"text/template"
	"bytes"
	"log"
)

func main() {

	type APIParameters struct {
		Username, Password string
	}

	const SOAPRequest = `<?xml version="1.0" encoding="UTF-8"?>
  <soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
    <soap:Header>
      <Action>
        ScheduleGetEvents
      </Action>
      <Security>
        <UsernameToken>
          <Username>{{.Username}}</Username>
          <Password>{{.Password}}</Password>
        </UsernameToken>
      </Security>
      <Timestamp>
        <Created>2010-08-12T14:45:00Z</Created>
        <Expires>2037-08-12T14:45:00Z</Expires>
      </Timestamp>
      <Locale>jp</Locale>
    </soap:Header>
    <soap:Body>
    <ScheduleGetEvents>
      <parameters start="2018-04-12T01:00:00" end="2018-04-20T10:00:00"/>
    </ScheduleGetEvents>
    </soap:Body>
  </soap:Envelope>
  `

	// コマンドラインから読み込む
	inputUsername     := flag.String("u", "test name", "user name")
	inputPassword     := flag.String("p", "test password", "user password")
	inputGaroonDomain := flag.String("d", "test domain", "garoon domain")
	flag.Parse()
	
	apiParameters := APIParameters{*inputUsername, *inputPassword}
	url           := fmt.Sprintf("https://%s/cgi-bin/cbgrn/grn.cgi/cbpapi/schedule/api?", *inputGaroonDomain)

	//リクエストbodyのテンプレートを適用
	t    := template.Must(template.New("SOAPRequest").Parse(SOAPRequest))
	body := bytes.NewBufferString("")
	err  := t.Execute(body, apiParameters)
	if err != nil {
		log.Panicf("template.Execute: %#v", err)
	}

	// Post
	response, err := http.Post(url, "text/xml; charset=utf-8", body)
	if err != nil {
		log.Panicf("http.Post: %#v", err)
	}

	// レスポンスを読み込む
	b, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Panicf("ioutil.ReadAll: %#v", err)
	}
	fmt.Printf(string(b))
	return
}
