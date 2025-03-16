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

variable "admin_id" {
  description = "value for admin id"
  type        = string
  sensitive   = true
}

variable "admin_passwd" {
  description = "value for admin password"
  type        = string
  sensitive   = true
}