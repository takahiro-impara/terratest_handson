output "main_vpc_id" {
  value = aws_vpc.vpc.id
}

output "main_vpc_cidr" {
  value = aws_vpc.vpc.cidr_block
}

output "main_subnet_cidr" {
  value = aws_subnet.subnet-a.cidr_block
}

output "subnet-a_id" {
  value = aws_subnet.subnet-a.id
}

output "public_ip" {
  value = aws_instance.demo.public_ip
}

output "ec2_tags_name" {
  value = aws_instance.demo.tags.Name
}

output "ec2_id" {
  value = aws_instance.demo.id
}