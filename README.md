# KSC Employee Mock API Server
A very simple mock API server, returning fake employee records

## Build
```bash
go build -o ksc-employee-mock .
```

## Run
```bash
./ksc-employee-mock
curl http://localhost:8080/v1/employees
```