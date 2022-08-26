from flask import Flask
from flask import make_response
import json

PORT = 8080
app = Flask(__name__)

@app.route('/')
def index():
    return 'hello world'

@app.route('/get-json')
def getjson():
    body = json.dumps({'msg': 'hello world'})
    resp = make_response(body)
    resp.headers['Content-Type'] = 'application/json'
    return resp

@app.route('/get/<id>')
def getId(id):
    msg = 'hello %s' % id
    body = json.dumps({'msg': msg})
    resp = make_response(body)
    resp.headers['Content-Type'] = 'application/json'
    return resp

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=PORT)