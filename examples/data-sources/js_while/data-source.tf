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
  condition = data.js_operation.i_lt_10.expression
  body      = [data.js_function_call.log_i.statement, data.js_increment.i.statement]
}
# => while (i < 10) {
#      console.log(i);
#      i++;
#    }
