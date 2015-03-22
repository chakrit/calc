#!/bin/sh

cat input.txt | go run *.go | diff - output.txt

