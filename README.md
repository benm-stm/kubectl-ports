# kubectl-ports
organized display of service ports using kubectl
Example display
```
$ kubectl ports svc -n ingress-controller
+--------------------------------+----------+----------+------+------------+----------+-------------+
|          SERVICE NAME          |   NAME   | PROTOCOL | PORT | TARGETPORT | NODEPORT | APPPROTOCOL |
+--------------------------------+----------+----------+------+------------+----------+-------------+
| haproxy-pub-kubernetes-ingress |
|                                | http     | TCP      |   80 | http       |    30836 |             |
|                                | https    | TCP      |  443 | https      |    32653 |             |
|                                | stat     | TCP      | 1024 | stat       |    32191 |             |
|                                | tcp-1883 | TCP      | 1883 |            |    32410 |             |
|                                | tcp-1884 | TCP      | 1884 |            |    30376 |             |
|                                | tcp-1885 | TCP      | 1885 |            |    31542 |             |
|                                | tcp-5671 | TCP      | 5671 |            |    32672 |             |
|                                | tcp-5672 | TCP      | 5672 |            |    30508 |             |
|                                | tcp-8000 | TCP      | 8000 |            |    31607 |             |
|                                | tcp-8443 | TCP      | 8443 |            |    32342 |             |
|                                | tcp-8883 | TCP      | 8883 |            |    32384 |             |
|                                | tcp-9000 | TCP      | 9000 |            |    30687 |             |
|                                | tcp-9443 | TCP      | 9443 |            |    31021 |             |
+--------------------------------+----------+----------+------+------------+----------+-------------+
+------------------------------------------------+------+----------+------+------------+----------+-------------+
|                  SERVICE NAME                  | NAME | PROTOCOL | PORT | TARGETPORT | NODEPORT | APPPROTOCOL |
+------------------------------------------------+------+----------+------+------------+----------+-------------+
| haproxy-pub-kubernetes-ingress-default-backend |
|                                                | http | TCP      | 8080 | http       |        0 | http        |
+------------------------------------------------+------+----------+------+------------+----------+-------------+
+--------------------------------+----------+----------+------+------------+----------+-------------+
|          SERVICE NAME          |   NAME   | PROTOCOL | PORT | TARGETPORT | NODEPORT | APPPROTOCOL |
+--------------------------------+----------+----------+------+------------+----------+-------------+
| haproxy-sub-kubernetes-ingress |
|                                | http     | TCP      |   80 | http       |    30102 |             |
|                                | https    | TCP      |  443 | https      |    31922 |             |
|                                | stat     | TCP      | 1024 | stat       |    32383 |             |
|                                | tcp-1883 | TCP      | 1883 |            |    32411 |             |
|                                | tcp-1884 | TCP      | 1884 |            |    31216 |             |
|                                | tcp-1885 | TCP      | 1885 |            |    30227 |             |
|                                | tcp-5671 | TCP      | 5671 |            |    32097 |             |
|                                | tcp-5672 | TCP      | 5672 |            |    31769 |             |
|                                | tcp-8000 | TCP      | 8000 |            |    30656 |             |
|                                | tcp-8443 | TCP      | 8443 |            |    31864 |             |
|                                | tcp-8883 | TCP      | 8883 |            |    30564 |             |
|                                | tcp-9000 | TCP      | 9000 |            |    30290 |             |
|                                | tcp-9443 | TCP      | 9443 |            |    32147 |             |
+--------------------------------+----------+----------+------+------------+----------+-------------+
+------------------------------------------------+------+----------+------+------------+----------+-------------+
|                  SERVICE NAME                  | NAME | PROTOCOL | PORT | TARGETPORT | NODEPORT | APPPROTOCOL |
+------------------------------------------------+------+----------+------+------------+----------+-------------+
| haproxy-sub-kubernetes-ingress-default-backend |
|                                                | http | TCP      | 8080 | http       |        0 | http        |
+------------------------------------------------+------+----------+------+------------+----------+-------------+
```

# Prerequisits
Go 1.18
# How to integrate
Build the project to generate a binary
```
$ cd kubectl-pots
$ go build
```
search for your kubectl plugins list to get the plugins dir
```
$kubectl plugin list
kubectl plugin list                                             

The following compatible plugins are available:

/home/raf/.krew/bin/kubectl-krew
/home/raf/.krew/bin/kubectl-ports
/home/raf/.krew/bin/kubectl-preflight
```

Copy the generated binary to your plugins location
```
$ cp kubectl-ports /home/raf/.krew/bin/
```
Source the plugin's autocomplete
```
$ source <(kubectl ports completion zsh)
```
