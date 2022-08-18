import { serve } from "https://deno.land/std@0.152.0/http/server.ts";

const port = 8080;

const getIdPattern = new URLPattern({ pathname: '/get/:id' });

const handler = (request: Request): Response => {
  const url = new URL(request.url);
  const uri = url.pathname;
  const method = request.method;

  // log each request
  console.log(`Request ${method} ${uri}`);

  // GET /
  if (uri === '/' && method === 'GET') {
    return new Response(`hello world`, { status: 200 });
  }

  // GET /get-json
  if (uri === '/get-json' && method === 'GET') {
    return new Response(
      JSON.stringify({ msg: 'hello world' }),
      {
        status: 200,
        headers: new Headers([
          ['Content-Type', 'application/json']
        ])
      });
  }

  // GET /get/:id
  const idMatch = getIdPattern.exec(request.url);
  if (idMatch && method === 'GET') {
    return new Response(
      JSON.stringify({ msg: `hello ${idMatch.pathname.groups.id}` }),
      {
        status: 200,
        headers: new Headers([
          ['Content-Type', 'application/json']
        ])
      });
  }

  // 404
  return new Response(``, { status: 404 });
};

console.log(`HTTP webserver running. Access it at: http://localhost:${port}/`);
await serve(handler, { port });