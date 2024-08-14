resource "js_const" "message" {
  name  = "message"
  value = "hello world"
}

resource "js_export" "message" {
  value = js_const.message.id
}
# => export message
