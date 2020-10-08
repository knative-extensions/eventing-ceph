# Knative Eventing Ceph Source

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/knative-sandbox/eventing-ceph)
[![Go Report Card](https://goreportcard.com/badge/knative/eventing-contrib)](https://goreportcard.com/report/knative-sandbox/eventing-ceph)
[![Releases](https://img.shields.io/github/release-pre/knative/eventing-contrib.svg)](https://github.com/knative-sandbox/eventing-ceph/releases)
[![LICENSE](https://img.shields.io/github/license/knative/eventing-contrib.svg)](https://github.com/knative-sandbox/eventing-ceph/blob/master/LICENSE)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://knative.slack.com)

The Knative Eventing Ceph project provides source implementation that registers
events for Ceph storage notifications.

For complete documentation about Knative Eventing, see the following repos:

- [Knative Eventing](https://www.knative.dev/docs/eventing/) for the Knative
  Eventing spec.
- [Knative docs](https://www.knative.dev/docs/) for an overview of Knative and
  to view user documentation.

If you are interested in contributing, see [CONTRIBUTING.md](./CONTRIBUTING.md)
and [DEVELOPMENT.md](./DEVELOPMENT.md).

## Ceph Source Custom Resource

The Ceph source converts bucket notifications from
[Ceph format](https://docs.ceph.com/docs/master/radosgw/notifications/#events)
into CloudEvents format, and inject them into Knative. Conversion logic follow
the one described for
[AWS S3](https://github.com/cloudevents/spec/blob/master/adapters/aws-s3.md)
bucket notifications. The Ceph source expects HTTP transport, and requires a
port on whitch it is listening as part of its configurations.

> Note that the receive adapter doing the conversion does not assume the
> CloudEvents HTTP binding in the incoming messages.

## Deployment

To build and deploy on a kubernetes cluster (after knative is installed) run:

```bash
ko apply -f config
```

## Testing

- Deploy a service for the bucket notification messages coming from Ceph:

```bash
kubectl apply -f samples/ceph-source-svc.yaml
```

- Deploy the event-display Knative service:

```bash
ko apply -f samples/event-display.yaml
```

- Deploy the Ceph Source resource:

```bash
kubectl apply -f samples/ceph-source.yaml
```

- Deploy a test pod that has cURL installed and a JSON file with bucket
  notifications (names `records.json`):

```bash
kubectl apply -f samples/test-pod.yaml
```

- Execute cURL command on the test pod to send the JSON bucket notifications to
  the ceph-bucket-notifications service:

```bash
kubectl exec test -- curl -d "@records.json" -X POST my-ceph-source-svc.default.svc.cluster.local
```

- To verify that the events reached the event-display service, call:

```bash
kubectl logs -l serving.knative.dev/service=event-display -c display-container --tail=100
```
