run "main" {
  assert {
    condition     = data.js_while.main.statement == "@js/raw:while((i<10)){(console.log)(i);i++}"
    error_message = ""
  }
}
