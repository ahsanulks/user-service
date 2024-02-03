FROM mcr.microsoft.com/devcontainers/go:1-1.19-bullseye

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>
RUN apt-get update

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/golang/mock/mockgen@latest
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go install github.com/cosmtrek/air@latest

# USER root

ENV GOPROXY="https://goproxy.io,direct"
