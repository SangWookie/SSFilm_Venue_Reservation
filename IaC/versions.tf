# Terraform Block
terraform {
  required_version = ">= 1.3"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.37"
    }
    local = {
      source  = "hashicorp/local"
      version = ">= 2.5"
    }
    null = {
      source  = "hashicorp/null"
      version = ">= 2.0"
    }
    tls = {
      source  = "hashicorp/tls"
      version = ">= 3.1"
    }
  }

  backend "s3" {
    bucket = "venue-reservation-terraform"
    key    = "dev/terraform.tfstate"
    region = "ap-northeast-2"

    # State Locking
    dynamodb_table = "venue_reservation_locking"
  }
}

# Provider Block
provider "aws" {
  region  = "ap-northeast-2"
  # profile = "default"
}
