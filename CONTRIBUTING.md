We 💚 Opensource!

Yes, because we feel that it’s the best way to build and improve a product. It allows people like you from across the globe to contribute and improve a product over time. And we’re super happy to see that you’d like to contribute to ZTKA.

We are always on the lookout for anything that can improve the product. Be it feature requests, issues/bugs, code or content, we’d love to see what you’ve got to make this better. If you’ve got anything exciting and would love to contribute, this is the right place to begin your journey as a contributor to ZTKA and the larger open source community.

**How to get started?**
The easiest way to start is to look at existing issues and see if there’s something there that you’d like to work on. You can filter issues with the label “Good first issue” which are relatively self sufficient issues and great for first time contributors.

Once you decide on an issue, please comment on it so that all of us know that you’re on it.

If you’re looking to add a new feature, raise a new issue and start a discussion with the community. Engage with the maintainers of the project and work your way through.

## Relay Server & Relay Agent

Relay server and agent for kubectl access.

### Development setup

Relay depends following code-repos:

`rcloud-base` : git@github.com:RafaySystems/rcloud-base.git

Check out the above code repos.

Go to rcloud-base

```bash
make teardown 
make setup
make run
```

### Rafay Relay Server

Get peer client token for relay.
Open the admindb postgress and execute below SQL to fetch the token.

```bash
admindb=# select token from sentry_bootstrap_agent_template where name='rafay-sentry-peering-client';
        token         
----------------------
 brhsrbbipt349qt0qad0
```

Use above token set the environment variable for relay-server

```bash
export RELAY_PEERING_TOKEN=brhsrbbipt349qt0qad0
export RAFAY_RELAY_PEERSERVICE=http://peering.sentry.rafay.local:10001 
export POD_NAME=relay-pod1
```

```go
go run main.go --mode server --log-level 3
```

**Rafay Relay Agent**

Get the config for rafay relay agent

`curl http://localhost:11000/infra/v3/project/:metadata.project/cluster/:metadata.name/download`

Use the values of relay-agent-config and apply in the cluster.

`kubectl -n rafay-system get cm/relay-agent-config -o yaml`

```yaml
apiVersion: v1
data:
  clusterID: 4qkolkn
  relays: '[{"token":"brht3cjipt35cb3kfjag","addr":"api.sentry.rafay.local:11000","endpoint":"*.core-connector.relay.rafay.local:443","name":"rafay-core-relay-agent","templateToken":"brhsrbbipt349qt0qae0"}]'
kind: ConfigMap
metadata:
  creationTimestamp: "2020-04-28T07:10:53Z"
  name: relay-agent-config
  namespace: rafay-system
  resourceVersion: "3127787"
  selfLink: /api/v1/namespaces/rafay-system/configmaps/relay-agent-config
  uid: e5fb5c16-c659-46d3-805f-3a2f3b73e9d7
```

`kubectl apply -f <file-name>`

**Run relay-agent**

`export POD_NAME=relay-agent-pod1`

`go run main.go --mode client --log-level 3`

Setup the local /etc/hosts file to resolve below hostname to localhost

```bash
127.0.0.1    peering.sentry.rafay.local

127.0.0.1    rx28oml.core-connector.relay.rafay.local
127.0.0.1    4qkolkn.core-connector.relay.rafay.local

127.0.0.1    rx28oml.user.relay.rafay.local
127.0.0.1     4qkolkn.user.relay.rafay.local
```

**Download kubeconfig**

`curl http://localhost:11000/v2/sentry/kubeconfig/user?opts.selector=rafay.dev/clusterName=c-4`

Save the contents to a file and use for kubectl

`kubectl --v=6 --kubeconfig /<PATH>/kubeconfig-c-5 get pod -A`

### Production

Relay server need following environment variables

- `POD_NAME` - name of the relay instaance (auto generated) POD_NAMESPACE - namespace of the relay instance
- `RAFAY_SENTRY` - URL to boostrap with sentry service RAFAY_RELAY_PEERSERVICE - To join gRPC with sentry peer service
- `RELAY_USER_DOMAIN` - domain suffix for relay user facing server RELAY_CONNECTOR_DOMAIN - domain suffix for relay connector facing domain
- `RELAY_TOKEN` - Peer Service Token bootstrapping RELAY_USER_TOKEN - Relay user facing server bootstrapping RELAY_CONNECTOR_TOKEN - Relay connector facing server bootstrapping
- `RELAY_LISTENIP` - specify the listen IP if you want servers to listen other than 0.0.0.0 (usefull in testing local)


Examples:

```bash
export RAFAY_SENTRY=http://sentry.rafay.dev:9000
export RAFAY_RELAY_PEERSERVICE=http://peering.sentry.rafay.local:7001

export RELAY_USER_DOMAIN="user.relay.rafay.local"
export RELAY_CONNECTOR_DOMAIN="core-connector.relay.rafay.local"

export RELAY_TOKEN=bqfvhabipt3a2g46986g
export RELAY_USER_TOKEN=bqfvhabipt3a2g46987g
export RELAY_CONNECTOR_TOKEN=bqfvhabipt3a2g469870
export RELAY_LISTENIP=127.0.10.1

```

## Code Structure

The following section lists out the code structure for prompt repo: 
https://github.com/RafayLabs/relay

## Need Help?

We’re there for you - the best part of being a part of an open source community. If you are stuck somewhere or are facing an issue or just don’t know how to get started, feel free to let us know.

You can reach out to us via our Slack Channel, Twitter, Discord etc.