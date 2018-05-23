# salestock
Toko Ijah API

Run : go run config.go controller.go main.go migrate.go model.go

API Endpoints : 
1. Item: 

    a. GET http://127.0.0.1:8080/api/v1/items
    
    b. GET http://127.0.0.1:8080/api/v1/items/:id
    
    c. POST http://127.0.0.1:8080/api/v1/items
    d. PUT http://127.0.0.1:8080/api/v1/items/:id
    e. DELETE http://127.0.0.1:8080/api/v1/items/:id
2. Item In: 
    a. GET http://127.0.0.1:8080/api/v1/itemsin
    b. GET http://127.0.0.1:8080/api/v1/itemsin/:id
    c. POST http://127.0.0.1:8080/api/v1/itemsin
    d. PUT http://127.0.0.1:8080/api/v1/itemsin/:id
    e. DELETE http://127.0.0.1:8080/api/v1/itemsin/:id
3. Item Out: 
    a. GET http://127.0.0.1:8080/api/v1/itemsout
    b. GET http://127.0.0.1:8080/api/v1/itemsout/:id
    c. POST http://127.0.0.1:8080/api/v1/itemsout
    d. PUT http://127.0.0.1:8080/api/v1/itemsout/:id
    e. DELETE http://127.0.0.1:8080/api/v1/itemsout/:id
4. Report:
    a. GET http://127.0.0.1:8080/api/v1/itemsreport
    b. GET http://127.0.0.1:8080/api/v1/salesreport
5. Export to CSV:
    a. GET http://127.0.0.1:8080/api/v1/items
    b. GET http://127.0.0.1:8080/api/v1/csvitemsin
    c. GET http://127.0.0.1:8080/api/v1/csvitemsout
    d. GET http://127.0.0.1:8080/api/v1/csvitemsreport
    e. GET http://127.0.0.1:8080/api/v1/csvsalesreport
