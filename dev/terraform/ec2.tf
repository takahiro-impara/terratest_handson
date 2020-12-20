
resource "aws_instance" "demo" {
  ami                         = "ami-cbf90ecb"
  instance_type               = "t2.micro"
  subnet_id                   = aws_subnet.subnet-a.id
  associate_public_ip_address = true
  security_groups             = [aws_security_group.secgroup.id]
  key_name = var.key_name
  root_block_device {
    volume_type = "gp2"
    volume_size = "8"
  }
  ebs_block_device {
    device_name = "/dev/sdf"
    volume_type = "gp2"
    volume_size = 8
  }

  user_data = <<EOF
IyEvYmluL2Jhc2gKc3VkbyB5dW0gaW5zdGFsbCBodHRwZCAteQpzdWRvIHNlcnZp
Y2UgaHR0cGQgc3RhcnQ=
EOF

  tags = {
    "Name" = "tf-test"
  }
}