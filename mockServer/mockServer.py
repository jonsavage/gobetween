from flask import Flask, jsonify

app = Flask(__name__)

load = 0

def increment_load():
  global load
  load += .1

def get_current_load():
  global load
  return load

@app.route("/")
def get_content():
  increment_load()
  print("content fetched")
  return "some content"

@app.route("/load", methods=['GET'])
def get_load():
  current_load = get_current_load()
  print("load fetched, current load:" + str(current_load))

  return jsonify({'load': current_load})
