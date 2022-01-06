# RESTful API Data Person

Ini adalah RESTful API untuk menyimpan data seseorang yaitu first_name, last_name, dan age. <br>
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

## Get All Data
### Request
`GET localhost:8080/persons`

     Accept: application/json
     
### Response 

     
