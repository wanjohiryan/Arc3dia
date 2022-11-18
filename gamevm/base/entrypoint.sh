#!/bin/bash -e

# Use VirtualGL to run wine with OpenGL if the GPU is available, otherwise use barebone wine
if [ -n "$(nvidia-smi --query-gpu=uuid --format=csv | sed -n 2p)" ]; then
  export VGL_DISPLAY="$DISPLAY"
  export VGL_REFRESHRATE="$REFRESH"
  vglrun +wm wine "${APPPATH}" &
else
  wine "${APPPATH}" &
fi

echo "Game Running. Press [Return] to exit."
read