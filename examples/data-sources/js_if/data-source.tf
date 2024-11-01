data "js_function_call" "log_true" {
  caller   = "console"
  function = "log"
  args     = [true]
}

data "js_function_call" "log_false" {
  caller   = "console"
  function = "log"
  args     = [false]
}

data "js_raw" "true" {
  value = "true"
}

data "js_if" "main" {
  condition = data.js_raw.true.content
  then      = [data.js_function_call.log_true.content]
  else      = [data.js_function_call.log_false.content]
}
# => if (true) {
#      console.log(true);
#    } else {
#      console.log(false);
#    }
