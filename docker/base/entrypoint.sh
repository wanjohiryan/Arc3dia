#!/bin/bash

#Create directories for selkies-js-interposer

export DISPLAY="${DISPLAY:-:0}"

# Configure joystick interposer
sudo mkdir -pm755 /dev/input
sudo touch /dev/input/{js0,js1,js2,js3}

#start systemd
source /sbin/init --log-level=err