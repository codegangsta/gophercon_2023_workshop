# Leaf Nodes

Let's now connect our local server to our the demo server as a leaf node. To do this, we will need to create our first server config for NATS.

Create a file called `leaf.conf` and fill it in with the following information:

```
leafnodes: {
  remotes: [
    { urls: "nats://demo.nats.io:7422" }
  ]
}
```

This will tell the server to make a leaf node connection to the demo server. Notice that remotes is an array? That's right, you can have multiple leaf node connections, pretty crazy right?

Now run the server with this configuration:

```
nats-server -c leaf.conf
```

Now when you run:

```
nats server req "gophercon.services" --replies=0
```

You should get a lot more responses back, but the service you are running locally has much lower RTT than the others! This is the power of location transparency and the easy ability to bring services closer to the end user.
