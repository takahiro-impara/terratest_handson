resource "aws_vpc" "vpc" {
  cidr_block = "10.10.0.0/16"
  tags = {
    "Name" = "TF-test-git"
  }
}

resource "aws_subnet" "subnet-a" {
  vpc_id     = aws_vpc.vpc.id
  cidr_block = "10.10.1.0/24"
  tags = {
    Name = "tf-subnet-a"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.vpc.id
  tags = {
    Name = "tf-igw"
  }
}

resource "aws_route_table" "rtb" {
  vpc_id = aws_vpc.vpc.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }
  tags = {
    "Name" = "tf-rtb"
  }
}

resource "aws_route_table_association" "rtb_ass" {
  subnet_id      = aws_subnet.subnet-a.id
  route_table_id = aws_route_table.rtb.id
}