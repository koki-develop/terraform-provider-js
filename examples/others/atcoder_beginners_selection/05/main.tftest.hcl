run "main" {
  assert {
    condition = local_file.main.content == chomp(
      <<-EOT
      // Code generated by JS.tf vdev (https://registry.terraform.io/providers/koki-develop/js/dev)
      function main(input){function add(a,b){return a+b}function sum(nums){return nums.reduce(add)}function digits(num){return num.toString().split("").map(Number)}function sumDigits(num){return sum(digits(num))}const nab=input.trim().split(" ").map(Number);let ans=0;for(let i=0;i<=nab[0];i++){const cmp=sumDigits(i);if(nab[1]<=cmp&&cmp<=nab[2]){ans+=i}}console.log(ans)}main(require("fs").readFileSync("/dev/stdin","utf8"))
      EOT
    )
    error_message = ""
  }
}
