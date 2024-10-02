# goblin

> [gobinaries](https://gobinaries.com/) alternative

Simply put it's a lot of code that's been picked up from the original
[gobinaries](https://github.com/tj/gobinaries) and the majority of the reason is
that most of the research for the work has been already done there.

The reason for another repo is that the development on gobinaries has been slow
for the past few months / years at this point and go has moved up 2 version and
people are still waiting for gobinaries to update itself.

**All credits to [tj](https://github.com/tj) for the idea and the initial
implementation.**

- [goblin](#goblin)
  - [Why not fork?](#why-not-fork)
  - [Highlights](#highlights)
  - [Roadmap](#roadmap)
  - [Authors](#authors)
  - [Usage](#usage)
  - [Exposed API](#exposed-api)
  - [Deploy your own](#deploy-your-own)
    - [Existing Image](#existing-image)
      - [Using Docker](#using-docker)
      - [Using Traditional Servers](#using-traditional-servers)
  - [Configuration](#configuration)
  - [License](#license)

## Why not fork?

To keep it short, it's fun to build the arch from scratch, helps you learn. Also
the mentality of both the authors differ.

(was easier to start from scratch then remove each blocking thing from the
original one)

## Highlights

- Easy to use - Users don't need go to install your CLI
- Works with most common package ( Raise an [issue](/issues) if you find it not
  working with something)
- Persistence ( can be made into a temporary cache to minimize storage costs)
- Self Hostable

## Roadmap

- [x] Cache a previously built version binary
  - Builds are cache for 6 hours due reduce storage costs, since it's run by a
    solo developer
- [ ] Add support for download binaries from existing Github Release artifacts

## Authors

[Reaper](https://github.com/barelyhuman), [Mvllow](https://github.com/mvllow)

## Usage

You can read about it on [https://goblin.run](https://goblin.run)

## Exposed API

<small> from v0.4.0 </small>
The server exposes a the following public routes usable for information

**`GET /version/<path>/<to>/<pkg>`**

- A `GET /version` request on a package path for example `github.com/barelyhuman/commitlog/v3` would give you the latest version resolved for it by comparing it on the goproxy and github's repo tags. This is the same algo used internally by Goblin and is available to you.

## Deploy your own

Since the entire reason for doing this was that delay on the original
implementation added a lot more handling and addition of scripts to my website
deployments than I liked.

I wouldn't want that to happen again, so I really recommend people to spin up
their own instances if they can afford to do so. If not, you can always use the
hosted version from me at [goblin.run](https://goblin.run)

**Note:the original code for gobinaries is equally simple to use and deploy but
you'll have to make a few tweaks to the original code to make it work in a
simpler fashion**

### Existing Image

The repository builds and publishes a `latest` and a semver tagged version on each release. You can use that if you do not wish to tweak or change anything in the original source code and build structure

```sh
$ docker run -p "3000:3000" ghcr.io/barelyhuman/goblin:latest
# change the domain to whatever you are using for it
$ docker run -e "GOBLIN_ORIGIN_URL=example.com" -p "3000:3000" ghcr.io/barelyhuman/goblin:latest
```

#### Using Docker

1. Setup docker or any other platform that would allow you to build and run
   docker images, if using services like Digital Ocean or AWS, you can use their
   container and docker image specific environments
2. Build the image

```sh
cd goblin
docker build -t goblin:latest .
```

3. And finally push the image to either of the environments as mentioned in
   point 1. If doing it on a personal compute instance, you can just install
   docker, do step 3 and then run the below command.

```sh
docker run -p "3000:3000" goblin:latest
```

#### Using Traditional Servers

Let's face it, docker can be heavy and sometimes it's easier to run these apps
separately as a simple service.

Much like most go lang projects, goblin can be built into a single binary and
run on any open port.

The repo comes with helper scripts to setup an ubuntu server with the required
stuff

1. Caddy for server
2. Go and Node for language support
3. PM2 as a process manager and start the process in the background

You can run it like so

```sh
./scripts/prepare-ubuntu.sh
```

If you already have all the above setup separately, you can modify the build
script and run that instead.

```sh
./scripts/build.sh
```

You'll have to create 2 `.env` files, one inside `www` and one at the root
`.env`

```sh
# .env

# token from github that allows authentication for resolving versions from go modules as github repositories
GITHUB_TOKEN=
#  the url that you want the server to use for creating scripts
ORIGIN_URL=
```

```sh
# www/.env

# the same url as ORIGIN_URL but added again because the static build needs it in the repo
GOBLIN_ORIGIN_URL=
```

running the `build.sh` should handle building with the needed env files and
restarting the server for you.

## Configuration

The server can be configured easily using environment variables

| KEY                   | default                    | description                                                                          | options         |
| --------------------- | -------------------------- | ------------------------------------------------------------------------------------ | --------------- |
| STORAGE_ENABLED       | `false`                    | Enable persistence                                                                   | `true`          |
| STORAGE_CLIENT_ID     | <empty>                    | CLIENT_ID of a S3 compatible storage                                                 |                 |
| STORAGE_CLIENT_SECRET | <empty>                    | CLIENT_SECRET of a S3 compatible storage                                             |                 |
| STORAGE_ENDPOINT      | <empty>                    | Endpoint value of an S3 compatible storage                                           |                 |
| STORAGE_BUCKET        | <empty>                    | Bucket name of the S3 compatible storage                                             |                 |
| STORAGE_BUCKET_PREFIX | <empty>                    | folder/namespace to store the files inside the bucket                                |                 |
| PORT                  | `3000`                     | Default port for running the application                                             |                 |
| ORIGIN_URL            | `http://localhost:${PORT}` | Default URL of the application                                                       |                 |
| GITHUB_TOKEN          | <empty>                    | Github authentication token for accessing github repositories and resolving versions |                 |
| CLEAR_CACHE_TIME      | <empty>                    | Duration used to clear and set expiry for stored binaries                            | `1m`,`30s`, etc |

## License

[MIT](/LICENSE)
