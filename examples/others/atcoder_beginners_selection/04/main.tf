#
# const abcx = input.trim().split("\n").map(Number)
#

data "js_const" "abcx" {
  name  = "abcx"
  value = data.js_function_call.input_trim_split_map.content
}

data "js_function_call" "input_trim_split_map" {
  caller   = data.js_function_call.input_trim_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_function_call" "input_trim_split" {
  caller   = data.js_function_call.input_trim.content
  function = "split"
  args     = ["\n"]
}

data "js_function_call" "input_trim" {
  caller   = data.js_function_param.input.id
  function = "trim"
}

data "js_raw" "number" {
  value = "Number"
}

# abcx[0]
data "js_index" "a" {
  ref   = data.js_const.abcx.id
  value = 0
}

# abcx[1]
data "js_index" "b" {
  ref   = data.js_const.abcx.id
  value = 1
}

# abcx[2]
data "js_index" "c" {
  ref   = data.js_const.abcx.id
  value = 2
}

# abcx[3]
data "js_index" "x" {
  ref   = data.js_const.abcx.id
  value = 3
}

#
# let count = 0
#

data "js_let" "count" {
  name  = "count"
  value = 0
}

#
# for (let i = 0; i <= abcx[0]; i++)
#

data "js_for" "i" {
  init      = data.js_let.for_i.content
  condition = data.js_operation.i_le_a.content
  update    = data.js_increment.for_i.content
  body      = [data.js_for.j.content]
}

data "js_let" "for_i" {
  name  = "i"
  value = 0
}

data "js_operation" "i_le_a" {
  left     = data.js_let.for_i.id
  operator = "<="
  right    = data.js_index.a.content
}

data "js_increment" "for_i" {
  ref = data.js_let.for_i.id
}

#
# for (let j = 0; j <= abcx[1]; j++)
#

data "js_for" "j" {
  init      = data.js_let.for_j.content
  condition = data.js_operation.j_le_b.content
  update    = data.js_increment.for_j.content
  body      = [data.js_for.k.content]
}

data "js_let" "for_j" {
  name  = "j"
  value = 0
}

data "js_operation" "j_le_b" {
  left     = data.js_let.for_j.id
  operator = "<="
  right    = data.js_index.b.content
}

data "js_increment" "for_j" {
  ref = data.js_let.for_j.id
}

#
# for (let k = 0; k <= abcx[2]; k++)
#

data "js_for" "k" {
  init      = data.js_let.for_k.content
  condition = data.js_operation.k_le_c.content
  update    = data.js_increment.for_k.content
  body      = [data.js_if.condition.content]
}

data "js_let" "for_k" {
  name  = "k"
  value = 0
}

data "js_operation" "k_le_c" {
  left     = data.js_let.for_k.id
  operator = "<="
  right    = data.js_index.c.content
}

data "js_increment" "for_k" {
  ref = data.js_let.for_k.id
}

#
# if (i * 500 + j * 100 + k * 50 === abcx[3]) { count++ }
#

data "js_if" "condition" {
  condition = data.js_operation.i_times_500_plus_j_times_100_plus_k_times_50_eq_x.content
  then      = [data.js_increment.count.content]
}

# i * 500 + j * 100 + k * 50 === abcx[3]
data "js_operation" "i_times_500_plus_j_times_100_plus_k_times_50_eq_x" {
  left     = data.js_operation.i_times_500_plus_j_times_100_plus_k_times_50.content
  operator = "==="
  right    = data.js_index.x.content
}

# i * 500 + j * 100 + k * 50
data "js_operation" "i_times_500_plus_j_times_100_plus_k_times_50" {
  left     = data.js_operation.i_times_500_plus_j_times_100.content
  operator = "+"
  right    = data.js_operation.k_times_50.content
}

# i * 500 + j * 100
data "js_operation" "i_times_500_plus_j_times_100" {
  left     = data.js_operation.i_times_500.content
  operator = "+"
  right    = data.js_operation.j_times_100.content
}

# i * 500
data "js_operation" "i_times_500" {
  left     = data.js_let.for_i.id
  operator = "*"
  right    = 500
}

# j * 100
data "js_operation" "j_times_100" {
  left     = data.js_let.for_j.id
  operator = "*"
  right    = 100
}

# k * 50
data "js_operation" "k_times_50" {
  left     = data.js_let.for_k.id
  operator = "*"
  right    = 50
}

data "js_increment" "count" {
  ref = data.js_let.count.id
}

#
# console.log(count)
#

data "js_function_call" "log_count" {
  caller   = "console"
  function = "log"
  args     = [data.js_let.count.id]
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_const.abcx.content,
    data.js_let.count.content,
    data.js_for.i.content,
    data.js_function_call.log_count.content,
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
