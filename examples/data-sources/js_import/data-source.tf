data "js_import" "as" {
  from = "path/to/module"
  as   = "name"
}
# => import * as name from "path/to/module"

data "js_import" "default" {
  from    = "path/to/module"
  as      = "name"
  default = true
}
# => import name from "path/to/module"
