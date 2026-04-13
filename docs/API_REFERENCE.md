# TaskDream API Reference

## Base URL
```
http://localhost:8080/api/v1
```

## Health Check (Public)
```
GET /health
Response: 200 OK
{
  "status": "ok"
}
```

---

## Authentication Endpoints (Public)

### Register
Create a new user account.
```
POST /auth/register
Content-Type: application/json

{
  "username": "string",
  "email": "string",
  "password": "string",
  "name": "string"
}

Response: 201 Created
{
  "access_token": "string",
  "user": {
    "id": 1,
    "username": "string",
    "email": "string",
    "name": "string"
  }
}
```

### Login
Authenticate and get access token.
```
POST /auth/login
Content-Type: application/json

{
  "email": "string",
  "password": "string"
}

Response: 200 OK
{
  "access_token": "string",
  "user": {
    "id": 1,
    "username": "string",
    "email": "string",
    "name": "string"
  }
}
```

### Refresh Token
Get a new access token.
```
POST /auth/refresh
Content-Type: application/json
Authorization: Bearer <token>

{
  "refresh_token": "string"
}

Response: 200 OK
{
  "access_token": "string"
}
```

---

## User Endpoints (Protected)

### Get Current User
```
GET /users/me
Authorization: Bearer <token>

Response: 200 OK
{
  "id": 1,
  "username": "string",
  "email": "string",
  "name": "string",
  "avatar_provider": "string",
  "avatar_file_id": 0,
  "is_active": true,
  "created": "2024-01-01T00:00:00Z",
  "updated": "2024-01-01T00:00:00Z"
}
```

### Update Current User
```
PUT /users/me
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "string",
  "avatar_provider": "string",
  "avatar_file_id": 0
}

Response: 200 OK
{
  "id": 1,
  "username": "string",
  "email": "string",
  "name": "string",
  "avatar_provider": "string",
  "avatar_file_id": 0,
  "is_active": true,
  "created": "2024-01-01T00:00:00Z",
  "updated": "2024-01-01T00:00:00Z"
}
```

---

## Project Endpoints (Protected)

### List Projects
```
GET /projects
Authorization: Bearer <token>

Response: 200 OK
[
  {
    "id": 1,
    "title": "string",
    "description": "string",
    "color": "#RRGGBB",
    "is_archived": false,
    "background_file_id": 0,
    "position": 0,
    "area_id": 0,
    "owner_id": 1,
    "created": "2024-01-01T00:00:00Z",
    "updated": "2024-01-01T00:00:00Z"
  }
]
```

### Create Project
```
POST /projects
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "description": "string",
  "color": "#RRGGBB",
  "background_file_id": 0,
  "area_id": 0
}

Response: 201 Created
{
  "id": 1,
  "title": "string",
  ...
}
```

### Get Project
```
GET /projects/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "id": 1,
  "title": "string",
  ...
}
```

### Update Project
```
PUT /projects/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "description": "string",
  "color": "#RRGGBB",
  "is_archived": false,
  "background_file_id": 0,
  "position": 0,
  "area_id": 0
}

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Delete Project
```
DELETE /projects/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "message": "project deleted"
}
```

---

## Task Endpoints (Protected)

### List Tasks
```
GET /tasks?project_id=1
Authorization: Bearer <token>

Response: 200 OK
[
  {
    "id": 1,
    "title": "string",
    "description": "string",
    "done": false,
    "priority": 3,
    "due_date": "2024-01-01T00:00:00Z",
    "project_id": 1,
    "bucket_id": 0,
    "labels": [1, 2, 3],
    "reminders": [],
    "repeat_after": "string",
    "attachments": [],
    "position": 0,
    "percent_done": 0,
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-01T00:00:00Z",
    "created_by_id": 1,
    "created": "2024-01-01T00:00:00Z",
    "updated": "2024-01-01T00:00:00Z"
  }
]
```

### Create Task
```
POST /tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "description": "string",
  "priority": 3,
  "due_date": "2024-01-01T00:00:00Z",
  "project_id": 1,
  "bucket_id": 0,
  "labels": [1, 2],
  "start_date": "2024-01-01T00:00:00Z",
  "end_date": "2024-01-01T00:00:00Z"
}

Response: 201 Created
{
  "id": 1,
  ...
}
```

### Get Task
```
GET /tasks/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Update Task
```
PUT /tasks/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "description": "string",
  "done": false,
  "priority": 3,
  "due_date": "2024-01-01T00:00:00Z",
  "bucket_id": 0,
  "labels": [1, 2],
  "position": 0,
  "percent_done": 50,
  "start_date": "2024-01-01T00:00:00Z",
  "end_date": "2024-01-01T00:00:00Z"
}

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Delete Task
```
DELETE /tasks/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "message": "task deleted"
}
```

---

## Task Comment Endpoints (Protected)

### List Task Comments
```
GET /tasks/:id/comments
Authorization: Bearer <token>

Response: 200 OK
[
  {
    "id": 1,
    "task_id": 1,
    "user_id": 1,
    "comment": "string",
    "created": "2024-01-01T00:00:00Z",
    "updated": "2024-01-01T00:00:00Z"
  }
]
```

### Create Task Comment
```
POST /tasks/:id/comments
Authorization: Bearer <token>
Content-Type: application/json

{
  "comment": "string"
}

Response: 201 Created
{
  "id": 1,
  ...
}
```

### Delete Task Comment
```
DELETE /comments/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "message": "comment deleted"
}
```

---

## Label Endpoints (Protected)

### List Labels
```
GET /labels
Authorization: Bearer <token>

Response: 200 OK
[
  {
    "id": 1,
    "title": "string",
    "description": "string",
    "hex_color": "#RRGGBB",
    "created_by_id": 1,
    "created": "2024-01-01T00:00:00Z",
    "updated": "2024-01-01T00:00:00Z"
  }
]
```

### Create Label
```
POST /labels
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "description": "string",
  "hex_color": "#RRGGBB"
}

Response: 201 Created
{
  "id": 1,
  ...
}
```

### Get Label
```
GET /labels/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Update Label
```
PUT /labels/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "description": "string",
  "hex_color": "#RRGGBB"
}

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Delete Label
```
DELETE /labels/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "message": "label deleted"
}
```

---

## Area Endpoints (Protected)

### List Areas
```
GET /areas
Authorization: Bearer <token>

Response: 200 OK
[
  {
    "id": 1,
    "title": "string",
    "description": "string",
    "color": "#RRGGBB",
    "created_by_id": 1,
    "created": "2024-01-01T00:00:00Z",
    "updated": "2024-01-01T00:00:00Z"
  }
]
```

### Create Area
```
POST /areas
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "description": "string",
  "color": "#RRGGBB"
}

Response: 201 Created
{
  "id": 1,
  ...
}
```

### Get Area
```
GET /areas/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Update Area
```
PUT /areas/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "description": "string",
  "color": "#RRGGBB"
}

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Delete Area
```
DELETE /areas/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "message": "area deleted"
}
```

---

## Bucket Endpoints (Protected)

### List Buckets
```
GET /buckets?project_id=1
Authorization: Bearer <token>

Response: 200 OK
[
  {
    "id": 1,
    "title": "string",
    "project_id": 1,
    "limit": 0,
    "position": 0,
    "created": "2024-01-01T00:00:00Z",
    "updated": "2024-01-01T00:00:00Z"
  }
]
```

### Create Bucket
```
POST /buckets
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "project_id": 1,
  "limit": 0,
  "position": 0
}

Response: 201 Created
{
  "id": 1,
  ...
}
```

### Update Bucket
```
PUT /buckets/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "limit": 0,
  "position": 0
}

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Delete Bucket
```
DELETE /buckets/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "message": "bucket deleted"
}
```

---

## Note Endpoints (Protected)

### List Notes
```
GET /notes?project_id=1
Authorization: Bearer <token>

Response: 200 OK
[
  {
    "id": 1,
    "title": "string",
    "content": "string",
    "project_id": 1,
    "created_by_id": 1,
    "created": "2024-01-01T00:00:00Z",
    "updated": "2024-01-01T00:00:00Z"
  }
]
```

### Create Note
```
POST /notes
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "content": "string",
  "project_id": 1
}

Response: 201 Created
{
  "id": 1,
  ...
}
```

### Get Note
```
GET /notes/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Update Note
```
PUT /notes/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "content": "string"
}

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Delete Note
```
DELETE /notes/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "message": "note deleted"
}
```

---

## Time Tracking Endpoints (Protected)

### List Time Tracking
```
GET /tasks/:id/time-tracking
Authorization: Bearer <token>

Response: 200 OK
[
  {
    "id": 1,
    "task_id": 1,
    "user_id": 1,
    "start": "2024-01-01T09:00:00Z",
    "end": "2024-01-01T10:30:00Z",
    "duration": 5400,
    "comment": "string",
    "created": "2024-01-01T00:00:00Z",
    "updated": "2024-01-01T00:00:00Z"
  }
]
```

### Create Time Tracking
```
POST /tasks/:id/time-tracking
Authorization: Bearer <token>
Content-Type: application/json

{
  "start": "2024-01-01T09:00:00Z",
  "end": "2024-01-01T10:30:00Z",
  "duration": 0,
  "comment": "string"
}

Response: 201 Created
{
  "id": 1,
  ...
}
```

### Update Time Tracking
```
PUT /time-tracking/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "start": "2024-01-01T09:00:00Z",
  "end": "2024-01-01T10:30:00Z",
  "duration": 0,
  "comment": "string"
}

Response: 200 OK
{
  "id": 1,
  ...
}
```

### Delete Time Tracking
```
DELETE /time-tracking/:id
Authorization: Bearer <token>

Response: 200 OK
{
  "message": "time tracking deleted"
}
```

---

## Error Responses

All endpoints return errors in this format:

```json
{
  "error": "Error message"
}
```

Common HTTP Status Codes:
- 200 OK - Success
- 201 Created - Resource created
- 400 Bad Request - Invalid input
- 401 Unauthorized - Missing/invalid token
- 404 Not Found - Resource not found
- 500 Internal Server Error - Server error

---

## Notes

- All protected endpoints require the `Authorization: Bearer <token>` header
- Tokens are obtained from `/auth/login` or `/auth/register`
- Default token expiration: 24 hours
- Dates/times use ISO 8601 format with UTC timezone
- All IDs are integers
- All colors use hex format (#RRGGBB)
