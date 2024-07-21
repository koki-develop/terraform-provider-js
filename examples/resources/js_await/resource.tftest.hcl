run "fetch" {
  assert {
    condition     = js_await.fetch.content == "@js/raw:await fetch(\"https://example.com\")"
    error_message = ""
  }
}
