run "i" {
  assert {
    condition     = data.js_decrement.i.statement == "@js/raw:i--"
    error_message = ""
  }
}

run "i_postfix" {
  assert {
    condition     = data.js_decrement.i_postfix.statement == "@js/raw:i--"
    error_message = ""
  }
}

run "i_prefix" {
  assert {
    condition     = data.js_decrement.i_prefix.statement == "@js/raw:--i"
    error_message = ""
  }
}
