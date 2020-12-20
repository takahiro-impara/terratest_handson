output "main_vpc_id" {
  value = aws_vpc.vpc.id
}

output "subnet-a_id" {
  value = aws_subnet.subnet-a.id
}

output "public_ip" {
  value = aws_instance.demo.public_ip
}