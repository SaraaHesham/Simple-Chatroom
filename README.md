# 💬 Simple Chatroom using Go RPC

A feature-rich **command-line chatroom** built with Go’s `net/rpc` package.
This project allows multiple clients to connect to a central server and exchange messages in real-time, demonstrating Go’s power in distributed systems, RPC communication, and concurrency.

## 🎥 Project Demo

Watch the demo video here:
👉 [**Chatroom Demo Video**](https://drive.google.com/drive/folders/130ojqKGmUSrGxhyIR9G00m1-yeg5F0nc?usp=sharing)

## 🧠 Project Concept

This chatroom illustrates a **Client–Server model** using **Remote Procedure Calls (RPC)**.

* **Server:**
  Acts as the central hub that receives and stores messages, and provides them to clients when requested.

* **Clients:**
  Connect to the server, send new messages via RPC, and fetch the chat history periodically to display real-time updates.

The result is a lightweight yet effective simulation of a distributed chat system.


## 🏗️ Architecture Diagram

```
                ┌─────────────────────┐
                │      Chat Server     │
                │ (Go RPC over TCP/IP) │
                │--------------------- │
                │  - Stores messages   │
                │  - Thread-safe list  │
                └─────────┬───────────┘
                          │
         ┌────────────────┼────────────────┐
         │                │                │
 ┌────────────┐     ┌────────────┐     ┌────────────┐
 │  Client 1  │     │  Client 2  │     │  Client 3  │
 │------------│     │------------│     │------------│
 │ - Send RPC │     │ - Send RPC │     │ - Send RPC │
 │ - Get RPC  │     │ - Get RPC  │     │ - Get RPC  │
 └────────────┘     └────────────┘     └────────────┘
```

## 📂 Folder Structure

```
simple-chatroom/
│
├── server/
│   └── server.go      # Server: handles RPC calls and stores messages
│
└── client/
    └── client.go      # Client: sends and receives messages via RPC
```

## ⚙️ Requirements

* Go 1.19 or higher
* Terminal or command prompt access

Verify installation:

```bash
go version
```

## 🚀 Getting Started

### 🖥️ Step 1 — Run the Server

```bash
cd server
go run chat_server.go
```

Output:

```
Chat server running on port 1234...
```

### 💬 Step 2 — Run Clients

Open multiple terminals (one for each user):

```bash
cd client
go run chat_client.go
```

You’ll be prompted:

```
Enter your name: Sara
Welcome, Sara
Type messages below. Type 'exit' to leave.
```

### 💡 Example Chat

**Sara’s terminal:**

```
You: Hi
X: Hello
You: How are you?
X: I'm good!
```

**X’s terminal:**

```
Sara: Hi
You: Hello
Sara: How are you?
You: I'm good!
```

## ⚙️ Implementation Details

### 🔹 Server Logic (`server.go`)

* Registers an RPC service using `net/rpc`.
* Maintains a **slice of strings** for message history.
* Uses `sync.Mutex` to handle concurrent access safely.
* Provides:

  * `SendMessage(msg string, reply *string)` — adds a message to history.
  * `GetMessages(dummy string, reply *[]string)` — returns all messages.

### 🔹 Client Logic (`client.go`)

* Connects to the server with `rpc.Dial("tcp", "localhost:1234")`.
* Uses two goroutines:

  1. **Sender loop:** reads user input and sends it as a message.
  2. **Receiver loop:** polls the server every 800 ms to show new messages.
* Prevents duplicate messages by tracking the last received index.
* Formats your own messages as `You:` for readability.


## 🧩 Features

| Feature                 | Description                                                   |
| ----------------------- | ------------------------------------------------------------- |
| 💬 Multi-client support | Multiple users can join the same chatroom.                    |
| 🔁 Real-time updates    | Each client automatically fetches new messages.               |
| 🧵 Thread-safe storage  | Server uses mutex locks for concurrency safety.               |
| 🧹 Graceful exit        | Type `exit` to leave the chat cleanly.                        |
| 🧠 Lightweight RPC      | Built entirely with Go’s standard library (no external deps). |


## 📚 Learning Outcomes

Through this project, you’ll learn:

* How RPCs enable distributed communication between programs.
* How Go’s concurrency model (goroutines + mutex) handles multiple users efficiently.
* The difference between local and remote function invocation.
* How to structure scalable client–server applications in Go.
