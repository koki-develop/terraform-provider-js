data "js_const" "message" {
  name  = "message"
  value = "hello world"
}

data "js_export" "message" {
  value = data.js_const.message.id
}
# => export message
