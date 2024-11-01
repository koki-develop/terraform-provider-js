data "js_raw" "number" {
  value = "Number"
}

#
# function sortDesc(nums) { return nums.sort(function(a, b) { return b - a }) }
#

data "js_function" "sort_desc" {
  name   = "sortDesc"
  params = [data.js_function_param.sort_desc_nums.id]
  body   = [data.js_return.sort_desc_nums_sort.content]
}

data "js_function_param" "sort_desc_nums" {
  name = "nums"
}

data "js_return" "sort_desc_nums_sort" {
  value = data.js_function_call.sort_desc_nums_sort.content
}

data "js_function_call" "sort_desc_nums_sort" {
  caller   = data.js_function_param.sort_desc_nums.id
  function = "sort"
  args     = [data.js_function.b_minus_a.content]
}

data "js_function" "b_minus_a" {
  params = [data.js_function_param.b_minus_a_a.id, data.js_function_param.b_minus_a_b.id]
  body   = [data.js_return.b_minus_a.content]
}

data "js_function_param" "b_minus_a_a" {
  name = "a"
}

data "js_function_param" "b_minus_a_b" {
  name = "b"
}

data "js_return" "b_minus_a" {
  value = data.js_operation.b_minus_a.content
}

data "js_operation" "b_minus_a" {
  left     = data.js_function_param.b_minus_a_b.id
  operator = "-"
  right    = data.js_function_param.b_minus_a_a.id
}

#
# const aa = input.trim().split("\n")[1].split(" ")
#

data "js_const" "aa" {
  name  = "aa"
  value = data.js_function_call.input_trim_split1_split.content
}

data "js_function_call" "input_trim_split1_split" {
  caller   = data.js_index.input_trim_split1.content
  function = "split"
  args     = [" "]
}

data "js_index" "input_trim_split1" {
  ref   = data.js_function_call.input_trim_split.content
  value = 1
}

data "js_function_call" "input_trim_split" {
  caller   = data.js_function_call.input_trim.content
  function = "split"
  args     = ["\n"]
}

data "js_function_call" "input_trim" {
  caller   = data.js_function_param.input.id
  function = "trim"
}

#
# const cards = aa.map(Number)
#

data "js_const" "cards" {
  name  = "cards"
  value = data.js_function_call.aa_map.content
}

data "js_function_call" "aa_map" {
  caller   = data.js_const.aa.id
  function = "map"
  args     = [data.js_raw.number.content]
}

#
# sortDesc(cards)
#

data "js_function_call" "sort_cards" {
  function = data.js_function.sort_desc.id
  args     = [data.js_const.cards.id]
}

#
# let alice = 0
#

data "js_let" "alice" {
  name  = "alice"
  value = 0
}

#
# let bob = 0
#

data "js_let" "bob" {
  name  = "bob"
  value = 0
}

#
# for(let i = 0; i < cards.length; i++)
#

data "js_for" "cards" {
  init      = data.js_let.for_i.content
  condition = data.js_operation.i_lt_cards_length.content
  update    = data.js_increment.for_i.content
  body      = [data.js_if.i_mod_2_eq_0.content]
}

data "js_let" "for_i" {
  name  = "i"
  value = 0
}

data "js_operation" "i_lt_cards_length" {
  left     = data.js_let.for_i.id
  operator = "<"
  right    = data.js_index.cards_length.content
}

data "js_index" "cards_length" {
  ref   = data.js_const.cards.id
  value = "length"
}

data "js_increment" "for_i" {
  ref = data.js_let.for_i.id
}

#
# if (i % 2 === 0) { alice += cards[i] } else { bob += cards[i] }
#

data "js_if" "i_mod_2_eq_0" {
  condition = data.js_operation.i_mod_2_eq_0.content
  then      = [data.js_operation.alice_plus_eq_cards_i.content]
  else      = [data.js_operation.bob_plus_eq_cards_i.content]
}

data "js_operation" "i_mod_2_eq_0" {
  left     = data.js_operation.i_mod_2.content
  operator = "==="
  right    = 0
}

data "js_operation" "i_mod_2" {
  left     = data.js_let.for_i.id
  operator = "%"
  right    = 2
}

# cards[i]
data "js_index" "cards_i" {
  ref   = data.js_const.cards.id
  value = data.js_let.for_i.id
}

# alice += cards[i]
data "js_operation" "alice_plus_eq_cards_i" {
  left     = data.js_let.alice.id
  operator = "+="
  right    = data.js_index.cards_i.content
}

# bob += cards[i]
data "js_operation" "bob_plus_eq_cards_i" {
  left     = data.js_let.bob.id
  operator = "+="
  right    = data.js_index.cards_i.content
}

#
# console.log(alice - bob)
#

data "js_function_call" "log_alice_minus_bob" {
  function = "console.log"
  args     = [data.js_operation.alice_minus_bob.content]
}

data "js_operation" "alice_minus_bob" {
  left     = data.js_let.alice.id
  operator = "-"
  right    = data.js_let.bob.id
}

#
# main
#

data "js_function" "main" {
  name   = "main"
  params = [data.js_function_param.input.id]
  body = [
    data.js_function.sort_desc.content,
    data.js_const.aa.content,
    data.js_const.cards.content,
    data.js_function_call.sort_cards.content,
    data.js_let.alice.content,
    data.js_let.bob.content,
    data.js_for.cards.content,
    data.js_function_call.log_alice_minus_bob.content,
  ]
}

data "js_function_param" "input" {
  name = "input"
}

#
# main(require("fs").readFileSync("/dev/stdin", "utf8"))
#

data "js_function_call" "main" {
  function = data.js_function.main.id
  args     = [data.js_function_call.read_stdin.content]
}

data "js_function_call" "require_fs" {
  function = "require"
  args     = ["fs"]
}

data "js_function_call" "read_stdin" {
  caller   = data.js_function_call.require_fs.content
  function = "readFileSync"
  args     = ["/dev/stdin", "utf8"]
}

#
# write to file
#

data "js_program" "main" {
  contents = [
    data.js_function.main.content,
    data.js_function_call.main.content,
  ]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = data.js_program.main.content
}
