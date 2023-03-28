# Base image
FROM debian

# Install dependencies
RUN apt-get update && apt-get install -y procps
RUN apt-get update && apt-get install -y ca-certificates gnupg2
RUN gpg -k

RUN gpg --no-default-keyring --keyring /usr/share/keyrings/k6-archive-keyring.gpg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
RUN echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" | tee /etc/apt/sources.list.d/k6.list
RUN apt-get update
RUN apt-get install k6


# Set working directory
WORKDIR /app

# Set default command
CMD [ "bash" ]
