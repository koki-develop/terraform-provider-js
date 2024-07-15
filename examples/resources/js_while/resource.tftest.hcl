run "main" {
  assert {
    condition     = js_while.main.content == "@js/raw:while(i<10){console.log(i);i++}"
    error_message = ""
  }
}
