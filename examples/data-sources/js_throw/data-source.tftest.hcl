run "error" {
  assert {
    condition     = data.js_throw.error.statement == "@js/raw:throw new Error(\"something went wrong\")"
    error_message = ""
  }
}
