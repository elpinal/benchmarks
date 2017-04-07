#!/bin/bash

kbytes=1024
echo $kbytes | awk '{printf "%.4f\n", $1 * 1024 / 10**9}'
