package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerHomepage(t *testing.T) {

}

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/helloworld", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloWorld(res, req)

	exp := "Hello World"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}
