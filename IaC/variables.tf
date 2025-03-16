variable "myregion" {
  type    = string
  default = "ap-northeast-2"
}

variable "accountId" {
  type    = string
  default = "796973485724"
}

variable "token_key" {
  description = "value for token key"
  type        = string
  sensitive   = true
}
