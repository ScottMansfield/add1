// needs to be the same package to access main()
package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()
	<-time.After(100 * time.Millisecond)
	os.Exit(m.Run())
}

func TestSmallInt(t *testing.T) {
	res, err := http.Get("http://localhost:7001/?num=1234")
	if err != nil {
		t.Fatalf("Error calling server: %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Fatalf("Response was not a 200. Error: %s", res.Status)
	}

	bodybytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response: %s", err.Error())
	}

	body := string(bodybytes)

	if body != "1235" {
		t.Fatalf("Got unexpected result: %s", body)
	}

	t.Log("Success")
}

func TestBigInt(t *testing.T) {
	res, err := http.Get("http://localhost:7001/?num=123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	if err != nil {
		t.Fatalf("Error calling server: %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Fatalf("Response was not a 200. Error: %s", res.Status)
	}

	bodybytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response: %s", err.Error())
	}

	body := string(bodybytes)

	if body != "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567891" {
		t.Fatalf("Got unexpected result: %s", body)
	}

	t.Log("Success")
}

func TestSmallFloat(t *testing.T) {
	res, err := http.Get("http://localhost:7001/?num=0.1234")
	if err != nil {
		t.Fatalf("Error calling server: %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Fatalf("Response was not a 200. Error: %s", res.Status)
	}

	bodybytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response: %s", err.Error())
	}

	body := string(bodybytes)

	if body != "1.1234" {
		t.Fatalf("Got unexpected result: %s", body)
	}

	t.Log("Success")
}

func TestBigFloat(t *testing.T) {
	res, err := http.Get("http://localhost:7001/?num=1.234e100")
	if err != nil {
		t.Fatalf("Error calling server: %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Fatalf("Response was not a 200. Error: %s", res.Status)
	}

	bodybytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response: %s", err.Error())
	}

	body := string(bodybytes)

	if body != "12340000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001" {
		t.Fatalf("Got unexpected result: %s", body)
	}

	t.Log("Success")
}

func TestBadNumberFormat(t *testing.T) {
	res, err := http.Get("http://localhost:7001/?num=abcd")
	if err != nil {
		t.Fatalf("Error calling server: %s", err.Error())
	}

	if res.StatusCode != 400 {
		t.Fatalf("Response was not a 200. Error: %s", res.Status)
	}

	t.Log("Success")
}

func TestHexNumber(t *testing.T) {
	res, err := http.Get("http://localhost:7001/?num=0xff")
	if err != nil {
		t.Fatalf("Error calling server: %s", err.Error())
	}

	// Ideally, I could parse hexadecimal, octal, and all other bases but oh well
	if res.StatusCode != 400 {
		t.Fatalf("Response was not a 200. Error: %s", res.Status)
	}

	t.Log("Success")
}
