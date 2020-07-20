#!/bin/sh

if ! test -f "harmony-key.pem" || ! test -f "harmony-key.pub"; then
    ./legato -g
fi