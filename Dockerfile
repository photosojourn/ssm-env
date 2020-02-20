FROM alpine:latest

COPY ssm-env-test .
RUN chmod 755 ssm-env-test
