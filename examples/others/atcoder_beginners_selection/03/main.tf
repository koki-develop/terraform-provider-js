#
# const lines = input.split("\n")
#

resource "js_const" "lines" {
  name  = "lines"
  value = js_function_call.input_split.content
}

resource "js_function_call" "input_split" {
  caller   = js_function_param.input.id
  function = "split"
  args     = ["\n"]
}

# lines[1]
data "js_index" "lines1" {
  ref   = js_const.lines.id
  value = 1
}

#
# const aa = lines[1].split(" ")
#

resource "js_const" "aa" {
  name  = "aa"
  value = js_function_call.lines1_split.content
}

resource "js_function_call" "lines1_split" {
  caller   = data.js_index.lines1.id
  function = "split"
  args     = [" "]
}

#
# let nums = aa.map(Number)
#

resource "js_let" "nums" {
  name  = "nums"
  value = js_function_call.aa_map_number.content
}

resource "js_function_call" "aa_map_number" {
  caller   = js_const.aa.id
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_raw" "number" {
  value = "Number"
}

#
# let count
#

resource "js_let" "count" {
  name  = "count"
  value = 0
}

#
# function isEven(num) { return num % 2 === 0 }
#

resource "js_function" "is_even" {
  name   = "isEven"
  params = [js_function_param.num.id]
  body   = [js_return.num_mod_2_eq_0.content]
}

resource "js_function_param" "num" {
  name = "num"
}

# return num % 2 === 0
resource "js_return" "num_mod_2_eq_0" {
  value = js_operation.num_mod_2_eq_0.content
}

resource "js_operation" "num_mod_2_eq_0" {
  left     = js_operation.num_mod_2.content
  operator = "==="
  right    = 0
}

resource "js_operation" "num_mod_2" {
  left     = js_function_param.num.id
  operator = "%"
  right    = 2
}

#
# function half(num) { return num / 2 }
#

resource "js_function" "half" {
  name   = "half"
  params = [js_function_param.num.id]
  body   = [js_return.num_div_2.content]
}

resource "js_return" "num_div_2" {
  value = js_operation.num_div_2.content
}

resource "js_operation" "num_div_2" {
  left     = js_function_param.num.id
  operator = "/"
  right    = 2
}

#
# while (nums.every(isEven)) {
#   nums = nums.map(half);
#   count++;
# }
#

resource "js_while" "count" {
  condition = js_function_call.nums_every.content
  body = [
    js_operation.half_nums.content,
    js_increment.count.content,
  ]
}

resource "js_function_call" "nums_every" {
  caller   = js_let.nums.id
  function = "every"
  args     = [js_function.is_even.id]
}

resource "js_operation" "half_nums" {
  left     = js_let.nums.id
  operator = "="
  right    = js_function_call.half_nums.content
}

resource "js_function_call" "half_nums" {
  caller   = js_let.nums.id
  function = "map"
  args     = [js_function.half.id]
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
    js_const.lines.content,
    js_const.aa.content,
    js_let.nums.content,
    js_let.count.content,
    js_function.is_even.content,
    js_function.half.content,
    js_while.count.content,
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
