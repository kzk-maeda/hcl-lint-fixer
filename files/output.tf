variable "env" {
  description = "Env"
  type        = string
}

variable "aws_assume_role" {
  description = "AWS Assume Role"
  type        = string
  default     = "arn:aws:iam::816691268740:role/atamaplus-terraform"
}

variable "aws_region" {
  description = "AWS Region"
  type        = string
  default     = "us-east-1"
}

variable "aws_region_tokyo" {
  description = "AWS Region Tokyo"
  type        = string
  default     = "ap-northeast-1"
}

variable "aws_region_osaka" {
  description = "AWS Region Osaka"
  type        = string
  default     = "ap-northeast-3"
}

variable "create_tokyo_region" {
  description = "Create Tokyo Region"
  type        = bool
  default     = false
}

variable "number" {
  description = "sample"
  type        = number
  default     = 2
}

variable "heroku_vpc_cidr_blocks_tokyo" {
  description = "Heroku VPC CIDR Blocks Tokyo"
  type        = list(string)
  default = [
    "10.1.144.0/20",
    "10.1.128.0/20",
    "10.1.16.0/20",
    "10.1.0.0/20"
  ]
}

variable "hoge" {
  description = "Hoge"
  type        = map(string)
  default     = { "foo" : "bar" }
}

variable "fuga" {
  description = "Fuga"
}

variable "main_db_instances_virginia" {
  description = "Configuration values for individual main DB instances in Virginia region"
  type        = list(map(string))
  default = [{
    instance_class               = "db.t3.large"
    promotion_tier               = 0
    preferred_maintenance_window = "sun:16:01-sun:16:31"
  }]
}

