#!/bin/bash

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
        -re
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