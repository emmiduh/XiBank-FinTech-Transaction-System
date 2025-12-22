from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/check', methods=['POST'])
def check_fraud():
    data = request.json
    amount = data.get('amount', 0)

    if amount > 5000:
        return jsonify({"fraud": True}), 200
    
    return jsonify({"fraud": False}), 200

if __name__ == '__main__':
    app.run(port=5000)