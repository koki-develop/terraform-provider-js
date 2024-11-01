run "hello" {
  assert {
    condition     = data.js_function.hello.content == "@js/raw:function hello(name){console.log(\"hello\",name)}"
    error_message = ""
  }
}

run "anonymous" {
  assert {
    condition     = data.js_function.anonymous.content == "@js/raw:function(name){console.log(\"hello\",name)}"
    error_message = ""
  }
}

run "async" {
  assert {
    condition     = data.js_function.async.content == "@js/raw:async function hello(name){console.log(\"hello\",name)}"
    error_message = ""
  }
}
