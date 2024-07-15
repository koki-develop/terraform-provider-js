run "str" {
  assert {
    condition     = js_let.str.content == "@js/raw:let str=\"hello\""
    error_message = ""
  }
}

run "num" {
  assert {
    condition     = js_let.num.content == "@js/raw:let num=10"
    error_message = ""
  }
}

run "bool" {
  assert {
    condition     = js_let.bool.content == "@js/raw:let bool=true"
    error_message = ""
  }
}

run "arr" {
  assert {
    condition     = js_let.arr.content == "@js/raw:let arr=[1,2,3]"
    error_message = ""
  }
}

run "obj" {
  assert {
    condition     = js_let.obj.content == "@js/raw:let obj={\"foo\":3,\"hoge\":\"fuga\"}"
    error_message = ""
  }
}
