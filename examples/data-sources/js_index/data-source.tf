data "js_const" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}

data "js_index" "arr_1" {
  ref   = data.js_const.arr.id
  value = 1
}
# => arr[1]

data "js_const" "obj" {
  name = "obj"
  value = {
    foo = "bar"
  }
}

data "js_index" "obj_foo" {
  ref   = data.js_const.obj.id
  value = "foo"
}
# => obj["foo"]
