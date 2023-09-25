# Installing the NATS server

So far we've been playing with the NATS demo server, but we may want to have our own local NATS server for better latency and more flexibility as we play with our systems. First up, lets install the nats-server for our operating system:

https://docs.nats.io/running-a-nats-service/introduction/installation

note: try not to use docker for this demonstration as we want to be running commands as local as possible, not that NATS doesn't work with docker, I would prefer us just not to have to debug overlay networks during this lab :) 

## Running the nats server

Running a nats server is super easy, just run:

`nats-server`

## Connecting to the NATS server

Now that the nats server is running, let's create a context for our local nats server and switch to it:

```sh
nats context save local --select
```

Also, we can change our go programs connection screen to `nats.DefaultURL`, which is where our local server is running.

Rerun your go program, and see if you can find your service:

`nats req gophercon.services "" --replies=0`

Notice something? First off, your RTT should now be super fast, which is great. But we also lost the other services from others in the room. This is because you are now running NATS locally. In the next exercise we will get your service back online, this time through a leaf node connection.
