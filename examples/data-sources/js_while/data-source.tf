data "js_let" "i" {
  name  = "i"
  value = 0
}

data "js_operation" "i_lt_10" {
  left     = data.js_let.i.id
  operator = "<"
  right    = 10
}

data "js_function_call" "log_i" {
  caller   = "console"
  function = "log"
  args     = [data.js_let.i.id]
}

data "js_increment" "i" {
  ref = data.js_let.i.id
}

data "js_while" "main" {
  condition = data.js_operation.i_lt_10.content
  body      = [data.js_function_call.log_i.content, data.js_increment.i.content]
}
# => while (i < 10) {
#      console.log(i);
#      i++;
#    }
