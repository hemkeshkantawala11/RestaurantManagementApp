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
items_collection = db['items']
@app.route('/api/order', methods=['POST'])
def process_order():
    order_data = request.json
    item_ids = order_data.pop('items', [])

    # Fetch item details from the database
    items = []
    total = 0
    for item_id in item_ids:
        item = items_collection.find_one({"_id": ObjectId(item_id)})
        if item:
            items.append(item)
            total += item.get('price', 0)
        else:
            return jsonify({"status": "error", "message": f"Item not found: {str(item_id)}"}), 400

    result = orders_collection.insert_one(order_data)
    order_data['_id'] = str(result.inserted_id)  # Convert ObjectId to string

    # Convert ObjectId in items to string
    for item in items:
        item['_id'] = str(item['_id'])


    print(order_data)
    print(type(order_data))

    return jsonify({"status": "success", "order": order_data}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)