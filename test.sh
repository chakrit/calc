#!/bin/sh

FILES="chars.go context.go main.go parse.go stack.go queue.go lex.go compile.go nodes.go tokens.go"
cat input.txt | go run $FILES | diff - output.txt

