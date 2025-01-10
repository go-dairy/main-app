#!/bin/sh

# wait for postgres to start
sleep 5

# run migrations
./migrator

# run the app
./go-dairy

# exit
exit 0