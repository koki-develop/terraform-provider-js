run "hello" {
  assert {
    condition     = data.js_function.hello.content == "@js/raw:function hello(){return \"hello world\"}"
    error_message = ""
  }
}
