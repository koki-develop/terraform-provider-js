run "main" {
  assert {
    condition = local_file.main.content == chomp(
      <<-EOT
      // Code generated by JS.tf vdev (https://registry.terraform.io/providers/koki-develop/js/dev)
      function main(input){const lines=input.trim().split("\n");const a=Number(lines[0]);const bc=lines[1].split(" ");const b=Number(bc[0]);const c=Number(bc[1]);const s=lines[2];console.log(a+b+c,s)}main(require("fs").readFileSync("/dev/stdin","utf8"))
      EOT
    )
    error_message = ""
  }
}
