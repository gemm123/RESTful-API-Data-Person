# RESTful API Data Person

Ini adalah RESTful API untuk menyimpan data seseorang yaitu firstName, lastName, dan age. <br>
API ini dibuat dengan bahasa golang dan beberapa dependency <br>
Database yang digunakan adalah MySql <br>

## Setup Dependeny
### MySql
     go get https://github.com/go-sql-driver/mysql 
### HTTP Router
     go get https://github.com/julienschmidt/httprouter
### Validation
     go get https://github.com/go-playground/validator

# RESTful API Example
Berikut adalah contoh request dan response tiap method yang ada pada API

## Get All Data
### Request
`GET localhost:8080/persons`

     Accept: application/json
     X-API-Key: KUNCIMASUK
     
### Response 

     "code": 200,
     "status": "OK",
     "data": []     

## Get Data by Id
### Request
`GET localhost:8080/persons/1`

     Accept: application/json
     X-API-Key: KUNCIMASUK

### Response 

     "code": 200,
     "status": "OK",
     "data": {"id": 1, "firstName": "nama depan", "lastName": "nama belakang", "age": "umur"}

## Create Data
### Request
`POST localhost:8080/persons`

     Accept: application/json
     X-API-Key: KUNCIMASUK
     {"firstName": "nama depan", "lastName": "nama belakang", "age": "umur"}

### Response 

     "code": 200,
     "status": "OK",
     "data": {"id": 1, "firstName": "nama depan", "lastName": "nama belakang", "age": "umur"}

## Update Data
### Request
`PUT localhost:8080/persons/1`

     Accept: application/json
     X-API-Key: KUNCIMASUK
     {"firstName": "nama depan", "lastName": "nama belakang", "age": "umur"}

### Response 

     "code": 200,
     "status": "OK",
     "data": {"id": 1, "firstName": "nama depan", "lastName": "nama belakang", "age": "umur"}
     
## Delete Data
`DELETE localhost:8080/persons/1`

     Accept: application/json
     X-API-Key: KUNCIMASUK
     
### Response 

     "code": 200,
     "status": "OK"
