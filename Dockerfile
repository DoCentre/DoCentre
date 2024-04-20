FROM golang:1.22

RUN apt-get --no-install-recommends install -y \
	make \
	&& rm -rf /var/lib/apt/lists/*
