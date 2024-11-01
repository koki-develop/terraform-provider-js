#
# function isEven(num) { return num % 2 === 0 }
#

data "js_function" "is_even" {
  name   = "isEven"
  params = [data.js_function_param.is_even_num.id]
  body   = [data.js_return.num_mod_2_eq_0.content]
}

data "js_function_param" "is_even_num" {
  name = "num"
}

# return num % 2 === 0
data "js_return" "num_mod_2_eq_0" {
  value = data.js_operation.num_mod_2_eq_0.content
}

data "js_operation" "num_mod_2_eq_0" {
  left     = data.js_operation.num_mod_2.content
  operator = "==="
  right    = 0
}

data "js_operation" "num_mod_2" {
  left     = data.js_function_param.is_even_num.id
  operator = "%"
  right    = 2
}

#
# function half(num) { return num / 2 }
#

data "js_function" "half" {
  name   = "half"
  params = [data.js_function_param.half_num.id]
  body   = [data.js_return.num_div_2.content]
}

data "js_function_param" "half_num" {
  name = "num"
}

data "js_return" "num_div_2" {
  value = data.js_operation.num_div_2.content
}

data "js_operation" "num_div_2" {
  left     = data.js_function_param.half_num.id
  operator = "/"
  right    = 2
}

#
# const lines = input.trim().split("\n")
#

data "js_const" "lines" {
  name  = "lines"
  value = data.js_function_call.input_trim_split.content
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

# lines[1]
data "js_index" "lines1" {
  ref   = data.js_const.lines.id
  value = 1
}

#
# const aa = lines[1].split(" ")
#

data "js_const" "aa" {
  name  = "aa"
  value = data.js_function_call.lines1_split.content
}

data "js_function_call" "lines1_split" {
  caller   = data.js_index.lines1.content
  function = "split"
  args     = [" "]
}

#
# let nums = aa.map(Number)
#

data "js_let" "nums" {
  name  = "nums"
  value = data.js_function_call.aa_map_number.content
}

data "js_function_call" "aa_map_number" {
  caller   = data.js_const.aa.id
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_raw" "number" {
  value = "Number"
}

#
# let count
#

data "js_let" "count" {
  name  = "count"
  value = 0
}

#
# while (nums.every(isEven)) {
#   nums = nums.map(half);
#   count++;
# }
#

data "js_while" "count" {
  condition = data.js_function_call.nums_every.content
  body = [
    data.js_operation.half_nums.content,
    data.js_increment.count.content,
  ]
}

data "js_function_call" "nums_every" {
  caller   = data.js_let.nums.id
  function = "every"
  args     = [data.js_function.is_even.id]
}

data "js_operation" "half_nums" {
  left     = data.js_let.nums.id
  operator = "="
  right    = data.js_function_call.half_nums.content
}

data "js_function_call" "half_nums" {
  caller   = data.js_let.nums.id
  function = "map"
  args     = [data.js_function.half.id]
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
    data.js_function.is_even.content,
    data.js_function.half.content,
    data.js_const.lines.content,
    data.js_const.aa.content,
    data.js_let.nums.content,
    data.js_let.count.content,
    data.js_while.count.content,
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
