//go:build js

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/emersion/go-message/mail"
	"io"
	"strings"
	"syscall/js"
)

type MailParseResult struct {
	MessageId string `json:"messageId"`
	Subject   string `json:"subject"`
	HTML      string `json:"html"`
	Text      string `json:"text"`
}

func ParseEmail(raw []byte) (*MailParseResult, error) {
	reader := bytes.NewReader(raw)
	mr, err := mail.CreateReader(reader)
	if err != nil {
		return nil, err
	}
	defer mr.Close()

	var result MailParseResult
	header := mr.Header

	result.MessageId = header.Get("Message-Id")
	result.Subject = header.Get("Subject")

	for {
		part, rErr := mr.NextPart()
		if rErr == io.EOF {
			break
		} else if rErr != nil {
			return nil, rErr
		}
		switch h := part.Header.(type) {
		case *mail.InlineHeader:
			contentType, _, _ := h.ContentType()
			if strings.HasPrefix(contentType, "text/html") {
				body, _ := io.ReadAll(part.Body)
				result.HTML = string(body)
			} else if strings.HasPrefix(contentType, "text/plain") {
				body, _ := io.ReadAll(part.Body)
				result.Text = string(body)
			}
		}
	}
	return &result, nil
}

func jsParseEmailWrapper(this js.Value, args []js.Value) any {
	jsArray := args[0]
	byteArray := make([]byte, jsArray.Get("length").Int())
	js.CopyBytesToGo(byteArray, jsArray)
	res, err := ParseEmail(byteArray)
	if err != nil {
		return js.ValueOf(nil)
	}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		return js.ValueOf(nil)
	}
	return js.ValueOf(string(jsonRes))
}

func main() {
	fmt.Printf("WASM loaded\n")
	js.Global().Set("parseEmail", js.FuncOf(jsParseEmailWrapper))
	select {}
}
