#!/bin/bash

for i in {0..2}; do 
    node ./example/index.js "818$i" &
    echo "$!"
done
