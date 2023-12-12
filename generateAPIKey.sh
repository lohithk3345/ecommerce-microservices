#!/bin/bash

command="openssl rand -base64 32"
result=$(eval "$command")

echo $result

