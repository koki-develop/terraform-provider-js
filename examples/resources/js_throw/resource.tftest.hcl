run "error" {
  assert {
    condition     = js_throw.error.content == "@js/raw:throw new Error(\"something went wrong\")"
    error_message = ""
  }
}
