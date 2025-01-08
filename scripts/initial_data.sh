#!/bin/bash

set -x # Aktiviert Debugging

#
## Global Variables
#
HOST="http://api:8000"

#
## Manufacturers
#
PATH=/api/v1/manufacturers

/usr/bin/curl $HOST$PATH -X POST -d '{"name": "eSUN"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Prusa"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "3DFS"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Anycubic"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Bambu"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Sunlu"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "3DJake"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "PrintWithSmile"}'
unset PATH


#
## Materials
#
PATH=/api/v1/materials

/usr/bin/curl $HOST$PATH -X POST -d '{"name": "PLA"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "PLA+"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "ecoPLA"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "PETG"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "rPETG"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "ABS"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "ABS+"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "ASA"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "TPU A95"}'
unset PATH


#
## Colors
#
PATH=/api/v1/colors

/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Weis"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Schwarz"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Rot"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Grün"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Blau"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Violet"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Grau"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Gold"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Cold White"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Orange"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Fire Engine Red"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Light Blue"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Hellgrün"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Transparent"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Transparent Grün"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Transparent Gelb"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Transparent Violet"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Cloudy Grey"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Silver"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Peak Green"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Gelb"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Olive Green"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Solid Black"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Solid White"}'
/usr/bin/curl $HOST$PATH -X POST -d '{"name": "Pink"}'
unset PATH