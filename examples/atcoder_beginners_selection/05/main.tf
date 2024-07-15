#
# const n, a, b
#

resource "js_const" "nab" {
  name  = "nab"
  value = js_function_call.input_map.content
}

resource "js_function_call" "input_split" {
  caller   = js_function_param.input.id
  function = "split"
  args     = [" "]
}

resource "js_function_call" "input_map" {
  caller   = js_function_call.input_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_raw" "number" {
  value = "Number"
}

data "js_index" "n" {
  ref   = js_const.nab.id
  value = 0
}

data "js_index" "a" {
  ref   = js_const.nab.id
  value = 1
}

data "js_index" "b" {
  ref   = js_const.nab.id
  value = 2
}

#
# sum
#

resource "js_let" "sum" {
  name  = "sum"
  value = 0
}

resource "js_for" "sum" {
  init      = js_let.for_i.content
  condition = js_operation.for_condition.content
  update    = js_increment.for_update.content
  body = [
    js_const.cmp.content,
    js_if.sum.content,
  ]
}

resource "js_let" "for_i" {
  name  = "i"
  value = 0
}

resource "js_operation" "for_condition" {
  left     = js_let.for_i.id
  operator = "<="
  right    = data.js_index.n.id
}

resource "js_increment" "for_update" {
  ref = js_let.for_i.id
}

resource "js_const" "cmp" {
  name  = "cmp"
  value = js_function_call.i_reduce.content
}

#
# split
#

resource "js_function_call" "i_string" {
  caller   = js_let.for_i.id
  function = "toString"
}

resource "js_function_call" "i_split" {
  caller   = js_function_call.i_string.content
  function = "split"
  args     = [""]
}

resource "js_function_call" "i_map" {
  caller   = js_function_call.i_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

#
# reduce
#

resource "js_function_call" "i_reduce" {
  caller   = js_function_call.i_map.content
  function = "reduce"
  args     = [js_function.i_reduce.content, 0]
}

resource "js_function" "i_reduce" {
  params = [js_function_param.acc.id, js_function_param.cur.id]
  body   = [js_return.acc_plus_cur.content]
}

resource "js_function_param" "acc" {
  name = "acc"
}

resource "js_function_param" "cur" {
  name = "cur"
}

resource "js_operation" "acc_plus_cur" {
  left     = js_function_param.acc.id
  operator = "+"
  right    = js_function_param.cur.id
}

resource "js_return" "acc_plus_cur" {
  value = js_operation.acc_plus_cur.content
}

resource "js_if" "sum" {
  condition = js_operation.a_le_cmp_and_cmp_le_b.content
  then      = [js_operation.sum_plus_i.content]
}

#
# a <= cmp && cmp <= b
#

resource "js_operation" "a_le_cmp" {
  left     = data.js_index.a.id
  operator = "<="
  right    = js_const.cmp.id
}

resource "js_operation" "cmp_le_b" {
  left     = js_const.cmp.id
  operator = "<="
  right    = data.js_index.b.id
}

resource "js_operation" "a_le_cmp_and_cmp_le_b" {
  left     = js_operation.a_le_cmp.content
  operator = "&&"
  right    = js_operation.cmp_le_b.content
}

#
# sum += i
#

resource "js_operation" "sum_plus_i" {
  left     = js_let.sum.id
  operator = "+="
  right    = js_let.for_i.id
}

#
# print sum
#

resource "js_function_call" "log_sum" {
  caller   = "console"
  function = "log"
  args     = [js_let.sum.id]
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_const.nab.content,
    js_let.sum.content,
    js_for.sum.content,
    js_function_call.log_sum.content,
  ]
}

resource "js_function_param" "input" {
  name = "input"
}

#
# call main
#

resource "js_function_call" "main" {
  function = js_function.main.id
  args     = [js_function_call.read_stdin.content]
}

resource "js_function_call" "require_fs" {
  function = "require"
  args     = ["fs"]
}

resource "js_function_call" "read_stdin" {
  caller   = js_function_call.require_fs.content
  function = "readFileSync"
  args     = ["/dev/stdin", "utf8"]
}

resource "js_program" "main" {
  contents = [
    js_function.main.content,
    js_function_call.main.content,
  ]
}

#
# write to file
#

resource "local_file" "main" {
  filename = "index.js"
  content  = js_program.main.content
}
