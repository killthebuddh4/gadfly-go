#!/bin/sh

for f in examples.*.fly; do
    echo; echo; echo;
    echo; echo; echo;
    echo "Running $f"
    echo; echo; echo;
    go run . "$f"
done