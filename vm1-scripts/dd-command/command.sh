#!/bin/sh

# This script generates random logs using the dd command and encodes them with base64.
dd if=/dev/urandom bs=10240 count=102400 | base64 | tee -a /logs/command-base64.log > /dev/null
