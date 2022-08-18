import { Application, Router } from "https://deno.land/x/oak/mod.ts";

const router = new Router();
router.get("/", (ctx) => {
  ctx.response.body = `hello world`;
});

const app = new Application();
app.use(router.routes());
app.use(router.allowedMethods());


console.log(`OAK webserver running. Access it at: http://localhost:8080/`);
app.listen({ port: 8080 });