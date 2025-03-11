resource "aws_lambda_function" "user_reservation_manager" {
  description   = "handles user reservation related requests"
  function_name = "user_reservation_manager"
  role          = aws_iam_role.sqs_lambda_role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"
}

resource "aws_lambda_function" "login_manager" {
  description   = "handles user reservation related requests"
  function_name = "login_manager"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"
}

resource "aws_lambda_function" "request_manager" {
  description   = "manages user reservations"
  function_name = "request_manager"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"
}

resource "aws_lambda_function" "mode_manager" {
  description   = "sets limitations for reservations"
  function_name = "mode_manager"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"
}

resource "aws_lambda_function" "stat_handler" {
  description   = "handles stat related requests"
  function_name = "stat_handler"
  role          = aws_iam_role.role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"
}

resource "aws_lambda_function" "reservation_queue_handler" {
  description   = "handles reservation queue"
  function_name = "reservation_queue_handler"
  role          = aws_iam_role.sqs_lambda_poll_role.arn
  handler       = "lambda.lambda_handler"
  runtime       = "python3.12"
}

resource "aws_lambda_event_source_mapping" "sqs_trigger" {
  event_source_arn = aws_sqs_queue.reservation_queue.arn
  function_name    = aws_lambda_function.reservation_queue_handler.function_name
  batch_size       = 5
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
