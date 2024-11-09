run "arr_1" {
  assert {
    condition     = data.js_index.arr_1.statement == "@js/raw:arr[1]"
    error_message = ""
  }
}

run "obj_foo" {
  assert {
    condition     = data.js_index.obj_foo.statement == "@js/raw:obj[\"foo\"]"
    error_message = ""
  }
}
