resource "aws_security_group" "secgroup" {
  name        = "sec-demo"
  description = "sec-demo"
  vpc_id      = aws_vpc.vpc.id
  ingress = [{
    cidr_blocks      = ["159.28.152.155/32"]
    description      = "allow_ssh"
    from_port        = 22
    ipv6_cidr_blocks = []
    prefix_list_ids  = []
    protocol         = "tcp"
    security_groups  = []
    self             = false
    to_port          = 22
    },
    {
      cidr_blocks      = ["0.0.0.0/0"]
      description      = "allow_http"
      from_port        = 80
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
      to_port          = 80
  }]
  egress = [{
    from_port        = 80
    to_port          = 80
    protocol         = "tcp"
    description      = "allow_all"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = []
    prefix_list_ids  = []
    security_groups  = []
    self             = false
  }]
}