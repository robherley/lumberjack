## create a yaml for the configmap:
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: elasticsearch-config
  namespace: lumberjack
data:
  host: <es host>
  port: <es port>
  ibmcloud_dbs.pem: |-
    -----BEGIN CERTIFICATE-----
    <es cert>
    -----END CERTIFICATE-----

---
apiVersion: v1
kind: Secret
type: Opaque
metadata:
  name: elasticsearch-creds
  namespace: lumberjack
data:
  username: <es-user>
  password: <es-pass>
```
## create a priv user
```
$ oc adm policy add-scc-to-user privileged -z fluent-bit
```