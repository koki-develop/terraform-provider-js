#
# function isOne(s) { return s === "1" }
#

data "js_function" "is_one" {
  name   = "isOne"
  params = [data.js_function_param.is_one_s.id]
  body   = [data.js_return.s_eq_1.content]
}

data "js_function_param" "is_one_s" {
  name = "s"
}

data "js_return" "s_eq_1" {
  value = data.js_operation.s_eq_1.content
}

data "js_operation" "s_eq_1" {
  left     = data.js_function_param.is_one_s.id
  right    = "1"
  operator = "==="
}

#
# const ss = input.trim().split("")
#

data "js_const" "ss" {
  name  = "ss"
  value = data.js_function_call.input_trim_split.content
}

data "js_function_call" "input_trim_split" {
  caller   = data.js_function_call.input_trim.content
  function = "split"
  args     = [""]
}

data "js_function_call" "input_trim" {
  caller   = data.js_function_param.input.id
  function = "trim"
}

#
# const count = ss.filter(filterOne).length
#

data "js_const" "count" {
  name  = "count"
  value = data.js_index.length.content
}

data "js_index" "length" {
  ref   = data.js_function_call.input_filter.content
  value = "length"
}

data "js_function_call" "input_filter" {
  caller   = data.js_const.ss.id
  function = "filter"
  args     = [data.js_function.is_one.id]
}

#
# console.log(count)
#

data "js_function_call" "log_count" {
  caller   = "console"
  function = "log"
  args     = [data.js_const.count.id]
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_function.is_one.content,
    data.js_const.ss.content,
    data.js_const.count.content,
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
