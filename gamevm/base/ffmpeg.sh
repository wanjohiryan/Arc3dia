#!/bin/bash

##TODO: use nvenc, x264 or vaapi encoders
# if [ -n "$(nvidia-smi --query-gpu=uuid --format=csv | sed -n 2p)" ]; then
    ffmpeg -r 60 -f x11grab -draw_mouse 0 -s 1920x1080 -i :99 -deadline realtime -quality realtime\ #get video from xvfb
        -f pulse -re -i default \ #get audio from pulseaudio source
        -f dash -ldash 1 \ #convert to mpeg-dash
    	-c:v libx264 \
        # -c:v h264_nvenc \
        -preset veryfast -tune zerolatency \
    	-c:a aac \
    	-b:a 128k -ac 2 -ar 44100 \
        #gotten from https://stream.twitch.tv/encoding/
    	-map v:0 -s:v:0 1920x1080 -r 60 -b:v:0 6M   \
    	-map v:0 -s:v:1 1280x720 -r 60 -b:v:1 4.5M   \
    	# -map v:0 -s:v:2 854x480 -r 60 -b:v:2 1.1M \ #be mindful of storage :)
    	# -map v:0 -s:v:3 640x360 -r 60 -b:v:3 365k \
    	-map 0:a \
    	-force_key_frames "expr:gte(t,n_forced*2)" \ #every 2 seconds, i think?
    	-sc_threshold 0 \
    	-streaming 1 \
    	-use_timeline 0 \
    	-seg_duration 2 -frag_duration 0.01 \
    	-frag_type duration \
    	/qwantify/media/playlist.mpd
