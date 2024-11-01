#
# const ab = input.trim().split(" ").map(Number);
#

data "js_const" "ab" {
  name  = "ab"
  value = data.js_function_call.input_trim_split_map_number.content
}

data "js_function_call" "input_trim_split_map_number" {
  caller   = data.js_function_call.input_trim_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_raw" "number" {
  value = "Number"
}

data "js_function_call" "input_trim_split" {
  caller   = data.js_function_call.input_trim.content
  function = "split"
  args     = [" "]
}

data "js_function_call" "input_trim" {
  caller   = data.js_function_param.input.id
  function = "trim"
}

# ab[0]
data "js_index" "a" {
  ref   = data.js_const.ab.id
  value = 0
}

# ab[1]
data "js_index" "b" {
  ref   = data.js_const.ab.id
  value = 1
}

#
# if (a * b % 2 === 0)
#

data "js_if" "even_or_odd" {
  condition = data.js_operation.a_times_b_mod_2_eq_0.content
  then      = [data.js_function_call.log_even.content]
  else      = [data.js_function_call.log_odd.content]
}

# a * b % 2 === 0
data "js_operation" "a_times_b_mod_2_eq_0" {
  left     = data.js_operation.a_times_b_mod_2.content
  right    = 0
  operator = "==="
}

# a * b % 2
data "js_operation" "a_times_b_mod_2" {
  left     = data.js_operation.a_times_b.content
  right    = 2
  operator = "%"
}

# a * b
data "js_operation" "a_times_b" {
  left     = data.js_index.a.content
  right    = data.js_index.b.content
  operator = "*"
}

# console.log("Even")
data "js_function_call" "log_even" {
  caller   = "console"
  function = "log"
  args     = ["Even"]
}

# console.log("Odd")
data "js_function_call" "log_odd" {
  caller   = "console"
  function = "log"
  args     = ["Odd"]
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_const.ab.content,
    data.js_if.even_or_odd.content,
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
