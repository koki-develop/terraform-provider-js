run "hello" {
  assert {
    condition     = data.js_function.hello.statement == "@js/raw:function hello(name){(console.log)(\"hello\",name)}"
    error_message = ""
  }
}

run "anonymous" {
  assert {
    condition     = data.js_function.anonymous.statement == "@js/raw:function(name){(console.log)(\"hello\",name)}"
    error_message = ""
  }
}

run "async" {
  assert {
    condition     = data.js_function.async.statement == "@js/raw:async function hello(name){(console.log)(\"hello\",name)}"
    error_message = ""
  }
}
