package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var websiteUrl string = "https://keyspot.app"
var apiUrl string = "https://database-driver-ifhogzjzbq-uc.a.run.app"

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

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return path + "/.keyspot", nil
}

func isConfigured() (string, error) {
	path, err := getConfigFilePath()

	if err != nil {
		return "", err
	}

	_, err = os.Stat(path)

	if os.IsNotExist(err) {
		return "", errors.New(fmt.Sprintf("No .keyspot file detected. The configure command needs to be run with a cli token before you can use these options. This can be done by acquiring a token from %s/account and running:\n\t$ keyspot configure <cli-token>", websiteUrl))
	}

	return path, nil
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

func apiPut(route string, jsonString string, token string) (string, error) {
	bearer := "Bearer " + token

	req, err := http.NewRequest("PUT", apiUrl+route, bytes.NewBuffer([]byte(jsonString)))

	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

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
	record := map[string]string{}

	json.Unmarshal([]byte(jsonString), &record)

	return &record
}

func documentFromJsonString(jsonString string) *JsonRecordDocument {
	document := JsonRecordDocument{}

	json.Unmarshal([]byte(jsonString), &document)

	return &document
}

func getToken(tokenPath string) (string, error) {
	buffer, err := ioutil.ReadFile(tokenPath)

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func getSubAndToken() (string, string, error) {
	configFilePath, err := isConfigured()

	if err != nil {
		return "", "", err
	}

	token, err := getToken(configFilePath)

	if err != nil {
		return "", "", errors.New("You need to configure your CLI tool.")
	}

	payload, err := parseJwtPayload(token)

	if err != nil {
		return "", "", err
	}

	return payload.Sub, token, nil
}
