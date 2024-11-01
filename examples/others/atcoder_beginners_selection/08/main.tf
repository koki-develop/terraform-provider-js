#
# const ny = input.trim().split(" ")
#

data "js_const" "ny" {
  name  = "ny"
  value = data.js_function_call.input_trim_split.content
}

data "js_function_call" "input_trim_split" {
  caller   = data.js_function_call.input_trim.content
  function = "split"
  args     = [" "]
}

data "js_function_call" "input_trim" {
  caller   = data.js_function_param.input.id
  function = "trim"
}

#
# const n = Number(ny[0])
#

data "js_const" "n" {
  name  = "n"
  value = data.js_function_call.number_n.content
}

data "js_function_call" "number_n" {
  function = "Number"
  args     = [data.js_index.n.content]
}

data "js_index" "n" {
  ref   = data.js_const.ny.id
  value = 0
}

#
# const y = Number(ny[1])
#

data "js_const" "y" {
  name  = "y"
  value = data.js_function_call.number_y.content
}

data "js_function_call" "number_y" {
  function = "Number"
  args     = [data.js_index.y.content]
}

data "js_index" "y" {
  ref   = data.js_const.ny.id
  value = 1
}

#
# for (let i = 0; i <= n; i++)
#

data "js_for" "i" {
  init      = data.js_let.for_i.content
  condition = data.js_operation.i_le_n.content
  update    = data.js_increment.i.content
  body      = [data.js_for.j.content]
}

data "js_let" "for_i" {
  name  = "i"
  value = 0
}

data "js_operation" "i_le_n" {
  left     = data.js_let.for_i.id
  operator = "<="
  right    = data.js_const.n.id
}

data "js_increment" "i" {
  ref = data.js_let.for_i.id
}

#
# for (let j = 0; j <= n - i; j++)
#

data "js_for" "j" {
  init      = data.js_let.for_j.content
  condition = data.js_operation.j_le_n_minus_i.content
  update    = data.js_increment.j.content

  body = [
    data.js_const.k.content,
    data.js_if.answer.content,
  ]
}

data "js_let" "for_j" {
  name  = "j"
  value = 0
}

data "js_operation" "j_le_n_minus_i" {
  left     = data.js_operation.j_le_n.content
  operator = "-"
  right    = data.js_let.for_i.id
}

data "js_operation" "j_le_n" {
  left     = data.js_let.for_j.id
  operator = "<="
  right    = data.js_const.n.id
}

data "js_increment" "j" {
  ref = data.js_let.for_j.id
}

#
# const k = n - i - j
#

data "js_const" "k" {
  name  = "k"
  value = data.js_operation.n_minus_i_minus_j.content
}

data "js_operation" "n_minus_i_minus_j" {
  left     = data.js_operation.n_minus_i.content
  operator = "-"
  right    = data.js_let.for_j.id
}

data "js_operation" "n_minus_i" {
  left     = data.js_const.n.id
  operator = "-"
  right    = data.js_let.for_i.id
}

#
# if (10000 * i + 5000 * j + 1000 * k === y)
#

data "js_if" "answer" {
  condition = data.js_operation._10000_times_i_plus_5000_times_j_plus_1000_times_k_eq_y.content
  then = [
    data.js_function_call.log_answer.content,
    data.js_return.answer.content,
  ]
}

data "js_operation" "_10000_times_i_plus_5000_times_j_plus_1000_times_k_eq_y" {
  left     = data.js_operation._10000_times_i_plus_5000_times_j_plus_1000_times_k.content
  operator = "==="
  right    = data.js_const.y.id
}

data "js_operation" "_10000_times_i_plus_5000_times_j_plus_1000_times_k" {
  left     = data.js_operation._10000_times_i_plus_5000_times_j.content
  operator = "+"
  right    = data.js_operation._1000_times_k.content
}

data "js_operation" "_10000_times_i_plus_5000_times_j" {
  left     = data.js_operation._10000_times_i.content
  operator = "+"
  right    = data.js_operation._5000_times_j.content
}

data "js_operation" "_10000_times_i" {
  left     = 10000
  operator = "*"
  right    = data.js_let.for_i.id
}

data "js_operation" "_5000_times_j" {
  left     = 5000
  operator = "*"
  right    = data.js_let.for_j.id
}

data "js_operation" "_1000_times_k" {
  left     = 1000
  operator = "*"
  right    = data.js_const.k.id
}

data "js_return" "answer" {}

#
# console.log(i, j, k)
#

data "js_function_call" "log_answer" {
  caller   = "console"
  function = "log"
  args     = [data.js_let.for_i.id, data.js_let.for_j.id, data.js_const.k.id]
}

#
# console.log(-1, -1, -1)
#

data "js_function_call" "log_no_answer" {
  caller   = "console"
  function = "log"
  args     = [-1, -1, -1]
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_const.ny.content,
    data.js_const.n.content,
    data.js_const.y.content,
    data.js_for.i.content,
    data.js_function_call.log_no_answer.content,
  ]
}

data "js_function_param" "input" {
  name = "input"
}

#
# main(require("fs").readFileSync("/dev/stdin", "utf8"))
#

data "js_function_call" "main" {
  function = data.js_function.main.id
  args     = [data.js_function_call.read_stdin.content]
}

data "js_function_call" "require_fs" {
  function = "require"
  args     = ["fs"]
}

data "js_function_call" "read_stdin" {
  caller   = data.js_function_call.require_fs.content
  function = "readFileSync"
  args     = ["/dev/stdin", "utf8"]
}

#
# write to file
#

data "js_program" "main" {
  contents = [
    data.js_function.main.content,
    data.js_function_call.main.content,
  ]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = data.js_program.main.content
}
