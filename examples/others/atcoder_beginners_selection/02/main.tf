#
# function isOne(s) { return s === "1" }
#

resource "js_function" "is_one" {
  name   = "isOne"
  params = [js_function_param.filter_s.id]
  body   = [js_return.s_eq_1.content]
}

resource "js_function_param" "filter_s" {
  name = "s"
}

resource "js_return" "s_eq_1" {
  value = js_operation.s_eq_1.content
}

resource "js_operation" "s_eq_1" {
  left     = js_function_param.filter_s.id
  right    = "1"
  operator = "==="
}

#
# const ss = input.split("")
#

resource "js_const" "ss" {
  name  = "ss"
  value = js_function_call.input_split.content
}

resource "js_function_call" "input_split" {
  caller   = js_function_param.input.id
  function = "split"
  args     = [""]
}

#
# const count = ss.filter(filterOne).length
#

resource "js_const" "count" {
  name  = "count"
  value = data.js_index.length.id
}

data "js_index" "length" {
  ref   = js_function_call.input_filter.content
  value = "length"
}

resource "js_function_call" "input_filter" {
  caller   = js_function_call.input_split.content
  function = "filter"
  args     = [js_function.is_one.id]
}

#
# console.log(count)
#

resource "js_function_call" "log_count" {
  caller   = "console"
  function = "log"
  args     = [js_const.count.id]
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_function.is_one.content,
    js_const.ss.content,
    js_const.count.content,
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
