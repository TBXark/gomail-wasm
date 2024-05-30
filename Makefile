.PHONY: build
build:
	cp "$$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./static/
	tinygo build -o static/mail.wasm -target wasm ./mail.go

.PHONY: test
test:
	go test -v ./...