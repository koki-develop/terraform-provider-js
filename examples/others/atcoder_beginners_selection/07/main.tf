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

#
# const dd = lines.slice(1)
#

data "js_const" "dd" {
  name  = "dd"
  value = data.js_function_call.lines_slice.content
}

data "js_function_call" "lines_slice" {
  caller   = data.js_const.lines.id
  function = "slice"
  args     = [1]
}

#
# const set = new Set(dd)
#

data "js_const" "set" {
  name  = "set"
  value = data.js_new.set.content
}

data "js_new" "set" {
  value = data.js_function_call.set.content
}

data "js_function_call" "set" {
  function = "Set"
  args     = [data.js_const.dd.id]
}

#
# console.log(set.size)
#

data "js_function_call" "log_set_size" {
  caller   = "console"
  function = "log"
  args     = [data.js_index.set_size.content]
}

data "js_index" "set_size" {
  ref   = data.js_const.set.id
  value = "size"
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_const.lines.content,
    data.js_const.dd.content,
    data.js_const.set.content,
    data.js_function_call.log_set_size.content,
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
