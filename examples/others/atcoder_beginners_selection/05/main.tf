data "js_raw" "number" {
  value = "Number"
}

#
# function add(a, b) { return a + b }
#

data "js_function" "add" {
  name   = "add"
  params = [data.js_function_param.add_a.id, data.js_function_param.add_b.id]
  body   = [data.js_return.a_plus_b.content]
}

data "js_function_param" "add_a" {
  name = "a"
}

data "js_function_param" "add_b" {
  name = "b"
}

data "js_return" "a_plus_b" {
  value = data.js_operation.a_plus_b.content
}

data "js_operation" "a_plus_b" {
  left     = data.js_function_param.add_a.id
  operator = "+"
  right    = data.js_function_param.add_b.id
}

#
# function sum(nums) { return nums.reduce(add) }
#

data "js_function" "sum" {
  name   = "sum"
  params = [data.js_function_param.sum_nums.id]
  body   = [data.js_return.nums_reduce.content]
}

data "js_function_param" "sum_nums" {
  name = "nums"
}

data "js_return" "nums_reduce" {
  value = data.js_function_call.nums_reduce.content
}

data "js_function_call" "nums_reduce" {
  caller   = data.js_function_param.sum_nums.id
  function = "reduce"
  args     = [data.js_function.add.id]
}

#
# function digits(num) { return num.toString().split("").map(Number) }
#

data "js_function" "digits" {
  name   = "digits"
  params = [data.js_function_param.digits_num.id]
  body   = [data.js_return.num_to_string_split_map.content]
}

data "js_function_param" "digits_num" {
  name = "num"
}

data "js_return" "num_to_string_split_map" {
  value = data.js_function_call.num_to_string_split_map.content
}

data "js_function_call" "num_to_string_split_map" {
  caller   = data.js_function_call.num_to_string_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_function_call" "num_to_string_split" {
  caller   = data.js_function_call.num_to_string.content
  function = "split"
  args     = [""]
}

data "js_function_call" "num_to_string" {
  caller   = data.js_function_param.digits_num.id
  function = "toString"
}

#
# function sumDigits(num) { return sum(digits(num)) }
#

data "js_function" "sum_digits" {
  name   = "sumDigits"
  params = [data.js_function_param.sum_digits_num.id]
  body   = [data.js_return.sum_digits.content]
}

data "js_function_param" "sum_digits_num" {
  name = "num"
}

data "js_return" "sum_digits" {
  value = data.js_function_call.sum_digits_num.content
}

data "js_function_call" "sum_digits_num" {
  function = data.js_function.sum.id
  args     = [data.js_function_call.digits_num.content]
}

data "js_function_call" "digits_num" {
  function = data.js_function.digits.id
  args     = [data.js_function_param.sum_digits_num.id]
}

#
# const nab = input.trim().split(" ").map(Number)
#

data "js_const" "nab" {
  name  = "nab"
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
  args     = [" "]
}

data "js_function_call" "input_trim" {
  caller   = data.js_function_param.input.id
  function = "trim"
}

# nab[0]
data "js_index" "n" {
  ref   = data.js_const.nab.id
  value = 0
}

# nab[1]
data "js_index" "a" {
  ref   = data.js_const.nab.id
  value = 1
}

# nab[2]
data "js_index" "b" {
  ref   = data.js_const.nab.id
  value = 2
}

#
# let sum = 0
#

data "js_let" "ans" {
  name  = "ans"
  value = 0
}

#
# for (let i = 0; i <= nab[0]; i++)
#

data "js_for" "sum" {
  init      = data.js_let.for_i.content
  condition = data.js_operation.for_condition.content
  update    = data.js_increment.for_update.content
  body = [
    data.js_const.cmp.content,
    data.js_if.sum.content,
  ]
}

data "js_let" "for_i" {
  name  = "i"
  value = 0
}

data "js_operation" "for_condition" {
  left     = data.js_let.for_i.id
  operator = "<="
  right    = data.js_index.n.content
}

data "js_increment" "for_update" {
  ref = data.js_let.for_i.id
}

#
# const cmp = sumDigits(i)
#

data "js_const" "cmp" {
  name  = "cmp"
  value = data.js_function_call.sum_digits_i.content
}

data "js_function_call" "sum_digits_i" {
  function = data.js_function.sum_digits.id
  args     = [data.js_let.for_i.id]
}

#
# if (a <= cmp && cmp <= b) { sum += i }
#

data "js_if" "sum" {
  condition = data.js_operation.a_le_cmp_and_cmp_le_b.content
  then      = [data.js_operation.sum_plus_i.content]
}

# a <= cmp && cmp <= b
data "js_operation" "a_le_cmp_and_cmp_le_b" {
  left     = data.js_operation.a_le_cmp.content
  operator = "&&"
  right    = data.js_operation.cmp_le_b.content
}

# a <= cmp
data "js_operation" "a_le_cmp" {
  left     = data.js_index.a.content
  operator = "<="
  right    = data.js_const.cmp.id
}

# cmp <= b
data "js_operation" "cmp_le_b" {
  left     = data.js_const.cmp.id
  operator = "<="
  right    = data.js_index.b.content
}

# sum += i
data "js_operation" "sum_plus_i" {
  left     = data.js_let.ans.id
  operator = "+="
  right    = data.js_let.for_i.id
}

#
# console.log(sum)
#

data "js_function_call" "log_sum" {
  caller   = "console"
  function = "log"
  args     = [data.js_let.ans.id]
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_function.add.content,
    data.js_function.sum.content,
    data.js_function.digits.content,
    data.js_function.sum_digits.content,
    data.js_const.nab.content,
    data.js_let.ans.content,
    data.js_for.sum.content,
    data.js_function_call.log_sum.content,
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
