resource "js_const" "anonymous" {
  name  = "anonymous"
  value = js_function.anonymous.content
}

resource "js_function" "anonymous" {
  params = [js_function_param.anonymous_message.id]
  body   = [js_function_call.log_message.content]
}

resource "js_function_param" "anonymous_message" {
  name = "message"
}

resource "js_function_call" "log_message" {
  function = "console.log"
  args     = [js_function_param.anonymous_message.id]
}

resource "js_program" "main" {
  contents = [
    js_const.anonymous.content,
  ]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = js_program.main.content
}
