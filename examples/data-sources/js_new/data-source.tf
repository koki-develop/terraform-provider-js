data "js_new" "error" {
  constructor = "Error"
  args        = ["something went wrong"]
}
# => new Error("something went wrong")
