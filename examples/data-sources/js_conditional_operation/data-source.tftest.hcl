run "main_statement" {
  assert {
    condition     = data.js_conditional_operation.main.statement == "@js/raw:(true?1:2)"
    error_message = ""
  }
}

run "main_expression" {
  assert {
    condition     = data.js_conditional_operation.main.expression == "@js/raw:(true?1:2)"
    error_message = ""
  }
}
