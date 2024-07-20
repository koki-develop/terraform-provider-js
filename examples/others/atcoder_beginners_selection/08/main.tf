#
# const ny = input.trim().split(" ")
#

resource "js_const" "ny" {
  name  = "ny"
  value = js_function_call.input_trim_split.content
}

resource "js_function_call" "input_trim_split" {
  caller   = js_function_call.input_trim.content
  function = "split"
  args     = [" "]
}

resource "js_function_call" "input_trim" {
  caller   = js_function_param.input.id
  function = "trim"
}

#
# const n = Number(ny[0])
#

resource "js_const" "n" {
  name  = "n"
  value = js_function_call.number_n.content
}

resource "js_function_call" "number_n" {
  function = "Number"
  args     = [data.js_index.n.id]
}

data "js_index" "n" {
  ref   = js_const.ny.id
  value = 0
}

#
# const y = Number(ny[1])
#

resource "js_const" "y" {
  name  = "y"
  value = js_function_call.number_y.content
}

resource "js_function_call" "number_y" {
  function = "Number"
  args     = [data.js_index.y.id]
}

data "js_index" "y" {
  ref   = js_const.ny.id
  value = 1
}

#
# for (let i = 0; i <= n; i++)
#

resource "js_for" "i" {
  init      = js_let.for_i.content
  condition = js_operation.i_le_n.content
  update    = js_increment.i.content
  body      = [js_for.j.content]
}

resource "js_let" "for_i" {
  name  = "i"
  value = 0
}

resource "js_operation" "i_le_n" {
  left     = js_let.for_i.id
  operator = "<="
  right    = js_const.n.id
}

resource "js_increment" "i" {
  ref = js_let.for_i.id
}

#
# for (let j = 0; j <= n - i; j++)
#

resource "js_for" "j" {
  init      = js_let.for_j.content
  condition = js_operation.j_le_n_minus_i.content
  update    = js_increment.j.content

  body = [
    js_const.k.content,
    js_if.answer.content,
  ]
}

resource "js_let" "for_j" {
  name  = "j"
  value = 0
}

resource "js_operation" "j_le_n_minus_i" {
  left     = js_operation.j_le_n.content
  operator = "-"
  right    = js_let.for_i.id
}

resource "js_operation" "j_le_n" {
  left     = js_let.for_j.id
  operator = "<="
  right    = js_const.n.id
}

resource "js_increment" "j" {
  ref = js_let.for_j.id
}

#
# const k = n - i - j
#

resource "js_const" "k" {
  name  = "k"
  value = js_operation.n_minus_i_minus_j.content
}

resource "js_operation" "n_minus_i_minus_j" {
  left     = js_operation.n_minus_i.content
  operator = "-"
  right    = js_let.for_j.id
}

resource "js_operation" "n_minus_i" {
  left     = js_const.n.id
  operator = "-"
  right    = js_let.for_i.id
}

#
# if (10000 * i + 5000 * j + 1000 * k === y)
#

resource "js_if" "answer" {
  condition = js_operation._10000_times_i_plus_5000_times_j_plus_1000_times_k_eq_y.content
  then = [
    js_function_call.log_answer.content,
    js_return.answer.content,
  ]
}

resource "js_operation" "_10000_times_i_plus_5000_times_j_plus_1000_times_k_eq_y" {
  left     = js_operation._10000_times_i_plus_5000_times_j_plus_1000_times_k.content
  operator = "==="
  right    = js_const.y.id
}

resource "js_operation" "_10000_times_i_plus_5000_times_j_plus_1000_times_k" {
  left     = js_operation._10000_times_i_plus_5000_times_j.content
  operator = "+"
  right    = js_operation._1000_times_k.content
}

resource "js_operation" "_10000_times_i_plus_5000_times_j" {
  left     = js_operation._10000_times_i.content
  operator = "+"
  right    = js_operation._5000_times_j.content
}

resource "js_operation" "_10000_times_i" {
  left     = 10000
  operator = "*"
  right    = js_let.for_i.id
}

resource "js_operation" "_5000_times_j" {
  left     = 5000
  operator = "*"
  right    = js_let.for_j.id
}

resource "js_operation" "_1000_times_k" {
  left     = 1000
  operator = "*"
  right    = js_const.k.id
}

resource "js_return" "answer" {}

#
# console.log(i, j, k)
#

resource "js_function_call" "log_answer" {
  caller   = "console"
  function = "log"
  args     = [js_let.for_i.id, js_let.for_j.id, js_const.k.id]
}

#
# console.log(-1, -1, -1)
#

resource "js_function_call" "log_no_answer" {
  caller   = "console"
  function = "log"
  args     = [-1, -1, -1]
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_const.ny.content,
    js_const.n.content,
    js_const.y.content,
    js_for.i.content,
    js_function_call.log_no_answer.content,
  ]
}

resource "js_function_param" "input" {
  name = "input"
}

#
# main(require("fs").readFileSync("/dev/stdin", "utf8"))
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

#
# write to file
#

resource "js_program" "main" {
  contents = [
    js_function.main.content,
    js_function_call.main.content,
  ]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = js_program.main.content
}
