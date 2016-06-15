#
# based on amd64/debian-jessie/Dockerfile
#

FROM aarch64/debian:jessie

# allow replacing httpredir mirror
ARG APT_MIRROR=httpredir.debian.org
RUN sed -i s/httpredir.debian.org/$APT_MIRROR/g /etc/apt/sources.list

RUN apt-get update && apt-get install -y apparmor bash-completion btrfs-tools build-essential curl ca-certificates debhelper dh-apparmor dh-systemd git libapparmor-dev libdevmapper-dev libltdl-dev  libsqlite3-dev pkg-config libsystemd-journal-dev --no-install-recommends && rm -rf /var/lib/apt/lists/*

ENV GO_VERSION 1.6.2
RUN curl -fSL "https://github.com/hypriot/golang-armbuilds/releases/download/v${GO_VERSION}/go${GO_VERSION}.linux-arm64.tar.gz" | tar xzC /usr/src
ENV PATH $PATH:/usr/src/go/bin

ENV AUTO_GOPATH 1

ENV DOCKER_BUILDTAGS apparmor pkcs11 selinux
ENV RUNC_BUILDTAGS apparmor selinux
