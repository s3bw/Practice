# Publish/Subscribe

Delivering messages to multiple consumers.

## Exchanges

- A _producer_ is a user application that sends messages.
- A _queue_ is a buffer that stores messages.
- A _consumer_ is a user application that receives messages.

The _exchange_ receives messages from a producer. An exchange is
a very simple thing. On one side it receives messages from
producers and the other side it pushes them to queues. The
exchange must know exactly what to do with a message it receives.
Should it be appended to a particular queue? Should it be
appended to many queues? Or should it get discarded. The rules
for that are defined by the _exchange type_.


