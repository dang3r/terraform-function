output "test_zip" {
  value = "${function_zip.f1.return}"
}

output "test_map" {
  value = "${function_map_from_list.f2.return}"
}

output "test_join" {
  value = "${template_file.foo.rendered}"
}
