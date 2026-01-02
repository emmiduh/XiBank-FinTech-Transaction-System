from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route("/health", methods=["GET"])
def health():
    """Health check endpoint."""
    return jsonify({"status": "ok"}), 200

@app.route("/checkfraud", methods=["POST"])
def check_fraud():
    """Fraud check endpoint."""
    try:
        data = request.get_json(force=True)
        amount = float(data.get("amount", 0))
    except (ValueError, TypeError):
        return jsonify({"error": "Invalid request data"}), 400

    if amount > 5000:
        return jsonify({
            "fraud": True,
            "rule": "amount_threshold",
            "threshold": 5000
        }), 200

    return jsonify({"fraud": False}), 200

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)

