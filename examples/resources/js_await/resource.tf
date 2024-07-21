resource "js_function_call" "fetch" {
  function = "fetch"
  args     = ["https://example.com"]
}

resource "js_await" "fetch" {
  value = js_function_call.fetch.content
}
# => await fetch("https://example.com")
