data "js_function_call" "error" {
  function = "Error"
  args     = ["something went wrong"]
}

data "js_new" "error" {
  value = data.js_function_call.error.content
}
# => new Error("something went wrong")
