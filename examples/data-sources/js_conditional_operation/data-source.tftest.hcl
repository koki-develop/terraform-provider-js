run "main" {
  assert {
    condition     = data.js_conditional_operation.main.content == "@js/raw:(true?1:2)"
    error_message = ""
  }
}
