#!/bin/bash
# start the server in a chroot.

nohup env GIN_MODE=release chroot --userspec=1000:1000 /var/chroot bin/wcserver
