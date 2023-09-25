# Creating your first stream

Now that we've learned a bit about JetStream, let's start playing with some of its concepts in our local server setup.

## Configuring your NATS server

First, we need to configure our NATS server to be JetStream compatible, meaning it will be able to Store JetStream assets. *Note: When clustering NATS servers, you can mix and match JetStream and non JetStream servers, depending on your needs*.

Add the following line to your configuration file.

```
jetstream {}
```

And restart your server.

## Naming your franchise

Okay so first things first, we are going to imagine that we are an e-commerce company that's distributed from various franchises. You get to now give your franchise an ID. This ID should not contain any punctuation or special characters beyond underscores:

### Good Examples
- codegangsta
- code_gangsta

### Bad Examples
- @codegangsta
- code.gangsta

## Creating your first stream

Let's create our first stream:

```sh
nats stream add
```

This will walk you through an interactive prompt with a lot of questions, I'll walk through them one by one. The most relevant pieces of information we need to select are:

```sh
? Stream Name orders
? Subjects orders.{id}.*
```

All the others can be left to it's default.

This command will then create a stream called "orders" that will store any messages published that match the "orders.{id}.>" subject.

We can view information about all our streams by typing:

```sh
nats stream ls
```

## Publishing messages to the stream

Since JetStream is built directly on top of NATS core, there is nothing special we need to do in order to publish messages to the stream, we can simply use a core NATS message. Let's add 1000 messages into the stream:

```sh
nats pub orders.codegangsta.sku1 --count 1000 '{ "id": "{{ID}}",  "created_at": "{{TimeStamp}}", "count": {{Count}} }'
```

This command makes use of some of the templating functions in the NATS cli to create a JSON object for each order.

Let's do the same thing, but add some orders for the sku2

```sh
nats pub orders.jeremy.sku2 --count 1000 '{ "id": "{{ID}}",  "created_at": "{{TimeStamp}}", "count": {{Count}} }'
```

Now that we have data, lets look at some info for our stream:

`nats s info`

And see what the message breakdown is by subject:

`nats s subjects`

You can also use the sub command to read items from a stream:

```sh
nats sub --stream orders
```

## Creating a Consumer

So far we've read messages from a stream using the `nats sub --stream` command, which creates what we call an "ephemeral" consumer under the hood. This is a consumer that we can create on the fly, and lives for as long as the connection does (along with some grace periods). This is great for ad-hock viewing or querying, but many data pipelines need to be able to pick up where they left off while crunching through a stream. Additionally, certain consumers may only be interested in a subset of a larger dataset, and can have parameters for having the server efficiently filter on subjects.

The NATS CLI also let's us create consumers. Let's go ahead and create our first durable consumer:

```sh
nats consumer add
```

We will give this a name of `sku1_order_processor`, and make this a pull consumer, meaning we can pull messages from the server in batches at our own pace.

```
? Consumer name sku1_order_processor
? Start policy (all, new, last, subject, 1h, msg sequence) all
? Acknowledgement policy explicit
? Replay policy instant
? Filter Stream by subject (blank for all) orders.codegangsta.sku1
? Maximum Allowed Deliveries -1
? Maximum Acknowledgements Pending 0
? Deliver headers only without bodies No
? Add a Retry Backoff Policy No
? Select a Stream orders
```

## Using a consumer

To view information about a consumer:

```sh
nats consumer info orders sku1_order_processor
```

Now let's start pulling batches of 100 messages:

```sh
nats comsumer next orders sku1_order_processor --count=100
```

Feel free to run this command a couple more times, then run:

```sh
NATS consumer info orders sku1_order_processor
```

To see how many messages we've processed
