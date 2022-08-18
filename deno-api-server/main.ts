import {
  Api,
  EMethod,
  Route,
  KeyMatch,
  EEvent,
  RequestEvent,
} from "https://deno.land/x/deno_api_server/mod.ts";

const port = 8080;

const api = new Api({ port });

// log each request
addEventListener(EEvent.BEFORE_REQUEST, (event) => {
  if (event instanceof RequestEvent) {
    //const url = new URL(event.request.url);
    const uri = event.request.url; // old deno variant
    const method = event.request.method;
    console.log(`Request ${method} ${uri}`);
  }
});


// GET /
api.addRoute(
  new Route(EMethod.GET, "/")
    .addPipe(({ response }) => {
      response.body = `hello world`
    }),
);

// GET /get-json
api.addRoute(
  new Route(EMethod.GET, "/get-json")
    .addPipe(({ response }) => {
      response.body = { msg: 'hello world' }
    }),
);


api.addRoute(
  new Route(EMethod.GET, new KeyMatch('/get/:id', { id: {} }))
    .addPipe(({ response, match }) => {
      const id = match.params.get('id');
      response.body = { msg: `hello ${id}` }
    }),
);


console.log(`deno-api-server webserver running. Access it at: http://localhost:${port}/`);
await api.listen();