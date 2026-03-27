# 01. REST API (Request–Response Pattern)

REST (Representational State Transfer) is the most common architecture for client–server communication on the Internet. It follows a simple **Request–Response** model.

---

## 1. How REST API Works

- **Client initiates everything:** The server only responds when a request is sent.
- **No request → no response:** The server cannot push data on its own.
- **Short-lived connection:** After responding, the connection is closed.

---

## 2. Core Principles

- **Stateless:**  
  Each request is independent. The server does not remember previous requests, so all necessary data (e.g., auth token) must be included every time.

- **Resource-Based:**  
  Everything is treated as a resource and identified by URLs (use nouns):
  - ❌ `/get-all-users`
  - ✅ `/users`

- **HTTP Methods:**

| Method | Action           | Example      | Success Code |
| ------ | ---------------- | ------------ | ------------ |
| GET    | Read (list)      | `/users`     | 200 OK       |
| GET    | Read (one)       | `/users/:id` | 200 OK       |
| POST   | Create           | `/users`     | 201 Created  |
| PUT    | Update (full)    | `/users/:id` | 200 OK       |
| PATCH  | Update (partial) | `/users/:id` | 200 OK       |
| DELETE | Delete           | `/users/:id` | 200 / 204    |

---

## 3. Pros & Cons

### Pros

- Easy to learn and widely adopted
- Highly scalable (stateless)
- Suitable for most web applications

### Cons

- Not ideal for real-time apps
- Repeated HTTP headers → extra overhead

---

## 4. Demo Setup

```bash
npm init -y
npm install express cors
```
