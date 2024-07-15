resource "js_return" "hello" {
  value = "hello world"
}

resource "js_function" "hello" {
  name = "hello"
  body = [js_return.hello.content]
}
# => function hello() {
#      return "hello world"
#    }
