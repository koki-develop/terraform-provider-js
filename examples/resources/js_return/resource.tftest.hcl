run "hello" {
  assert {
    condition     = js_function.hello.content == "@js/raw:function hello(){return \"hello world\"}"
    error_message = ""
  }
}
