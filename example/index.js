const http = require("http")

const port = Number(process.argv.at(-1));

const server = http.createServer((req, res) => {
  console.log(`[${port}][${new Date().toISOString()}] Received ::: ${req.method} - ${req.url}`);
  res.writeHead(200, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({
    data: 'Hello World!',
  }));
});

server.listen(port);
