run "main" {
  assert {
    condition     = data.js_for.main.statement == "@js/raw:for(let i=0;(i<10);i++){console.log(i)}"
    error_message = ""
  }
}
