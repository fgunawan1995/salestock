# salestock
Toko Ijah API

Run : go run config.go controller.go main.go migrate.go model.go

API Endpoints : 
1. Item: 
    GET http://127.0.0.1:8080/api/v1/items
    GET http://127.0.0.1:8080/api/v1/items/:id
    POST http://127.0.0.1:8080/api/v1/items
    PUT http://127.0.0.1:8080/api/v1/items/:id
    DELETE http://127.0.0.1:8080/api/v1/items/:id
2. Item In: 
    GET http://127.0.0.1:8080/api/v1/itemsin
    GET http://127.0.0.1:8080/api/v1/itemsin/:id
    POST http://127.0.0.1:8080/api/v1/itemsin
    PUT http://127.0.0.1:8080/api/v1/itemsin/:id
    DELETE http://127.0.0.1:8080/api/v1/itemsin/:id
3. Item Out: 
    GET http://127.0.0.1:8080/api/v1/itemsout
    GET http://127.0.0.1:8080/api/v1/itemsout/:id
    POST http://127.0.0.1:8080/api/v1/itemsout
    PUT http://127.0.0.1:8080/api/v1/itemsout/:id
    DELETE http://127.0.0.1:8080/api/v1/itemsout/:id
4. Report:
    GET http://127.0.0.1:8080/api/v1/itemsreport
    GET http://127.0.0.1:8080/api/v1/salesreport
5. Export to CSV:
    GET http://127.0.0.1:8080/api/v1/items
    GET http://127.0.0.1:8080/api/v1/csvitemsin
    GET http://127.0.0.1:8080/api/v1/csvitemsout
    GET http://127.0.0.1:8080/api/v1/csvitemsreport
    GET http://127.0.0.1:8080/api/v1/csvsalesreport
