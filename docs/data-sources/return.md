---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_return Data Source - terraform-provider-js"
subcategory: ""
description: |-
  
---

# js_return (Data Source)



## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `value` (Dynamic) Expression whose value is to be returned.

### Read-Only

- `statement` (String)
