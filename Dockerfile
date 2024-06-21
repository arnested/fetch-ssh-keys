FROM alpine:3.20.1@sha256:b89d9c93e9ed3597455c90a0b88a8bbb5cb7188438f70953fede212a0c4394e0 AS certificates

FROM scratch

COPY fetch-ssh-keys /fetch-ssh-keys
COPY --from=certificates /etc/ssl/certs/ /etc/ssl/certs/

ENTRYPOINT ["/fetch-ssh-keys"]
