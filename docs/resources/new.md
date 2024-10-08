---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_new Resource - terraform-provider-js"
subcategory: ""
description: |-
  The js_new resource creates a new operation.
---

# js_new (Resource)

The `js_new` resource creates a new operation.

## Example Usage

```terraform
resource "js_function_call" "error" {
  function = "Error"
  args     = ["something went wrong"]
}

resource "js_new" "error" {
  value = js_function_call.error.content
}
# => new Error("something went wrong")
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `value` (String) The value of the operation.

### Read-Only

- `content` (String) The content of the operation.
