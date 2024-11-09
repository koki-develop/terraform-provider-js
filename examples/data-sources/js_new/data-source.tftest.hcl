run "error" {
  assert {
    condition     = data.js_new.error.statement == "@js/raw:new Error(\"something went wrong\")"
    error_message = ""
  }
}
