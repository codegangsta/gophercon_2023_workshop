# KV and Object Store

## KV Overview

`nats kv add my_bucket --history 5`

`nats kv put my_bucket foo bar`
`nats kv put my_bucket foo baz`

`nats kv get my_bucket foo`

`nats kv history my_bucket foo`

`nats kv watch my_bucket`

`nats kv put my_bucket foo bat`

`nats kv del my_bucket foo`

Then try watching with go
