resource "js_function_call" "hello_world" {
  caller   = "console"
  function = "log"
  args     = ["hello world"]
}

resource "js_program" "main" {
  contents = [js_function_call.hello_world.content]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = js_program.main.content
}
