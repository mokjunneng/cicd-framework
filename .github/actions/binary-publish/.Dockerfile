# Base image
FROM golang:1.16.3-alpine3.13

# Copies code file repository to filesystem
COPY entrypoint.sh /entrypoint.sh

# change permission to execute script
RUN chmod +x /entrypoint.sh

# File to execute when the docker container starts up
ENTRYPOINT ["/entrypoint.sh"]
