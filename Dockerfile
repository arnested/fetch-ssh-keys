FROM alpine:3.22.0@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715 AS certificates

FROM scratch

COPY fetch-ssh-keys /fetch-ssh-keys
COPY --from=certificates /etc/ssl/certs/ /etc/ssl/certs/

ENTRYPOINT ["/fetch-ssh-keys"]
