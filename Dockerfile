
FROM alpine:3.9 as alpine
RUN apk add -U --no-cache ca-certificates
RUN adduser -D -g '' appuser

FROM alpine:3.9
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=alpine /etc/passwd /etc/passwd

ADD release/ /gobin/

#Cloud Run setup
ENV PORT 8080

ENTRYPOINT ["/gobin/app"]