run "main" {
  assert {
    condition     = data.js_if.main.statement == "@js/raw:if(true){console.log(true)}else{console.log(false)}"
    error_message = ""
  }
}
