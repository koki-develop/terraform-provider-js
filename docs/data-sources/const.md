---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_const Data Source - terraform-provider-js"
subcategory: ""
description: |-
  
---

# js_const (Data Source)



## Example Usage

```terraform
data "js_const" "str" {
  name  = "str"
  value = "hello"
}
# => const str = "hello";

data "js_const" "num" {
  name  = "num"
  value = 10
}
# => const num = 10;

data "js_const" "bool" {
  name  = "bool"
  value = true
}
# => const bool = true;

data "js_const" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}
# => const arr = [1, 2, 3];

data "js_const" "obj" {
  name = "obj"
  value = {
    hoge = "fuga"
    foo  = 3
  }
}
# => const obj = { "hoge": "fuga", "foo": 3 };
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of variable to declare.
- `value` (Dynamic) Initial value of the variable.

### Read-Only

- `expression` (String)
- `id` (String) The ID of this resource.
- `statement` (String)
