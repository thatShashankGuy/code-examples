const WebSocket = require("ws");

const server = new WebSocket.Server({ port: 8080 });

// Queue to store the names of connected clients
let clientQueue = [];

/**
 * Node.js uses an event-driven model, and `server.on()` is an event listener.
 * We are listening for the "connection" event, which occurs when a client successfully connects to our WebSocket server.
 * 
 * The "connection" event callback takes two parameters:
 * - `socket`: This represents the individual client that connected.
 * - `request`: (Optional) This contains metadata about the client request, such as headers.
 */
server.on("connection", (socket, request) => {
    /**
     * Here we extract the name of the chat client from the headers of the request.
     * The client should provide their name when connecting.
     * If no name is provided, we default it to "unknown client".
     */
    let clientName = request.headers["name"] || "unknown client";

    // Add the new client to the queue
    clientQueue.push(clientName);
    console.log("New client connected:", clientName);

    // Notify all other clients that a new client has joined
    server.clients.forEach((client) => {
        if (client !== socket && client.readyState === WebSocket.OPEN) {
            client.send(`${clientName} has joined the lobby`);
        }
    });

    /**
     * The `server.clients` property holds a `Set` of all currently connected clients.
     * We convert this `Set` to an array and filter out the current `socket` (which represents the client making the request).
     * 
     * If the chat room is empty (i.e., no other clients are connected), we send a message to the current client
     * letting them know they are alone in the lobby.
     */
    const chatRoom = Array.from(server.clients).filter(client => client !== socket);

    if (chatRoom.length === 0) {
        socket.send("No one is in the lobby");
    }

    /**
     * The `socket` listens for a "message" event, which is triggered whenever the client sends a message.
     * 
     * In the callback function, we receive the `message` sent by the client.
     * We then iterate over all connected clients and send the message to everyone except the sender.
     * 
     * We also check `client.readyState` to ensure that the connection is still open before sending the message.
     */
    socket.on("message", (message) => {
        server.clients.forEach((client) => {
            if (client !== socket && client.readyState === WebSocket.OPEN) {
                client.send(`${clientName}: ${message}`);
            }
        });
    });

    /**
     * The server also listens for the "close" event to handle client disconnection.
     * When a client disconnects, we remove them from the queue and notify all other clients.
     */
    socket.on("close", () => {
        // Remove the client from the queue
        clientQueue = clientQueue.filter(name => name !== clientName);
        console.log(`${clientName} got disconnected`);

        // Notify all other clients that a client has left
        server.clients.forEach((client) => {
            if (client.readyState === WebSocket.OPEN) {
                client.send(`${clientName} has left the lobby`);
            }
        });
    });
});

console.log('WebSocket server is running on ws://localhost:8080');
