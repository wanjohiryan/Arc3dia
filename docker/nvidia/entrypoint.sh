#!/bin/bash
source /opt/bash-utils/logger.sh

INFO "Starting ffmpeg"
/usr/local/bin/ffmpeg.sh &
sleep 10

INFO "Starting arc3dia server" #FIXME: this is for testing purposes only #-cert /certs/cert.crt -key /certs/private.key -dash /media/playlist.mpd
/usr/local/bin/arc3dia -key /certs/key.pem -cert /certs/cert.pem -dash /media/playlist.mpd 2>&1 &
sleep 5

wait -n

jobs -p | xargs --no-run-if-empty kill
wait

exit