from flask import Flask, jsonify

app = Flask(__name__)

a = 0


@app.route('/get', methods=['GET'])
def root():
    globals()['a'] += 1
    return jsonify(a)


if __name__ == '__main__':
    app.run('0.0.0.0', 60001)
