resource "js_function_call" "log_name" {
  caller   = "console"
  function = "log"
  args     = ["hello", js_function_param.name.id]
}

resource "js_function_param" "name" {
  name = "name"
}

resource "js_function" "hello" {
  name   = "hello"
  params = [js_function_param.name.id]
  body   = [js_function_call.log_name.content]
}
# => function hello(name) {
#      console.log("hello", name);
#    }

resource "js_function" "anonymous" {
  params = [js_function_param.name.id]
  body   = [js_function_call.log_name.content]
}

resource "js_const" "anonymous" {
  name  = "anonymous"
  value = js_function.hello.content
}
# => const anonymous = function(name) {
#      console.log("hello", name);
#    };
