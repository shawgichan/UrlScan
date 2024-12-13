FROM debian:bullseye-slim

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    wget \
    gnupg \
    apt-transport-https \
    ca-certificates \
    curl \
    procps \
    net-tools \
    && rm -rf /var/lib/apt/lists/*

RUN curl https://repo.powerdns.com/FD380FBB-pub.asc | apt-key add - \
    && echo "deb [arch=amd64] http://repo.powerdns.com/debian bullseye-rec-51 main" > /etc/apt/sources.list.d/powerdns.list

RUN apt-get update && apt-get install -y \
    pdns-recursor \
    && rm -rf /var/lib/apt/lists/*
    
RUN useradd -m pdnsuser

RUN mkdir -p /var/run/pdns-recursor /project \
    && chown -R pdnsuser:pdnsuser /var/run/pdns-recursor /project


RUN chown -R pdnsuser:pdnsuser /project

WORKDIR /project

EXPOSE 53/udp 53/tcp

COPY docker/powerdns/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh && chown pdnsuser:pdnsuser /entrypoint.sh

USER pdnsuser

ENTRYPOINT ["/entrypoint.sh"]
