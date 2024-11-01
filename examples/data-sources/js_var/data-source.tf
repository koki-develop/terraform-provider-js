data "js_var" "str" {
  name  = "str"
  value = "hello"
}
# => var str = "hello";

data "js_var" "num" {
  name  = "num"
  value = 10
}
# => var num = 10;

data "js_var" "bool" {
  name  = "bool"
  value = true
}
# => var bool = true;

data "js_var" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}
# => var arr = [1, 2, 3];

data "js_var" "obj" {
  name = "obj"
  value = {
    hoge = "fuga"
    foo  = 3
  }
}
# => var obj = { "hoge": "fuga", "foo": 3 };
