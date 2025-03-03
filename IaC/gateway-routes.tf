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

resource "aws_apigatewayv2_route" "reservation" {
  api_id    = aws_apigatewayv2_api.mainGW.id
  route_key = "ANY /reservations"

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
