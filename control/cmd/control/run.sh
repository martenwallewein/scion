# !/bin/bash
set -e
/root/dispatcher --config $1 & 
/root/control --config $2 