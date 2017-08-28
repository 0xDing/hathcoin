# HathCoin 蛤丝币

[![Build Status](https://travis-ci.org/borisding1994/hathcoin.svg)](https://travis-ci.org/borisding1994/hathcoin) [![Docker Pulls](https://img.shields.io/docker/pulls/borisding/hathcoin.svg)](https://hub.docker.com/r/borisding/hathcoin/) [![+1s](https://img.shields.io/badge/%CE%98..%CE%98-%2B1s-green.svg)](https://zh.wikipedia.org/wiki/%E8%86%9C%E8%9B%A4%E6%96%87%E5%8C%96)

>© 2017 borisding<i@boris.tech> | Licensed under the Apache License

![Big News](https://ipfs.io/ipfs/QmbKHv4r5buzSD1GApRrHf6zgQY5eX4mPw7ATiFSxubS16)

HathCoin is an experimental digital currency. Long live the man who changed china. 

## Quick start

* Docker (recommend)
  ```shell
  # TBD
  docker run -d -p 8081:8081 borisding/hathcoin
  ```
* Binary ([download release-archive](https://github.com/borisding1994/hathcoin/releases))
  ```shell
  ./hathcoin daemon
  ```

## Development Guide
> [Bazel](https://www.bazel.build/) and [dep](https://github.com/golang/dep) is required.

```shell
# download dependencies
make dep-install
 
# run test
make test
 
# build binary
make build
```

## RPC API

