run "alert" {
  assert {
    condition     = js_function_call.alert.content == "@js/raw:alert(\"hello world\")"
    error_message = ""
  }
}

run "console_log" {
  assert {
    condition     = js_function_call.console_log.content == "@js/raw:console.log(\"hello\",\"world\")"
    error_message = ""
  }
}
