#
# const cards
#

resource "js_const" "cards" {
  name  = "cards"
  value = js_function_call.input_split_1_split_map.content
}

resource "js_function_call" "input_split" {
  caller   = js_function_param.input.id
  function = "split"
  args     = ["\n"]
}

data "js_index" "input_split_1" {
  ref   = js_function_call.input_split.content
  value = 1
}

resource "js_function_call" "input_split_1_split" {
  caller   = data.js_index.input_split_1.content
  function = "split"
  args     = [" "]
}

resource "js_function_call" "input_split_1_split_map" {
  caller   = js_function_call.input_split_1_split.content
  function = "map"
  args     = [data.js_raw.number.content]
}

data "js_raw" "number" {
  value = "Number"
}

#
# sort cards
#

resource "js_function_call" "sort_cards" {
  caller   = js_const.cards.id
  function = "sort"
  args     = [js_function.sort_cards.content]
}

resource "js_function" "sort_cards" {
  params = [js_function_param.sort_cards_a.id, js_function_param.sort_cards_b.id]
  body   = [js_return.b_minus_a.content]
}

resource "js_function_param" "sort_cards_a" {
  name = "a"
}

resource "js_function_param" "sort_cards_b" {
  name = "b"
}

resource "js_return" "b_minus_a" {
  value = js_operation.b_minus_a.content
}

resource "js_operation" "b_minus_a" {
  left     = js_function_param.sort_cards_b.id
  operator = "-"
  right    = js_function_param.sort_cards_a.id
}

#
# let alice, bob
#

resource "js_let" "alice" {
  name  = "alice"
  value = 0
}

resource "js_let" "bob" {
  name  = "bob"
  value = 0
}

#
# for loop
#

resource "js_for" "cards" {
  init      = js_let.for_i.content
  condition = js_operation.for_condition.content
  update    = js_increment.for_update.content
  body      = [js_if.cards.content]
}

resource "js_let" "for_i" {
  name  = "i"
  value = 0
}

resource "js_operation" "for_condition" {
  left     = js_let.for_i.id
  operator = "<"
  right    = data.js_index.cards_length.id
}

data "js_index" "cards_length" {
  ref   = js_const.cards.id
  value = "length"
}

resource "js_increment" "for_update" {
  ref = js_let.for_i.id
}

data "js_index" "cards_i" {
  ref   = js_const.cards.id
  value = js_let.for_i.id
}

resource "js_if" "cards" {
  condition = js_operation.i_mod_2_eq_0.content
  then      = [js_operation.add_card_to_alice.content]
  else      = [js_operation.add_card_to_bob.content]
}

resource "js_operation" "i_mod_2" {
  left     = js_let.for_i.id
  operator = "%"
  right    = 2
}

resource "js_operation" "i_mod_2_eq_0" {
  left     = js_operation.i_mod_2.content
  operator = "==="
  right    = 0
}

resource "js_operation" "add_card_to_alice" {
  left     = js_let.alice.id
  operator = "+="
  right    = data.js_index.cards_i.id
}

resource "js_operation" "add_card_to_bob" {
  left     = js_let.bob.id
  operator = "+="
  right    = data.js_index.cards_i.id
}

#
# print
#

resource "js_function_call" "log_result" {
  caller   = "console"
  function = "log"
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
    js_const.cards.content,
    js_function_call.sort_cards.content,
    js_let.alice.content,
    js_let.bob.content,
    js_for.cards.content,
    js_function_call.log_result.content,
  ]
}

resource "js_function_param" "input" {
  name = "input"
}

#
# call main
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

resource "js_program" "main" {
  contents = [
    js_function.main.content,
    js_function_call.main.content,
  ]
}

#
# write to file
#

resource "local_file" "main" {
  filename = "index.js"
  content  = js_program.main.content
}
