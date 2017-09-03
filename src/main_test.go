package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	b64 "encoding/base64"
)


const ValidLogin = "bob:WhatAboutBob?!"
const InvalidLogin = "aaaa:bbbbb"

func checkBody(t *testing.T, r *http.Response, body string) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		t.Error("reading reponse body: %v, want %q", err, body)
	}
	if g, w := string(b), body; g != w {
		t.Errorf("request body mismatch: got %q, want %q", g, w)
	}
}

func TestBlankIndex(t *testing.T) {
	server := httptest.NewServer(GetMainEngine())
	defer server.Close()

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 404 {
		t.Fatalf("Index should not exist, but it returns non-404: %d\n", resp.StatusCode)
	}
	
}

func TestAuthRequired(t *testing.T) {
	server := httptest.NewServer(GetMainEngine())
	defer server.Close()

	var jsonStr = []byte(`{}`)
	req, _ := http.NewRequest("POST", server.URL + "/api/v1/wc", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 401 {
		t.Fatalf("Authentication is not enabled on the /wc endpoint: %d\n", resp.StatusCode)
	}
}


func TestAuthBadPW(t *testing.T) {
	server := httptest.NewServer(GetMainEngine())
	defer server.Close()

	var jsonStr = []byte(`{}`)
	req, _ := http.NewRequest("POST", server.URL + "/api/v1/wc", bytes.NewBuffer(jsonStr))

	loginData := InvalidLogin
	loginEnc := b64.StdEncoding.EncodeToString([]byte(loginData))
	
	req.Header.Set("Authorization", "Basic " + loginEnc)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 403 {
		t.Fatalf("Authentication with invalid credentials should not work: %d\n", resp.StatusCode)
	}
	
}

func TestAuthValidPW(t *testing.T) {
	server := httptest.NewServer(GetMainEngine())
	defer server.Close()

	var jsonStr = []byte(`{"input":"foo bar"}`)
	req, _ := http.NewRequest("POST", server.URL + "/api/v1/wc", bytes.NewBuffer(jsonStr))

	loginData := ValidLogin
	loginEnc := b64.StdEncoding.EncodeToString([]byte(loginData))
	
	req.Header.Set("Authorization", "Basic " + loginEnc)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Authentication with valid credentials should work: %d\n", resp.StatusCode)
	}
	
}

func TestWordCount(t *testing.T) {
	server := httptest.NewServer(GetMainEngine())
	defer server.Close()

	var jsonStr = []byte(`{"input":"foo bar baz"}`)
	req, _ := http.NewRequest("POST", server.URL + "/api/v1/wc", bytes.NewBuffer(jsonStr))

	loginData := ValidLogin
	loginEnc := b64.StdEncoding.EncodeToString([]byte(loginData))
	
	req.Header.Set("Authorization", "Basic " + loginEnc)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Wordcount endpoint should return 200 but returned %d\n", resp.StatusCode)
	}

	expected := "{\"count\":3,\"words\":{\"bar\":1,\"baz\":1,\"foo\":1}}"
	checkBody(t, resp, expected)
}
