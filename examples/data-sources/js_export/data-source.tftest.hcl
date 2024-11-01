run "main" {
  assert {
    condition     = data.js_export.message.content == "@js/raw:export message"
    error_message = ""
  }
}
