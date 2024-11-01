data "js_let" "str" {
  name  = "str"
  value = "hello"
}
# => let str = "hello";

data "js_let" "num" {
  name  = "num"
  value = 10
}
# => let num = 10;

data "js_let" "bool" {
  name  = "bool"
  value = true
}
# => let bool = true;

data "js_let" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}
# => let arr = [1, 2, 3];

data "js_let" "obj" {
  name = "obj"
  value = {
    hoge = "fuga"
    foo  = 3
  }
}
# => let obj = { "hoge": "fuga", "foo": 3 };
