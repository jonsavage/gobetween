#!/bin/sh

rm *.xlsx

export FLASK_APP=mockServer.py
flask run -h localhost -p 9000 &
flask run -h localhost -p 9001 &
# flask run -h localhost -p 9002 &
# flask run -h localhost -p 9003 &
# flask run -h localhost -p 9004 &
# flask run -h localhost -p 9005 &
# flask run -h localhost -p 9006 &
# flask run -h localhost -p 9007 &
# flask run -h localhost -p 9008 &
# flask run -h localhost -p 9009 &
# flask run -h localhost -p 9010 &

sleep 1

./loadCreator.sh

./stop.sh
