# Kubernetes resource for [Concourse.ci](https://concourse.ci/)

---
This resource is maintained by [Simon Groenborg](https://github.com/groenborg) and [David Johannes Christensen](https://github.com/Sharor)
---

## Acknowledgements to jcderr
The first Kubernetes resource was written by [jcderr](https://github.com/jcderr/concourse-kubernetes-resource), and we originally tried to use his resource instead of this repository. However it was unmaintained, and was breaking sporadically which caused us to initially fork the repository. 

However even with some patching, there was a fundamental problem in the approach to the resource for our use case, which did not allow us to run kubectl commands as needed. Thus this resource was born - however it is heavily inspired by jcderr's work. 

# Resource source configuration
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

ha