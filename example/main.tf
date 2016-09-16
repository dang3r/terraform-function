resource "function_zip" "f1" {
  list1 = "${var.list1}"
  list2 = "${var.list2}"
}

resource "function_map_from_list" "f2" {
  list = "${var.list3}"
}

resource "template_file" "foo" {
  template = "${lol}"

  vars {
    lol = "${join("_",keys("${function_zip.f1.return}"))}"
  }
}
