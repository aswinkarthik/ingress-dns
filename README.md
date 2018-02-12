# Ingress-dns (Work-In-Progress)

A tool to register an ingress host with its ingress controller in Consul

## Overview

- `IngressController`s are used to route incoming traffic based on `rules` defined as `Ingress` resources.
- When `Ingress` is used to setup Name based virtual hosts, there is a need to add the `host` to a DNS server.
- `ingress-dns` helps in registering the `host` with the `Service` IpAddress of the `IngressController` to `Consul`


## Pre-Requisites

- Kubernetes with Helm
- Consul
- Nginx-Ingress-Controller

## Installation

We will use helm to install Consul and ingress controller

```bash
# To install consul
$ helm install --name dns stable/consul -f kubernetes/consul-values.yaml

# To install Ingress controller
$ helm install --name gateway stable/nginx-ingress
```
## Building

```bash
# To install dependencies using glide
$ glide install

# To build application
$ go build
```

## Running

```bash
# Create config.yml file
$ cp -v config.yml.sample config.yml

# Make necessary changes to config.yml before starting
# To start
$ ./ingress-dns start

# Or you can use docker
$ docker build . -t ingress-dns
$ docker run --rm -d ingress-dns
```
