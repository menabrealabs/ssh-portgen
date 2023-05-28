# SSH Port Number Generator
Generate an SSH port number based on an SHA-256 digest of the hostname.

## Use Case
The command, `ssh-portgen` allows you to calculate a deterministic port number for SSH or any other service based on the name of the host of the service. Primarily, this makes it easy to assign a unique SSH port number to each host, which makes it much more difficult to identify SSH ports during port scans. It is not a high level of server hardening, but it is one additional line of defense to make scripted fingerprinting of your servers more difficult for automated port scans. You can lookup the port for a host by passing the hostname of the host as an argument to the command. The command is cross-platform and should work in any CLI environment on Linux, Mac (darwin), Windows, or the three major flavors of BSD.

# Usage
To generate a port number for the local host:
```
> ssh-portgen
Hostname: calamity
SHA2 Digest: a776fba1ad9fd5713cb1d0f4b4c772c4ba4a1f853012444b9e28a800a423a9b6
SSH port number: 1253
```
To generate/lookup a port number for a remote host:
```
> ssh-portgen calamity
Hostname: calamity
SHA2 Digest: a776fba1ad9fd5713cb1d0f4b4c772c4ba4a1f853012444b9e28a800a423a9b6
SSH port number: 1253
```
To select the 3rd and 23rd byte of the digest to generate the port number:
```
> ssh-portgen -i 3/23
Hostname: calamity
SHA2 Digest: a776fba1ad9fd5713cb1d0f4b4c772c4ba4a1f853012444b9e28a800a423a9b6
SSH port number: 1936
```
To return only the raw port number for use in a back-tick expression:
```
> echo `./ssh-portgen -r`
1253
```

# Installation
## Compile from Source
You can compile any platform and architecture that is supported by Go (see the distro list below). You will need Go 1.20 or higher to compile the source code:
```
go build -o ssh-portgen
```
Or you can cross-compile for a different OS and/or architecture by setting the GOOS and GOARCH environment variables:
```
env GOOS=linux GOARCH=amd64 go build -o ssh-portgen
```
## Supported Platforms
### AIX
ppc64
### Android
386, amd64, arm, arm64
### Mac / Darwin
amd64, arm64
### DragonFly (BSD)
amd64
### FreeBSD
386, amd64, arm, arm64, riscv64
### Illumos
amd64
### iOS
amd64, arm64
### Javascript
js/wasm
### Linux
386, amd64, arm, arm64, loong64, mips, mips64, mips64le, mipsle, ppc64, ppc64le, riscv64, s390x
### NetBSD
386, amd64, arm, arm64
### OpenBSD
386, amd64, arm, arm64, mips64
### Plan9
386, amd64, arm
### Solaris
amd64
### Windows
386, amd64, arm, arm64

## Available Binaries for Download
[Get Latest release](https://github.com/menabrealabs/ssh-portgen/releases/tag/v1.0.0)

Linux: amd64/x86_64, arm64, ppc64le, ppc64, s390x, riscv64  
Darwin/Mac: amd64/x86_64, arm64  
Windows: amd64/x86_64, arm64  
FreeBSD: amd64/x86_64, arm64, riscv64  
OpenBSD: amd64/x86_64, arm64, mips64  
NetBSD: amd64/x86_64, arm64

