/**
 * @typedef {Object} MailParseResult
 * @property {string} messageId
 * @property {string} subject
 * @property {string} html
 * @property {string} text
 */

/**
 * Parse email
 * @param {Uint8Array} raw - email raw data
 * @param {Promise<ArrayBuffer> | function(): Promise<ArrayBuffer>} wasmLoader - wasm loader
 * @returns {Promise<MailParseResult>}
 */
export async function parseMail(raw, wasmLoader) {
    await import("./static/wasm_exec.js" )
    const go = new Go()
    const loader = wasmLoader || defaultLoader
    const wasm = typeof loader === "function" ? await loader() : await loader
    const result = await WebAssembly.instantiate(wasm, go.importObject)
    go.run(result.instance)
    const mailJson = parseEmail(raw)
    return JSON.parse(mailJson)
}

/**
 * Wasm loader
 * @param {string} url
 * @returns {Promise<ArrayBuffer>}
 */
export async function urlLoader(url) {
    const response = await fetch(url)
    return await response.arrayBuffer()
}

/**
 * Default wasm loader
 * @returns {Promise<ArrayBuffer>}
 */
export async function defaultLoader() {
    return urlLoader("https://raw.githubusercontent.com/TBXark/gomail-wasm/master/static/mail.wasm")
}