run "i" {
  assert {
    condition     = data.js_increment.i.content == "@js/raw:i++"
    error_message = ""
  }
}

run "i_postfix" {
  assert {
    condition     = data.js_increment.i_postfix.content == "@js/raw:i++"
    error_message = ""
  }
}

run "i_prefix" {
  assert {
    condition     = data.js_increment.i_prefix.content == "@js/raw:++i"
    error_message = ""
  }
}
