resource "js_function" "hello" {
  name   = "hello"
  params = [js_function_param.hello_name.id]
  body   = [js_function_call.log_message.content]
}

resource "js_function_param" "hello_name" {
  name = "name"
}

resource "js_function_call" "log_message" {
  caller   = "console"
  function = "log"
  args     = ["hello", js_function_param.hello_name.id]
}

resource "js_function_call" "hello" {
  function = js_function.hello.id
  args     = ["world"]
}

resource "js_program" "main" {
  contents = [
    js_function.hello.content,
    js_function_call.hello.content,
  ]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = js_program.main.content
}
