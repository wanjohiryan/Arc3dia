#!/bin/bash -e
 #Add VirtualGL directories to path
export PATH="${PATH}:/opt/VirtualGL/bin"

# Use VirtualGL to run wine with OpenGL if the GPU is available, otherwise use barebone wine
if [ -n "$(nvidia-smi --query-gpu=uuid --format=csv | sed -n 2p)" ]; then
  export VGL_DISPLAY="${VGL_DISPLAY:-egl}"
  export VGL_REFRESHRATE="$REFRESH"
  cd "${APPPATH}" && vglrun +wm wine "${APPFILE}" 
else
  cd "${APPPATH}" && wine "${APPFILE}"
fi
