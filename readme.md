# goblin

> [gobinaries](https://gobinaries.com/) alternative

Simply put it's a lot of code that's been picked up from the original [gobinaries](https://github.com/tj/gobinaries)
and the majority of the reason is that most of the research for the work has been already done there.

The reason for another repo is that the development on gobinaries has been slow for the past few months / years at this point
and go has moved up 2 version and people are still waiting for gobinaries to update itself.

**All credits to [tj](github.com/tj) for the idea and the initial implementation.**

## Why not fork?

While the utils and build use the same concept, it's not gobinaries.

Also the original is tied to GCP, Apex Logs and other services thus making it hard for anyone to spin up their own version of the binary service. Which isn't what I wish, aka the direction of the project is different in this case

That and it's fun to build the arch from scratch, helps you learn.

# Note

WIP! HACKY TO THE CORE.
The current repo isn't even usable, you can join in and help speed up the process but right now the below is what it can do

## Working Features

- [x] Curl Request scripts and Renders
- [x] Create a local build of the hardcoded package
- [x] Upload it to a storage service (commented out in initial versions)
- [ ] Fetch from the storage service (part of the request scripts)
- [ ] Finally host this thing

## Authors

Right now, just [Reaper](https://github.com/barelyhuman)

## Usage

Coming soon

## License

[MIT](/LICENSE)
