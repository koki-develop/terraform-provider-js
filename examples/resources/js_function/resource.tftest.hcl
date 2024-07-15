run "hello" {
  assert {
    condition     = js_function.hello.content == "@js/raw:function hello(name){console.log(\"hello\",name)}"
    error_message = ""
  }
}
