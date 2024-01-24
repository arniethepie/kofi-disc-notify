## Description
Notification of new kofi subscriber via discord, written in Go

**This is not complete, just committing something for now**

## Usage
- `make` and test using sample Kofi data 
```
curl -X POST 127.0.0.1:8080/kofi -d '{
  "verification_token": "ad5a7655-c913-4725-9ee3-f5d0410ea5af",
  "message_id": "4afbccad-e983-4034-8759-d4c525ee8c48",
  "timestamp": "2024-01-24T19:19:42Z",
  "type": "Subscription",
  "is_public": true,
  "from_name": "Jo Example",
  "message": "Good luck with the integration!",
  "amount": "3.00",
  "url": "https://ko-fi.com/Home/CoffeeShop?txid=00000000-1111-2222-3333-444444444444",
  "email": "jo.example@example.com",
  "currency": "USD",
  "is_subscription_payment": true,
  "is_first_subscription_payment": true,
  "kofi_transaction_id": "00000000-1111-2222-3333-444444444444",
  "shop_items": null,
  "tier_name": "bronze",
  "shipping": null
}'
```
- `make build` to generate binary, then run it wherever you like

