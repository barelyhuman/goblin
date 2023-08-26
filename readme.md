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

## Why not fork?

To keep it short, it's fun to build the arch from scratch, helps you learn. Also
the mentality of both the authors differ.

(was easier to start from scratch then remove each blocking thing from the
original one)

## Highlights

- Easy to use - Users don't need go to install your CLI
- Works with most common package ( Raise an [issue](/issues) if you find it not
  working with something)
- Self Hostable

## Roadmap

- [ ] Cache a previously built version binary
- [ ] Add support for download binaries from existing Github Release artifacts

## Authors

[Reaper](https://github.com/barelyhuman), [Mvllow](https://github.com/mvllow)

## Usage

You can read about it on [https://goblin.run](https://goblin.run)

## Deploy your own

Since the entire reason for doing this was that delay on the original
implementation added a lot more handling and addition of scripts to my website
deployments than I liked.

I wouldn't want that to happen again, so I really recommend people to spin up
their own instances if they can afford to do so. If not, you can always use the
hosted version from me at [goblin.barelyhuman.xyz](https://goblin.run)

**Note:the original code for gobinaries is equally simple to use and deploy but
you'll have to make a few tweaks to the original code to make it work in a
simpler fashion**

Let's start

1. Clone the code

```sh
git clone https://github.com/barelyhuman/goblin
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

# the same url as ORIGIN_URL but added again because vite needs it in the repo
VITE_GOBLIN_ORIGIN_URL=
```

running the `build.sh` should handle building with the needed env files and
restarting the server for you.

## License

[MIT](/LICENSE)
