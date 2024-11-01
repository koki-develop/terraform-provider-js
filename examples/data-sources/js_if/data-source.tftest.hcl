run "main" {
  assert {
    condition     = data.js_if.main.content == "@js/raw:if(true){console.log(true)}else{console.log(false)}"
    error_message = ""
  }
}
