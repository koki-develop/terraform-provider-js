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
  statements = [
    data.js_const.message.statement,
    data.js_function_call.log_message.statement,
  ]
}
# => const message = "hello world";
#    console.log(message);
