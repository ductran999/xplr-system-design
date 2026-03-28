# 02. Short Polling (Periodic Requests)

Short Polling is an early workaround used to build **real-time applications** when constrained by the nature of **REST APIs** (request → response only).

---

## 1. How Short Polling Works

If in **REST API (01)** users must manually refresh (F5) to check for updates, then **Short Polling (02)** simply automates that refresh in the background.

**Flow:**

1. The client (browser) sets up a timer (e.g., `setInterval` in JavaScript).
2. Every X seconds (e.g., 3s), the client sends a request to the server:
   > "Any new data?"
3. The server checks its database/memory.
4. The server **responds immediately**, whether there is new data or not.
5. The connection is closed. Repeat after X seconds.

---

## 2. Pros & Cons

### ✅ Pros
- **Very easy to implement:**  
  Frontend uses `setInterval` + `fetch`. Backend remains a normal REST API.
- **Works everywhere:**  
  No special protocols needed. Works through proxies and firewalls.
- **Feels like real-time:**  
  If polling interval is short, users feel like data updates instantly.

### ❌ Cons
- **Server overload (self-DDoS):**
  - Example: 1 request every 3s → 20 requests/min/user  
  - 10,000 users → **200,000 requests/min**
  - Can easily overwhelm servers

- **Massive resource waste (empty responses):**
  - Example: Chat app, no new message for 1 hour  
  - Client still sends **1,200 useless requests**
  - Wastes bandwidth, CPU, and battery

- **Latency:**
  - Not truly real-time  
  - If data updates at second 1 but client polls at second 3 → 2s delay

---

## 3. When to Use Short Polling

Despite its drawbacks, it’s still useful in some cases:

1. Internal tools (e.g., admin dashboards) with few users (<50)
2. Long polling intervals (e.g., check updates every 5 minutes)

---
