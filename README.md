## RUN
    `cd cmd/webapp`
    `go build -o ./../../webapp`
    `cd ../..`
    `./webapp`

## Docker Build
    `docker build --tag parsley_webapp .`

## UT
    `/cmd/webapp/routes/routes_test.go`
    `/internal/patient/service_test.go`

## RUN UT
    `go test ./cmd/webapp/routes/`
    `go test ./internal/patient/`

## Functional Test
    `/test/atf/src/automation/suite/patient_functional_test.go`

## RUN Functional Test
    `./webapp`
    `go test ./test/atf/src/automation/suite`
