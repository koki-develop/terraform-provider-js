#
# const abcx = input.trim().split("\n").map(Number)
#

resource "js_const" "abcx" {
  name  = "abcx"
  value = js_function_call.input_trim_split_map.content
}

resource "js_function_call" "input_trim_split_map" {
  caller   = js_function_call.input_trim_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

resource "js_function_call" "input_trim_split" {
  caller   = js_function_call.input_trim.content
  function = "split"
  args     = ["\n"]
}

resource "js_function_call" "input_trim" {
  caller   = js_function_param.input.id
  function = "trim"
}

data "js_raw" "number" {
  value = "Number"
}

# abcx[0]
data "js_index" "a" {
  ref   = js_const.abcx.id
  value = 0
}

# abcx[1]
data "js_index" "b" {
  ref   = js_const.abcx.id
  value = 1
}

# abcx[2]
data "js_index" "c" {
  ref   = js_const.abcx.id
  value = 2
}

# abcx[3]
data "js_index" "x" {
  ref   = js_const.abcx.id
  value = 3
}

#
# let count = 0
#

resource "js_let" "count" {
  name  = "count"
  value = 0
}

#
# for (let i = 0; i <= abcx[0]; i++)
#

resource "js_for" "i" {
  init      = js_let.for_i.content
  condition = js_operation.i_le_a.content
  update    = js_increment.for_i.content
  body      = [js_for.j.content]
}

resource "js_let" "for_i" {
  name  = "i"
  value = 0
}

resource "js_operation" "i_le_a" {
  left     = js_let.for_i.id
  operator = "<="
  right    = data.js_index.a.id
}

resource "js_increment" "for_i" {
  ref = js_let.for_i.id
}

#
# for (let j = 0; j <= abcx[1]; j++)
#

resource "js_for" "j" {
  init      = js_let.for_j.content
  condition = js_operation.j_le_b.content
  update    = js_increment.for_j.content
  body      = [js_for.k.content]
}

resource "js_let" "for_j" {
  name  = "j"
  value = 0
}

resource "js_operation" "j_le_b" {
  left     = js_let.for_j.id
  operator = "<="
  right    = data.js_index.b.id
}

resource "js_increment" "for_j" {
  ref = js_let.for_j.id
}

#
# for (let k = 0; k <= abcx[2]; k++)
#

resource "js_for" "k" {
  init      = js_let.for_k.content
  condition = js_operation.k_le_c.content
  update    = js_increment.for_k.content
  body      = [js_if.condition.content]
}

resource "js_let" "for_k" {
  name  = "k"
  value = 0
}

resource "js_operation" "k_le_c" {
  left     = js_let.for_k.id
  operator = "<="
  right    = data.js_index.c.id
}

resource "js_increment" "for_k" {
  ref = js_let.for_k.id
}

#
# if (i * 500 + j * 100 + k * 50 === abcx[3]) { count++ }
#

resource "js_if" "condition" {
  condition = js_operation.i_times_500_plus_j_times_100_plus_k_times_50_eq_x.content
  then      = [js_increment.count.content]
}

# i * 500 + j * 100 + k * 50 === abcx[3]
resource "js_operation" "i_times_500_plus_j_times_100_plus_k_times_50_eq_x" {
  left     = js_operation.i_times_500_plus_j_times_100_plus_k_times_50.content
  operator = "==="
  right    = data.js_index.x.id
}

# i * 500 + j * 100 + k * 50
resource "js_operation" "i_times_500_plus_j_times_100_plus_k_times_50" {
  left     = js_operation.i_times_500_plus_j_times_100.content
  operator = "+"
  right    = js_operation.k_times_50.content
}

# i * 500 + j * 100
resource "js_operation" "i_times_500_plus_j_times_100" {
  left     = js_operation.i_times_500.content
  operator = "+"
  right    = js_operation.j_times_100.content
}

# i * 500
resource "js_operation" "i_times_500" {
  left     = js_let.for_i.id
  operator = "*"
  right    = 500
}

# j * 100
resource "js_operation" "j_times_100" {
  left     = js_let.for_j.id
  operator = "*"
  right    = 100
}

# k * 50
resource "js_operation" "k_times_50" {
  left     = js_let.for_k.id
  operator = "*"
  right    = 50
}

resource "js_increment" "count" {
  ref = js_let.count.id
}

#
# console.log(count)
#

resource "js_function_call" "log_count" {
  caller   = "console"
  function = "log"
  args     = [js_let.count.id]
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_const.abcx.content,
    js_let.count.content,
    js_for.i.content,
    js_function_call.log_count.content,
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
