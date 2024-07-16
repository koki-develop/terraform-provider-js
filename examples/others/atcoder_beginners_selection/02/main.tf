#
# input.split
#

resource "js_function_call" "input_split" {
  caller   = js_function_param.input.id
  function = "split"
  args     = [""]
}

#
# input.filter
#

resource "js_function_call" "input_filter" {
  caller   = js_function_call.input_split.content
  function = "filter"
  args     = [js_function.filter.content]
}

resource "js_function" "filter" {
  params = [js_function_param.filter_s.id]
  body   = [js_return.s_eq_1.content]
}

resource "js_function_param" "filter_s" {
  name = "s"
}

#
# return s === "1"
#

resource "js_return" "s_eq_1" {
  value = js_operation.s_eq_1.content
}

resource "js_operation" "s_eq_1" {
  left     = js_function_param.filter_s.id
  right    = "1"
  operator = "==="
}

#
# length
#

data "js_index" "length" {
  ref   = js_function_call.input_filter.content
  value = "length"
}

#
# print
#

resource "js_function_call" "log_count" {
  caller   = "console"
  function = "log"
  args     = [data.js_index.length.id]
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body   = [js_function_call.log_count.content]
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
