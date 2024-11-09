run "fetch_statement" {
  assert {
    condition     = data.js_await.fetch.statement == "@js/raw:await fetch(\"https://example.com\")"
    error_message = ""
  }
}

run "fetch_expression" {
  assert {
    condition     = data.js_await.fetch.expression == "@js/raw:await fetch(\"https://example.com\")"
    error_message = ""
  }
}
