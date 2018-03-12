#!/bin/bash

kill -9 $(pgrep testweb)
cd ~/newweb/
git pull https://github.com/xiejin740640431/testweb.git
cd testweb
./testweb