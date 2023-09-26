# Connecting the Room Part III

So hopefully we've had a lot of fun in our previous iterations of connecting the room. So far I hope you've been impressed by how nomadic your micro-services can be. One of the things we really wanted to solve with JetStream was to try and bring as much of that location transparency into the data as well, which is no easy task, since data has to live somewhere and in most cases needs to be consistent.

In this last exercise connecting the room, We are going to attempt to connect all of our streams together into a mega stream.

First, let's update our `jetstream` stanza to define a jetstream domain. JetStream domains are basically just a way to differentiate buckets of JetStream assets, this is important since we are using leaf nodes.

```
jetstream {
    domain: "{your id here}"
}
```

You will want to use your previously used franchise identifier here.

Once configured, let's restart our server.

```sh
nats-server -c server.conf
```

And now I can create a stream that starts muxing some of your orders streams from your leaf node, this stream lives on the demo server in Texas, but can allow anyone to access your data, even when your nodes go offline:

```sh
nats context select demo
nats stream create global_orders --source orders
...
```

Okay, so I've added one stream, lets add another by editing the NATS stream:

```sh
nats stream edit -i
```

Now we have a muxed stream in the cloud. Feel free to stop your server, comment out your leaf node connection or disable wifi, add some messages to your own stream, and reconnect, we can see your messages populating in the cloud stream.
