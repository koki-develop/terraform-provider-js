---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "js_for Resource - terraform-provider-js"
subcategory: ""
description: |-
  The js_for resource creates a for loop.
---

# js_for (Resource)

The `js_for` resource creates a for loop.

## Example Usage

```terraform
resource "js_let" "i" {
  name  = "i"
  value = 0
}

resource "js_operation" "i_lt_10" {
  left     = js_let.i.id
  operator = "<"
  right    = 10
}

resource "js_increment" "i" {
  ref = js_let.i.id
}

resource "js_function_call" "log_i" {
  caller   = "console"
  function = "log"
  args     = [js_let.i.id]
}

resource "js_for" "main" {
  init      = js_let.i.content
  condition = js_operation.i_lt_10.content
  update    = js_increment.i.content
  body      = [js_function_call.log_i.content]
}
# => for (let i = 0; i < 10; i++) {
#      console.log(i);
#    }
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `body` (List of String) The body of the for loop.
- `condition` (String) The condition expression.
- `init` (String) The initialization expression.
- `update` (String) The update expression.

### Read-Only

- `content` (String) The content of the for loop.
