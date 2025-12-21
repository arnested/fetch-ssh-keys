FROM scratch

COPY fetch-ssh-keys /fetch-ssh-keys

ENTRYPOINT ["/fetch-ssh-keys"]
