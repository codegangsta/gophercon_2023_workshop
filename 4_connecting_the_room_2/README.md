# Connecting the Room Part II

Now that we've all created our favorites service, let's run them and make sure they connect to our demo server:

`go run main.go`

Now let's see who has connected:

`nats req gophercon.services "" --replies=0`

When I set replies to 0, it will wait until the timeout period before exiting, giving time for everyone to respond. This acts as a good way to ping all the nodes in a service, for instance, or just a single node in each service. All is possible with some clever queue group naming.

Now that we have our available services, why don't we go ahead and have a little race, who can find what X's favorite color is. Go!

Spend a few minutes to play with the request response architecture in the cli and get to know the folks you are sitting next to.
