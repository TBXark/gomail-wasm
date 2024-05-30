# GoMail Wasm

This is a simple mail parser written in Go and compiled to WebAssembly.

## Usage

```javascript
import {parseMail} from 'gomail-wasm'

const mail = parseMail(raw)

console.log(mail.messageId)
console.log(mail.subject)
console.log(mail.html)
console.log(mail.text)

```

## License

**GoMail Wasm** is released under the MIT license. [See LICENSE](LICENSE) for details.

