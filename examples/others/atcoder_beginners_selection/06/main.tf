data "js_raw" "number" {
  value = "Number"
}

#
# function sortDesc(nums) { return nums.sort(function(a, b) { return b - a }) }
#

resource "js_function" "sort_desc" {
  name   = "sortDesc"
  params = [js_function_param.sort_desc_nums.id]
  body   = [js_return.sort_desc_nums_sort.content]
}

resource "js_function_param" "sort_desc_nums" {
  name = "nums"
}

resource "js_return" "sort_desc_nums_sort" {
  value = js_function_call.sort_desc_nums_sort.content
}

resource "js_function_call" "sort_desc_nums_sort" {
  caller   = js_function_param.sort_desc_nums.id
  function = "sort"
  args     = [js_function.b_minus_a.content]
}

resource "js_function" "b_minus_a" {
  params = [js_function_param.b_minus_a_a.id, js_function_param.b_minus_a_b.id]
  body   = [js_return.b_minus_a.content]
}

resource "js_function_param" "b_minus_a_a" {
  name = "a"
}

resource "js_function_param" "b_minus_a_b" {
  name = "b"
}

resource "js_return" "b_minus_a" {
  value = js_operation.b_minus_a.content
}

resource "js_operation" "b_minus_a" {
  left     = js_function_param.b_minus_a_b.id
  operator = "-"
  right    = js_function_param.b_minus_a_a.id
}

#
# const aa = input.trim().split("\n")[1].split(" ")
#

resource "js_const" "aa" {
  name  = "aa"
  value = js_function_call.input_trim_split1_split.content
}

resource "js_function_call" "input_trim_split1_split" {
  caller   = data.js_index.input_trim_split1.id
  function = "split"
  args     = [" "]
}

data "js_index" "input_trim_split1" {
  ref   = js_function_call.input_trim_split.content
  value = 1
}

resource "js_function_call" "input_trim_split" {
  caller   = js_function_call.input_trim.content
  function = "split"
  args     = ["\n"]
}

resource "js_function_call" "input_trim" {
  caller   = js_function_param.input.id
  function = "trim"
}

#
# const cards = aa.map(Number)
#

resource "js_const" "cards" {
  name  = "cards"
  value = js_function_call.aa_map.content
}

resource "js_function_call" "aa_map" {
  caller   = js_const.aa.id
  function = "map"
  args     = [data.js_raw.number.content]
}

#
# sortDesc(cards)
#

resource "js_function_call" "sort_cards" {
  function = js_function.sort_desc.id
  args     = [js_const.cards.id]
}

#
# let alice = 0
#

resource "js_let" "alice" {
  name  = "alice"
  value = 0
}

#
# let bob = 0
#

resource "js_let" "bob" {
  name  = "bob"
  value = 0
}

#
# for(let i = 0; i < cards.length; i++)
#

resource "js_for" "cards" {
  init      = js_let.for_i.content
  condition = js_operation.i_lt_cards_length.content
  update    = js_increment.for_i.content
  body      = [js_if.i_mod_2_eq_0.content]
}

resource "js_let" "for_i" {
  name  = "i"
  value = 0
}

resource "js_operation" "i_lt_cards_length" {
  left     = js_let.for_i.id
  operator = "<"
  right    = data.js_index.cards_length.id
}

data "js_index" "cards_length" {
  ref   = js_const.cards.id
  value = "length"
}

resource "js_increment" "for_i" {
  ref = js_let.for_i.id
}

#
# if (i % 2 === 0) { alice += cards[i] } else { bob += cards[i] }
#

resource "js_if" "i_mod_2_eq_0" {
  condition = js_operation.i_mod_2_eq_0.content
  then      = [js_operation.alice_plus_eq_cards_i.content]
  else      = [js_operation.bob_plus_eq_cards_i.content]
}

resource "js_operation" "i_mod_2_eq_0" {
  left     = js_operation.i_mod_2.content
  operator = "==="
  right    = 0
}

resource "js_operation" "i_mod_2" {
  left     = js_let.for_i.id
  operator = "%"
  right    = 2
}

# cards[i]
data "js_index" "cards_i" {
  ref   = js_const.cards.id
  value = js_let.for_i.id
}

# alice += cards[i]
resource "js_operation" "alice_plus_eq_cards_i" {
  left     = js_let.alice.id
  operator = "+="
  right    = data.js_index.cards_i.id
}

# bob += cards[i]
resource "js_operation" "bob_plus_eq_cards_i" {
  left     = js_let.bob.id
  operator = "+="
  right    = data.js_index.cards_i.id
}

#
# console.log(alice - bob)
#

resource "js_function_call" "log_alice_minus_bob" {
  function = "console.log"
  args     = [js_operation.alice_minus_bob.content]
}

resource "js_operation" "alice_minus_bob" {
  left     = js_let.alice.id
  operator = "-"
  right    = js_let.bob.id
}

#
# main
#

resource "js_function" "main" {
  name   = "main"
  params = [js_function_param.input.id]
  body = [
    js_function.sort_desc.content,
    js_const.aa.content,
    js_const.cards.content,
    js_function_call.sort_cards.content,
    js_let.alice.content,
    js_let.bob.content,
    js_for.cards.content,
    js_function_call.log_alice_minus_bob.content,
  ]
}

resource "js_function_param" "input" {
  name = "input"
}

#
# main(require("fs").readFileSync("/dev/stdin", "utf8"))
#

resource "js_function_call" "main" {
  function = js_function.main.id
  args     = [js_function_call.read_stdin.content]
}

resource "js_function_call" "require_fs" {
  function = "require"
  args     = ["fs"]
}

resource "js_function_call" "read_stdin" {
  caller   = js_function_call.require_fs.content
  function = "readFileSync"
  args     = ["/dev/stdin", "utf8"]
}

#
# write to file
#

resource "js_program" "main" {
  contents = [
    js_function.main.content,
    js_function_call.main.content,
  ]
}

resource "local_file" "main" {
  filename = "index.js"
  content  = js_program.main.content
}
