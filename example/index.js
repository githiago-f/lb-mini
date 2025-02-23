const http = require("http")

const port = Number(process.argv.at(-1));
let counter = 0;

const server = http.createServer((req, res) => {
  const remote = req.socket.remoteAddress;

  console.log(`[${port}][${counter++}] Received ::: ${remote}`);
  
  res.writeHead(200, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({
    data: 'Hello World!',
  }));
});

server.listen(port);
