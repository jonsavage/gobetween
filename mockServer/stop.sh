#!/bin/sh
echo "stop.sh"
curl localhost:9000/close
curl localhost:9001/close
curl localhost:9002/close
curl localhost:9003/close
curl localhost:9004/close
curl localhost:9005/close
curl localhost:9006/close
curl localhost:9007/close
curl localhost:9008/close
curl localhost:9009/close
curl localhost:9010/close

kill $(pgrep -f "flask")
