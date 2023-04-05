<p align="center">
    <img src="./assets/head.png" alt="head" width="150"/>
</p>
<h1 align="center">Stealcamoor</h1>

A monitoring service to notify you when a new memory from your favorite influencor drops in [Stealcam](https://www.stealcam.com/).

## Install

```console
git clone https://github.com/0xmichalis/stealcamoor
cd stealcamoor
make build
```

## Setup email connection

Create a Mailtrap account and follow [these instructions](https://www.youtube.com/watch?v=g5o0ixCi4tg) to set up the account.

## Run

Copy the example env file and ensure it has the desired configuration
```console
cp .env.example .env
# make updates to .env
make run
```
