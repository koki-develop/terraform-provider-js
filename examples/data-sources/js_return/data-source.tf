data "js_return" "hello" {
  value = "hello world"
}

data "js_function" "hello" {
  name = "hello"
  body = [data.js_return.hello.statement]
}
# => function hello() {
#      return "hello world"
#    }
