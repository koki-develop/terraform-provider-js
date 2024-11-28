run "alert" {
  assert {
    condition     = data.js_function_call.alert.statement == "@js/raw:(alert)(\"hello world\")"
    error_message = ""
  }
}

run "console_log" {
  assert {
    condition     = data.js_function_call.console_log.statement == "@js/raw:(console.log)(\"hello\",\"world\")"
    error_message = ""
  }
}
