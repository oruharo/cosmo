# Workspace

`Workspace` is a instance of DevEnvironment defined as Kubernetes CRD.

When you create Workspace, actual Workspace Pod will be started.

It is not an only container but its deployment, networking and storage of DevEnvironment.

In Kubernetes term, it is a set of Kubernetes resources such as Deployment, Service, PersistentVolumeClaim and so on.

```yaml
apiVersion: cosmo-workspace.github.io/v1alpha1
kind: Workspace
metadata:
  name: ws1
  namespace: cosmo-user-tom # User namespace which must start with 'cosmo-user-'
spec:
  template:
    name: nodejs-code-server-template # Name of WorkspaceTemplate
  network:
  - customHostPrefix: main # WebIDE server routing rule
    httpPath: /
    portNumber: 8080
    protocol: http
    public: false
```

## WorkspaceTemplate

WorkspaceTemplate is a set of Kubernetes manifests for Workspace.

<details>
<summary>cosmo-template.yaml</summary>

```yaml
# Generated by cosmoctl template command
apiVersion: cosmo-workspace.github.io/v1alpha1
kind: Template
metadata:
  annotations:
    workspace.cosmo-workspace.github.io/deployment: code-server
    workspace.cosmo-workspace.github.io/service: code-server
    workspace.cosmo-workspace.github.io/service-main-port: http
  labels:
    cosmo-workspace.github.io/type: workspace
  name: code-server-example
spec:
  description: Example Workspace Template for code-server
  requiredVars:
  - default: "20"
    var: STORAGE_SIZE_GB
  rawYaml: |
    apiVersion: v1
    kind: Service
    metadata:
      labels:
        app.kubernetes.io/instance: code-server
        app.kubernetes.io/managed-by: Helm
        app.kubernetes.io/name: code-server
        cosmo-workspace.github.io/instance: '{{INSTANCE}}'
        cosmo-workspace.github.io/template: '{{TEMPLATE}}'
        helm.sh/chart: code-server-1.0.5
      name: '{{INSTANCE}}-code-server'
      namespace: '{{NAMESPACE}}'
    spec:
      ports:
      - name: http
        port: 8080
        protocol: TCP
        targetPort: http
      selector:
        app.kubernetes.io/instance: code-server
        app.kubernetes.io/name: code-server
        cosmo-workspace.github.io/instance: '{{INSTANCE}}'
        cosmo-workspace.github.io/template: '{{TEMPLATE}}'
      type: ClusterIP
    ---
    apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
      labels:
        app.kubernetes.io/instance: code-server
        app.kubernetes.io/managed-by: Helm
        app.kubernetes.io/name: code-server
        cosmo-workspace.github.io/instance: '{{INSTANCE}}'
        cosmo-workspace.github.io/template: '{{TEMPLATE}}'
        helm.sh/chart: code-server-1.0.5
      name: '{{INSTANCE}}-data'
      namespace: '{{NAMESPACE}}'
    spec:
      accessModes:
      - ReadWriteOnce
      resources:
        requests:
          storage: '{{STORAGE_SIZE_GB}}Gi'
      volumeMode: Filesystem
      storageClassName: gp2
    ---
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      labels:
        app.kubernetes.io/instance: code-server
        app.kubernetes.io/managed-by: Helm
        app.kubernetes.io/name: code-server
        cosmo-workspace.github.io/instance: '{{INSTANCE}}'
        cosmo-workspace.github.io/template: '{{TEMPLATE}}'
        helm.sh/chart: code-server-1.0.5
      name: '{{INSTANCE}}-code-server'
      namespace: '{{NAMESPACE}}'
    spec:
      replicas: 1
      selector:
        matchLabels:
          app.kubernetes.io/instance: code-server
          app.kubernetes.io/name: code-server
          cosmo-workspace.github.io/instance: '{{INSTANCE}}'
          cosmo-workspace.github.io/template: '{{TEMPLATE}}'
      strategy:
        type: Recreate
      template:
        metadata:
          labels:
            app.kubernetes.io/instance: code-server
            app.kubernetes.io/name: code-server
            cosmo-workspace.github.io/instance: '{{INSTANCE}}'
            cosmo-workspace.github.io/template: '{{TEMPLATE}}'
        spec:
          containers:
          - args:
            - --auth=none
            image: codercom/code-server:3.12.0
            imagePullPolicy: Always
            livenessProbe:
              httpGet:
                path: /
                port: http
            name: code-server
            ports:
            - containerPort: 8080
              name: http
              protocol: TCP
            readinessProbe:
              httpGet:
                path: /
                port: http
            resources: {}
            securityContext:
              runAsUser: 1000
            volumeMounts:
            - mountPath: /home/coder
              name: data
          initContainers:
          - command:
            - sh
            - -c
            - |
              chown -R 1000:1000 /home/coder
            image: busybox:latest
            imagePullPolicy: IfNotPresent
            name: init-chmod-data
            securityContext:
              runAsUser: 0
            volumeMounts:
            - mountPath: /home/coder
              name: data
          securityContext:
            fsGroup: 1000
          serviceAccountName: default
          volumes:
          - name: data
            persistentVolumeClaim:
              claimName: '{{INSTANCE}}-data'
```

</details><br>

Template may be little bit large and K8s YAMLs are in single string properties, but don't worry.

As you can see `# Generated by cosmoctl template command` top of the yaml, WorkspaceTemplate can be generated via `cosmoctl`.

### Create WorkspaceTemplate by Kustomize

1.  Prepare the set of Kubernetes YAML by your own.

    All you have to do is to prepare your own Kubernetes YAMLs that is deployable.

    For example, you prepare like here.

    ```sh
    $ ls /tmp/code-server-example
    kustomization.yaml   deployment.yaml   service.yaml   pvc.yaml
    ```

    <details>
    <summary>kustomization.yaml</summary>

    ```yaml
    apiVersion: kustomize.config.k8s.io/v1beta1
    kind: Kustomization

    namespace: default

    resources:
    - deployment.yaml
    - service.yaml
    - pvc.yaml
    ```

    </details><br>

    <details>
    <summary>service.yaml</summary>

    ```yaml
    apiVersion: v1
    kind: Service
    metadata:
      name: code-server
    spec:
      ports:
      - name: http
        port: 8080
        protocol: TCP
        targetPort: http
      selector:
        app: code-server
      type: ClusterIP
    ```

    </details><br>

    <details>
    <summary>pvc.yaml</summary>

    ```yaml
    apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
      name: data
    spec:
      accessModes:
      - ReadWriteOnce
      resources:
        requests:
          storage: 20Gi
      volumeMode: Filesystem
      storageClassName: gp2
    ```

    </details><br>

    <details>
    <summary>deployment.yaml</summary>

    ```yaml
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: code-server
    spec:
      replicas: 1
      selector:
        matchLabels:
          app: code-server
      template:
        metadata:
          labels:
            app: code-server
        spec:
          containers:
          - args:
            - --auth=none
            image: codercom/code-server:3.12.0
            imagePullPolicy: Always
            livenessProbe:
              httpGet:
                path: /
                port: http
            name: code-server
            ports:
            - containerPort: 8080
              name: http
              protocol: TCP
            readinessProbe:
              httpGet:
                path: /
                port: http
            resources: {}
            securityContext:
              runAsUser: 1000
            volumeMounts:
            - mountPath: /home/coder
              name: data
          initContainers:
          - command:
            - sh
            - -c
            - |
              chown -R 1000:1000 /home/coder
            image: busybox:latest
            imagePullPolicy: IfNotPresent
            name: init-chmod-data
            securityContext:
              runAsUser: 0
            volumeMounts:
            - mountPath: /home/coder
              name: data
          securityContext:
            fsGroup: 1000
          serviceAccountName: default
          volumes:
          - name: data
            persistentVolumeClaim:
              claimName: data
    ```
    </details><br>

    Be sure it can be deployed

    ```sh
    kustomize build . | kubectl apply -f - --dry-run=server
    ```

    > ✅Note:
    >
    > Be sure it can be deployed as a single WebIDE deployment without any networking or storage troubles!
    >
    > Delete them if you applied the manifest actually. `kustomize build . | kubectl delete -f -`

2.  Generate WorkspaceTemplate

    Pass kustomize-generated manifest to `cosmoctl template gen` command by stdin.

    ```sh
    kustomize build . | cosmoctl tmpl gen --workspace -o cosmo-template.yaml
    ```

    > ✅Note: 
    >
    > Be sure to execute with `--workspace` option.


### Create WorkspaceTemplate by Helm

1.  Prepare Helm chart for the WorkspaceTemplate.

    Here we use the COSMO development code-server helm charts.
    https://github.com/cosmo-workspace/cosmo-dev

    ```sh
    helm repo add cosmo https://cosmo-workspace.github.io/charts
    ```

    Deploy

    ```sh
    helm template code-server-example cosmo/dev-code-server | kubectl apply -f - --dry-run=server
    ``` 

    > ✅Note:
    >
    > Be sure it can be deployed as a single WebIDE deployment without any networking or storage troubles!
    >

2.  Generate WorkspaceTemplate

    Pass Helm-generated manifests to `cosmoctl template gen` command by stdin.

    ```sh
    helm template code-server-example cosmo/dev-code-server | cosmoctl tmpl gen --workspace -o cosmo-template.yaml
    ```

    > ✅Note: 
    >
    > In this example chart `cosmo/dev-code-server`, the error will occur:
    >  
    > `Error: type workspace validation failed: failed to specify the service port`
    >
    > To recognize the WebIDE URL, the service port included in the given Kubernetes manifests is read.
    > If there are multiple service ports, the WebIDE port must be specified by flag.
    >
    > ```sh
    > helm template code-server-example cosmo/dev-code-server | cosmoctl tmpl gen --workspace \
    >     --workspace-main-service-port-name http
    > ```


## Annotations

| Annotatio keys | Avairable values(default) | Description | cosmoctl option |
|:--|:--|:--|:--|
| `workspace.cosmo-workspace.github.io/deployment` | Deployment name(automatically recognized in input) | Deployment name of WebIDE Container | `--workspace-deployment-name` |
| `workspace.cosmo-workspace.github.io/service` | Service port name(automatically recognized in input) | Service name which is WebIDE Serivce | `--workspace-service-name` |
| `workspace.cosmo-workspace.github.io/service-main-port` | Service port name(automatically recognized in input) | Service port name which is for WebIDE URL | `--workspace-main-service-port-name` |
| `cosmo-workspace.github.io/required-useraddons` | comma-separated UserAddon names(None)  | User who use this Template must be attached all of the UserAddons specified in this annotation | `--required-useraddons` |


### More infomation

When you create `Workspace`, you can also see the Kubernetes resource `Instance` is created.

See [TEMPLATE-ENGINE.md](https://github.com/cosmo-workspace/cosmo/blob/main/docs/TEMPLATE-ENGINE.md) for deepdive into `Instance`.