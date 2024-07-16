run "str" {
  assert {
    condition     = js_var.str.content == "@js/raw:var str=\"hello\""
    error_message = ""
  }
}

run "num" {
  assert {
    condition     = js_var.num.content == "@js/raw:var num=10"
    error_message = ""
  }
}

run "bool" {
  assert {
    condition     = js_var.bool.content == "@js/raw:var bool=true"
    error_message = ""
  }
}

run "arr" {
  assert {
    condition     = js_var.arr.content == "@js/raw:var arr=[1,2,3]"
    error_message = ""
  }
}

run "obj" {
  assert {
    condition     = js_var.obj.content == "@js/raw:var obj={\"foo\":3,\"hoge\":\"fuga\"}"
    error_message = ""
  }
}
