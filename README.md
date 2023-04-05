# Stealcamoor

A monitoring service to notify you when a new memory from your favorite influencor drops in [Stealcam](https://www.stealcam.com/).

## Install

```console
git clone https://github.com/0xmichalis/stealcamoor
cd stealcamoor
make build
```

## Setup email connection

Create a Mailtrap account and follow [these instructions](https://mailtrap.io/blog/golang-send-email/#Sending-emails-with-a-third-party-SMTP-server) to set up the account.

## Run

Copy the example env file and ensure it has the desired configuration
```console
cp .env.example .env
# make updates to .env
make run
```
