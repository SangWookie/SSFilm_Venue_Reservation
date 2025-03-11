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
  billing_mode   = "PAY_PER_REQUEST"
  stream_enabled = false

  attribute {
    name = "venueDate"
    type = "S"
  }
}

resource "aws_dynamodb_table" "venue_info" {
  name           = "venue_info"
  hash_key       = "venue"
  billing_mode   = "PAY_PER_REQUEST"
  stream_enabled = false

  attribute {
    name = "venue"
    type = "S"
  }
}

variable "venues" {
  type = list(object({
    venue       = string
    venueKor    = string
    allowPolicy = string
  }))
  default = [
    { venue = "studio", venueKor = "스튜디오", allowPolicy = "auto" },
    { venue = "mastering1", venueKor = "마스터링룸 1", allowPolicy = "auto" },
    { venue = "mastering2", venueKor = "마스터링룸 2", allowPolicy = "auto" },
    { venue = "meeting", venueKor = "회의실", allowPolicy = "auto" },
    { venue = "mixing", venueKor = "믹싱룸/ADR룸", allowPolicy = "auto" },
    { venue = "editing", venueKor = "편집실", allowPolicy = "manual" },
    { venue = "lounge", venueKor = "과방", allowPolicy = "manual" },
  ]
}

resource "aws_dynamodb_table_item" "venue_init" {
  for_each   = { for item in var.venues : item.venue => item }
  table_name = aws_dynamodb_table.venue_info.name
  hash_key   = "venue"

  item = jsonencode({
    "venue"       = { "S" = each.value.venue }
    "venueKor"    = { "S" = each.value.venueKor }
    "allowPolicy" = { "S" = each.value.allowPolicy }
  })
}
