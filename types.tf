variable "string" {
  type = string
}

variable "number" {
  type = number
}

variable "bool" {
  type = bool
}

variable "list_string" {
  type = list(string)
}

variable "list_number" {
  type = list(number)
}

variable "list_map_string" {
  type = list(map(string))
}

variable "map_string" {
  type = map(string)
}
