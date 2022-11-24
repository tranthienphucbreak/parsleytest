RUN
    cd cmd/webapp
    go build -o ./../../webapp
    cd ../..
    ./webapp

Docker Build
    docker build --tag parsley_webapp .
