run "i" {
  assert {
    condition     = js_decrement.i.content == "@js/raw:i--"
    error_message = ""
  }
}

run "i_postfix" {
  assert {
    condition     = js_decrement.i_postfix.content == "@js/raw:i--"
    error_message = ""
  }
}

run "i_prefix" {
  assert {
    condition     = js_decrement.i_prefix.content == "@js/raw:--i"
    error_message = ""
  }
}