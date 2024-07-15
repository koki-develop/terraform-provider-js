run "main" {
  assert {
    condition     = js_for.main.content == "@js/raw:for(let i=0;i<10;i++){console.log(i)}"
    error_message = ""
  }
}
