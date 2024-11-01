data "js_const" "str" {
  name  = "str"
  value = "hello"
}
# => const str = "hello";

data "js_const" "num" {
  name  = "num"
  value = 10
}
# => const num = 10;

data "js_const" "bool" {
  name  = "bool"
  value = true
}
# => const bool = true;

data "js_const" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}
# => const arr = [1, 2, 3];

data "js_const" "obj" {
  name = "obj"
  value = {
    hoge = "fuga"
    foo  = 3
  }
}
# => const obj = { "hoge": "fuga", "foo": 3 };
