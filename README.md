# Overview

## Introduction

This repository is in support of Kamailio GH issue 4603:

https://github.com/kamailio/kamailio/issues/4603

This demonstration works by sending an 8 kB payload out via EVAPI in response to a SIP `OPTIONS` request. An EVAPI client service reads up to 8000 bytes and then stops, causing buffer backpressure in the sender.

For simplicity, Kamailio has been limited to a single UDP SIP worker and a single EVAPI dispatcher process.

## Dependencies

SIP Swiss Army Knife (`sipsak`), easily obtained from various sources and package managers. This is needed to send `OPTIONS` requests easily from the command line.

## Build and initialisation

To build and run:

```
$ docker compose up --build
```

Then, in a separate window:

```
$ while : ; do 
  sleep 0.1;
  sipsak -T -H 127.0.0.1 -p 127.0.0.1 -s sip:s@sip-proxy 2>/dev/null; 
done
```

Monitor Kamailio output:

```
sip-proxy-1     |  1(8) ERROR: <script>: Total bytes relayed out: 2441216
sip-proxy-1     |  1(8) ERROR: <script>: Total bytes relayed out: 2449408
sip-proxy-1     |  1(8) ERROR: <script>: Total bytes relayed out: 2457600
sip-proxy-1     |  1(8) ERROR: <script>: Total bytes relayed out: 2465792
```

... and `sipsak` output, which for some time remains normal and indicates responses to `OPTIONS` requests:

```
address: 16777343, rport: 0
address: 16777343, rport: 0, username: 's', domain: 'sip-proxy'
0: ?? (3.012 ms) SIP/2.0 483 Too Many Hops
0: ?? (1.952 ms) SIP/2.0 200 OK
        without Contact header
```

... then, at some point--that point is probably platform-dependent--stops:

```
nothing received, select returned error
address: 16777343, rport: 0
address: 16777343, rport: 0, username: 's', domain: 'sip-proxy'
```