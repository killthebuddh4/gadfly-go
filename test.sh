#!/bin/sh

for f in tests/*.fly; do
    echo; echo; echo;
    go run . "$f"
done

for f in examples/*.fly; do
    echo; echo;
    echo "#######################################################"
    echo; echo;
    echo "Running Example: $f"
    echo; echo;
    echo "#######################################################"
    echo; echo;
    go run . "$f"
done