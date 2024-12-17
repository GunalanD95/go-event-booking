# Go Events API

## Overview

**Go Events** is a RESTful API built with **Go (Golang)** and the **Gin framework** for managing events. It provides endpoints to **create**, **read**, **update**, **delete**, and **list events**, along with user authentication features like **signup** and **login**.

---

## Features

- User authentication using **JWT (JSON Web Tokens)**.
- CRUD operations for events.
- Secure endpoints requiring authentication.
- Lightweight and fast API using the Gin web framework.

---

## Endpoints

### Event Management

| Method   | Endpoint          | Description               |
|----------|-------------------|---------------------------|
| `GET`    | `/events`         | Fetch all events          |
| `POST`   | `/events`         | Create a new event        |
| `GET`    | `/events/:id`     | Fetch details of an event |
| `PUT`    | `/events/:id`     | Update an existing event  |
| `DELETE` | `/events/:id`     | Delete an event           |

### User Authentication

| Method   | Endpoint          | Description               |
|----------|-------------------|---------------------------|
| `POST`   | `/signup`         | Register a new user       |
| `POST`   | `/login`          | Login and generate JWT    |

---

## Requirements

- **Go** (v1.21+)
- **Gin** (Web Framework)
- **SQLite/PostgreSQL** (Database)
