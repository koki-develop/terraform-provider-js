resource "js_function_call" "console_log" {
  caller   = "console"
  function = "log"
  args     = ["hello", "world"]
}
# => console.log("hello", "world")
