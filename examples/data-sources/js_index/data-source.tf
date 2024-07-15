data "js_index" "example_index" {
  ref   = "example"
  value = 1
}
# => example[1]

data "js_index" "example_property" {
  ref   = "example"
  value = "name"
}
# => example["name"]
