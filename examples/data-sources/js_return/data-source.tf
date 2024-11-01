data "js_return" "hello" {
  value = "hello world"
}

data "js_function" "hello" {
  name = "hello"
  body = [data.js_return.hello.content]
}
# => function hello() {
#      return "hello world"
#    }
