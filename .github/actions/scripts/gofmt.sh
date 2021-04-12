#!/bin/bash

cd $1
ls -p | grep -v vendor | grep "/$" | xargs gofmt -l -s
