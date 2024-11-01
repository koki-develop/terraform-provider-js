run "as" {
  assert {
    condition     = data.js_import.as.content == "@js/raw:import * as name from \"path/to/module\""
    error_message = ""
  }
}

run "default" {
  assert {
    condition     = data.js_import.default.content == "@js/raw:import name from \"path/to/module\""
    error_message = ""
  }
}
