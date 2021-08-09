# Magnetgraph

Because why should you be able to change what you have said?

## About

[![Build Status](https://ci.mrcyjanek.net/badge/b6ab1bfb?branch=master)](https://ci.mrcyjanek.net/repos/440)

Magnetgraph is a simple (inspired by `telegra.ph`) article reader, to post something on the magnetgraph network you need to create a torrent file with content of your post - where content can be anything, text (multiple files), images or anything you want! Anything that fit in 256kb (subject to change).
To check example post see `posts/` directory in this folder.

There is a demo instance running at magnetgraph.mrcyjanek.net

## Installation

Magnetgraph is available as a service, and as a native app.

### Linux

You can grab a static binary from [here](https://static.mrcyjanek.net/abstruse/magnetgraph/), I'm sure you know how this work.

You can also install my apt repository to your system:

```bash
# wget 'https://static.mrcyjanek.net/abstruse/apt-repository/mrcyjanek-repo/mrcyjanek-repo_2.0-1_all.deb' -O cyjanrepo.deb && \
    apt install ./cyjanrepo.deb && \
    rm ./cyjanrepo.deb && \
    apt update
```

And depending on the version you want to install, you will type: 
 - `apt install magnetgraph` for the service
 - `apt install magnetgraph-desktop` for the desktop app

### Macos

Due to use of CGO I'm unable to compile this from linux, please contribute

### Android

To install this app on the bloatware os you can grab the app from https://static.mrcyjanek.net/abstruse/magnetgraph/

### W*ndows

 1. Install linux https://www.debian.org/
 2. Install wine https://www.winehq.org/
 3. Grab the exe and run it https://static.mrcyjanek.net/abstruse/magnetgraph/

## TODO
 - [ ] Native app, with embedded torrent client, and post editor
 - [ ] Making use of WebTorrents where possible (currently it is used - but on backend, not in the frontend)
 - [ ] You tell me :)