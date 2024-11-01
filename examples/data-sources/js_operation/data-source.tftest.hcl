run "a_plus_b" {
  assert {
    condition     = data.js_operation.a_plus_b.content == "@js/raw:a+b"
    error_message = ""
  }
}

run "a_minus_b" {
  assert {
    condition     = data.js_operation.a_minus_b.content == "@js/raw:a-b"
    error_message = ""
  }
}

run "a_times_b" {
  assert {
    condition     = data.js_operation.a_times_b.content == "@js/raw:a*b"
    error_message = ""
  }
}

run "a_div_b" {
  assert {
    condition     = data.js_operation.a_div_b.content == "@js/raw:a/b"
    error_message = ""
  }
}

run "a_mod_b" {
  assert {
    condition     = data.js_operation.a_mod_b.content == "@js/raw:a%b"
    error_message = ""
  }
}

run "a_lt_b" {
  assert {
    condition     = data.js_operation.a_lt_b.content == "@js/raw:a<b"
    error_message = ""
  }
}

run "a_lte_b" {
  assert {
    condition     = data.js_operation.a_lte_b.content == "@js/raw:a<=b"
    error_message = ""
  }
}

run "a_gt_b" {
  assert {
    condition     = data.js_operation.a_gt_b.content == "@js/raw:a>b"
    error_message = ""
  }
}

run "a_gte_b" {
  assert {
    condition     = data.js_operation.a_gte_b.content == "@js/raw:a>=b"
    error_message = ""
  }
}

run "a_eq_b" {
  assert {
    condition     = data.js_operation.a_eq_b.content == "@js/raw:a===b"
    error_message = ""
  }
}

run "assign_a_to_b" {
  assert {
    condition     = data.js_operation.assign_a_to_b.content == "@js/raw:a=b"
    error_message = ""
  }
}
