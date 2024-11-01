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
# const a = Number(lines[0])
#

data "js_const" "a" {
  name  = "a"
  value = data.js_function_call.number_a.content
}

data "js_function_call" "number_a" {
  function = "Number"
  args     = [data.js_index.lines0.content]
}

data "js_index" "lines0" {
  ref   = data.js_const.lines.id
  value = 0
}

#
# const bc = lines[1].split(" ")
#

data "js_const" "bc" {
  name  = "bc"
  value = data.js_function_call.lines1_split.content
}

data "js_function_call" "lines1_split" {
  caller   = data.js_index.lines1.content
  function = "split"
  args     = [" "]
}

data "js_index" "lines1" {
  ref   = data.js_const.lines.id
  value = 1
}

#
# const b = Number(bc[0])
#

data "js_const" "b" {
  name  = "b"
  value = data.js_function_call.number_bc0.content
}

data "js_function_call" "number_bc0" {
  function = "Number"
  args     = [data.js_index.bc0.content]
}

data "js_index" "bc0" {
  ref   = data.js_const.bc.id
  value = 0
}

#
# const c = Number(bc[1])
#

data "js_const" "c" {
  name  = "c"
  value = data.js_function_call.number_bc1.content
}

data "js_function_call" "number_bc1" {
  function = "Number"
  args     = [data.js_index.bc1.content]
}

data "js_index" "bc1" {
  ref   = data.js_const.bc.id
  value = 1
}

#
# const s = lines[2]
#

data "js_const" "s" {
  name  = "s"
  value = data.js_index.lines2.content
}

data "js_index" "lines2" {
  ref   = data.js_const.lines.id
  value = 2
}

#
# console.log(a + b + c, s)
#

data "js_function_call" "log" {
  caller   = "console"
  function = "log"
  args     = [data.js_operation.a_plus_b_plus_c.content, data.js_const.s.id]
}

data "js_operation" "a_plus_b_plus_c" {
  left     = data.js_operation.a_plus_b.content
  operator = "+"
  right    = data.js_const.c.id
}

data "js_operation" "a_plus_b" {
  left     = data.js_const.a.id
  operator = "+"
  right    = data.js_const.b.id
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_const.lines.content,
    data.js_const.a.content,
    data.js_const.bc.content,
    data.js_const.b.content,
    data.js_const.c.content,
    data.js_const.s.content,
    data.js_function_call.log.content,
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
