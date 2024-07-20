# Install Janction Node Guide

The janction node can run on Android, Windows, macos, and linux.

## Running on Android

1. You can get our apk installation package through the following link (Tip: remove the .1 suffix)
> Download Link: https://janction-origin-1324956105.cos.ap-seoul.myqcloud.com/janction-v1.0.apk.1

2.  Install the app
3. Enter the Janction app and click wallet connect
4. Then click the login button, which will redirect you to the wallet app of your choice, then approve and sign in

## Run on desktop platforms
###  Prerequisites

Please install the [Prerequisites](./Prerequisites.md) before following these install instructions.



## Download Docker images, Run Node

To help you run Janction node, we have created a Docker image

You can easily run the Janction node by just using the following command:


1. Pull the jct image
```bash
docker pull jct:0.4
```
2. Check if you have successfully pulled the image 
```bash
docker images
```
If the pull is successful, you will see
` jct             0.4             38b1ba30d0f6`

3. Run Janction Node
```bash
docker run -it --name jct-node jct:0.4 -e PRIVATE_KEY="0x0abc...."
```
The `PRIVATE_KEY` parameter is the wallet private key, the user login account, and your points reward will be issued to this account

## Check your node status

You can view your node online status, job execution status, and points reward history in Janction Genesis

* If you need help, post your questions and share your logs on the **run-node-faq** channel on [Discord](https://discord.com/invite/) or on [Telegram](https://).
