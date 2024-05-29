package main

import (
	"fmt"
	"net/http"
	"testing"
)

//func TestParseEmail(t *testing.T) {
//	raw, err := os.ReadFile("static/test.eml")
//	if err != nil {
//		t.Fatal(err)
//	}
//	res, err := ParseEmail(raw)
//	if err != nil {
//		t.Fatal(err)
//	}
//	log.Printf("MessageId: %s\n", res.MessageId)
//	log.Printf("Subject: %s\n", res.Subject)
//	log.Printf("HTML: %s\n", res.HTML)
//	log.Printf("Text: %s\n", res.Text)
//}

func TestParseEmail(t *testing.T) {
	fmt.Printf("Starting server: http://localhost:9999\n")
	_ = http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))
}
