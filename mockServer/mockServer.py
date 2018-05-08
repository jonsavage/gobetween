from flask import Flask, jsonify, request
import xlsxwriter, time, threading, uuid

app = Flask(__name__)

load = 0

workbook = xlsxwriter.Workbook(str(uuid.uuid4()) + ".xlsx")
worksheet = workbook.add_worksheet("worksheet")
worksheet.write('A1','time')
worksheet.write('B1','load')
worksheet_counter = 1
current_time = time.time()

def increment_worksheet_counter():
  global worksheet_counter
  print("counter:")
  print(worksheet_counter)
  worksheet_counter = worksheet_counter + 1

def increment_load():
  global load
  load += 1

def decrement_load():
  global load
  if (load >= 1):
    load -= 1

def get_current_load():
  global load
  return load

@app.route("/")
def get_content():
  increment_load()
  print("get content")
  return "load: " + str(get_current_load())

@app.route("/load", methods=['GET'])
def get_load():
  current_load = get_current_load()
  print("load fetched, current load:" + str(current_load))

  return jsonify({'load': current_load})

@app.route("/close", methods=['GET'])
def close_workbook():
  chart = workbook.add_chart({"type": "line"})

  chart.add_series({
    'name':str(request.host)[-4:],
    'categories': '=worksheet!$A2:$A$' + str(worksheet_counter),
    'values': '=worksheet!$B2:$B$' + str(worksheet_counter)
  })

  worksheet.write("D1", "average")
  worksheet.write_formula("E1", '=AVERAGE($B2:$B$' + str(worksheet_counter) + ')')

  print("worksheet_counter")
  print(str(worksheet_counter))

  worksheet.insert_chart("D5", chart)

  workbook.close()
  return jsonify({'close': 'close'})

#########################################################################################
def f(f_stop):
  print("f")





  worksheet.write(worksheet_counter,0,time.time() - current_time)
  worksheet.write(worksheet_counter,1,get_current_load())
  # global worksheet_counter
  increment_worksheet_counter()





  if (get_current_load() > 0):
    decrement_load()

  if not f_stop.is_set():
    # call f() again in 60 seconds
    threading.Timer(1, f, [f_stop]).start()

f_stop = threading.Event()
# start calling f now and every 60 sec thereafter
f(f_stop)
