resource "aws_dynamodb_table" "current_reservation" {
  name           = "current_reservation"
  hash_key       = "reservationId"
  billing_mode   = "PAY_PER_REQUEST"
  stream_enabled = false

  attribute {
    name = "reservationId"
    type = "S"
  }

  attribute {
    name = "venueDate"
    type = "S"
  }

  global_secondary_index {
    name = "venueDate-index"
    hash_key = "venueDate"
    projection_type = "ALL"
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
    sortOrder   = number
  }))
  default = [
    { venue = "editing", venueKor = "편집실", allowPolicy = "manual", sortOrder = 7 },
    { venue = "lounge", venueKor = "과방", allowPolicy = "manual", sortOrder = 6 },
    { venue = "mixing", venueKor = "믹싱룸/ADR룸", allowPolicy = "auto", sortOrder = 5 },
    { venue = "meeting", venueKor = "회의실", allowPolicy = "auto", sortOrder = 4 },
    { venue = "studio", venueKor = "스튜디오", allowPolicy = "auto", sortOrder = 3 },
    { venue = "mastering2", venueKor = "마스터링룸 2", allowPolicy = "auto" , sortOrder = 2 },
    { venue = "mastering1", venueKor = "마스터링룸 1", allowPolicy = "auto", sortOrder = 1 },
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
    "sortOrder"   = { "N" = tostring(each.value.sortOrder) }
  })
}
