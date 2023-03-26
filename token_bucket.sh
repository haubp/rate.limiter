#!/bin/sh

echo "Execute algorithm Token Bucket"

redis-cli

set counter 0 EX 60
