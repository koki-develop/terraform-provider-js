run "main" {
  assert {
    condition     = data.js_export.message.statement == "@js/raw:export message"
    error_message = ""
  }
}
