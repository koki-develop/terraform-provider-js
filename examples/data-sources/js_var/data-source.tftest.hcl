run "str" {
  assert {
    condition     = data.js_var.str.statement == "@js/raw:var str=\"hello\""
    error_message = ""
  }
}

run "num" {
  assert {
    condition     = data.js_var.num.statement == "@js/raw:var num=10"
    error_message = ""
  }
}

run "bool" {
  assert {
    condition     = data.js_var.bool.statement == "@js/raw:var bool=true"
    error_message = ""
  }
}

run "arr" {
  assert {
    condition     = data.js_var.arr.statement == "@js/raw:var arr=[1,2,3]"
    error_message = ""
  }
}

run "obj" {
  assert {
    condition     = data.js_var.obj.statement == "@js/raw:var obj={\"foo\":3,\"hoge\":\"fuga\"}"
    error_message = ""
  }
}
