We 💚 Opensource!

Yes, because we feel that it’s the best way to build and improve a product. It allows people like you from across the globe to contribute and improve a product over time. And we’re super happy to see that you’d like to contribute to Paralus dashboard.

We are always on the lookout for anything that can improve the product. Be it feature requests, issues/bugs, code or content, we’d love to see what you’ve got to make this better. If you’ve got anything exciting and would love to contribute, this is the right place to begin your journey as a contributor to Paralus dashboard and the larger open source community.

**How to get started?**
The easiest way to start is to look at existing issues and see if there’s something there that you’d like to work on. You can filter issues with the label [“[Good first issue](https://github.com/paralus/relay/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22)” which are relatively self sufficient issues and great for first time contributors.

Once you decide on an issue, please comment on it so that all of us know that you’re on it.

If you’re looking to add a new feature, [raise a new issue](https://github.com/paralus/relay/issues/new) and start a discussion with the community. Engage with the maintainers of the project and work your way through.

# Paralus relay-server / relay-agent

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

## DCO Sign off

All authors to the project retain copyright to their work. However, to ensure
that they are only submitting work that they have rights to, we are requiring
everyone to acknowledge this by signing their work.

Any copyright notices in this repo should specify the authors as "the
paralus contributors".

To sign your work, just add a line like this at the end of your commit message:

```
Signed-off-by: Joe Bloggs <joe@example.com>
```

This can easily be done with the `--signoff` option to `git commit`.
You can also mass sign-off a whole PR with `git rebase --signoff master`, replacing
`master` with the branch you are creating a pull request against, if not master.

By doing this you state that you can certify the following (from https://developercertificate.org/):

```
Developer Certificate of Origin
Version 1.1
Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129
Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.
Developer's Certificate of Origin 1.1
By making a contribution to this project, I certify that:
(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or
(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or
(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.
(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

## Need Help?

If you are interested to contribute to prompt but are stuck with any of the steps, feel free to reach out to us. Please [create an issue](https://github.com/paralus/relay/issues/new) in this repository describing your issue and we'll take it up from there.

You can reach out to us via our [Slack Channel](https://join.slack.com/t/paralus/shared_invite/zt-1a9x6y729-ySmAq~I3tjclEG7nDoXB0A) or [Twitter](https://twitter.com/paralus_).