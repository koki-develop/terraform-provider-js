data "js_let" "num" {
  name  = "i"
  value = 0
}

data "js_decrement" "i" {
  ref = data.js_let.num.id
}
# => i--

data "js_decrement" "i_postfix" {
  ref  = data.js_let.num.id
  type = "postfix"
}
# => i--

data "js_decrement" "i_prefix" {
  ref  = data.js_let.num.id
  type = "prefix"
}
# => --i