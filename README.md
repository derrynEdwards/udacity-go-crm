Udacity Go Nanodegree Project - CRM
===================================

Overview
--------

Simple REST API Built in Go that acts as a CRM Backend.

Running
-------
1. Clone the Repository  
`git clone https://github.com/derrynEdwards/udacity-go-crm`  
2. Run the code
`go run *.go`  
3. Test endpoints on `http://localhost:3000`  

Available Endpoints
-------------------
__GET__ `http://localhost:3000/customers`  
__GET__ `http://localhost:3000/customers/{id}`  
__POST__ `http://localhost:3000/customers`  
__PUT__ `http://localhost:3000/customers/{id}`  
__DELETE__ `http://localhost:3000/customers/{id}`  

Sample Endpoints
----------------
__GET__ `http://localhost:3000/customers`  

Run:  
```
curl http://localhost:3000/customers
```  
Response:  
```
{
    "1": {
        "name": "John",
        "role": "CEO",
        "email": "john@udatech.com",
        "phone": "+11234567890",
        "contacted": false
    },
    "2": {
        "name": "Jane",
        "role": "COO",
        "email": "jane@udatech.com",
        "phone": "+11234567908",
        "contacted": true
    },
    "3": {
        "name": "Carl",
        "role": "CEO",
        "email": "carl@mycompany.com",
        "phone": "+11234567809",
        "contacted": true
    },
    "4": {
        "name": "Kane",
        "role": "CTO",
        "email": "kane@kanecorp.com",
        "phone": "+11234567788",
        "contacted": false
    }
}
```

__GET__ `http://localhost:3000/customers/{id}`  

Run:  
```
curl http://localhost:3000/customers/1
```  
Response:  
```
{
    "name": "John",
    "role": "CEO",
    "email": "john@udatech.com",
    "phone": "+11234567890",
    "contacted": false
}
```  

__POST__ `http://localhost:3000/customers` 

Run:  
```
curl -X POST http://localhost:3000/customers \
-H 'Content-Type: application/json' \
-d '{"5": {"name":"Edwards","role":"Director of IT","email":"edwards@it.com","phone":"+11234567788","contacted":false}}'
```  
Response: 
```
{
    "1": {
        "name": "John",
        "role": "CEO",
        "email": "john@udatech.com",
        "phone": "+11234567890",
        "contacted": false
    },
    "2": {
        "name": "Jane",
        "role": "COO",
        "email": "jane@udatech.com",
        "phone": "+11234567908",
        "contacted": true
    },
    "3": {
        "name": "Carl",
        "role": "CEO",
        "email": "carl@mycompany.com",
        "phone": "+11234567809",
        "contacted": true
    },
    "4": {
        "name": "Kane",
        "role": "CTO",
        "email": "kane@kanecorp.com",
        "phone": "+11234567788",
        "contacted": false
    },
    "5": {
        "name": "Edwards",
        "role": "Director of IT",
        "email": "edwards@it.com",
        "phone": "+11234567788",
        "contacted": false
    }
}
``` 

__PUT__ `http://localhost:3000/customers/{id}`  

Run:  
```
curl -X PUT http://localhost:3000/customers/4 \
-H 'Content-Type: application/json' \
-d '{"4": {"name":"Edwards","role":"Director of IT","email":"edwards@it.com","phone":"+11234567788","contacted":false}}'
```  
Response: 
```
{
    "1": {
        "name": "John",
        "role": "CEO",
        "email": "john@udatech.com",
        "phone": "+11234567890",
        "contacted": false
    },
    "2": {
        "name": "Jane",
        "role": "COO",
        "email": "jane@udatech.com",
        "phone": "+11234567908",
        "contacted": true
    },
    "3": {
        "name": "Carl",
        "role": "CEO",
        "email": "carl@mycompany.com",
        "phone": "+11234567809",
        "contacted": true
    },
    "4": {
        "name": "Edwards",
        "role": "Director of IT",
        "email": "edwards@it.com",
        "phone": "+11234567788",
        "contacted": false
    }
}
```

__DELETE__ `http://localhost:3000/customers/1` 

Run:  
```
curl -X DELETE http://localhost:3000/customers/1
```  
Response:
```
{
    "2": {
        "name": "Jane",
        "role": "COO",
        "email": "jane@udatech.com",
        "phone": "+11234567908",
        "contacted": true
    },
    "3": {
        "name": "Carl",
        "role": "CEO",
        "email": "carl@mycompany.com",
        "phone": "+11234567809",
        "contacted": true
    },
    "4": {
        "name": "Kane",
        "role": "CTO",
        "email": "kane@kanecorp.com",
        "phone": "+11234567788",
        "contacted": false
    }
}
```