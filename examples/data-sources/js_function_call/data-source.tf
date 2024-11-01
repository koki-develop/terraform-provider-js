data "js_function_call" "alert" {
  function = "alert"
  args     = ["hello world"]
}

data "js_function_call" "console_log" {
  caller   = "console"
  function = "log"
  args     = ["hello", "world"]
}
# => console.log("hello", "world")
