FROM kio.ee/base/go:1.19 as builder
WORKDIR /go/src/app
COPY cmd cmd
COPY Makefile Makefile
COPY go.mod go.mod
RUN make small-binary
RUN chmod ug+x bin/void-app

FROM kio.ee/base/abi:edge as runner
COPY --from=builder --chown=appuser:appgroup /go/src/app/bin/void-app /void-app
USER appuser
ENTRYPOINT ["/void-app"]
EXPOSE 8080
