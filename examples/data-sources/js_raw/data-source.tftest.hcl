run "example" {
  assert {
    condition     = data.js_raw.example.content == "@js/raw:console.log('Hello, World!');"
    error_message = ""
  }
}
