resource "js_function_call" "error" {
  function = "Error"
  args     = ["something went wrong"]
}

resource "js_new" "error" {
  value = js_function_call.error.content
}

resource "js_throw" "error" {
  value = js_new.error.content
}
# => throw new Error("something went wrong")
