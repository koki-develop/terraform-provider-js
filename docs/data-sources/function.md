---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_function Data Source - terraform-provider-js"
subcategory: ""
description: |-
  
---

# js_function (Data Source)



## Example Usage

```terraform
data "js_function_call" "log_name" {
  caller   = "console"
  function = "log"
  args     = ["hello", data.js_function_param.name.id]
}

data "js_function_param" "name" {
  name = "name"
}

data "js_function" "hello" {
  name   = "hello"
  params = [data.js_function_param.name.id]
  body   = [data.js_function_call.log_name.statement]
}
# => function hello(name) {
#      console.log("hello", name);
#    }

data "js_function" "anonymous" {
  params = [data.js_function_param.name.id]
  body   = [data.js_function_call.log_name.statement]
}

data "js_const" "anonymous" {
  name  = "anonymous"
  value = data.js_function.hello.expression
}
# => const anonymous = function(name) {
#      console.log("hello", name);
#    };

data "js_function" "async" {
  name   = "hello"
  async  = true
  params = [data.js_function_param.name.id]
  body   = [data.js_function_call.log_name.statement]
}
# => async function hello(name) {
#      console.log("hello", name);
#    }
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `async` (Boolean) Whether function is async.
- `body` (List of String) Body of function.
- `name` (String) Name of function.
- `params` (List of String) Parameters of function.

### Read-Only

- `expression` (String)
- `id` (String) The ID of this resource.
- `statement` (String)
