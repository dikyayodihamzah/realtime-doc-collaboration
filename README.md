
---

# Real-Time Document Collaboration Backend

This is a simple backend service for real-time document collaboration, similar to Google Docs or Microsoft Word 365. It allows multiple users to edit a shared document and receive updates in real-time using WebSockets.

## Features

- Real-time collaboration on a shared document.
- Broadcasts document changes to all connected clients.
- Synchronizes the document state across all clients.
- Uses WebSockets for efficient bi-directional communication.

## Prerequisites

- Go 1.16 or newer
- [gorilla/websocket](https://github.com/gorilla/websocket) package for WebSocket support. Install it with:

  ```bash
  go get github.com/gorilla/websocket
  ```

## Project Structure

```
.
├── README.md                           # Project documentation
├── go.mod
├── go.sum
├── app
│   ├── index.html                      # Static HTML sample
│   ├── main.go                         # Main file
│   └── service                         # Bussiness logic layer
│       ├── hubsrv
│       │   └── hub_service.go          # Hub handlers
│       └── wssrv
│           └── websocket_service.go    # WebSocket handlesrs
└── pkg                                 # Project packages
    ├── config                          # Configuration packages
    │   ├── hub.go
    │   └── websocket.go
    └── model                           # Model (struct obejct) packages
        ├── connection.go
        ├── document.go
        └── hub.go
```

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/realtime-doc-collab.git
cd realtime-doc-collab
```

### 2. Install Dependencies

Run the following command to install the WebSocket package:

```bash
go get github.com/gorilla/websocket
```

### 3. Run the Server

Execute the following command to start the WebSocket server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

### 4. Connect to the Server

To connect to the WebSocket server, open a WebSocket client (such as Postman) or create a simple HTML/JavaScript frontend. Connect to the server using the WebSocket URL:

```
ws://localhost:8080/ws
```

Once connected, any messages sent by one client will be broadcast to all connected clients, allowing for collaborative editing.

## Testing with JavaScript

You can test the WebSocket server with the following JavaScript code. This example connects to the WebSocket server and logs incoming messages to the console.

```javascript
const socket = new WebSocket("ws://localhost:8080/ws");

socket.onmessage = (event) => {
    console.log("Received:", event.data);
};

socket.onopen = () => {
    socket.send("Hello, World!"); // Sends a message to the server
};
```

### Sample Frontend (Optional)

For a simple frontend interface, you can use `app/index.html`. Open this file in a browser and connect multiple clients (e.g., in different tabs) to see real-time document collaboration in action. This file can be modified according to project needs.

## Code Explanation

### main.go

- **Document struct**: Holds the shared document’s content, using a mutex to protect concurrent access.
- **Hub struct**: Manages active WebSocket connections, handling registration, unregistration, and broadcasting messages to all connected clients.
- **Connection struct**: Represents a single WebSocket connection and includes methods for reading and writing messages.
- **WebSocket Handlers**:
  - `ServeWebSocket`: Upgrades the HTTP connection to WebSocket and registers it with the hub.
  - `ReadPump`: Reads messages from clients and sends them to the hub for broadcasting.
  - `WritePump`: Sends broadcast messages from the hub to the client.

## How It Works

1. **Client Connections**: Clients connect via the `/ws` WebSocket endpoint.
2. **Document Synchronization**: The server maintains a single document state and broadcasts changes to all connected clients.
3. **Message Flow**: When a client sends a message (document change), the server updates its document state and broadcasts the change to all other clients.

## Future Improvements

- Add user authentication to identify users.
- Handle more complex document operations like specific text changes instead of the entire document.
- Add database support for persistent storage of the document.
- Implement more granular access control (e.g., who can edit or view the document).

## License

This project is open-source and available under the MIT License.

---