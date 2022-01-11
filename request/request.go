package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"imgurproxy/config"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

var cfg = config.GetConfig()

func Make(path, authorization string) (*http.Response) {
	var protocol string
	if cfg.Imgur.UseHttps {
		protocol = "https:"
	} else {
		protocol = "http:"
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("%v//%v/%v/%v", protocol, cfg.Imgur.ApiDomain, strconv.Itoa(cfg.Imgur.ApiVersion), path), nil)
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}

	if len(authorization) != 0 {
		request.Header.Add("Authorization", authorization)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

func Oauth(path string, body map[string]string) OAuth {
	var protocol string
	if cfg.Imgur.UseHttps {
		protocol = "https:"
	} else {
		protocol = "http:"
	}
	
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for p, v := range body {
		writer.WriteField(p, v)
	}
	err := writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%v//%v/oauth2/%v", protocol, cfg.Imgur.ApiDomain, path), payload)
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	b, _ := ioutil.ReadAll(response.Body)

	oauth := OAuth{}
	json.Unmarshal(b, &oauth)

	return oauth
}

func GetResource(path, authorization string) Resource {
	var protocol string
	if cfg.Imgur.UseHttps {
		protocol = "https:"
	} else {
		protocol = "http:"
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("%v//%v/%v/%v", protocol, cfg.Imgur.ApiDomain, strconv.Itoa(cfg.Imgur.ApiVersion), path), nil)
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}

	if len(authorization) != 0 {
		request.Header.Add("Authorization", authorization)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	resource := Resource{}
	json.Unmarshal(body, &resource)

	return resource
}

func GetResourcesArray(path, authorization string) Resources {
	var protocol string
	if cfg.Imgur.UseHttps {
		protocol = "https:"
	} else {
		protocol = "http:"
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("%v//%v/%v/%v", protocol, cfg.Imgur.ApiDomain, strconv.Itoa(cfg.Imgur.ApiVersion), path), nil)
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}

	if len(authorization) != 0 {
		request.Header.Add("Authorization", authorization)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	resources := Resources{}
	json.Unmarshal(body, &resources)

	return resources
}

func GetAccount(path, authorization string) Account {
	var protocol string
	if cfg.Imgur.UseHttps {
		protocol = "https:"
	} else {
		protocol = "http:"
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("%v//%v/%v/%v", protocol, cfg.Imgur.ApiDomain, strconv.Itoa(cfg.Imgur.ApiVersion), path), nil)
	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}

	if len(authorization) != 0 {
		request.Header.Add("Authorization", authorization)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	account := Account{}
	json.Unmarshal(body, &account)

	return account
}