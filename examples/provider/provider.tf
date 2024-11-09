data "js_function_call" "hello_world" {
  caller   = "console"
  function = "log"
  args     = ["hello world"]
}

data "js_program" "main" {
  statements = [data.js_function_call.hello_world.statement]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = data.js_program.main.content
}
