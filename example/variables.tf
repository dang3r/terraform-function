variable "list1" {
  type    = "list"
  default = ["Toronto", "San-Jose", "Boston"]
}

variable "list2" {
  type    = "list"
  default = ["Maple-Leafs", "Sharks", "Bruins"]
}

variable "map1" {
  type = "map"
  default = {
    "Justin" = "Trudeau"
    "George" = "Bush"
    "Vladimir" = "Putin"
  }
}

variable "list3" {
  type = "list"

  default = [
    "George",
    "Washington",
    "Albert",
    "Einstein",
    "Ed",
    "Witten",
    "Kurt",
    "Godel",
  ]
}
