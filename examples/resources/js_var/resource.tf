resource "js_var" "str" {
  name  = "str"
  value = "hello"
}
# => var str = "hello";

resource "js_var" "num" {
  name  = "num"
  value = 10
}
# => var num = 10;

resource "js_var" "bool" {
  name  = "bool"
  value = true
}
# => var bool = true;

resource "js_var" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}
# => var arr = [1, 2, 3];

resource "js_var" "obj" {
  name = "obj"
  value = {
    hoge = "fuga"
    foo  = 3
  }
}
# => var obj = { "hoge": "fuga", "foo": 3 };
