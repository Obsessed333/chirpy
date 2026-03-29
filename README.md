#  Chirpy API Documentation

Welcome to the **Chirpy** API. This is a RESTful service built in Go for a microblogging platform, featuring user authentication, "chirp" management, and administrative metrics.

---

## Authentication & Security

The API uses **JWT (JSON Web Tokens)** for secure access. Private endpoints require a `Bearer Token` in the `Authorization` header.

### Login
**`POST /api/login`**
*   **Description**: Authenticates a user and returns an Access Token and a Refresh Token.
*   **Request Body**:
    ```json
    {
      "email": "user@example.com",
      "password": "securepassword",
      "expires_in_seconds": 3600
    }
    ```

### Token Management

| Endpoint | Method | Description |
| :--- | :--- | :--- |
| `/api/refresh` | `POST` | Exchange a **Refresh Token** for a new **Access Token**. |
| `/api/revoke` | `POST` | Revoke a Refresh Token (Logout). |

---

## Chirp Endpoints

Chirps are the core of the platform. They are limited to **140 characters** and filtered for profanity.

### Create a Chirp
**`POST /api/chirps`**
*   **Auth Required**: ✅ Yes
*   **Request Body**:
    ```json
    {
      "body": "This is a great new chirp!"
    }
    ```
*   **Success Response (201 Created)**:
    ```json
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "body": "This is a great new chirp!",
      "author_id": "12345"
    }
    ```

### Retrieve Chirps
*   **`GET /api/chirps`**: Returns all chirps. Supports optional query parameters:
    *   `author_id`: Filter by a specific user.
    *   `sort`: `asc` (default) or `desc`.
*   **`GET /api/chirps/{id}`**: Fetch a single chirp by its unique UUID.

### Delete a Chirp
**`DELETE /api/chirps/{id}`**
*   **Auth Required**: ✅ Yes (Must be the author).
*   **Success**: `204 No Content`.

---

##  User Management


| Method | Endpoint | Description |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/users` | Register a new user account. |
| `PUT` | `/api/users` | Update email or password. |
| `POST` | `/api/polka/upgrade` | Upgrade a user to **Chirpy Red** status. |

---

## Administration & Metrics

Used for monitoring server health and development testing.

### Server Metrics
**`GET /admin/metrics`**
Returns an HTML page showing the total number of requests received by the server.

### Database Reset
**`POST /admin/reset`**
> **DANGER**: This endpoint wipes all users and chirps from the database. Use only in development/testing environments.

---

### Error Codes Reference
*   **`200 OK`**: Request successful.
*   **`201 Created`**: Resource created successfully.
*   **`401 Unauthorized`**: Token is missing, invalid, or expired.
*   **`403 Forbidden`**: You do not have permission to modify this resource.
*   **`404 Not Found`**: The requested ID does not exist.