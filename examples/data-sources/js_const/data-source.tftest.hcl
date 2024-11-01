run "str" {
  assert {
    condition     = data.js_const.str.content == "@js/raw:const str=\"hello\""
    error_message = ""
  }
}

run "num" {
  assert {
    condition     = data.js_const.num.content == "@js/raw:const num=10"
    error_message = ""
  }
}

run "bool" {
  assert {
    condition     = data.js_const.bool.content == "@js/raw:const bool=true"
    error_message = ""
  }
}

run "arr" {
  assert {
    condition     = data.js_const.arr.content == "@js/raw:const arr=[1,2,3]"
    error_message = ""
  }
}

run "obj" {
  assert {
    condition     = data.js_const.obj.content == "@js/raw:const obj={\"foo\":3,\"hoge\":\"fuga\"}"
    error_message = ""
  }
}
