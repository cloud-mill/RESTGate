# RESTGate

## The Cloud-Native L-7 Gateway to Seamless API Integration 🚀

--- 

## Features:

- Minimalism: 6 Golang files only
- Cloud-Native: RESTGate is engineered from the ground up for seamless integration into Kubernetes clusters. It
  specializes in intelligent ingress management and optimized routing for RESTful API traffic
- Battle-Tested: It has proven its resilience, reliability by serving world-dependent platforms that are both traffic
  intensive and computationally intensive
- Blazingly Fast: RESTGate offers low-latency and high-throughput capabilities to meet the most demanding API routing
  needs. Say goodbye to bottlenecks.
- Production-Ready: A few clicks are all it takes to make RESTGate work like magic in your cluster.

## Pre-Requisite

RESTGate requires just a single environment variable: `RESTGATE_CONFIG_PATH`. This variable specifies the path to a YAML
configuration file, which both configures RESTGate and defines the services and routes to which RESTGate will direct
traffic.

## example yaml file:

```
port: 7000
observed_services:
  - name: account-service
    service_url: "http://account-service:8000"
    routes:
      - name: account_service_health_check
        methods: "GET,OPTIONS"
        pattern: "/account-service/healthz"
        
      - name: create_new_account
        methods: "POST,OPTIONS"
        description: "create a new account."
        pattern: "/account"

      - name: get_account_info
        methods: "GET,OPTIONS"
        description: "get information about an account."
        pattern: "/account"

      - name: update_account
        methods: "POST,OPTIONS"
        description: "update an account's information."
        pattern: "/account"

      - name: delete_account
        methods: "DELETE,OPTIONS"
        description: "delete an account."
        pattern: "/account"

  - name: project-service
    service_url: "http://localhost:8002"
    routes:
      - name: project_service_health_check
        methods: "GET,OPTIONS"
        pattern: "/project-service/healthz"

      - name: create_new_project
        methods: "POST,OPTIONS"
        description: "create a new project."
        pattern: "/project"

      - name: get_project_info
        methods: "GET,OPTIONS"
        description: "get information about a project."
        pattern: "/project"

      - name: update_project
        methods: "POST,OPTIONS"
        description: "update a project's information."
        pattern: "/project"

      - name: delete_project
        methods: "DELETE,OPTIONS"
        description: "delete a project."
        pattern: "/project"
```

## yaml file explanation:

```
port: 7000 # the port RESTGate will be running on
observed_services: # the microservices running in your k8s cluster
  - name: account-service # microservice name, no specific requirements, it serves like a comment purpose
    service_url: "http://account-service:8000" # this following k8s DNS convention
    routes: # RESTFUL API routes
      - name: account_service_health_check # name for this API route, no specific requirements, it serves like a comment purpose
        methods: "GET,OPTIONS" # comma-separated api route methods
        pattern: "/account-service/healthz" # this implies, there's http://account-service:8000/account-service/healthz route available 
```

## If you doubt it, dockerise it!

1, prepare your `RESTGate.yaml`

2, build the docker image:

```
docker build --build-arg RESTATE_CONFIG_PATH=/etc/RESTGate/RESTGate.yaml -t restgate:latest .  
```

3, run the docker image

```
docker run -v whatever/your/path/to/RESTGate.yaml:/etc/RESTGate/RESTGate.yaml -p 7000:7000 restgate:latest
```

^^^ this mounts your prepared `RESTGate.yaml` to `/etc/RESTGate/RESTGate.yaml` inside the container, and assumes you
specified RESTGate running port `7000` in `RESTGate.yaml`

## Let's get it rolling into k8s cluster!

I will provide some example k8s manifest YAML files, but feel free to modify, custom and adapt to fit into your use
cases :)

1, define `restgate-service.yaml`:

```
apiVersion: v1
kind: Service
metadata:
  name: restgate
  namespace: default
spec:
  selector:
    run: restgate
  type: NodePort
  ports:
    - protocol: TCP
      port: 80
      targetPort: 7000
```

2,

```
kubectl apply -f restgate-service.yaml
```

3, provision RESTGate.yaml as a config map in k8s cluster:

```
kubectl create configmap restgate-config --from-file=whatever/your/path/to/RESTGate.yaml
```

3, define `restgate-deployment.yaml`:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: restgate
  namespace: default
spec:
  selector:
    matchLabels:
      run: restgate
  template:
    metadata:
      labels:
        run: restgate
    spec:
      containers:
          image: .../restgate:latest # docker image registry URI to previously built docker image
          imagePullPolicy: Always
          name: restgate
          ports:
            - containerPort: 7000
              protocol: TCP
          volumeMounts:
            - mountPath: /etc/RESTGate/RESTGate.yaml
              name: restgate-config
```

4,

```
kubectl apply -f restgate-deployment.yaml
```

5, hook up with your ingress!

`restgate-ingress.yaml`:

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: restgate-ingress
spec:
  defaultBackend:
    service:
      name: restgate
      port:
        number: 80
```

then:

```
kubectl apply -f restgate-ingress.yaml
```

## Contributing

At Catache, we have fundamental philosophy: Giving back more to the world than we take from it. Catache is built on
cutting-edge open-source software, and we're deeply committed to being active contributors, supporters, and sponsors
within the open-source community.

RESTGate features a minimalist design, making it exceptionally easy to modify, customize, and adapt to your specific use
cases. As a Layer 7 Gateway service, RESTGate offers extensive capabilities for handling incoming API traffic. From
logging and monitoring to implementing security controls or filters, the possibilities are vast. We're excited to see
how you can enhance RESTGate!

Just simply fork it, make a PR, our team will review your contribution and get back to you as soon as possible!

🥇 Every contributor will be honorably mentioned, and their contributions will be immortalized in the git commit history
of
this repository.

## Raise Issues

If you have questions or encounter issues with RESTGate, please feel free to open a GitHub issue. Our team will address
it as soon as possible!
