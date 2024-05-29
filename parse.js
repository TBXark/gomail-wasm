import "./static/wasm_exec.js"
import { mailWasmBase64 } from "./static/mailwasm.js"

export async function parseMail(raw) {
    const go = new Go()
    const wasmStrBuffer = atob(mailWasmBase64)
    const wasmCodeArray = new Uint8Array(wasmStrBuffer.length);
    for (let i = 0; i < wasmStrBuffer.length; i++) {
        wasmCodeArray[i] = wasmStrBuffer.charCodeAt(i);
    }
    const result = await WebAssembly.instantiate(wasmCodeArray, go.importObject)
    go.run(result.instance)
    const mailJson = parseEmail(raw)
    return  JSON.parse(mailJson)
}