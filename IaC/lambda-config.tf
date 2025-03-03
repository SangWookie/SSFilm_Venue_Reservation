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

  source_code_hash = data.archive_file.login_manager_zip.output_base64sha256
}

locals {
  lambda_functions = {
    "user_reservation_manager" = aws_lambda_function.user_reservation_manager.function_name
    "login_manager"            = aws_lambda_function.login_manager.function_name
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
