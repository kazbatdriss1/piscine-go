#!/bin/bash

(ls -l | sed -e '1,1d')| awk 'FNR%2'