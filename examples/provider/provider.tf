data "js_function_call" "hello_world" {
  caller   = "console"
  function = "log"
  args     = ["hello world"]
}

data "js_program" "main" {
  contents = [data.js_function_call.hello_world.content]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = data.js_program.main.content
}
