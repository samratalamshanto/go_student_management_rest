# go_student_management_rest

Key Points:
1. Basic CRUD in GO
2. Eager Loading and Lazy Loading
3. Relation between entities
4. CommonResponse for All rest API
5. Postman Collections



## Common Response For All Rest API:

1. **Success:**
```json
{
    "code": 200,
    "success": true,
    "message": "Successfully get the data.",
    "msg": "Successfully get the data.",
    "data": [],
    "error": null,
    "err": null
}
```

2. **Failed/Error:**
```json
 {
  "code": 500,
  "success": false,
  "message": "Failed to get data.",
  "msg": "Failed to get data.",
  "data": null,
  "error": {},
  "err": {}
}
```
  

**Example: Get All Students List with Nested Course Details, Teacher List** 
```json
{
  "code": 200,
  "success": true,
  "message": "Successfully get the data.",
  "msg": "Successfully get the data.",
  "data": {
    "id": 1,
    "created_at": "2025-03-15T19:43:36.857931+06:00",
    "created_by": 0,
    "created_by_name": "",
    "updated_at": "2025-03-15T19:43:36.857931+06:00",
    "updated_by": 0,
    "updated_by_name": "",
    "active": true,
    "name": "student22",
    "email": "student22@gmail.com",
    "class": 1,
    "rollNo": 1,
    "courses": [
      {
        "id": 1,
        "created_at": "2025-03-15T19:43:36.866201+06:00",
        "created_by": 0,
        "created_by_name": "",
        "updated_at": "2025-03-15T19:43:36.866201+06:00",
        "updated_by": 0,
        "updated_by_name": "",
        "active": true,
        "courseName": "math",
        "courseCode": "math-001",
        "description": "math biginer level",
        "students": null,
        "teachers": [
          {
            "id": 1,
            "created_at": "2025-03-15T20:46:11.813948+06:00",
            "created_by": 0,
            "created_by_name": "",
            "updated_at": "2025-03-15T20:46:11.813948+06:00",
            "updated_by": 0,
            "updated_by_name": "",
            "active": true,
            "name": "teacher1",
            "email": "teacher1@gmail.com",
            "courses": null,
            "payslips": null
          }
        ]
      }
    ],
    "fees_invoices": []
  },
  "error": null,
  "err": null
}
```