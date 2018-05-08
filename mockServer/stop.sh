curl localhost:9000/close
curl localhost:9001/close
curl localhost:9002/close
curl localhost:9003/close

kill $(pgrep -f "flask")
