data "js_new" "error" {
  constructor = "Error"
  args        = ["something went wrong"]
}

data "js_throw" "error" {
  value = data.js_new.error.expression
}
# => throw new Error("something went wrong")
