socialradar-api
===============

Social Radar API

Current API
===========

GET /questions
Response:
```javascript
[
  {
    "id": "1",
    "order": "1",
    "text": "What are your favorite menus?"
  },
  {
    "id": "2",
    "order": "2",
    "text": "What sports do you play?"
  }
]
```

GET /question/{id}
Response:
```javascript
[
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
```

POST /user/{id}/answer
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

GET /user/{id}/answer

