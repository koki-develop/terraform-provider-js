resource "js_const" "str" {
  name  = "str"
  value = "hello"
}

resource "js_const" "num" {
  name  = "num"
  value = 10
}

resource "js_const" "bool" {
  name  = "bool"
  value = true
}

resource "js_const" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}

resource "js_const" "obj" {
  name = "obj"
  value = {
    hoge = "fuga"
    foo  = 3
  }
}

resource "js_program" "main" {
  contents = [
    js_const.str.content,
    js_const.num.content,
    js_const.bool.content,
    js_const.arr.content,
    js_const.obj.content,
  ]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = js_program.main.content
}
