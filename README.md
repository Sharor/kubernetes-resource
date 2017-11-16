# Kubernetes resource for [Concourse.ci](https://concourse.ci/)

---
This resource is maintained by [Simon Groenborg](https://github.com/groenborg) and [David Johannes Christensen](https://github.com/Sharor)
---

## Acknowledgements to jcderr
The first Kubernetes resource was written by [jcderr](https://github.com/jcderr/concourse-kubernetes-resource), and we originally tried to use his resource instead of this repository. However it was unmaintained, and was breaking sporadically which caused us to initially fork the repository. 

However even with some patching, there was a fundamental problem in the approach to the resource for our use case, which did not allow us to run kubectl commands as needed. Thus this resource was born - however it is heavily inspired by jcderr's work. 

## Import the resource to your Concourse 
Simply add the following under resource_types: 
```yaml
resource_types:
- name: kubernetes
 type: docker-image
 source:
   repository: 
   tag: ''
```

This gets the resource running. 

# Source configuration
* `cluster_url` *Required* The url pointing to the Kubernetes Master API service. For example `https://192.168.99.100:8443` in [minikube](https://github.com/kubernetes/minikube). 

* `certificate_authority` *Required w/https cluster* The certificate for the cluster. For example ca.crt in [minikube](https://github.com/kubernetes/minikube). 

Format (credentials file): 
```yaml
    certificate_authority: |
      -----BEGIN CERTIFICATE-----
      Whole bunch of text and numbers
      -----END CERTIFICATE-----
  ```

* `client_certificate` *Required w/https cluster* The client certificate for admin in the cluster. For example apiserver.crt in [minikube](https://github.com/kubernetes/minikube). 

Format (credentials file): 
```yaml
    client_certificate: |
      -----BEGIN CERTIFICATE-----
      Whole bunch of text and numbers
      -----END CERTIFICATE-----
  ```

* `client_key` *Required w/https cluster* The certificate for the cluster. For example apiserver.key in [minikube](https://github.com/kubernetes/minikube). 

Format (credentials file): 
```yaml
    client_key: |
      -----BEGIN RSA PRIVATE KEY-----
      Whole bunch of text and numbers
      -----END RSA PRIVATE KEY-----
  ```

* `namespace` *Optional* Only execute on the specified namespace. Technically this can be specified in the out behaviour as well, but will then not work on `in resource`. 

# Behaviour
### `check`: Runs kubectl against the cluster.
Runs kubectl get pods, to ensure connection to cluster works. 

### `in`: Show resources in cluster.
Runs kubectl, and returns all resources in the current namespace (default if no namespace chosen). 

### `out`: Update Kubernetes cluster configuration.
Runs kubectl with the given command, thereby changing the cluster. 

#### Parameters
* `command` *Required* The command to execute on the cluster. For example: `get pods` translates to `kubectl get pods` (in default namespace). 


Work in progress. . . 


