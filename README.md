# Paralus relay-server / relay-agent
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fparalus%2Frelay.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fparalus%2Frelay?ref=badge_shield)


Relay server and agent for kubectl access.

## Development Setup

Relay depends on [paralus](https://github.com/paralus/paralus) and you need to check out this repo.

### Using docker-compose

Run following Docker Compose command to setup all requirements like Postgres db, Kratos etc. for core.

This will start up postgres and elasticsearch as well as kratos and run the kratos migrations. It will also run all the necessary migrations. It also starts up a mail slurper for you to use Kratos.

  `docker-compose --env-file ./env.example up -d`

Start paralus:

  `go run github.com/paralus/paralus`

### Paralus Relay Server

Get peer client token for relay.

Open the admindb postgress and execute below SQL to fetch the token.

```bash
admindb=# select token from sentry_bootstrap_agent_template where name='paralus-sentry-peering-client';
        token         
----------------------
 brhsrbbipt349qt0qad0

```

Use above token set the environment variable for relay-server

```bash
export RELAY_PEERING_TOKEN=brhsrbbipt349qt0qad0
export PARALUS_RELAY_PEERSERVICE=http://peering.sentry.paralus.local:10001 
export POD_NAME=relay-pod1
go run main.go --mode server --log-level 3
```

### Paralus Relay Agent

Get the config for paralus relay agent

```bash
curl http://localhost:11000/infra/v3/project/:metadata.project/cluster/:metadata.name/download
```

Use the values of relay-agent-config and apply in the cluster.

```bash
kubectl -n paralus-system get cm/relay-agent-config -o yaml

apiVersion: v1
data:
  clusterID: 4qkolkn
  relays: '[{"token":"brht3cjipt35cb3kfjag","addr":"api.sentry.paralus.local:11000","endpoint":"*.core-connector.relay.paralus.local:443","name":"paralus-core-relay-agent","templateToken":"brhsrbbipt349qt0qae0"}]'
kind: ConfigMap
metadata:
  creationTimestamp: "2020-04-28T07:10:53Z"
  name: relay-agent-config
  namespace: paralus-system
  resourceVersion: "3127787"
  selfLink: /api/v1/namespaces/paralus-system/configmaps/relay-agent-config
  uid: e5fb5c16-c659-46d3-805f-3a2f3b73e9d7
```

```bash
kubectl apply -f <file-name>
```

Run relay-agent

```bash
export POD_NAME=relay-agent-pod1
go run main.go --mode client --log-level 3
```

Setup the local /etc/hosts file to resolve below hostname to localhost

```bash
127.0.0.1	peering.sentry.paralus.local

127.0.0.1	rx28oml.core-connector.relay.paralus.local
127.0.0.1	4qkolkn.core-connector.relay.paralus.local

127.0.0.1	rx28oml.user.relay.paralus.local
127.0.0.1 4qkolkn.user.relay.paralus.local
```

Download kubeconfig

```bash
curl http://localhost:11000/v2/sentry/kubeconfig/user?opts.selector=paralus.dev/clusterName=c-4
```

Save the contents to a file and use for kubectl

```bash
kubectl --v=6 --kubeconfig /<PATH>/kubeconfig-c-5 get pod -A
```

## Examples

```bash
export PARALUS_SENTRY=http://sentry.paralus.dev:9000
export PARALUS_RELAY_PEERSERVICE=http://peering.sentry.paralus.local:7001

export RELAY_USER_DOMAIN="user.relay.paralus.local"
export RELAY_CONNECTOR_DOMAIN="core-connector.relay.paralus.local"

export RELAY_TOKEN=bqfvhabipt3a2g46986g
export RELAY_USER_TOKEN=bqfvhabipt3a2g46987g
export RELAY_CONNECTOR_TOKEN=bqfvhabipt3a2g469870
export RELAY_LISTENIP=127.0.10.1
```

## Community & Support

- Visit [Paralus website](https://paralus.io) for the complete documentation and helpful links.
- Join our [Slack channel](https://join.slack.com/t/paralus/shared_invite/zt-1a9x6y729-ySmAq~I3tjclEG7nDoXB0A) to post your queries and discuss features.
- Tweet to [@paralus_](https://twitter.com/paralus_/) on Twitter.
- Create [GitHub Issues](https://github.com/paralus/relay/issues) to report bugs or request features.

## Contributing

The easiest way to start is to look at existing issues and see if there’s something there that you’d like to work on. You can filter issues with the label “Good first issue” which are relatively self sufficient issues and great for first time contributors.

Once you decide on an issue, please comment on it so that all of us know that you’re on it.

If you’re looking to add a new feature, raise a [new issue](https://github.com/paralus/relay/issues) and start a discussion with the community. Engage with the maintainers of the project and work your way through.


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fparalus%2Frelay.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fparalus%2Frelay?ref=badge_large)