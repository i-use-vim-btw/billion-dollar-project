# Use the official Milvus base image
FROM milvusdb/milvus:latest

# Set environment variables
ENV ETCD_USE_EMBED=true
ENV ETCD_DATA_DIR=/var/lib/milvus/etcd
ENV ETCD_CONFIG_PATH=/milvus/configs/embedEtcd.yaml
ENV COMMON_STORAGETYPE=local

# Create necessary directories and configuration files
RUN mkdir -p /var/lib/milvus /milvus/configs

# Copy configuration files to the container
COPY embedEtcd.yaml /milvus/configs/embedEtcd.yaml

# Expose necessary ports
EXPOSE 19530 9091 2379

# Healthcheck command
HEALTHCHECK --interval=30s --timeout=20s --start-period=90s --retries=3 CMD curl -f http://localhost:9091/healthz || exit 1

# Run Milvus
CMD ["milvus", "run", "standalone"]

