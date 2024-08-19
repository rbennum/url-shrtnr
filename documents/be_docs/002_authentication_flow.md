# Tech Specification Document

PM Docs Phase 002: [here](../pm_docs/phase_002.md).

## 1. Title and Overview

Title: User Authentication Service.

Overview: This document outlines the design and implementation of a User Authentication Service for our product. The service will handle user registration, login, password management, and session handling using JWTs (JSON Web Tokens) for authentication. It will support email-based login only for now.

## 2. Scope

In-Scope:

- User registration via email and password.
- Login and password management.
- Token-based authentication (JWT).
- Session management and expiry.

Out-of-Scope:

- OAuth integration (Google, Apple, GitHub)
- Multi-factor Authentication (MFA)

## 3. Architecture Overview

System Architecture Diagram:

*TODO: insert image here*.

Technology Stack:

- Backend: Golang (Gin Framework)
- Database: PostgreSQL
- Authentication: JWT

## 4. Detailed Design

Components:

1. Authentication Module: Handles user registration, login, log out, and token generation.

APIs/Endpoints:

*TODO: insert detailed request and response JSON*.

- POST /api/v1/register
- POST /api/v1/login
- GET /api/v1/logout

Data Models:

| Column | Type | Description |
| ------ | ------ | ------ |
|id|UUID|Primary key|
|email|String|User's email|
|password_hash|String|Encrypted password|
|created_at|DateTime|Account creation date|
|updated_at|DateTime|Last update timestamp|

## 5. Use Cases/Requirements

Use Case Scenarios:

- Scenario 1: User Registration
  - The user enters their email and password.
  - The system hashes the password and saves it in the database.
- Scenario 2: User Login
  - The user provides their credentials.
  - If valid, the system generates a JWT and returns it to the user for subsequent API requests.

Functional Requirements:

- Secure password hashing using bcrypt.
- JWT-based stateless authentication.

Non-Functional Requirements:

- JWTs should expire after 24 hours.

## 6. Workflow and Process Flow

Workflow Diagrams:

- User Login Workflow:

  - User sends login credentials.
  - System verifies credentials.
  - On success, generates JWT token.
  - User stores JWT for future API calls.

- User Register Workflow:

  - User sends email and password.
  - On success:
    - System hashes the password and saves it in the database.
    - System then generates JWT token as a response.
  - On failure:
    - System returns error.

- User Log Out Workflow:

  - User starts log out flow.
  - *TODO: check the best way to do log out flow*

Sequence Diagrams:

- *TODO: fill this one*.

## 7. Error Handling

Error Codes and Messages:

- 400 Bad Request
- 401 Unauthorized
- 500 Internal Server Error

Fallback Strategies:

- On request failure, retry mechanism will be implemented with a maximum of 3 retries before returning a failure response.

## 8. Database Design

ER Diagrams:

- *TODO: fill this one*.

Indexes and Optimization:

- *TODO: fill this one*.

## 9. Security Considerations

Authentication/Authorization:

- Passwords will be hashed using bcrypt with a salt.
- All tokens will be signed using the HS256 algorithm for JWT.

Data Encryption:

- All sensitive data will be encrypted in transit using TLS/SSL.
- Passwords will never be stored in plaintext.

Vulnerabilities:

- CSRF protection using token-based validation.
- Rate limiting to prevent brute-force attacks.

## 10. Testing Strategy

Unit Tests:

- Test all controllers for valid input/output, including edge cases.

Integration Tests:

- Test interactions between the authentication service and the database.

Load/Performance Testing:

- Target performance: handle 500 concurrent login requests with a response time of less than 200ms.

## 11. Deployment Strategy

CI/CD Pipeline: no CI/CD pipeline implemented.

Environments:

- Production: live environment with real/fake user data.

Rollback Plan:

- A blue-green deployment strategy will be employed to ensure safe rollbacks in case of production failures.

## 12. Monitoring and Logging

Monitoring Tools: no monitoring tools implemented.

Logging: File-based logging.

## 13. Dependencies

External Dependencies: none.

Library Dependencies:

- Gin Framework
- Golang
- Bcrypt Library

## 14. Conclusion

Summary:

The User Authentication Service will provide a secure, scalable, and easy-to-use system for user management and session handling. By implementing this system, we can improve security and usability, while ensuring that the architecture can scale to meet our growing user base.

Next Steps:

- Begin implementation of the auth module.
