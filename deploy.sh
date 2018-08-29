#!/bin/sh

scp -i ~/.ssh/davegarred.pem wine_image.tar  ubuntu@ec2-18-236-228-216.us-west-2.compute.amazonaws.com:.
ssh -i ~/.ssh/davegarred.pem ubuntu@ec2-18-236-228-216.us-west-2.compute.amazonaws.com <<EOF
  sudo docker stop wine
  sudo docker rm wine
  sudo docker rmi wine
  sudo docker load -i wine_image.tar
  ./wine.sh
  exit
EOF
