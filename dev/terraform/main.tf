provider "aws" {
  region     = "ap-northeast-1"
  //access_key = var.aws_key
  //secret_key = var.aws_secret
}

terraform {
  required_version = "~> 0.14.0"
}

resource "aws_vpc" "vpc" {
  cidr_block = "10.10.0.0/16"
  tags = {
    "Name" = "TF-test-git"
  }
}
output "TestVal" {
  value = "test value!"
}

output "main_vpc_id" {
  value  = aws_vpc.vpc.id
}