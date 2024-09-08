from flask import Flask, request, jsonify
from pymongo import MongoClient
from bson import ObjectId
import os
from dotenv import load_dotenv


load_dotenv()

app = Flask(__name__)

mongo_url = os.getenv('MONGO')

client = MongoClient(mongo_url)
db = client['mvc_db']
orders_collection = db['orders']

@app.route('/api/order', methods=['POST'])
def process_order():
    order_data = request.json
    order_data['total'] = sum(item['price'] for item in order_data['items'])
    result = orders_collection.insert_one(order_data)
    order_data['_id'] = str(result.inserted_id)  # Convert ObjectId to string
    return jsonify({"status": "success", "order": order_data}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)