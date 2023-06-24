#!/bin/bash
source /opt/bash-utils/logger.sh

INFO "Starting arc3dia server" #FIXME: this is for testing purposes only #-cert /certs/cert.crt -key /certs/private.key -dash /media/playlist.mpd
/usr/local/bin/arc3dia -key /certs/key.pem -cert /certs/cert.pem >> /dev/null 2>&1 &

wait -n

jobs -p | xargs --no-run-if-empty kill
wait

exit