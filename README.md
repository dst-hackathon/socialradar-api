socialradar-api
===============

Social Radar API

Note: The API consumer need to provide the security token in HTTP request header (except the signup and logon API).

The request HTTP header name is "token". Consumer will get token when they successfully called logon API.

Current API
===========

GET /questions
--------------
Response:
```javascript
[
  {
    "id": "1",
    "order": "1",
    "tag": "food",
    "text": "What are your favorite menus?"
  },
  {
    "id": "2",
    "order": "2",
    "tag": "sport",
    "text": "What sports do you play?"
  }
]
```

GET /questions/{id}
------------------
Response:
```javascript
{
  "id": "1",
  "text": "What are your favorite menus?",
  "tag": "food",
  "order": "1",
  "categories": [
    {
      "id": "1",
      "text": "Thai",
      "order": "1",
      "options": [
        {
          "id": "1",
          "order": "1",
          "text": "Pad Thai"
        },
        {
          "id": "2",
          "order": "2",
          "text": "Tom Yum Kung"
        },
        {
          "id": "3",
          "order": "3",
          "text": "Kao Pad"
        }
      ]
    },
    {
      "id": "2",
      "text": "Japanese",
      "order": "2",
      "options": [
        {
          "id": "6",
          "order": "1",
          "text": "Sashimi"
        },
        {
          "id": "4",
          "order": "1",
          "text": "Sushi"
        },
        {
          "id": "5",
          "order": "1",
          "text": "Ramen"
        }
      ]
    },
    {
      "id": "3",
      "text": "Chinese",
      "order": "3",
      "options": []
    },
    {
      "id": "4",
      "text": "Western",
      "order": "4",
      "options": [
        {
          "id": "7",
          "order": "1",
          "text": "T-bone steak"
        },
        {
          "id": "8",
          "order": "2",
          "text": "Fish \u0026 chip"
        },
        {
          "id": "9",
          "order": "3",
          "text": "Irish stew"
        }
      ]
    }
  ]
}
```

POST /users/{id}/answer
----------------------
Request: 
```javascript
{
  "1": {  // Question ID
    "1": [1, 2],  // "Selected Category": [Selected Option ID, ....]
    "2": []
  },
  "2": {
    "5": [10]
  }
}
```

Response:
```javascript
{
	"success": true
}
```

GET /users/{id}/answer
---------------------
Response: 
```javascript
{
  "1": {  // Question ID
    "1": [1, 2],  // "Selected Category": [Selected Option ID, ....]
    "2": []
  },
  "2": {
    "5": [10]
  }
}
```

POST /users/{id}/avatar (won't use, merged with signup)
-------------------------------------------------------
Request:
Multipart form
form data key name = file

Response:
```javascript
{
    "Status": "Success",
    "Filename": "22.png"
}
```

GET /users/{id}/avatar
---------------------
Response:
Raw image

POST /login
---------------------
Request: 
```javascript
{
  "email": "my_email@gmail.com",
  "password": "my_password"
}
```

Response:
```javascript
{
  "id": "1",
  "success": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IiJ9.hdZsKHrlV9xEcv7KtRjGKLxlGC6QqfBctsaSEMqXL6k"
}
```

POST /signup
---------------------
Request:
Multipart Form with folowing form data

file: the avartar file

email: "user@email.com"

password: "23x03dkc$"

Response:
```javascript
{
    "success": ""
}
```

GET /users/{id}/friendsuggestions
---------------------
Response:
```javascript
[
  {
    "email": "mail4@gmail.com",
    "id": "4",
    "weight": "6"
  },
  {
    "email": "mail1@gmail.com",
    "id": "1",
    "weight": "3"
  },
  {
    "email": "mail2@gmail.com",
    "id": "2",
    "weight": "1"
  },
  {
    "email": "mail3@gmail.com",
    "id": "3",
    "weight": "1"
  }
]
```
