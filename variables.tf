variable "env" {
  type = string
}
variable "aws_assume_role" {
  default = "arn:aws:iam::816691268740:role/atamaplus-terraform"
}

variable "aws_region" {
  default = "us-east-1"
}

variable "aws_region_tokyo" {
  default = "ap-northeast-1"
}

variable "aws_region_osaka" {
  default = "ap-northeast-3"
}

variable "create_tokyo_region" {
  default = false
}

variable "number" {
  description = "sample"
  default = 2
}

variable "heroku_vpc_cidr_blocks_tokyo" {
  default = [
    "10.1.144.0/20",
    "10.1.128.0/20",
    "10.1.16.0/20",
    "10.1.0.0/20"
  ]
}

variable "hoge" {
}
