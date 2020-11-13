import os
import paypayopa

api_key = os.environ['API_KEY']
api_secret = os.environ['API_SECRET']
merchant_id = os.environ['MERCHANT_ID']

client = paypayopa.Client(auth=(api_key, api_secret), production_mode=False)
print(client.get_version())

client.set_assume_merchant(merchant_id)

request = {
    "merchantPaymentId": "002",
    "codeType": "ORDER_QR",
    "redirectUrl": "http://foobar.com",
    "redirectType":"WEB_LINK",
    "orderDescription":"Example - Mune Cake shop",
    "orderItems": [{
        "name": "Moon cake",
        "category": "pasteries",
        "quantity": 1,
        "productId": "67678",
        "unitPrice": {
            "amount": 1,
            "currency": "JPY"
        }
    }],
    "amount": {
        "amount": 1,
        "currency": "JPY"
    },
}

resp = client.Code.create_qr_code(request)
print(resp)
