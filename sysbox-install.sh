#!/bin/bash

#Make the amd gpu available inside the container
#chmod -R 777 /dev/dri

#
#This script is for anyone who has trouble installing sybox on their system
#

#This works with AWS-linux
sudo yum install jq git -y

#Install docker
sudo amazon-linux-extras install docker -y
sudo service docker start
sudo usermod -a -G docker ec2-user

#make docker autostart
sudo chkconfig docker on

#reboot if necessary
sudo reboot

#clone the sybox repo
git clone --recursive https://github.com/nestybox/sysbox.git

#run make
cd sysbox && sudo make sysbox-static

#then install the built packages
sudo make install

#start the sysbox systemd service
sudo ./scr/sysbox

#configure docker to use sysbox as a runtime !jq must be installed for this to work
sudo ./scr/docker-cfg --sysbox-runtime=enable

sudo cat /etc/docker/daemon.json

# if everything was correcctly installed you should see, otherwise this will be empty
#{
#   "runtimes": {
#      "sysbox-runc": {
#          "path": "/usr/bin/sysbox-runc"
#      }
#  }
#}