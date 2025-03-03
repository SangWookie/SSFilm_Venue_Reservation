resource "aws_apigatewayv2_api" "mainGW" {
  name          = "venue-reservation-http"
  description   = "Main API gateway for the project"
  protocol_type = "HTTP"
  cors_configuration {
    allow_origins = ["*"]
    allow_methods = ["*"]
  }
}

resource "aws_apigatewayv2_stage" "dev" {
  api_id      = aws_apigatewayv2_api.mainGW.id
  description = "stage for development environment"
  name        = "dev"
  auto_deploy = true
}

# IAM
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "role" {
  name               = "myrole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}
