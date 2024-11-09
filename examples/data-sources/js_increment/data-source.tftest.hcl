run "i" {
  assert {
    condition     = data.js_increment.i.statement == "@js/raw:i++"
    error_message = ""
  }
}

run "i_postfix" {
  assert {
    condition     = data.js_increment.i_postfix.statement == "@js/raw:i++"
    error_message = ""
  }
}

run "i_prefix" {
  assert {
    condition     = data.js_increment.i_prefix.statement == "@js/raw:++i"
    error_message = ""
  }
}
