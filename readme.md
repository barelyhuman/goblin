# goblin

> [gobinaries](https://gobinaries.com/) alternative

Simply put it's a lot of code that's been picked up from the original [gobinaries](https://github.com/tj/gobinaries)
and the majority of the reason is that most of the research for the work has been already done there.

The reason for another repo is that the development on gobinaries has been slow for the past few months / years at this point
and go has moved up 2 version and people are still waiting for gobinaries to update itself.

**All credits to [tj](github.com/tj) for the idea and the initial implementation.**

## Why not fork?

To keep it short, it's fun to build the arch from scratch, helps you learn.
Also the mentality of both the authors differ.

(was easier to start from scratch then remove each blocking thing from the original one)

## Features

- [x] Easy install Script
- [x] Go Lang 1.16 (1.17 - Coming soon)
- [ ] Binary Build Caching

## Authors

[Reaper](https://github.com/barelyhuman), [Mvllow](https://github.com/mvllow)

## Usage

You can read about it on https://goblin.reaper.im

## Deploy your own

Since the entire reason for doing this was that delay on the original implementation added a lot more handling and addition of scripts to my website deployments than I liked. 

I wouldn't want that to happen again, so I really recommend people to spin up their own instances if they can afford to do so. If not, you can always use the hosted version from me at [goblin.reaper.im](https://goblin.reaper.im)

**Note:the original code for gobinaries is equally simple to use and deploy but you'll have to make a few tweaks to the original code to make it work in a simpler fashion**

Let's start

1. Clone the code 
```sh 
git clone https://github.com/barelyhuman/goblin
```
2. Setup docker or any other platform that would allow you to build and run docker images, if using services like Digital Ocean or AWS, you can use their container and docker image specific environments 
3. Build the image 
```sh
cd goblin
docker build -t goblin:latest .
```
4. And finally push the image to either of the environments as mentioned in point 2. If doing it on a personal compute instance, you can just install docker, do step 3 and then run the below command. 
```sh
docker run -p "3000:3000" goblin:latest
```

## License

[MIT](/LICENSE)
