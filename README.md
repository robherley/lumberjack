# lumberjack
Pipe k8s/openshift pod logs into an elasticsearch instance using fluentbit and
pull back logs using websockets.

Example Environment:
```
ENV=<development|production>
PORT=<port to run server on>
ELASTIC_HOST=<elastic host>
ELASTIC_PORT=<elastic port>
ELASTIC_USER=<elastic user>
ELASTIC_PASS=<elastic password>
```

## fluent-bit daemonset
yaml configs located in `fluent-bit/`