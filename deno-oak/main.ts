import { Application, Router } from "https://deno.land/x/oak/mod.ts";

const port = 8080;

const router = new Router();


// GET /
router.get("/", (ctx) => {
  ctx.response.body = `hello world`;
});

// GET /get-json
router.get("/get-json", (ctx) => {
  ctx.response.body = { msg: `hello world` };
});


// GET /get/:id
router.get("/get/:id", (ctx) => {
  ctx.response.body = { msg: `hello ${ctx.params.id}` };
});

const app = new Application();

// log each request
app.use((ctx, next) => {
  const url = new URL(ctx.request.url);
  const uri = url.pathname;
  const method = ctx.request.method;

  console.log(`Request ${method} ${uri}`);
  next();
})

app.use(router.routes());
app.use(router.allowedMethods({
  throw: true,
}));



console.log(`OAK webserver running. Access it at: http://localhost:${port}/`);
app.listen({ port });