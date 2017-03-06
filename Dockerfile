FROM scratch

MAINTAINER Christoph Witzko <docker@christophwitzko.com>

COPY hello-hostname /

EXPOSE 5000

ENTRYPOINT ["/hello-hostname"]
