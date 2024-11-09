---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_throw Data Source - terraform-provider-js"
subcategory: ""
description: |-
  
---

# js_throw (Data Source)



## Example Usage

```terraform
data "js_new" "error" {
  constructor = "Error"
  args        = ["something went wrong"]
}

data "js_throw" "error" {
  value = data.js_new.error.content
}
# => throw new Error("something went wrong")
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `value` (String) Expression to throw.

### Read-Only

- `content` (String)