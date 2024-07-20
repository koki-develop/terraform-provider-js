data "js_raw" "number" {
  value = "Number"
}

#
# function add(a, b) { return a + b }
#

resource "js_function" "add" {
  name   = "add"
  params = [js_function_param.add_a.id, js_function_param.add_b.id]
  body   = [js_return.a_plus_b.content]
}

resource "js_function_param" "add_a" {
  name = "a"
}

resource "js_function_param" "add_b" {
  name = "b"
}

resource "js_return" "a_plus_b" {
  value = js_operation.a_plus_b.content
}

resource "js_operation" "a_plus_b" {
  left     = js_function_param.add_a.id
  operator = "+"
  right    = js_function_param.add_b.id
}

#
# function sum(nums) { return nums.reduce(add) }
#

resource "js_function" "sum" {
  name   = "sum"
  params = [js_function_param.sum_nums.id]
  body   = [js_return.nums_reduce.content]
}

resource "js_function_param" "sum_nums" {
  name = "nums"
}

resource "js_return" "nums_reduce" {
  value = js_function_call.nums_reduce.content
}

resource "js_function_call" "nums_reduce" {
  caller   = js_function_param.sum_nums.id
  function = "reduce"
  args     = [js_function.add.id]
}

#
# function digits(num) { return num.toString().split("").map(Number) }
#

resource "js_function" "digits" {
  name   = "digits"
  params = [js_function_param.digits_num.id]
  body   = [js_return.num_to_string_split_map.content]
}

resource "js_function_param" "digits_num" {
  name = "num"
}

resource "js_return" "num_to_string_split_map" {
  value = js_function_call.num_to_string_split_map.content
}

resource "js_function_call" "num_to_string_split_map" {
  caller   = js_function_call.num_to_string_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

resource "js_function_call" "num_to_string_split" {
  caller   = js_function_call.num_to_string.content
  function = "split"
  args     = [""]
}

resource "js_function_call" "num_to_string" {
  caller   = js_function_param.digits_num.id
  function = "toString"
}

#
# function sumDigits(num) { return sum(digits(num)) }
#

resource "js_function" "sum_digits" {
  name   = "sumDigits"
  params = [js_function_param.sum_digits_num.id]
  body   = [js_return.sum_digits.content]
}

resource "js_function_param" "sum_digits_num" {
  name = "num"
}

resource "js_return" "sum_digits" {
  value = js_function_call.sum_digits_num.content
}

resource "js_function_call" "sum_digits_num" {
  function = js_function.sum.id
  args     = [js_function_call.digits_num.content]
}

resource "js_function_call" "digits_num" {
  function = js_function.digits.id
  args     = [js_function_param.sum_digits_num.id]
}

#
# const nab = input.split(" ").map(Number)
#

resource "js_const" "nab" {
  name  = "nab"
  value = js_function_call.input_split_map.content
}

resource "js_function_call" "input_split_map" {
  caller   = js_function_call.input_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

resource "js_function_call" "input_split" {
  caller   = js_function_param.input.id
  function = "split"
  args     = [" "]
}

# nab[0]
data "js_index" "n" {
  ref   = js_const.nab.id
  value = 0
}

# nab[1]
data "js_index" "a" {
  ref   = js_const.nab.id
  value = 1
}

# nab[2]
data "js_index" "b" {
  ref   = js_const.nab.id
  value = 2
}

#
# let sum = 0
#

resource "js_let" "ans" {
  name  = "ans"
  value = 0
}

#
# for (let i = 0; i <= nab[0]; i++)
#

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

#
# const cmp = sumDigits(i)
#

resource "js_const" "cmp" {
  name  = "cmp"
  value = js_function_call.sum_digits_i.content
}

resource "js_function_call" "sum_digits_i" {
  function = js_function.sum_digits.id
  args     = [js_let.for_i.id]
}

#
# if (a <= cmp && cmp <= b) { sum += i }
#

resource "js_if" "sum" {
  condition = js_operation.a_le_cmp_and_cmp_le_b.content
  then      = [js_operation.sum_plus_i.content]
}

# a <= cmp && cmp <= b
resource "js_operation" "a_le_cmp_and_cmp_le_b" {
  left     = js_operation.a_le_cmp.content
  operator = "&&"
  right    = js_operation.cmp_le_b.content
}

# a <= cmp
resource "js_operation" "a_le_cmp" {
  left     = data.js_index.a.id
  operator = "<="
  right    = js_const.cmp.id
}

# cmp <= b
resource "js_operation" "cmp_le_b" {
  left     = js_const.cmp.id
  operator = "<="
  right    = data.js_index.b.id
}

# sum += i
resource "js_operation" "sum_plus_i" {
  left     = js_let.ans.id
  operator = "+="
  right    = js_let.for_i.id
}

#
# console.log(sum)
#

resource "js_function_call" "log_sum" {
  caller   = "console"
  function = "log"
  args     = [js_let.ans.id]
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_function.add.content,
    js_function.sum.content,
    js_function.digits.content,
    js_function.sum_digits.content,
    js_const.nab.content,
    js_let.ans.content,
    js_for.sum.content,
    js_function_call.log_sum.content,
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
