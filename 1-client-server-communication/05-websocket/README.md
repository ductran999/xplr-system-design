# 05. WebSocket (Full-Duplex Communication)

WebSocket is the most powerful technology today for real-time applications. It no longer follows the traditional request–response (HTTP) model, but instead establishes a **persistent, full-duplex TCP connection** between the client and the server.

---

## 1. How WebSocket Works (The "Upgrade" Trick)

WebSocket does not start immediately.

1. The client first sends a normal HTTP request to the server, but includes special headers:  
   `"Connection: Upgrade"` and `"Upgrade: websocket"`.

2. This implicitly tells the server:  
   *"Hey, let’s stop using HTTP and switch to the WebSocket protocol!"*

3. If the server agrees, it responds with status code:  
   `101 Switching Protocols`.

4. From that moment on, HTTP is no longer used.  
   The browser and server establish a **bidirectional communication channel** (starting with `ws://` instead of `http://`).  
   Both sides can **send and receive data at any time**.

---

## 2. Comparison with SSE and Long Polling

| Criteria | Long Polling (03) | SSE (04) | WebSocket (05) |
| :--- | :--- | :--- | :--- |
| **Communication Direction** | Two-way (via repeated open/close) | One-way (Server → Client) | **Two-way (Full Duplex)** |
| **Connection Lifecycle** | Closed after each response | Kept open indefinitely | Kept open indefinitely |
| **Protocol** | HTTP (`http://`) | HTTP (`http://`) | WebSocket (`ws://` or `wss://`) |
| **Data Types** | Text / JSON | Text only (Event Stream) | **Text + Binary (images, audio, files)** |

---

## 3. When SHOULD You Use WebSocket?

Use WebSocket whenever the client needs to send data continuously at high frequency:

- Multiplayer games (movement, shooting, real-time actions)
- Real-time chat applications (Zalo, Messenger)
- Collaborative tools (shared whiteboards, Google Docs editing)
- Bidirectional audio/video streaming  
  *(WebRTC often uses WebSocket for signaling)*

---