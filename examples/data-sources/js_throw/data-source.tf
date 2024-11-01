data "js_function_call" "error" {
  function = "Error"
  args     = ["something went wrong"]
}

data "js_new" "error" {
  value = data.js_function_call.error.content
}

data "js_throw" "error" {
  value = data.js_new.error.content
}
# => throw new Error("something went wrong")
