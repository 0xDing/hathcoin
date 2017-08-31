# HathCoin 蛤丝币(Work In Process)

[![Build Status](https://travis-ci.org/borisding1994/hathcoin.svg)](https://travis-ci.org/borisding1994/hathcoin) [![Docker Pulls](https://img.shields.io/docker/pulls/borisding/hathcoin.svg)](https://hub.docker.com/r/borisding/hathcoin/) [![Go Report Card](https://goreportcard.com/badge/github.com/borisding1994/hathcoin)](https://goreportcard.com/report/github.com/borisding1994/hathcoin) [![+1s](https://img.shields.io/badge/%CE%98..%CE%98-%2B1s-green.svg)](https://en.wikipedia.org/wiki/Moha_culture)

>© 2017 borisding [<i@boris.tech>](https://github.com/borisding1994/hathcoin) | Licensed under the Apache License

HathCoin is an experimental digital currency, Just for learning blockchain and golang.

"Hath(蛤丝)" is a Chinese [Internet meme](https://en.wikipedia.org/wiki/Internet_meme).

## What's about "Hath" ?

It’s a kind of web culture in China today. Moha (膜蛤), literally "admiring toad" or "toad worship", is an internet meme spoofing Jiang Zemin, former General Secretary of the Communist Party of China and paramount leader of China. It originated among the netizens in mainland China and has become a subculture on the Chinese internet. In the culture, Jiang is nicknamed ha, or "toad", because of his amphibious resemblance. 

Another nickname for Jiang is "elder" or "senior", for he once called himself an "elder" or "senior" when he was berating a Hong Kong journalist Sharon Cheung who questioned him. A video clip recording this event spread on the internet and led to the rise of the culture around 2014, when Hong Kong was experiencing a period of political instability. Initially, netizens extracted Jiang's quotes from the video and imitated his wording and tone, for parody and insult. However, as the culture develops, some imitations have taken to carrying affection toward him. The quotes for imitation have also evolved to include what he said during his leadership, and in his personal life.

Netizens who moha (worship the toad) call themselves "toad fans" or "Hath" (蛤絲), or "mogicians" (膜法師) which is a wordplay of mofashi (魔法师) in Mandarin.

**Long live the man who changed china.**

[Read more](https://en.wikipedia.org/wiki/Moha_culture)

![Big News](https://ipfs.io/ipfs/QmbKHv4r5buzSD1GApRrHf6zgQY5eX4mPw7ATiFSxubS16)


## Quick start

* Docker (recommend)
  ```shell
  # TBD
  docker run -d -p 8081:8081 borisding/hathcoin
  ```
* Binary ([download release-archive](https://github.com/borisding1994/hathcoin/releases))
  ```shell
  ./hathcoin server
  ```

## Development Guide

```shell
# download dependencies
make dep-install
 
# run test
make test
 
# build binary
make build
```

> 很惭愧，就做了一点微小的工作，谢谢大家
