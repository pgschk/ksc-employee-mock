# KSC Employee Mock API Server
A very simple mock API server, returning fake employee records

## Build
```bash
go build -o ksc-employee-mock .
```

## Run

### Commandline
```bash
./ksc-employee-mock
curl http://localhost:8080/v1/employees
```

### Docker
```bash
docker run -p 8080:8080 ghcr.io/pgschk/ksc-employee-mock
curl http://localhost:8080/v1/employees
```

### Kubernetes
```bash
kubectl create -n <namespace> -f ./k8s/deployment.yaml
kubectl port-forward -n <namespace> svc/ksc-employee-mock 8080:80
curl http://localhost:8080/v1/employees
```