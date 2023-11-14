curl -H POST localhost:9091/report -vv -d '{
  "docs": {
    "name": "name_for_documents",
    "department": {
      "code": "uuid_code",
      "time": 123123123,
      "employee": {
        "name": "Ivan",
        "surname": "Polich",
        "code": "uuidv4"
      }
    },
    "price": {
      "categoryA": "1.0",
      "categoryB": "2.0",
      "categoryC": "3.0"
    },
    "owner": {
      "uuid": "uuid",
      "secret": "dsfdwr32fd0fdspsod"
    },
    "data": {
      "transaction": {
        "type": "CODE",
        "uuid": "df23erd0sfods0fw",
        "pointCode": "01"
      }
    },
    "delivery": {
      "company": "TTC",
      "address": {
        "code": "01",
        "country": "uk",
        "street": "Main avenue",
        "apartment": "1A"
      }
    },
    "goods": [
      {
        "name": "toaster v12",
        "amount": 15,
        "code": "12312reds12313e1"
      }
    ]
  }
}'