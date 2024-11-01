#
# const s = input.trim().replaceAll("eraser", "").replaceAll("erase", "").replaceAll("dreamer", "").replaceAll("dream", "")
#

data "js_function_call" "input_trim" {
  caller   = data.js_function_param.input.id
  function = "trim"
}

data "js_function_call" "remove_eraser" {
  caller   = data.js_function_call.input_trim.content
  function = "replaceAll"
  args     = ["eraser", ""]
}

data "js_function_call" "remove_erase" {
  caller   = data.js_function_call.remove_eraser.content
  function = "replaceAll"
  args     = ["erase", ""]
}

data "js_function_call" "remove_dreamer" {
  caller   = data.js_function_call.remove_erase.content
  function = "replaceAll"
  args     = ["dreamer", ""]
}

data "js_function_call" "remove_dream" {
  caller   = data.js_function_call.remove_dreamer.content
  function = "replaceAll"
  args     = ["dream", ""]
}

data "js_const" "s" {
  name  = "s"
  value = data.js_function_call.remove_dream.content
}

#
# if (s === "") { console.log("YES") } else { console.log("NO") }
#

data "js_if" "log_answer" {
  condition = data.js_operation.s_eq_empty.content
  then      = [data.js_function_call.log_yes.content]
  else      = [data.js_function_call.log_no.content]
}

data "js_operation" "s_eq_empty" {
  left     = data.js_const.s.id
  operator = "==="
  right    = ""
}

data "js_function_call" "log_yes" {
  caller   = "console"
  function = "log"
  args     = ["YES"]
}

data "js_function_call" "log_no" {
  caller   = "console"
  function = "log"
  args     = ["NO"]
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_const.s.content,
    data.js_if.log_answer.content,
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
