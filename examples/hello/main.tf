data "js_function" "console_log" {
  name = "console.log"
}

resource "js_function_call" "hello_world" {
  function = data.js_function.console_log.id
  args     = ["hello world"]
}

resource "js_program" "main" {
  contents = [js_function_call.hello_world.content]
}

output "result" {
  value = js_program.main.content
}
