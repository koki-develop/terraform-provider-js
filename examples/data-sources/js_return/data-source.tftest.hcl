run "hello" {
  assert {
    condition     = data.js_function.hello.statement == "@js/raw:function hello(){return \"hello world\"}"
    error_message = ""
  }
}
