run "main" {
  assert {
    condition = local_file.main.content == chomp(
      <<-EOT
    // Code generated by JS.tf vdev (https://registry.terraform.io/providers/koki-develop/js/dev)
    console.log("hello world")
    EOT
    )
    error_message = ""
  }
}