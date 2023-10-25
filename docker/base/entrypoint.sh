#!/bin/bash

#Create directories for selkies-js-interposer

export DISPLAY="${DISPLAY:-:0}"

WARP_SERVER_HOST="${WARP_HOST:-"localhost"}"
WARP_SERVER_PORT="${WARP_PORT:-4443}"
WARP_ADDRESS="${WARP_ADDRESS:-$WARP_SERVER_HOST:$WARP_SERVER_PORT}"

# Generate a random 16 character name by default. for each container
URL_NAME="${NAME:-$(head /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | head -c 16)}"

#Full server url
export WARP_FULL_URL="${URL:-"https://$WARP_ADDRESS/$URL_NAME"}"

# Configure joystick interposer
sudo mkdir -pm755 /dev/input
sudo touch /dev/input/{js0,js1,js2,js3}

#start systemd
source /sbin/init --log-level=err