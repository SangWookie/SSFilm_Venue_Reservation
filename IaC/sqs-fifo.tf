resource "aws_sqs_queue" "reservation_queue" {
  name                        = "reservation-queue.fifo"
  fifo_queue                  = true
  deduplication_scope         = "messageGroup"
  content_based_deduplication = true
}
