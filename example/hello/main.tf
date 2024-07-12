data "js_function" "console_log" {
  name = "console.log"
}

resource "js_function_call" "hello_world" {
  function = data.js_function.console_log.id
  args     = ["hello world"]
} # => console.log("hello world")

output "content" {
  value = js_function_call.hello_world.content
}
