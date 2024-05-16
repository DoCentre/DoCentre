FROM golang:1.22

ARG UID=1000
ARG GID=1000
ARG USER=user

RUN apt-get --no-install-recommends install -y \
	make \
	&& rm -rf /var/lib/apt/lists/*

RUN groupadd -g $GID $USER \
    && useradd -u $UID -g $USER -s /bin/bash -m $USER

USER $USER

WORKDIR /home/$USER
