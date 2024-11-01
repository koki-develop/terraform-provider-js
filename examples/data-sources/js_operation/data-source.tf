data "js_let" "a" {
  name  = "a"
  value = 0
}

data "js_let" "b" {
  name  = "b"
  value = 10
}

data "js_operation" "a_plus_b" {
  left     = data.js_let.a.id
  operator = "+"
  right    = data.js_let.b.id
}
# => a + b

data "js_operation" "a_minus_b" {
  left     = data.js_let.a.id
  operator = "-"
  right    = data.js_let.b.id
}
# => a - b

data "js_operation" "a_times_b" {
  left     = data.js_let.a.id
  operator = "*"
  right    = data.js_let.b.id
}
# => a * b

data "js_operation" "a_div_b" {
  left     = data.js_let.a.id
  operator = "/"
  right    = data.js_let.b.id
}
# => a / b

data "js_operation" "a_mod_b" {
  left     = data.js_let.a.id
  operator = "%"
  right    = data.js_let.b.id
}
# => a % b

data "js_operation" "a_lt_b" {
  left     = data.js_let.a.id
  operator = "<"
  right    = data.js_let.b.id
}
# => a < b

data "js_operation" "a_lte_b" {
  left     = data.js_let.a.id
  operator = "<="
  right    = data.js_let.b.id
}
# => a <= b

data "js_operation" "a_gt_b" {
  left     = data.js_let.a.id
  operator = ">"
  right    = data.js_let.b.id
}
# => a > b

data "js_operation" "a_gte_b" {
  left     = data.js_let.a.id
  operator = ">="
  right    = data.js_let.b.id
}
# => a > b

data "js_operation" "a_eq_b" {
  left     = data.js_let.a.id
  operator = "==="
  right    = data.js_let.b.id
}
# => a === b

data "js_operation" "assign_a_to_b" {
  left     = data.js_let.a.id
  operator = "="
  right    = data.js_let.b.id
}
# => a = b
