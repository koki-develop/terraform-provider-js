resource "js_const" "message" {
  name  = "message"
  value = "hello world"
}

resource "js_function_call" "log_message" {
  caller   = "console"
  function = "log"
  args     = [js_const.message.id]
}

resource "js_program" "main" {
  contents = [
    js_const.message.content,
    js_function_call.log_message.content,
  ]
}
# => const message = "hello world";
#    console.log(message);
