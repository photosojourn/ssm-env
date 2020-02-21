FROM golang:1.13-alpine AS build

#Install git
RUN apk add --no-cache git
RUN go get github.com/aws/aws-sdk-go
RUN go get github.com/photosojourn/ssm-env
WORKDIR /go/src/github.com/photosojourn/ssm-env
RUN go build -o /bin/ssm-env

FROM golang:1.13-alpine
COPY --from=build /bin/ssm-env /bin/ssm-env
ENTRYPOINT ["/bin/ssm-env"]