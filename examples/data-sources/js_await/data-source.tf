data "js_function_call" "fetch" {
  function = "fetch"
  args     = ["https://example.com"]
}

data "js_await" "fetch" {
  value = data.js_function_call.fetch.expression
}
# => await fetch("https://example.com")
