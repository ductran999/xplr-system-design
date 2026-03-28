# 04. SSE (Server-Sent Events)

SSE is a standard HTML5 protocol that allows a client to open a **one-way HTTP connection** from the server to the client. This connection is kept open (Keep-Alive) indefinitely, enabling the server to **push data** to the client at any time.

---

## Key Difference from Long Polling:

- **Long Polling:** Receive data → Close connection → Open a new connection.  
  *(High overhead due to repeated network setup)*

- **SSE:** Receive data → Keep the connection open → Wait for more data.  
  *(Only establishes the connection once)*

---

## Key Characteristics:

1. **Unidirectional (One-way):**  
   The server sends data, and the client listens.  
   *(Unlike WebSocket, which is bidirectional)*

2. **Text-only:**  
   Only supports text data (String/JSON), not binary data (e.g., files or audio).

3. **Auto-Reconnect:**  
   The browser automatically handles reconnection if the connection is lost.  
   *(No need to implement reconnection logic manually)*

4. **Strict Format:**  
   Server responses must follow this format:  
   - Prefix: `data: `  
   - Suffix: two newline characters `\n\n`

---

## When to Use SSE?

Use SSE when you need a continuous one-way data stream:

- Stock / Crypto price tickers (live updates)
- Live sports score updates
- User notifications
- File upload/download progress (progress bar)