resource "js_let" "a" {
  name  = "a"
  value = 0
}

resource "js_let" "b" {
  name  = "b"
  value = 10
}

resource "js_operation" "a_plus_b" {
  left     = js_let.a.id
  operator = "+"
  right    = js_let.b.id
}
# => a + b

resource "js_operation" "a_minus_b" {
  left     = js_let.a.id
  operator = "-"
  right    = js_let.b.id
}
# => a - b

resource "js_operation" "a_times_b" {
  left     = js_let.a.id
  operator = "*"
  right    = js_let.b.id
}
# => a * b

resource "js_operation" "a_div_b" {
  left     = js_let.a.id
  operator = "/"
  right    = js_let.b.id
}
# => a / b

resource "js_operation" "a_mod_b" {
  left     = js_let.a.id
  operator = "%"
  right    = js_let.b.id
}
# => a % b

resource "js_operation" "a_lt_b" {
  left     = js_let.a.id
  operator = "<"
  right    = js_let.b.id
}
# => a < b

resource "js_operation" "a_lte_b" {
  left     = js_let.a.id
  operator = "<="
  right    = js_let.b.id
}
# => a <= b

resource "js_operation" "a_gt_b" {
  left     = js_let.a.id
  operator = ">"
  right    = js_let.b.id
}
# => a > b

resource "js_operation" "a_gte_b" {
  left     = js_let.a.id
  operator = ">="
  right    = js_let.b.id
}
# => a > b

resource "js_operation" "a_eq_b" {
  left     = js_let.a.id
  operator = "==="
  right    = js_let.b.id
}
# => a === b

resource "js_operation" "assign_a_to_b" {
  left     = js_let.a.id
  operator = "="
  right    = js_let.b.id
}
# => a = b
