FROM alpine:3.21.2@sha256:56fa17d2a7e7f168a043a2712e63aed1f8543aeafdcee47c58dcffe38ed51099 AS certificates

FROM scratch

COPY fetch-ssh-keys /fetch-ssh-keys
COPY --from=certificates /etc/ssl/certs/ /etc/ssl/certs/

ENTRYPOINT ["/fetch-ssh-keys"]
