resource "js_function_call" "error" {
  function = "Error"
  args     = ["something went wrong"]
}

resource "js_new" "error" {
  value = js_function_call.error.content
}
# => new Error("something went wrong")
