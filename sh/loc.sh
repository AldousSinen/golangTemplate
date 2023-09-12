#!/bin/bash

port="$(cat .env-LOC | grep -m1 'PORT' | cut -c 8-)"
echo "PORT = $port" >> .env