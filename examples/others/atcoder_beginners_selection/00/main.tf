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

#
# const a = Number(lines[0])
#

resource "js_const" "a" {
  name  = "a"
  value = js_function_call.number_a.content
}

resource "js_function_call" "number_a" {
  function = "Number"
  args     = [data.js_index.lines0.id]
}

data "js_index" "lines0" {
  ref   = js_const.lines.id
  value = 0
}

#
# const bc = lines[1].split(" ")
#

resource "js_const" "bc" {
  name  = "bc"
  value = js_function_call.lines1_split.content
}

resource "js_function_call" "lines1_split" {
  caller   = data.js_index.lines1.id
  function = "split"
  args     = [" "]
}

data "js_index" "lines1" {
  ref   = js_const.lines.id
  value = 1
}

#
# const b = Number(bc[0])
#

resource "js_const" "b" {
  name  = "b"
  value = js_function_call.number_bc0.content
}

resource "js_function_call" "number_bc0" {
  function = "Number"
  args     = [data.js_index.bc0.id]
}

data "js_index" "bc0" {
  ref   = js_const.bc.id
  value = 0
}

#
# const c = Number(bc[1])
#

resource "js_const" "c" {
  name  = "c"
  value = js_function_call.number_bc1.content
}

resource "js_function_call" "number_bc1" {
  function = "Number"
  args     = [data.js_index.bc1.id]
}

data "js_index" "bc1" {
  ref   = js_const.bc.id
  value = 1
}

#
# const s = lines[2]
#

resource "js_const" "s" {
  name  = "s"
  value = data.js_index.lines2.id
}

data "js_index" "lines2" {
  ref   = js_const.lines.id
  value = 2
}

#
# console.log(a + b + c, s)
#

resource "js_function_call" "log" {
  caller   = "console"
  function = "log"
  args     = [js_operation.a_plus_b_plus_c.content, js_const.s.id]
}

resource "js_operation" "a_plus_b_plus_c" {
  left     = js_operation.a_plus_b.content
  operator = "+"
  right    = js_const.c.id
}

resource "js_operation" "a_plus_b" {
  left     = js_const.a.id
  operator = "+"
  right    = js_const.b.id
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_const.lines.content,
    js_const.a.content,
    js_const.bc.content,
    js_const.b.content,
    js_const.c.content,
    js_const.s.content,
    js_function_call.log.content,
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
