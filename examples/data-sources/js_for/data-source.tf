data "js_let" "i" {
  name  = "i"
  value = 0
}

data "js_operation" "i_lt_10" {
  left     = data.js_let.i.id
  operator = "<"
  right    = 10
}

data "js_increment" "i" {
  ref = data.js_let.i.id
}

data "js_function_call" "log_i" {
  caller   = "console"
  function = "log"
  args     = [data.js_let.i.id]
}

data "js_for" "main" {
  init      = data.js_let.i.statement
  condition = data.js_operation.i_lt_10.expression
  update    = data.js_increment.i.statement
  body      = [data.js_function_call.log_i.statement]
}
# => for (let i = 0; i < 10; i++) {
#      console.log(i);
#    }
