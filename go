#! /usr/bin/env bash

ROLL=$((1 + RANDOM % 10))  # ROLL \in [1, 10]

if [ "$ROLL" -eq "1" ]; then
    cat ~/.tmp/.help
else
    `cat ~/.tmp/.cache` "$@"
fi
