rm *.xlsx

export FLASK_APP=mockServer.py
flask run -h localhost -p 9000 &
flask run -h localhost -p 9001 &
flask run -h localhost -p 9002 &
flask run -h localhost -p 9003 &
