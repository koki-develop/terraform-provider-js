#
# const lines = input.trim().split("\n")
#

resource "js_const" "lines" {
  name  = "lines"
  value = js_function_call.input_trim_split.content
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

#
# const dd = lines.slice(1)
#

resource "js_const" "dd" {
  name  = "dd"
  value = js_function_call.lines_slice.content
}

resource "js_function_call" "lines_slice" {
  caller   = js_const.lines.id
  function = "slice"
  args     = [1]
}

#
# const set = new Set(dd)
#

resource "js_const" "set" {
  name  = "set"
  value = js_new.set.content
}

resource "js_new" "set" {
  value = js_function_call.set.content
}

resource "js_function_call" "set" {
  function = "Set"
  args     = [js_const.dd.id]
}

#
# console.log(set.size)
#

resource "js_function_call" "log_set_size" {
  caller   = "console"
  function = "log"
  args     = [data.js_index.set_size.id]
}

data "js_index" "set_size" {
  ref   = js_const.set.id
  value = "size"
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_const.lines.content,
    js_const.dd.content,
    js_const.set.content,
    js_function_call.log_set_size.content,
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
