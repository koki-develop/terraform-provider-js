run "main" {
  assert {
    condition     = js_export.message.content == "@js/raw:export message"
    error_message = ""
  }
}
