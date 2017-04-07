#!/bin/bash

kbytes=1024
echo "scale = 4 ; $kbytes * 1024 / 1000000000" | bc | sed 's/^\./0./'
