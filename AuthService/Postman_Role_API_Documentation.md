# AuthService Role Management APIs - Postman Documentation

## Overview

This document provides comprehensive examples for testing all Role Management APIs in your AuthService using Postman.

## Base Configuration

- **Base URL**: `http://localhost:3001` (default port from your config)
- **Content-Type**: `application/json` for all requests

## Environment Variables Setup

Before testing, set up these environment variables in Postman:

| Variable    | Default Value           | Description                          |
| ----------- | ----------------------- | ------------------------------------ |
| `base_url`  | `http://localhost:3001` | Base URL for AuthService             |
| `role_id`   | `1`                     | Role ID for testing                  |
| `user_id`   | `1`                     | User ID for role assignment          |
| `jwt_token` | `your_jwt_token_here`   | JWT token for authenticated requests |

## API Endpoints

### 1. Get All Roles

**GET** `/roles`

**Description**: Retrieves all roles from the system

**Headers**:

```
Content-Type: application/json
```

**Example Response**:

```json
{
  "success": true,
  "message": "Role fetched successfully",
  "data": [
    {
      "id": 1,
      "name": "admin",
      "description": "Administrator role with full system access",
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
    }
  ]
}
```

### 2. Get Role by ID

**GET** `/roles/{id}`

**Description**: Retrieves a specific role by its ID

**Path Parameters**:

- `id` (required): Role ID

**Example Request**: `GET /roles/1`

**Example Response**:

```json
{
  "success": true,
  "message": "Role fetched successfully",
  "data": {
    "id": 1,
    "name": "admin",
    "description": "Administrator role with full system access",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
}
```



**DELETE** `/roles/{id}`

**Description**: Deletes a role by ID

**Path Parameters**:

- `id` (required): Role ID to delete

**Example Request**: `DELETE /roles/1`

**Example Response**:

```json
{
  "success": true,
  "message": "Role deleted successfully",
  "data": null
}
```


```

### 3. Get All Role Permissions

**GET** `/role-permissions`

**Description**: Retrieves all role-permission mappings in the system

**Example Response**:

```json
{
  "success": true,
  "message": "all permission fetched successfully",
  "data": [
    {
      "id": 1,
      "role_id": 1,
      "permission_id": 1,
      "role_name": "admin",
      "permission_name": "read_users",
      "assigned_at": "2025-01-01T00:00:00Z"
    }
  ]
}

```


- JWT Authentication (`JWTAuthMiddleware`)
- Admin role required (`RequireAllRoles("admin")`)

**Example Request**: `POST /roles/1/assign/2`

**Example Res

## Error Responses

### Common Error Format

```json
{
  "success": false,
  "message": "Error description",
  "error": "Detailed error information"
}
```

### Common Error Codes

- `400 Bad Request`: Invalid request body, validation errors, missing parameters
- `401 Unauthorized`: Missing or invalid JWT token
- `403 Forbidden`: Insufficient permissions (not admin role)
- `404 Not Found`: Role not found
- `500 Internal Server Error`: Server-side errors

### Example Error Responses

**Validation Error**:

```json
{
  "success": false,
  "message": "validation failed",
  "error": "Key: 'CreateRoleRequestDTO.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

**Role Not Found**:

```json
{
  "success": false,
  "message": "Role not found",
  "error": "role with ID 999 not found"
}
```

## Testing Workflow

### Recommended Testing Order:

1. **Create Role** - Create a test role
2. **Get All Roles** - Verify the role was created
3. **Get Role by ID** - Test specific role retrieval

4. **Get All Role Permissions** - Check system-wide role permissions



