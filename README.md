# Knative Eventing Ceph Source

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/knative-sandbox/eventing-ceph)
[![Go Report Card](https://goreportcard.com/badge/knative/eventing-contrib)](https://goreportcard.com/report/knative-sandbox/eventing-ceph)
[![Releases](https://img.shields.io/github/release-pre/knative/eventing-contrib.svg)](https://github.com/knative-sandbox/eventing-ceph/releases)
[![LICENSE](https://img.shields.io/github/license/knative/eventing-contrib.svg)](https://github.com/knative-sandbox/eventing-ceph/blob/master/LICENSE)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://knative.slack.com)

The Knative Eventing Ceph project provides source implementation that
registers events for Ceph storage notifications.

For complete documentation about Knative Eventing, see the following repos:

- [Knative Eventing](https://www.knative.dev/docs/eventing/) for the Knative
  Eventing spec.
- [Knative docs](https://www.knative.dev/docs/) for an overview of Knative and
  to view user documentation.

If you are interested in contributing, see [CONTRIBUTING.md](./CONTRIBUTING.md)
and [DEVELOPMENT.md](./DEVELOPMENT.md).

# Ceph to Knative Receive Adapters

The following receive adapters convert bucket notifications from
[Ceph format](https://docs.ceph.com/docs/master/radosgw/notifications/#events)
into CloudEvents format, and inject them into Knative. Adapter logic follow the
one described for
[AWS S3](https://github.com/cloudevents/spec/blob/master/adapters/aws-s3.md)
bucket notifications.

There are 4 different transport options:

## HTTP

This receive adapter expects bucket notifications from Ceph over HTTP transport,
as payload in POST messages.

> Note that this adapted does not assume the CloudEvents HTTP binding in the
> incoming messages.

### Testing

- Create a namespace for the tests:

```
kubectl apply -f samples/ceph-eventing-ns.yaml
```

- Deploy a service for the bucket notification messages coming from Ceph:

```
kubectl apply -f samples/ceph-bucket-notification-svc.yaml -n ceph-eventing
```

- Deploy the ceph-event-display Knative service together with a channel and
  subscribe it to the channel:

```
ko apply -f samples/ceph-display-resources.yaml -n ceph-eventing
```

- Build and deploy the Ceph source adapter with SinkBinding to inject the sink:

```
ko apply -f samples/ceph-source.yaml -n ceph-eventing
```

- Deploy a test pod that has cURL installed and a JSON file with bucket
  notifications (names `records.json`):

```
kubectl apply -f samples/test-pod.yaml -n ceph-eventing
```

- Execute cURL command on the test pod to send the JSON bucket notifications to
  the ceph-bucket-notifications service:

```
kubectl exec test -n ceph-eventing -- curl -d "@records.json" -X POST ceph-bucket-notifications.ceph-eventing.svc.cluster.local
```

- To verify that the events reached the ceph-event-display service, call:

```
kubectl logs -l serving.knative.dev/service=ceph-event-display -n ceph-eventing -c ceph-display-container --tail=100
```

## PubSub (TODO)

This receive adapter is pulling bucket notifications from a special "pubsub"
zone in Ceph, and acking them once they are successfully sent to Knative. More
information about this mechanism could be found
[here](https://docs.ceph.com/docs/master/radosgw/pubsub-module/).

> Note that this adapter needs access to the Ceph cluster to pull the
> notifications and ack them.

## AMQP0.9.1 (TODO)

This receive adapter is subscribed to a list of AMQP topics, on which bucket
notifications are expected.

> Note that this adapted does not assume the CloudEvents AMQP binding in the
> incoming messages.

## Kafka (TODO)

This receive adapter is subscribed to a list of Kafka topics, on which bucket
notifications are expected.

> Note that this adapted does not assume the CloudEvents Kafka binding in the
> incoming messages.
