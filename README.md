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
Use this endpoint in order to fetch the full list of customers available in the database.  

__GET__ `http://localhost:3000/customers/{id}`  
Use this endpoint in order to fetch a specific customer from the list of customers in the database.
The `{id}` is the unique ID assigned to the customer entry.  

__POST__ `http://localhost:3000/customers`  
Use this endpoint with a POST request in order to add a customer to the list of customers in the database.  
The required format of the payload to send:  
```
{
    "name": <string>,
    "role": <string>,
    "email": <string>,
    "phone": <string>,
    "contacted": <bool>
}
```
If successful, it will return a status `201` and the full list of customers.  

__PUT__ `http://localhost:3000/customers/{id}`  
Use this endpoint with a PUT request in order to update a customer entry which already exists in the database.  
The `{id}` is the unique ID assigned to the customer entry.  
The required format of the payload to send:  
```
{
    "name": <string>,
    "role": <string>,
    "email": <string>,
    "phone": <string>,
    "contacted": <bool>
}
```

__DELETE__ `http://localhost:3000/customers/{id}`  
USe this endpoint with a DELETE request in order to remove a customer entry which already exists in the database.  
The `{id}` is the unique ID assigned to the customer entry.  

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
-d '{"name":"Edwards","role":"Director of IT","email":"edwards@it.com","phone":"+11234567788","contacted":false}'
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
-d '{"name":"Edwards","role":"Director of IT","email":"edwards@it.com","phone":"+11234567788","contacted":false}'
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