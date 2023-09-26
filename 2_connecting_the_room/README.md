# Connecting the Room Part I

Now that you have NATS CLI installed and in your PATH, we are going to use it to connect everyone in the room.

First thing we should familiarize ourselves with is the concept of a `nats context`. A `nats context` simply represents an identity to execute NATS CLI commands against. Contexts can point to different NATS servers, with different identities and authentication methods. If you are a Kubernetes user, you may be familiar with the concept of contexts in `kubectl`, these contexts are very similar.

## Setting up a NATS context

Since we haven't even touched a NATS server yet, we are going to pray to the demo gods and utilize the `demo.nats.io` server for our early experiments. Let's create a context to point to that server:

`nats context save demo --select`

`nats context edit demo`

And change the url to point to "nats://demo.nats.io:4222"

## Calling the guestbook service

Now to test that this works, let's see if you can call a service I built and am hosting on this demo server, called the gophercon guestbook.

`nats req gophercon.guestbook`

If you can ping that endpoint, congratulations! You have successfully connected to the NATS demo server. Next up, why don't you try and figure out how to sign the guest book with the NATS CLI.

## Hosting your own service

Hosting your own service on NATS infrastructure is also just as easy, just use the `nats reply` command. Let's collectively create a rolecall service that has a one to many relationship:

`nats reply gophercon.rolecall "YOUR NAME HERE" --queue=""`

Once everyone has their reply set, I can then call role and see who is here!

`nats req gophercon.rolecall "Who is here?"`

Unlike many RPC solutions, NATS let you mix and match semantics from pub sub architectures, making for some really neat patterns.

This is a fun pattern, but a more common one is to load balance between multiple instances of a service. Let's change our service a little bit.

`nats reply gophercon.lottery "YOUR NAME HERE" --queue="lottery"`

And I can select the winner by once again using `nats req`

`nats req gophercon.lottery "Who is the winner?"`

BONUS POINTS: Do you know how you'd be able to increase your odds at winning the NATS lottery?

## Poor man's chat

Sometimes your services need more than just request/reply semantics, and pub/sub is a perfect choice for alternatives.

I'm going to create a quick chat room subscription using the NATS CLI:

`nats sub "gophercon.chatroom.>"`

You can participate in this chat room by publishing:

`nats pub gophercon.chatroom.[handle]` "Your Message Here"
