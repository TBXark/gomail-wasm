.PHONY: build
build:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static/
	GOOS=js GOARCH=wasm go build -o static/mail.wasm

.PHONY: base64
base64:
	@echo "export const mailWasmBase64=\`$$(base64 --input=static/mail.wasm)\`" > static/mailwasm.js

.PHONY: test
test:
	go test -v ./...