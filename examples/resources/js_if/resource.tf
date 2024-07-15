resource "js_function_call" "log_true" {
  caller   = "console"
  function = "log"
  args     = [true]
}

resource "js_function_call" "log_false" {
  caller   = "console"
  function = "log"
  args     = [false]
}

data "js_raw" "true" {
  value = "true"
}

resource "js_if" "main" {
  condition = data.js_raw.true.content
  then      = [js_function_call.log_true.content]
  else      = [js_function_call.log_false.content]
}
# => if (true) {
#      console.log(true);
#    } else {
#      console.log(false);
#    }
