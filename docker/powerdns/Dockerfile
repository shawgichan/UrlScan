FROM debian:bullseye-slim

ENV DEBIAN_FRONTEND=noninteractive

# Install required packages
RUN apt-get update && apt-get install -y \
    wget \
    gnupg \
    dnsutils \
    apt-transport-https \
    ca-certificates \
    curl \
    procps \
    net-tools \
    git \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Add PowerDNS repository
RUN curl https://repo.powerdns.com/FD380FBB-pub.asc | apt-key add - \
    && echo "deb [arch=amd64] http://repo.powerdns.com/debian bullseye-rec-51 main" > /etc/apt/sources.list.d/powerdns.list

# Install PowerDNS Recursor
RUN apt-get update && apt-get install -y \
    pdns-recursor \
    && rm -rf /var/lib/apt/lists/*

# Install Go
ENV GO_VERSION=1.21.5
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz

# Add Go to PATH
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/project/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Create user and set permissions
RUN useradd -m pdnsuser

RUN mkdir -p /var/run/pdns-recursor /project/go \
    && chown -R pdnsuser:pdnsuser /var/run/pdns-recursor /project

WORKDIR /project

EXPOSE 53/udp 53/tcp

COPY docker/powerdns/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh && chown pdnsuser:pdnsuser /entrypoint.sh

USER pdnsuser

ENTRYPOINT ["/entrypoint.sh"]