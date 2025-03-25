resource "aws_apigatewayv2_api" "mainGW" {
  name          = "venue-reservation-http"
  description   = "Main API gateway for the project"
  protocol_type = "HTTP"
  cors_configuration {
    allow_origins = ["*"]
    allow_methods = ["*"]
    allow_headers = ["*"]
  }
}

resource "aws_apigatewayv2_authorizer" "authorizer" {
  api_id                            = aws_apigatewayv2_api.mainGW.id
  name                              = "authorizer"
  authorizer_type                   = "REQUEST"
  identity_sources                  = ["$request.header.auth_token"]
  authorizer_uri                    = aws_lambda_function.authorizer.invoke_arn
  authorizer_payload_format_version = "2.0"
  enable_simple_responses           = true
  authorizer_result_ttl_in_seconds  = 0
}

resource "aws_apigatewayv2_stage" "dev" {
  api_id      = aws_apigatewayv2_api.mainGW.id
  description = "stage for development environment"
  name        = "dev"
  auto_deploy = true
  access_log_settings {
    destination_arn = "arn:aws:logs:ap-northeast-2:796973485724:log-group:/api-gateway/venue-reservation-http/dev"
    format = jsonencode({
      httpMethod : "$context.httpMethod"
      ip : "$context.identity.sourceIp"
      protocol : "$context.protocol"
      requestId : "$context.requestId"
      requestTime : "$context.requestTime"
      responseLatency : "$context.responseLatency"
      responseLength : "$context.responseLength"
      routeKey : "$context.routeKey"
      status : "$context.status"
    })
  }
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

resource "aws_iam_policy" "cloudwatch_write" {
  name        = "CloudWatchWritePolicy"
  description = "Allows Lambda to write logs to CloudWatch"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "logs:CreateLogGroup"
        Resource = "*"
      },
      {
        Effect   = "Allow"
        Action   = "logs:CreateLogStream"
        Resource = "*"
      },
      {
        Effect   = "Allow"
        Action   = "logs:PutLogEvents"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role" "role" {
  name               = "myrole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role" "sqs_lambda_role" {
  name               = "sqsLambdaRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role" "sqs_lambda_poll_role" {
  name               = "sqsLambdaPollRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_policy" "sqs_send" {
  name        = "LambdaSQSSendMessagePolicy"
  description = "Allows Lambda to send messages to SQS"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "sqs:SendMessage"
        Resource = aws_sqs_queue.reservation_queue.arn
      }
    ]
  })
}

resource "aws_iam_policy" "sqs_poll" {
  name        = "LambdaSQSPollPolicy"
  description = "Allows Lambda to poll messages from SQS"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "sqs:ReceiveMessage",
          "sqs:DeleteMessage",
          "sqs:GetQueueAttributes"
        ]
        Resource = aws_sqs_queue.reservation_queue.arn
      }
    ]
  })
}

resource "aws_iam_policy" "dynamodb_full_access" {
  name        = "DynamoDBFullAccessPolicy"
  description = "Allows all DynamoDB actions"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "dynamodb:*"
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_sqs_poll_attach" {
  role       = aws_iam_role.sqs_lambda_poll_role.name
  policy_arn = aws_iam_policy.sqs_poll.arn
}

resource "aws_iam_role_policy_attachment" "lambda_dynamodb_attach" {
  role       = aws_iam_role.sqs_lambda_poll_role.name
  policy_arn = aws_iam_policy.dynamodb_full_access.arn
}

resource "aws_iam_role_policy_attachment" "lambda_sqs_attach" {
  role       = aws_iam_role.sqs_lambda_role.name
  policy_arn = aws_iam_policy.sqs_send.arn
}

resource "aws_iam_role_policy_attachment" "lambda_user_dynamodb" {
  role       = aws_iam_role.sqs_lambda_role.name
  policy_arn = aws_iam_policy.dynamodb_full_access.arn
}

locals {
  roles = [
    aws_iam_role.role.name,
    aws_iam_role.sqs_lambda_role.name,
    aws_iam_role.sqs_lambda_poll_role.name
  ]
}

resource "aws_iam_role_policy_attachment" "lambda_cloudwatch_attach" {
  for_each   = toset(local.roles)
  role       = each.value
  policy_arn = aws_iam_policy.cloudwatch_write.arn
}
