
terraform {
}

data "template_file" "example" {
  template = var.example
}

resource "local_file" "example" {
  content  = data.template_file.example.rendered
  filename = "example.txt"
}
