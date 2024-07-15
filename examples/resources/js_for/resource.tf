resource "js_let" "i" {
  name  = "i"
  value = 0
}

resource "js_operation" "i_lt_10" {
  left     = js_let.i.id
  operator = "<"
  right    = 10
}

resource "js_increment" "i" {
  ref = js_let.i.id
}

resource "js_function_call" "log_i" {
  caller   = "console"
  function = "log"
  args     = [js_let.i.id]
}

resource "js_for" "main" {
  init      = js_let.i.content
  condition = js_operation.i_lt_10.content
  update    = js_increment.i.content
  body      = [js_function_call.log_i.content]
}
# => for (let i = 0; i < 10; i++) {
#      console.log(i);
#    }
