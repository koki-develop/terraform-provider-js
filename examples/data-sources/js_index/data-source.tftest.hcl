run "example_index" {
  assert {
    condition     = data.js_index.example_index.content == "@js/raw:example[1]"
    error_message = ""
  }
}

run "example_property" {
  assert {
    condition     = data.js_index.example_property.content == "@js/raw:example[\"name\"]"
    error_message = ""
  }
}
