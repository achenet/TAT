package main

import (
	"fmt"
	"testing"
)

func TestNewRequestFromJson(t *testing.T) {
	request, err := NewRequestFromJson("test1.json")
	if err != nil {
		t.Error("Error getting new request from json:", err.Error())
	}
	fmt.Println(request)
}
