resource "js_let" "num" {
  name  = "i"
  value = 0
}

resource "js_increment" "i" {
  ref = js_let.num.id
}
# => i++

resource "js_increment" "i_postfix" {
  ref  = js_let.num.id
  type = "postfix"
}
# => i++

resource "js_increment" "i_prefix" {
  ref  = js_let.num.id
  type = "prefix"
}
# => ++i
