.PHONY: build
build:
	cp "$$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./static/
	tinygo build -o static/mail.wasm -target wasm ./mail.go

.PHONY: base64
base64:
	@echo "export const mailWasmBase64=\`$$(base64 --input=static/mail.wasm)\`" > static/mailwasm.js

.PHONY: test
test:
	go test -v ./...