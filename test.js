const WebSocket = require("ws");
const fetch = require("node-fetch");

const backend = "http://localhost:8080";
const roomId = "room-1";

async function createUser(username) {
  await fetch(`${backend}/signup`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username,
      password: "test123",
    }),
  });

  const res = await fetch(`${backend}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username,
      password: "test123",
    }),
  });

  const data = await res.json();
  return data.user.id;
}

async function createClient(num) {
  const username = `test${num}`;
  const userId = await createUser(username);

  const ws = new WebSocket(
    `ws://localhost:8080/ws/joinRoom/${roomId}?userId=${userId}&username=${username}`
  );

  ws.on("open", () => console.log(`✅ ${username} joined room`));

  ws.on("message", (msg) => {
    console.log(`${username} received:`, msg.toString());
  });

  return ws;
}

// ✅ Create 12 users, auto join the room
(async () => {
  const clients = [];

  for (let i = 1; i <= 12; i++) {
    clients.push(await createClient(i));
    await new Promise((r) => setTimeout(r, 300)); // small delay
  }

  console.log("✅ 12 users joined.");

  // Optional: bots send a test message
  setTimeout(() => {
    clients[0].send("hello, I am user 1");
  }, 2000);
})();
