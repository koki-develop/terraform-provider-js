#
# const ab = input.trim().split(" ").map(Number);
#

resource "js_const" "ab" {
  name  = "ab"
  value = js_function_call.input_trim_split_map_number.content
}

resource "js_function_call" "input_trim_split_map_number" {
  caller   = js_function_call.input_trim_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_raw" "number" {
  value = "Number"
}

resource "js_function_call" "input_trim_split" {
  caller   = js_function_call.input_trim.content
  function = "split"
  args     = [" "]
}

resource "js_function_call" "input_trim" {
  caller   = js_function_param.input.id
  function = "trim"
}

# ab[0]
data "js_index" "a" {
  ref   = js_const.ab.id
  value = 0
}

# ab[1]
data "js_index" "b" {
  ref   = js_const.ab.id
  value = 1
}

#
# if (a * b % 2 === 0)
#

resource "js_if" "even_or_odd" {
  condition = js_operation.a_times_b_mod_2_eq_0.content
  then      = [js_function_call.log_even.content]
  else      = [js_function_call.log_odd.content]
}

# a * b % 2 === 0
resource "js_operation" "a_times_b_mod_2_eq_0" {
  left     = js_operation.a_times_b_mod_2.content
  right    = 0
  operator = "==="
}

# a * b % 2
resource "js_operation" "a_times_b_mod_2" {
  left     = js_operation.a_times_b.content
  right    = 2
  operator = "%"
}

# a * b
resource "js_operation" "a_times_b" {
  left     = data.js_index.a.id
  right    = data.js_index.b.id
  operator = "*"
}

# console.log("Even")
resource "js_function_call" "log_even" {
  caller   = "console"
  function = "log"
  args     = ["Even"]
}

# console.log("Odd")
resource "js_function_call" "log_odd" {
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
