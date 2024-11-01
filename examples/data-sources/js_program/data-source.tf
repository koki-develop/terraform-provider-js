data "js_const" "message" {
  name  = "message"
  value = "hello world"
}

data "js_function_call" "log_message" {
  caller   = "console"
  function = "log"
  args     = [data.js_const.message.id]
}

data "js_program" "main" {
  contents = [
    data.js_const.message.content,
    data.js_function_call.log_message.content,
  ]
}
# => const message = "hello world";
#    console.log(message);
