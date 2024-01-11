---
theme: css/synadia.css
---

<!-- slide bg="./images/Aurora-5k.jpg" -->
<grid drag="100 75" drop="topleft" >
# **Supercharged** Micro-services with **NATS** and **Go**
</grid>

<grid drag="100 50" drop="bottomleft" >
<split gap="1">
![Image|100](https://avatars.githubusercontent.com/u/178316?v=4) <!-- element class="bio" style="align-self:flex-end;" -->
::: block
Jeremy Saenz <!-- element style="padding-top:0.5rem;" -->
<br>
Senior Software Engineer at Synadia
:::
</split>
</grid>

---

# Agenda

---

##### Agenda

## Today's Workshop

- 4 hours, broken down into 4 sections
- Lecture + Exercises
- 45 minutes work, 15 minutes break

---

##### Agenda

## Today's Workshop

- Hour 1: **Rethinking Connectivity**
- Hour 2: **Building Micro-services in Go**
- Hour 3: **The NATS Server**
- Hour 4: **JetStream and Persistence**
- Bonus: **Authentication and Authorization**

---

## About Me

---

##### About Me

::: block
<split gap="1">
![Image|80](https://avatars.githubusercontent.com/u/178316?v=4) <!-- element class="bio" style="align-self:flex-end;" -->

### Jeremy Saenz <!-- element style="padding-top:1.75rem;" -->

</split>
:::

- Long time Gopher (@thecodegangsta)
- Author of **Martini**, **CLI**, **Negroni**, and more...
- Working **@synadia** on **NATS**
- Moved from Engineer -> Product and back again

---

# 🙏 Pray to the Demo Gods

---

#### Lecture:

## Rethinking Connectivity

---

##### Rethinking Connectivity

## Why Rethink Connectivity?

**Multi-cloud** and **Edge** computing is driving a massive transformation

---

##### Rethinking Connectivity

## Limitations of Today's Technology

- **DNS/hostnames/IP** based discovery
- **Pull based** request/reply semantics
- **Perimeter based** security
- **Location-dependent** backends
- Many layers built on **HTTP 1:1** communication

---

##### Rethinking Connectivity

## Introducing NATS

---

##### Rethinking Connectivity

## Introducing NATS

NATS is an **open source**, **high performance** messaging system and **connective fabric**.

It aims to **simplify** the number of technologies you use for your services to communicate, while also **empowering** you to build systems that are **globally available**, **multi-cloud**, **multi-geo**, and **highly adaptive** to change and scale.

---

##### Rethinking Connectivity

# Introducing NATS

- Location-independent addressing
- M:N communications
- Push and pull based
- Decentralized and secure multi-tenancy
- Intelligent persistence
- Global scale

---

##### Rethinking Connectivity

## NATS Architecture

- **Server:** simple, small, easy to deploy Go binary
- **Client:** 40+ client libraries in various languages

---

##### Rethinking Connectivity

## NATS Architecture

- **Core NATS** - High performance messaging. Temporal coupling.
- **JetStream** - Flexible, modern streaming and persistence. Temporal decoupling.

---

##### Rethinking Connectivity

## Core NATS

- Fire and forget message publishing
- Very fast - Scales to millions of msg/s on a single instance
- Flexible subject based addressing with wildcards
- Payload agnostic

---

##### Rethinking Connectivity

## Core NATS

- **Request** and **Reply**
- **Publish** and **Subscribe**
- **Fan In** and **Fan Out**
- **Load Balancing** via **Queue Groups**

---

##### Rethinking Connectivity

## Core NATS Demo

---

### Exercise #1:

## Install the NATS CLI

---

##### Install the NATS CLI

## Now it's your turn!

```bash
$ go install \
    github.com/nats-io/natscli/nats@latest

$ nats help
```

---

### Exercise #2:

## Connecting the Room Part I

---

##### Connecting the Room Part I

## Setting up a NATS context

```
$ nats context save --select demo \
    --server "nats://demo.nats.io:4222"

#   Server URLs: nats://demo.nats.io:4222
#         Path: ...
#   Connection: OK
```

---

##### Connecting the Room Part I

## Calling the guestbook service

```
$ nats req gophercon.guestbook ""
```

---

##### Connecting the Room Part I

## Hosting your own service

```
$ nats reply gophercon.rollcall \
    "YOUR NAME HERE" --queue ""
```

---

##### Connecting the Room Part I

## Play the queue group lottery

```
$ nats reply gophercon.lottery \
    "YOUR NAME HERE" --queue "lottery"
```

---

##### Connecting the Room Part I

## Poor-mans chat

```
$ nats sub "gophercon.chatroom.>"
```

And in another window:

```
$ nats pub gophercon.chatroom.[handle] \
    "Your Message Here"
```

---

### Lecture:

## NATS For Micro-service Architectures

---

##### NATS For Micro-service Architectures

## What makes a good architecture?

- Resilient
- Secure
- Observable
- Extensible
- Adaptive to change

---

##### NATS For Micro-service Architectures

## What NATS gives us

- Resilience
- Secure multi-tenancy
- Location transparency
- Observability
- Multi-pattern development

---

##### NATS For Micro-service Architectures

## Resilience

- **Clients** self heal and reconnect to available servers automatically
- **Servers** protect themselves at all costs
- **Failover** to other Geos/Clouds is **automatic**
- **Load balancing** comes for **free**

---

##### NATS For Micro-service Architectures

## Secure Multi-tenancy

- **Decentralize** authentication and authorization
- **Isolate** NATS environments via **Accounts**
- **Share** streams and services between accounts
- **Enforce** resource limits for tenants
- **Create** permissions for each service without server changes

---

##### NATS For Micro-service Architectures

## Location Transparency

Location transparency is a key characteristic of service-oriented architecture.

Consumers of a service do not know a service's location until they locate it in the registry.

The lookup and dynamic binding to a service at runtime allows the service implementation to move from location to location without the client's knowledge.

---

##### NATS For Micro-service Architectures

## Location Transparency

- Free **Service Discovery** via subject based addressing
- **Easily move** services between cloud providers
- **Automatically** get routed to the closest responder
- **Traffic Shaping** and **Subject mapping**

---

##### NATS For Micro-service Architectures

## Observability

- **Observe traffic** in real time
- **Gather metrics** on each of your services automatically via `nats micro`
- **Filter metrics ingestion** via subjects

---

##### NATS For Micro-service Architectures

## Multi-pattern development

- Synchronous **Request** and **Reply**
- Asynchronous **Publish** and **Subscribe**
- **Streaming** with NATS JetStream
  - Key/Value and Object store
- All with multi-language support!

---

#### Exercise #3:

## Building a Go Micro-service

---

##### Building a Go MicroService

## NATS `micro` package

- Set of conventions for microservices
- Service discovery
- Load balancing
- Metrics and monitoring
- Metadata

---

##### Building a Go MicroService

## Start a `micro` from NATS CLI

```bash
$ nats micro serve hello \
    -H Owner="Your Name"
```

... And make a request

```bash
$ nats req hello.echo ""
```

---

##### Building a Go MicroService

## Start a new Go project

```bash
$ mkdir my_service && cd my_service
$ go mod init github.com/[ME]/my_service
$ go get github.com/nats-io/nats.go@latest
```

---

#### Exercise #4:

## Connecting the Room Part II

---

#### Lecture:

## NATS Server Topologies

---

##### NATS Server Topologies

## Global Scale and Diversity

- **Single Server** - Millions of messages per sec. ~70GiB throughput
- **Clusters and Superclusters** - Fully meshed groups of servers that can span the globe
- **Leaf Nodes** - Extend a NATS system with your own private island

---

##### NATS Server Topologies

## An Illustrated Guide

---

#### Exercise 5:

## Installing and Connecting to a NATS Server

---

#### Exercise 6:

## Leaf Nodes

---

#### Lecture:

## NATS JetStream

---

##### NATS JetStream

## What is JetStream?

JetStream is a next-gen persistence layer built on top of NATS Core that allows temporal decoupling between subscribers and publishers.

It is multi-tenant, highly configurable and globally scalable.

---

##### NATS JetStream

## What is JetStream?

- **Secure** data streams with **multiple consumer models**
- **Multiple streaming patterns** supported
- **Digital twins**/**replicated data**
- **Mux** and **Demux** data
- **Materialized views:** key/value and object stores

---

##### NATS JetStream

## JetStream Demo

https://nats-whiteboard.onrender.com/?room=rhxghz9f

---

##### NATS JetStream

## An Illustrated Guide

---

#### Exercise 7:

## Creating your first Stream

---

#### Exercise 8:

## Key/Value and Object Store

---

#### Exercise 9:

## Connecting the Room Part III

---

---

## Closing Ceremony

---

#### Bonus Lecture:

## Authentication and Authorization

---

---

---

---
