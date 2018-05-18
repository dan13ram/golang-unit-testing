package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestHelloHandler(t *testing.T) {
	server := httptest.NewServer(&helloHandler{})
	defer server.Close()

	random := randSeq(10)

	var tests = []struct {
		name     string
		path     string
		expected string
	}{
		{"HelloWorld", "/", "Hello, World!"},
		{"Hello" + random, "/" + random, "Hello, " + random + "!"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(server.URL + tc.path)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != 200 {
				t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
			}
			actual, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if tc.expected != string(actual) {
				t.Errorf("%s test failed. Expected '%s', but got: '%s'\n", tc.name, tc.expected, actual)
			} else {
				t.Logf("%s test passed", tc.name)
			}
		})
	}
}
