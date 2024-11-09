run "str" {
  assert {
    condition     = data.js_const.str.statement == "@js/raw:const str=\"hello\""
    error_message = ""
  }
}

run "num" {
  assert {
    condition     = data.js_const.num.statement == "@js/raw:const num=10"
    error_message = ""
  }
}

run "bool" {
  assert {
    condition     = data.js_const.bool.statement == "@js/raw:const bool=true"
    error_message = ""
  }
}

run "arr" {
  assert {
    condition     = data.js_const.arr.statement == "@js/raw:const arr=[1,2,3]"
    error_message = ""
  }
}

run "obj" {
  assert {
    condition     = data.js_const.obj.statement == "@js/raw:const obj={\"foo\":3,\"hoge\":\"fuga\"}"
    error_message = ""
  }
}
