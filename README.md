
# LDAP Authentication Service

This project is a Go-based LDAP Authentication Service that provides user authentication via an LDAP server. It is designed with modular components for better maintainability and scalability.

---

## Features

- **LDAP Integration**: Authenticate users with LDAP.
- **Configuration Management**: Easily configure the service using `config.yaml`.
- **Modular Design**: Separate layers for services, controllers, and repositories.
- **RESTful API**: Exposes endpoints for authentication operations.

---

## Project Structure

```plaintext
ldap-auth-service/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   ├── config.yaml             # Configuration file
│   └── config.go               # Configuration loader
├── controllers/
│   └── auth_controller.go      # Handles HTTP requests
├── models/
│   └── user.go                 # User data model
├── repositories/
│   └── ldap_repository.go      # LDAP interaction logic
├── routes/
│   └── routes.go               # API route definitions
├── services/
│   └── auth_service.go         # Business logic for authentication
```

---

## Prerequisites

- **Go**: Version 1.18 or higher.
- **LDAP Server**: An operational LDAP server.
- **Configuration**: Update `config.yaml` with your LDAP server details.

---

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ldap-auth-service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Update the configuration in `config/config.yaml` to match your environment.

---

## Configuration

The `config/config.yaml` file contains all necessary configurations:

```yaml
ldap:
  host: "ldap.example.com"
  port: 389
  base_dn: "dc=example,dc=com"
  bind_dn: "cn=admin,dc=example,dc=com"
  bind_password: "password"
server:
  port: 8080
```

---

## Running the Application

Start the service with:

```bash
go run cmd/main.go
```

The application will be accessible at `http://localhost:8080`.

---

## Endpoints

### POST `/auth/login`

**Description**: Authenticates a user with LDAP.

**Request Body**:
```json
{
  "username": "user1",
  "password": "password123"
}
```

**Response**:
- **200 OK**: Authentication successful.
- **401 Unauthorized**: Authentication failed.

**Sample Response**:
```json
{
  "message": "Authentication successful",
  "user": {
    "username": "user1",
    "email": "user1@example.com"
  }
}
```

---

## Development

### Adding New Features

1. Create new models in `models/`.
2. Add the required business logic in `services/`.
3. Integrate logic into a controller in `controllers/`.
4. Define new routes in `routes/routes.go`.

### Running Tests

To run tests for the service:
```bash
go test ./...
```

---

## Contribution

Feel free to open issues or submit pull requests for improvements and bug fixes. Follow these guidelines:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add new feature"
   ```
4. Push and open a pull request.


---

## Contact

For inquiries or support, contact [Your Name](mailto:info@hajimohammadi.net).
# ldap-auth-service
