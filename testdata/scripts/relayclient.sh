#curl -kv --cert ../certs/clientcerts/relayserver1-ABCD-123456.relay.paralus.dev.crt --key ../certs/clientcerts/relayserver1-ABCD-123456.relay.paralus.dev.key https://cluster1.kubectl.relay.paralus.dev/


#curl -kv --cert ../certs/clientcerts/star.namespace-admin-sa.crt --key ../certs/clientcerts/star.namespace-admin-sa.key https://cluster1.kubectl.relay.paralus.dev/api?timeout=32s


#curl  --http1.1 -kv --cert ../certs/clientcerts/star.namespace-admin-sa.crt --key ../certs/clientcerts/star.namespace-admin-sa.key https://cluster1.kubectl.relay.paralus.dev/apis/rbac.authorization.k8s.io/v1/?timeout=32s


#curl  --http1.1 -kv --cert ../certs/clientcerts/star.namespace-admin-sa.crt --key ../certs/clientcerts/star.namespace-admin-sa.key https://cluster1.kubectl.relay.paralus.dev/

curl --http1.1 -kv --cert ../certs/clientcerts/star.namespace-admin-sa.crt --key ../certs/clientcerts/star.namespace-admin-sa.key POST --data '{"username":"xyz","password":"xyz"}' https://cluster1.kubectl.relay.paralus.dev/apis/rbac.authorization.k8s.io/v1/?timeout=32s

#curl -kv --cert ../certs/clientcerts/star.namespace-admin-sa.crt --key ../certs/clientcerts/star.namespace-admin-sa.key POST --data '{"username":"xyz","password":"xyz"}' https://cluster1.kubectl.relay.paralus.dev/apis/rbac.authorization.k8s.io/v1?timeout=32s

#curl --http1.1 -kv --cert ../certs/clientcerts/star.namespace-admin-sa.crt --key ../certs/clientcerts/star.namespace-admin-sa.key POST --data '{"username":"xyz","password":"xyz"}' https://cluster1.kubectl.relay.paralus.dev/api/v1/namespaces/default/pods/myapp1-pod/exec?command=%2Fbin%2Fsh&container=myapp1-container&stdin=true&stdout=true&tty=true \
