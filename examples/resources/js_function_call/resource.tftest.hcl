run "console_log" {
  assert {
    condition     = js_function_call.console_log.content == "@js/raw:console.log(\"hello\",\"world\")"
    error_message = ""
  }
}
