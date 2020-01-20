package main

import (
	"net/http"
	"testing"
)

func TestGetPosts(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	http.HandleFunc(getPosts)
}
