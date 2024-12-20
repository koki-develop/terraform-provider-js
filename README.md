<h1 align="center">
JS.tf
</h1>

<p align="center">
<b>
<i>
JavaScript × Terraform
</i>
</b>
</p>

<p align="center">
The Next Generation AltJS.
</p>

<p align="center">
<a href="https://github.com/koki-develop/terraform-provider-js/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/terraform-provider-js" alt="GitHub release (latest by date)"></a>
<a href="https://github.com/koki-develop/terraform-provider-js/actions/workflows/test.yml"><img src="https://img.shields.io/github/actions/workflow/status/koki-develop/terraform-provider-js/test.yml?logo=github" alt="GitHub Workflow Status"></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/terraform-provider-js"><img src="https://goreportcard.com/badge/github.com/koki-develop/terraform-provider-js" alt="Go Report Card"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/terraform-provider-js" alt="LICENSE"></a>
</p>

```hcl
data "js_function_call" "hello_world" {
  caller   = "console"
  function = "log"
  args     = ["hello world"]
}

data "js_program" "main" {
  statements = [data.js_function_call.hello_world.statement]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = data.js_program.main.content
}
```

```console
$ terraform init
$ terraform apply
$ node index.js
hello world
```

# Getting Started

JS.tf is a Terraform provider that allows you to write JavaScript code in Terraform configuration files.
To use it, add the following provider configuration.

```hcl
terraform {
  required_providers {
    js = {
      source = "koki-develop/js"
    }
  }
}

provider "js" {}
```

Next, run `terraform init`.

```console
$ terraform init
```

That's it. You are ready to use JS.tf.

# Examples

- [`examples/`](./examples)

# Documentation

- [Terraform Registry](https://registry.terraform.io/providers/koki-develop/js/latest/docs)

# LICENSE

[MIT](./LICENSE)
