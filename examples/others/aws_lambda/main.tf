resource "aws_lambda_function" "main" {
  function_name = "hello-jstf"
  role          = aws_iam_role.assume_role.arn
  publish       = true

  runtime          = "nodejs20.x"
  source_code_hash = data.archive_file.source_code.output_base64sha256
  filename         = data.archive_file.source_code.output_path
  handler          = "index.handler"
}

resource "aws_iam_role" "assume_role" {
  name               = "iam_for_lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

data "archive_file" "source_code" {
  type                    = "zip"
  output_path             = "index.zip"
  source_content_filename = "index.mjs"
  source_content          = data.js_program.main.content
}

data "js_program" "main" {
  statements = [data.js_export.handler.statement]
}

data "js_function" "handler" {
  name   = "handler"
  async  = true
  params = [data.js_function_param.event.id]
  body = [
    data.js_function_call.log_event.statement,
    data.js_return.handler.statement,
  ]
}

data "js_function_param" "event" {
  name = "event"
}

data "js_function_call" "log_event" {
  caller   = "console"
  function = "log"
  args     = ["event:", data.js_function_param.event.id]
}

data "js_return" "handler" {
  value = {
    message = "Hello JS.tf!"
  }
}

data "js_export" "handler" {
  value = data.js_function.handler.statement
}
