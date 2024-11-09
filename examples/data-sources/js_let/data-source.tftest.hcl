run "str" {
  assert {
    condition     = data.js_let.str.statement == "@js/raw:let str=\"hello\""
    error_message = ""
  }
}

run "num" {
  assert {
    condition     = data.js_let.num.statement == "@js/raw:let num=10"
    error_message = ""
  }
}

run "bool" {
  assert {
    condition     = data.js_let.bool.statement == "@js/raw:let bool=true"
    error_message = ""
  }
}

run "arr" {
  assert {
    condition     = data.js_let.arr.statement == "@js/raw:let arr=[1,2,3]"
    error_message = ""
  }
}

run "obj" {
  assert {
    condition     = data.js_let.obj.statement == "@js/raw:let obj={\"foo\":3,\"hoge\":\"fuga\"}"
    error_message = ""
  }
}
