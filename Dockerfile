FROM alpine:3.19.1@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b AS certificates

FROM scratch

COPY fetch-ssh-keys /fetch-ssh-keys
COPY --from=certificates /etc/ssl/certs/ /etc/ssl/certs/

ENTRYPOINT ["/fetch-ssh-keys"]
