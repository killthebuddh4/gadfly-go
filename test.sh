#!/bin/sh

for f in tests/*.fly; do
    echo; echo; echo;
    go run . "$f"
done