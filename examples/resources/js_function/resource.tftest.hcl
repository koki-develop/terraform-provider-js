run "hello" {
  assert {
    condition     = js_function.hello.content == "@js/raw:function hello(name){console.log(\"hello\",name)}"
    error_message = ""
  }
}

run "anonymous" {
  assert {
    condition     = js_function.anonymous.content == "@js/raw:function(name){console.log(\"hello\",name)}"
    error_message = ""
  }
}
