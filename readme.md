## you will need 
https://github.com/golang-migrate/migrate 
https://github.com/julienschmidt/httprouter 
https://github.com/joho/godotenv 
Mysql driver installed 

## How to start 
1. creat database in mysql
2. migrate go to root and migrate 
```bash
migrate -path . -database mysql://{username}:{password}@tcp({dbhost}:{dbport})/{dbname} up
```
3. create env file based on env-examples
4. run the app 
```bash
go run [-race] main.go
```

## Test
1. download and install [postman](https://www.getpostman.com/downloads/)
2. import file simpleCRUD.postman_collection
3. test all route
