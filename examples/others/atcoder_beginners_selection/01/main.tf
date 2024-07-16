#
# const ab
#

resource "js_const" "ab" {
  name  = "ab"
  value = js_function_call.input_map_number.content
}

resource "js_function_call" "input_split" {
  caller   = js_function_param.input.id
  function = "split"
  args     = [" "]
}

resource "js_function_call" "input_map_number" {
  caller   = js_function_call.input_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_raw" "number" {
  value = "Number"
}

data "js_index" "a" {
  ref   = js_const.ab.id
  value = 0
}

data "js_index" "b" {
  ref   = js_const.ab.id
  value = 1
}

#
# a * b % 2 === 0
#

resource "js_operation" "a_times_b" {
  left     = data.js_index.a.id
  right    = data.js_index.b.id
  operator = "*"
}

resource "js_operation" "a_times_b_mod_2" {
  left     = js_operation.a_times_b.content
  right    = 2
  operator = "%"
}

resource "js_operation" "a_times_b_mod_2_eq_0" {
  left     = js_operation.a_times_b_mod_2.content
  right    = 0
  operator = "==="
}

#
# even or odd
#

resource "js_if" "even_or_odd" {
  condition = js_operation.a_times_b_mod_2_eq_0.content
  then      = [js_function_call.even.content]
  else      = [js_function_call.odd.content]
}

resource "js_function_call" "even" {
  caller   = "console"
  function = "log"
  args     = ["Even"]
}

resource "js_function_call" "odd" {
  caller   = "console"
  function = "log"
  args     = ["Odd"]
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_const.ab.content,
    js_if.even_or_odd.content,
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
