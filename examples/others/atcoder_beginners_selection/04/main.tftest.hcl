run "main" {
  assert {
    condition = local_file.main.content == chomp(
      <<-EOT
      // Code generated by JS.tf vdev (https://registry.terraform.io/providers/koki-develop/js/dev)
      function main(input){const abcx=input.trim().split("\n").map(Number);let count=0;for(let i=0;i<=abcx[0];i++){for(let j=0;j<=abcx[1];j++){for(let k=0;k<=abcx[2];k++){if(i*500+j*100+k*50===abcx[3]){count++}}}}console.log(count)}main(require("fs").readFileSync("/dev/stdin","utf8"))
      EOT
    )
    error_message = ""
  }
}
