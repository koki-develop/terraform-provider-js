---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_let Resource - terraform-provider-js"
subcategory: ""
description: |-
  The js_let resource defines a let statement.
---

# js_let (Resource)

The `js_let` resource defines a let statement.

## Example Usage

```terraform
resource "js_let" "str" {
  name  = "str"
  value = "hello"
}
# => let str = "hello";

resource "js_let" "num" {
  name  = "num"
  value = 10
}
# => let num = 10;

resource "js_let" "bool" {
  name  = "bool"
  value = true
}
# => let bool = true;

resource "js_let" "arr" {
  name  = "arr"
  value = [1, 2, 3]
}
# => let arr = [1, 2, 3];

resource "js_let" "obj" {
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

- `name` (String) The name of the let statement.

### Optional

- `value` (Dynamic) The value of the let statement.

### Read-Only

- `content` (String) The content of the let statement.
- `id` (String) The id of the let statement.
