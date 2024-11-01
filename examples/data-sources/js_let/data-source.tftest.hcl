run "str" {
  assert {
    condition     = data.js_let.str.content == "@js/raw:let str=\"hello\""
    error_message = ""
  }
}

run "num" {
  assert {
    condition     = data.js_let.num.content == "@js/raw:let num=10"
    error_message = ""
  }
}

run "bool" {
  assert {
    condition     = data.js_let.bool.content == "@js/raw:let bool=true"
    error_message = ""
  }
}

run "arr" {
  assert {
    condition     = data.js_let.arr.content == "@js/raw:let arr=[1,2,3]"
    error_message = ""
  }
}

run "obj" {
  assert {
    condition     = data.js_let.obj.content == "@js/raw:let obj={\"foo\":3,\"hoge\":\"fuga\"}"
    error_message = ""
  }
}
