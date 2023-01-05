#!/bin/bash

# if no hwenc required, noop
[[ -z "$HWENC" ]] && exit 0

if [[ -z "$RENDER_GID" ]]; then
  RENDER_GID=$(stat -c "%g" /dev/dri/render* | tail -n 1)
  # is /dev/dri passed to the container?
  [[ -z "$RENDER_GID" ]] && exit 1
fi

# note that this could conceivably be a security risk...
cnt_group=$(getent group "$RENDER_GID" | cut -d: -f1)
if [[ -z "$cnt_group" ]]; then
  groupadd -g "$RENDER_GID" hwrender
  cnt_group=hwrender
fi
usermod -a -G "$cnt_group" "$USER"
