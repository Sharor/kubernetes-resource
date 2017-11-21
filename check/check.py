from datetime import datetime
import sys
import json
from kubernetes import client, config

#Remember to close stdin

def check():
    """check"""
    data = json.load(sys.stdin)

    if not data["source"]["cluster_url"].startswith("https"):
        sys.stderr.write("missing encryption for connection")

    config.load_kube_config()
    v1 = client.CoreV1Api()
    pods = v1.list_pod_for_all_namespaces(watch=False, timeout_seconds=30)
    for i in pods.items:
        sys.stderr.write("%s\t%s\t%s\n" % (i.status.pod_ip, i.metadata.namespace, i.metadata.name))

    output = {
        "version": [ datetime.now().isoformat()]
    }

    print(json.dumps(output))

if __name__ == "__main__":
    check()
