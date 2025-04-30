# Watermill Basic PubSub

```console
$ go build
$ ./watermill-basic-pubsub -h
Usage of ./watermill-basic-pubsub:
  -debug
        debug mode
  -logger string
        logger adapter
  -trace
        trace mode
$ ./watermill-basic-pubsub
[uuid]: 2153cec2-1872-49d5-935e-64e9e63ec21f [payload]: Do the laundry
[uuid]: a27bd4dd-f9de-40e0-901c-9e5655f8e233 [payload]: Go to workout
[uuid]: 3054a460-bb01-47a8-80ee-40560b9164ac [payload]: Prepare the breakfast
[watermill] 2025/04/30 16:04:14.111695 pubsub.go:306:   level=INFO  msg="Pub/Sub closed" pubsub_uuid=Hd4comZng5u5tgjK4b4oWj
$ ./watermill-basic-pubsub -debug
[watermill] 2025/04/30 16:04:21.483480 pubsub.go:127:   level=DEBUG msg="Waiting for subscribers ack" message_uuid=735f6587-1e5f-4791-a124-db15118d6b0c pubsub_uuid=YVUZGV6euCZnsXBFrsVfzb
[uuid]: 735f6587-1e5f-4791-a124-db15118d6b0c [payload]: Do the laundry
[watermill] 2025/04/30 16:04:21.483602 pubsub.go:127:   level=DEBUG msg="Waiting for subscribers ack" message_uuid=e726f584-deef-464a-a833-000f6608194f pubsub_uuid=YVUZGV6euCZnsXBFrsVfzb
[uuid]: e726f584-deef-464a-a833-000f6608194f [payload]: Go to workout
[watermill] 2025/04/30 16:04:21.483647 pubsub.go:127:   level=DEBUG msg="Waiting for subscribers ack" message_uuid=1c4152c9-3947-4a8a-aee1-54de93469e3c pubsub_uuid=YVUZGV6euCZnsXBFrsVfzb
[uuid]: 1c4152c9-3947-4a8a-aee1-54de93469e3c [payload]: Prepare the breakfast
[watermill] 2025/04/30 16:04:21.483679 pubsub.go:303:   level=DEBUG msg="Closing Pub/Sub, waiting for subscribers" pubsub_uuid=YVUZGV6euCZnsXBFrsVfzb
[watermill] 2025/04/30 16:04:21.483690 pubsub.go:331:   level=DEBUG msg="Closing subscriber, waiting for sending lock" pubsub_uuid=YVUZGV6euCZnsXBFrsVfzb
[watermill] 2025/04/30 16:04:21.483696 pubsub.go:337:   level=DEBUG msg="GoChannel Pub/Sub Subscriber closed" pubsub_uuid=YVUZGV6euCZnsXBFrsVfzb
[watermill] 2025/04/30 16:04:21.483721 pubsub.go:306:   level=INFO  msg="Pub/Sub closed" pubsub_uuid=YVUZGV6euCZnsXBFrsVfzb
```
