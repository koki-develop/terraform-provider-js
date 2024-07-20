#
# const s = input.trim().replaceAll("eraser", "").replaceAll("erase", "").replaceAll("dreamer", "").replaceAll("dream", "")
#

resource "js_function_call" "input_trim" {
  caller   = js_function_param.input.id
  function = "trim"
}

resource "js_function_call" "remove_eraser" {
  caller   = js_function_call.input_trim.content
  function = "replaceAll"
  args     = ["eraser", ""]
}

resource "js_function_call" "remove_erase" {
  caller   = js_function_call.remove_eraser.content
  function = "replaceAll"
  args     = ["erase", ""]
}

resource "js_function_call" "remove_dreamer" {
  caller   = js_function_call.remove_erase.content
  function = "replaceAll"
  args     = ["dreamer", ""]
}

resource "js_function_call" "remove_dream" {
  caller   = js_function_call.remove_dreamer.content
  function = "replaceAll"
  args     = ["dream", ""]
}

resource "js_const" "s" {
  name  = "s"
  value = js_function_call.remove_dream.content
}

#
# if (s === "") { console.log("YES") } else { console.log("NO") }
#

resource "js_if" "log_answer" {
  condition = js_operation.s_eq_empty.content
  then      = [js_function_call.log_yes.content]
  else      = [js_function_call.log_no.content]
}

resource "js_operation" "s_eq_empty" {
  left     = js_const.s.id
  operator = "==="
  right    = ""
}

resource "js_function_call" "log_yes" {
  caller   = "console"
  function = "log"
  args     = ["YES"]
}

resource "js_function_call" "log_no" {
  caller   = "console"
  function = "log"
  args     = ["NO"]
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_const.s.content,
    js_if.log_answer.content,
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
