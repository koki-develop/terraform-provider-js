#
# input.split("\n")[1].split(" ").map(Number)
#

resource "js_function_call" "input_split" {
  caller   = js_function_param.input.id
  function = "split"
  args     = ["\n"]
}

data "js_index" "input_split1" {
  ref   = js_function_call.input_split.content
  value = 1
}

resource "js_function_call" "input_split1_split" {
  caller   = data.js_index.input_split1.id
  function = "split"
  args     = [" "]
}

resource "js_function_call" "input_split1_split_map_number" {
  caller   = js_function_call.input_split1_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_raw" "number" {
  value = "Number"
}

#
# const nums
#

resource "js_let" "nums" {
  name  = "nums"
  value = js_function_call.input_split1_split_map_number.content
}

#
# let count
#

resource "js_let" "count" {
  name  = "count"
  value = 0
}

#
# num % 2 === 0
#

resource "js_operation" "num_mod_2" {
  left     = js_function_param.num.id
  operator = "%"
  right    = 2
}

resource "js_operation" "num_mod_2_eq_0" {
  left     = js_operation.num_mod_2.content
  operator = "==="
  right    = 0
}

#
# every even
#

resource "js_function" "every_even" {
  params = [js_function_param.num.id]
  body   = [js_return.num_mod_2_eq_0.content]
}

resource "js_function_param" "num" {
  name = "num"
}

resource "js_return" "num_mod_2_eq_0" {
  value = js_operation.num_mod_2_eq_0.content
}

resource "js_function_call" "nums_every" {
  caller   = js_let.nums.id
  function = "every"
  args     = [js_function.every_even.content]
}

#
# nums half
#

resource "js_function_call" "nums_map" {
  caller   = js_let.nums.id
  function = "map"
  args = [
    js_function.nums_half.content,
  ]
}

resource "js_function" "nums_half" {
  params = [js_function_param.num.id]
  body   = [js_return.num_half.content]
}

resource "js_operation" "nums_half" {
  left     = js_function_param.num.id
  operator = "/"
  right    = 2
}

resource "js_return" "num_half" {
  value = js_operation.nums_half.content
}

resource "js_operation" "nums_half_assign" {
  left     = js_let.nums.id
  operator = "="
  right    = js_function_call.nums_map.content
}

#
# increment
#

resource "js_operation" "count_increment" {
  left     = js_let.count.id
  operator = "+="
  right    = 1
}

#
# count
#

resource "js_while" "count" {
  condition = js_function_call.nums_every.content
  body = [
    js_operation.nums_half_assign.content,
    js_operation.count_increment.content,
  ]
}

#
# print
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
    js_let.nums.content,
    js_let.count.content,
    js_while.count.content,
    js_function_call.log_count.content,
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
