resource "aws_apigatewayv2_integration" "reservation_proxy" {
  api_id           = aws_apigatewayv2_api.mainGW.id
  integration_type = "AWS_PROXY"

  connection_type        = "INTERNET"
  description            = "for reservation requests"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.user_reservation_manager.invoke_arn
  passthrough_behavior   = "WHEN_NO_MATCH"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "reservation_proxy" {
  api_id    = aws_apigatewayv2_api.mainGW.id
  route_key = "GET /reservations/{proxy+}"

  target = "integrations/${aws_apigatewayv2_integration.reservation_proxy.id}"
}

resource "aws_apigatewayv2_integration" "reservation" {
  api_id           = aws_apigatewayv2_api.mainGW.id
  integration_type = "AWS_PROXY"

  connection_type        = "INTERNET"
  description            = "for reservation requests"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.user_reservation_manager.invoke_arn
  passthrough_behavior   = "WHEN_NO_MATCH"
  payload_format_version = "2.0"
}

variable "reservation_methods" {
  type    = list(string)
  default = ["GET", "POST"]
}

resource "aws_apigatewayv2_route" "reservation" {
  for_each = toset(var.reservation_methods)

  api_id    = aws_apigatewayv2_api.mainGW.id
  route_key = "${each.value} /reservations"

  target = "integrations/${aws_apigatewayv2_integration.reservation.id}"
}

resource "aws_apigatewayv2_integration" "login" {
  api_id           = aws_apigatewayv2_api.mainGW.id
  integration_type = "AWS_PROXY"

  connection_type        = "INTERNET"
  description            = "login requests"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.login_manager.invoke_arn
  passthrough_behavior   = "WHEN_NO_MATCH"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "login" {
  api_id    = aws_apigatewayv2_api.mainGW.id
  route_key = "POST /login"

  target = "integrations/${aws_apigatewayv2_integration.login.id}"
}

resource "aws_apigatewayv2_integration" "admin_request" {
  api_id           = aws_apigatewayv2_api.mainGW.id
  integration_type = "AWS_PROXY"

  connection_type        = "INTERNET"
  description            = "admin requests"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.request_manager.invoke_arn
  passthrough_behavior   = "WHEN_NO_MATCH"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "admin_request" {
  api_id    = aws_apigatewayv2_api.mainGW.id
  route_key = "ANY /admin/requests/{proxy+}"

  target             = "integrations/${aws_apigatewayv2_integration.admin_request.id}"
  authorizer_id      = aws_apigatewayv2_authorizer.authorizer.id
  authorization_type = "CUSTOM"
}

resource "aws_apigatewayv2_integration" "mode_manager" {
  api_id           = aws_apigatewayv2_api.mainGW.id
  integration_type = "AWS_PROXY"

  connection_type        = "INTERNET"
  description            = " requests"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.mode_manager.invoke_arn
  passthrough_behavior   = "WHEN_NO_MATCH"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "mode_manager" {
  api_id    = aws_apigatewayv2_api.mainGW.id
  route_key = "PUT /admin/mode/{proxy+}"

  target             = "integrations/${aws_apigatewayv2_integration.mode_manager.id}"
  authorizer_id      = aws_apigatewayv2_authorizer.authorizer.id
  authorization_type = "CUSTOM"
}

resource "aws_apigatewayv2_integration" "stat_handler" {
  api_id           = aws_apigatewayv2_api.mainGW.id
  integration_type = "AWS_PROXY"

  connection_type        = "INTERNET"
  description            = " requests"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.stat_handler.invoke_arn
  passthrough_behavior   = "WHEN_NO_MATCH"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "stat_handler" {
  api_id    = aws_apigatewayv2_api.mainGW.id
  route_key = "GET /admin/statistic"

  target             = "integrations/${aws_apigatewayv2_integration.stat_handler.id}"
  authorizer_id      = aws_apigatewayv2_authorizer.authorizer.id
  authorization_type = "CUSTOM"
}
