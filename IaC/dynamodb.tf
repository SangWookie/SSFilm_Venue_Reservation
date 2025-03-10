resource "aws_dynamodb_table" "current_reservation" {
  name           = "current_reservation"
  hash_key       = "venueDate"
  range_key      = "time"
  billing_mode   = "PAY_PER_REQUEST"
  stream_enabled = false

  attribute {
    name = "venueDate"
    type = "S"
  }

  attribute {
    name = "time"
    type = "S"
  }
}

resource "aws_dynamodb_table" "pending_reservation" {
  name           = "pending_reservation"
  hash_key       = "requestId"
  billing_mode   = "PAY_PER_REQUEST"
  stream_enabled = false

  attribute {
    name = "requestId"
    type = "S"
  }
}

resource "aws_dynamodb_table" "reservation_limit" {
  name           = "reservation_limit"
  hash_key       = "venueDate"
  range_key      = "time"
  billing_mode   = "PAY_PER_REQUEST"
  stream_enabled = false

  attribute {
    name = "venueDate"
    type = "S"
  }

  attribute {
    name = "time"
    type = "S"
  }
}