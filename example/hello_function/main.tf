data "js_function" "console_log" {
  name = "console.log"
}

resource "js_function" "hello" {
  name   = "hello"
  params = [js_function_param.hello_name.id]
  body   = [js_function_call.log_message.content]
} # => function hello(name){console.log("hello", name)}

resource "js_function_param" "hello_name" {
  name = "name"
}

resource "js_function_call" "log_message" {
  function = data.js_function.console_log.id
  args     = ["hello", js_function_param.hello_name.id]
} # => console.log("hello", name)

output "content" {
  value = js_function.hello.content
}
