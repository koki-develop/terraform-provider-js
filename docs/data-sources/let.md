---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_let Data Source - terraform-provider-js"
subcategory: ""
description: |-
  
---

# js_let (Data Source)



## Example Usage

```terraform
data "js_let" "str" {
  name  = "str"
  value = "hello"
}
# => let str = "hello";

data "js_let" "num" {
  name  = "num"
  value = 10
}
# => let num = 10;

data "js_let" "bool" {
  name  = "bool"
  value = true
}
# => let bool = true;

data "js_let" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}
# => let arr = [1, 2, 3];

data "js_let" "obj" {
  name = "obj"
  value = {
    hoge = "fuga"
    foo  = 3
  }
}
# => let obj = { "hoge": "fuga", "foo": 3 };
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of variable to declare.

### Optional

- `value` (Dynamic) Initial value of the variable.

### Read-Only

- `content` (String)
- `id` (String) The ID of this resource.