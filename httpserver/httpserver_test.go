package httpserver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/hello?name=TestHTTP")
	if err != nil {
		t.Fatalf("failed to get /hello: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	want := "Hello, TestHTTP!"
	if string(body) != want {
		t.Errorf("unexpected response: got %q, want %q", string(body), want)
	}
}

func TestHelloHandler_DefaultName(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/hello")
	if err != nil {
		t.Fatalf("failed to get /hello: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	want := "Hello, World!"
	if string(body) != want {
		t.Errorf("unexpected response: got %q, want %q", string(body), want)
	}
}

func TestPingHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(pingHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Fatalf("failed to get /ping: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "pong" {
		t.Errorf("unexpected response: got %q, want %q", string(body), "pong")
	}
}

func TestReverseHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(reverseHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/reverse?text=abcde")
	if err != nil {
		t.Fatalf("failed to get /reverse: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "edcba" {
		t.Errorf("unexpected response: got %q, want %q", string(body), "edcba")
	}
}

func TestSumHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(sumHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/sum?a=3&b=4")
	if err != nil {
		t.Fatalf("failed to get /sum: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "7" {
		t.Errorf("unexpected response: got %q, want %q", string(body), "7")
	}
}

func TestSumHandler_InvalidInput(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(sumHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/sum?a=foo&b=bar")
	if err != nil {
		t.Fatalf("failed to get /sum: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", resp.StatusCode)
	}
}
