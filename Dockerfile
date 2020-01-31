FROM golang:latest AS base

MAINTAINER devansh42
RUN mkdir /tmp/runner
COPY . /tmp/runner
WORKDIR /tmp/runner
RUN go build -o runner .

RUN mkdir /runner
ARG TF_ORG_TOKEN
FROM alpine
WORKDIR /tmp/momo-runner
COPY --from=base /tmp/runner/runner /runner/ 
CMD ["-org_token=${TF_ORG_TOKEN}"]
ENTRYPOINT ["/runner/runner"]  

