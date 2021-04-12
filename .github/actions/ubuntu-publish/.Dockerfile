FROM ubuntu:latest
RUN apt update -y
RUN apt install -y gnupg dput dh-make devscripts lintian
