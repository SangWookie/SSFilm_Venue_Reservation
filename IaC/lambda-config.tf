# Lambda
data "archive_file" "user_reservation_manager_zip" {
  type        = "zip"
  source_dir  = "../lambda_functions/user_reservation_manager"
  output_path = "../lambda_functions/user_reservation_manager.zip"
}

resource "aws_lambda_function" "user_reservation_manager" {
  description   = "handles user reservation related requests"
  filename      = data.archive_file.user_reservation_manager_zip.output_path
  function_name = "user_reservation_manager"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"

  # comment this line to upload source code only once(untrack changes)
  source_code_hash = data.archive_file.user_reservation_manager_zip.output_base64sha256
}

data "archive_file" "login_manager_zip" {
  type        = "zip"
  source_dir  = "../lambda_functions/login_manager"
  output_path = "../lambda_functions/login_manager.zip"
}

resource "aws_lambda_function" "login_manager" {
  description   = "handles user reservation related requests"
  filename      = data.archive_file.login_manager_zip.output_path
  function_name = "login_manager"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"

  # comment this line to upload source code only once(untrack changes)
  source_code_hash = data.archive_file.login_manager_zip.output_base64sha256
}

data "archive_file" "request_manager_zip" {
  type        = "zip"
  source_dir  = "../lambda_functions/request_manager"
  output_path = "../lambda_functions/request_manager.zip"
}

resource "aws_lambda_function" "request_manager" {
  description   = "manages user reservations"
  filename      = data.archive_file.request_manager_zip.output_path
  function_name = "request_manager"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"

  # comment this line to upload source code only once(untrack changes)
  source_code_hash = data.archive_file.request_manager_zip.output_base64sha256
}

data "archive_file" "mode_manager_zip" {
  type        = "zip"
  source_dir  = "../lambda_functions/mode_manager"
  output_path = "../lambda_functions/mode_manager.zip"
}

resource "aws_lambda_function" "mode_manager" {
  description   = "sets limitations for reservations"
  filename      = data.archive_file.mode_manager_zip.output_path
  function_name = "mode_manager"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"

  # comment this line to upload source code only once(untrack changes)
  source_code_hash = data.archive_file.mode_manager_zip.output_base64sha256
}

data "archive_file" "stat_handler_zip" {
  type        = "zip"
  source_dir  = "../lambda_functions/stat_handler"
  output_path = "../lambda_functions/stat_handler.zip"
}

resource "aws_lambda_function" "stat_handler" {
  description   = "handles stat related requests"
  filename      = data.archive_file.stat_handler_zip.output_path
  function_name = "stat_handler"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"

  # comment this line to upload source code only once(untrack changes)
  source_code_hash = data.archive_file.stat_handler_zip.output_base64sha256
}

locals {
  lambda_functions = {
    "user_reservation_manager" = aws_lambda_function.user_reservation_manager.function_name
    "login_manager"            = aws_lambda_function.login_manager.function_name
    "request_manager"          = aws_lambda_function.request_manager.function_name
    "mode_manager"             = aws_lambda_function.mode_manager.function_name
    "stat_handler"             = aws_lambda_function.stat_handler.function_name
  }
}

resource "aws_lambda_permission" "apigw_lambda" {
  for_each      = local.lambda_functions
  statement_id  = "AllowExecutionFromAPIGateway-${each.key}"
  action        = "lambda:InvokeFunction"
  function_name = each.value
  principal     = "apigateway.amazonaws.com"

  # More: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html
  source_arn = "${aws_apigatewayv2_api.mainGW.execution_arn}/*"
}
