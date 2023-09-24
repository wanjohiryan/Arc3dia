#!/bin/bash

#Start XServer and Pulseaudio from here? or systemd will do that for us?
set -xeuo pipefail

#Fallback values
WARP_SERVER_URL="${WARP_SERVER_URL:-"https://localhost"}"
WARP_SERVER_PORT="${WARP_SERVER_PORT:-4443}"

#Full server url
WARP_SERVER_FULL_URL="${WARP_SERVER_URL}:${WARP_SERVER_PORT}"

#Start dbus for pulseaudio
/etc/init.d/dbus start

#Change time zone from environment variable
ln -snf "/usr/share/zoneinfo/$TZ" /etc/localtime && echo "$TZ" | sudo tee /etc/timezone >/dev/null
# Default display is :0 across the container
export DISPLAY=":0"
# Run Xvfb server with required extensions
/usr/bin/Xvfb "${DISPLAY}" -screen 0 "8192x4096x${CDEPTH}" -dpi "${DPI}" &

# Wait for X11 to start
echo "Waiting for X socket"
until [ -S "/tmp/.X11-unix/X${DISPLAY/:/}" ]; do sleep 1; done
echo "X socket is ready"

pulseaudio --system --log-level=info --disallow-module-loading --disallow-exit --exit-idle-time=-1 &
sleep 20

CMD=(
    ffmpeg
    -hide_banner
    -loglevel error
    #screen image size
    -s 1920x1080
    #video fps
    -r 30
    #grab x11 display
    -f x11grab
        -i ${DISPLAY}
    #capture pulse audio
    -f pulse
        -ac 2
        -i default
    -c:v libx264 
        -preset veryfast
        -tune zerolatency
        -profile main
        -pix_fmt yuv420p #let us use 4:2:0 for now 
        #FIXME: add full color 4:4:4
    -c:a aac
        -b:a 128k
        -ar 44100
        -ac 2
    -map v:0 -s:v:0 1280x720 -b:v:0 3M
    #FIXME: add this later
    #-map v:0 -s:v:1 854x480  -b:v:1 1.1M
    #-map v:0 -s:v:2 640x360  -b:v:2 365k
    -map 1:a
    -streaming 1
    -movflags empty_moov+frag_every_frame+separate_moof+omit_tfhd_offset

    ./playlist.mpd
    #Then send this to our remote/local warp-relay server
    #- | RUST_LOG=warp=info /usr/bin/warp -i - -u $WARP_SERVER_FULL_URL
)

exec "${CMD[@]}"