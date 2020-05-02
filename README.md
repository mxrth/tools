# Cross Platform CLI Tools

## Installation

You can download pre-built executables for linux and windows (64bit) on the releases page.

If you have a somewhat recent go toolchain (>= 1.13) installed you can run

`go get github.com/mxrth/tools/cmd/<tool>`

to install \<tool\>.

## Tools

* slice  - slice and dice files and put them back together again
* aes - aes encryption/decryption in various modes (ecb, cbc, ctr, gcm,...)
* rsa - rsa encryption/decryption
* ecc - elliptic curve crypto
* [pphgen - generate passphrases](https://github.com/mxrth/tools/tree/master/cmd/pphgen)


## Supported Platforms

All tools are tested and used on Linux and Windows. They should run on any platform supported by Go though.
The goal is that these tools not merely run on Linux and Windows, but work naturally with a POSIX-Shell (bash, mostly) and PowerShell.


## CLI Conventions