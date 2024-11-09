---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_decrement Data Source - terraform-provider-js"
subcategory: ""
description: |-
  
---

# js_decrement (Data Source)



## Example Usage

```terraform
data "js_let" "num" {
  name  = "i"
  value = 0
}

data "js_decrement" "i" {
  ref = data.js_let.num.id
}
# => i--

data "js_decrement" "i_postfix" {
  ref  = data.js_let.num.id
  type = "postfix"
}
# => i--

data "js_decrement" "i_prefix" {
  ref  = data.js_let.num.id
  type = "prefix"
}
# => --i
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ref` (String) Reference to decrement.

### Optional

- `type` (String) Type of decrement to perform. (Valid values: `prefix`, `postfix`)

### Read-Only

- `content` (String)