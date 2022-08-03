# !/bin/bash
set -e
/root/dispatcher --config $1 & 
/root/daemon --config $2 