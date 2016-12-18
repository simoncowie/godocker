FROM golang
 
ADD . /go/src/github.com/simoncowie/godocker
RUN go install github.com/simoncowie/godocker

EXPOSE 8080

ENTRYPOINT /go/bin/godocker