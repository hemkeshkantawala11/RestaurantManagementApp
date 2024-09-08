const express = require('express');
const bodyParser = require('body-parser');
const PDFDocument = require('pdfkit');
const fs = require('fs');
const path = require('path');
const { MongoClient } = require('mongodb');
require('dotenv').config(); // Add this line to load environment variables

const app = express();
app.use(bodyParser.json());

const client = new MongoClient(process.env.MONGO);
let receiptsCollection;

client.connect().then(() => {
    const db = client.db('mvc_db');
    receiptsCollection = db.collection('receipts');
    console.log('Connected to MongoDB');
}).catch(err => console.error(err));

app.post('/api/receipt', (req, res) => {
    const { customer_name, items, total } = req.body;
    console.log("Req body : ", req.body);
    console.log('Generating receipt for:', req.body.customer_name);
    console.log('Items:', req.body.items);
    console.log('Total amount:', req.body.total);
    const doc = new PDFDocument();
    const filePath = path.join(__dirname, `receipt_${Date.now()}.pdf`);
    doc.pipe(fs.createWriteStream(filePath));

    doc.fontSize(25).text('Receipt', { align: 'center' });
    doc.fontSize(16).text(`Customer Name: ${customer_name}`);
    items.forEach(item => {
        if (item.price !== undefined && typeof item.price === 'number') {
            doc.text(`${item.name}: $${item.price.toFixed(2)}`);
        } else {
            doc.text(`${item.name}: Price not available`);
        }
    });
    doc.text(`Total: $${total.toFixed(2)}`);
    doc.end();

    // Store receipt data in MongoDB
    const receiptData = {
        customer_name,
        items,
        total,
        filePath
    };

    receiptsCollection.insertOne(receiptData)
        .then(result => res.json({ status: 'success', filePath }))
        .catch(err => res.status(500).json({ status: 'error', error: err.message }));
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Node.js server running on port ${PORT}`);
});