#!/bin/bash
sudo yum install httpd -y
sudo service httpd start
sudo touch /var/www/html/index.html
sudo chmod -R 777 /var/www/html/index.html
sudo echo "Hello World">/var/www/html/index.html