# syntax=docker/dockerfile:1
FROM golang:1.19-alpine
ENV PORT 8080
ENV HOSTDIR 0.0.0.0

EXPOSE 8080
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
COPY . ./
RUN go build -o /main
CMD [ "/main" ]

# # Use the official Prometheus image as a base
# FROM prom/prometheus:v2.33.1

# # Set the working directory in the container
# WORKDIR /etc/prometheus

# # Copy your configuration file into the container
# COPY prometheus.yml /etc/prometheus/prometheus.yml

# # Expose the port Prometheus runs on
# # EXPOSE 9090

# # Run Prometheus
# CMD [ "--config.file=/etc/prometheus/prometheus.yml", \
#       "--storage.tsdb.path=/prometheus", \
#       "--web.enable-lifecycle", \
#       "--web.enable-admin-api" ]