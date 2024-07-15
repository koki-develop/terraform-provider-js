resource "js_let" "i" {
  name  = "i"
  value = 0
}

resource "js_operation" "i_lt_10" {
  left     = js_let.i.id
  operator = "<"
  right    = 10
}

resource "js_function_call" "log_i" {
  caller   = "console"
  function = "log"
  args     = [js_let.i.id]
}

resource "js_increment" "i" {
  ref = js_let.i.id
}

resource "js_while" "main" {
  condition = js_operation.i_lt_10.content
  body      = [js_function_call.log_i.content, js_increment.i.content]
}
# => while (i < 10) {
#      console.log(i);
#      i++;
#    }
