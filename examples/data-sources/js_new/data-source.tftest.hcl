run "error" {
  assert {
    condition     = data.js_new.error.content == "@js/raw:new Error(\"something went wrong\")"
    error_message = ""
  }
}
