package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var websiteUrl string = "https://keyspot.app"
var apiUrl string = "https://database-driver-ifhogzjzbq-uc.a.run.app"
var configFilePath string = os.Getenv("HOME") + "/.keyspot"

type JwtPayload struct {
	Sub string
	Iat int
}

type JsonRecordDocument struct {
	Id     string            `json:"_id"`
	Name   string            `json:"name"`
	Sub    string            `json:"sub"`
	Record map[string]string `json:"record"`
}

func parseJwtPayload(token string) (*JwtPayload, error) {
	payload := &JwtPayload{}
	payloadJsonString, err := jwt.DecodeSegment(strings.Split(token, ".")[1])

	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(payloadJsonString), payload)

	return payload, nil
}

func apiGet(route string, token string) (string, error) {
	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", apiUrl+route, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string([]byte(body)), nil
}

func recordFromJsonString(jsonString string) *map[string]string {
	document := JsonRecordDocument{}

	json.Unmarshal([]byte(jsonString), &document)

	return &document.Record
}

func getToken(tokenPath string) (string, error) {
	buffer, err := ioutil.ReadFile(tokenPath)

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
