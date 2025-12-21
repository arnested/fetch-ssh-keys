FROM scratch

ARG TARGETPLATFORM
COPY $TARGETPLATFORM/fetch-ssh-keys /fetch-ssh-keys

ENTRYPOINT ["/fetch-ssh-keys"]
