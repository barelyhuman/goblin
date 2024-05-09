# docker-compose-setup

The example setup is for running an instance of goblin on your personal server
using Docker + Docker Compose and Caddy.

The Setup process involves a few simple steps.

1. Create a Caddy Network bridge

```sh
; docker network create -d bridge caddy_network
```

> [!NOTE]: Replace `caddy_network` with something else if you already have a
> similar name for a different cluster and would like to avoid merging the
> existing caddy services with this one. If merging them is not an issue, please
> continue with the exiting name.

2. Modify the `Caddyfile` to make sure it has the right domains, the example has
   `goblin.run`.
3. Run the Caddy Service, you only need to run this once and just verify it's
   running by using `docker compose ps`

```sh
docker compose -f ./docker-compose.yml up -d
```

4. Now you can run the `goblin` service using the following

```sh
docker compose -f ./docker-compose.goblin.yml up -d
```

You can verify the status of the running containers by running the following

```sh
docker compose ps
```

You should have the services `caddy-1` and `goblin-1` running and will manage
the domain mapping for you.
