FROM alpine:3.20.0@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd AS certificates

FROM scratch

COPY fetch-ssh-keys /fetch-ssh-keys
COPY --from=certificates /etc/ssl/certs/ /etc/ssl/certs/

ENTRYPOINT ["/fetch-ssh-keys"]
