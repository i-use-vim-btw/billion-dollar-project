from pymilvus import connections, FieldSchema, CollectionSchema, DataType, Collection
import random

# Define the service DNS name and port
namespace = input("enter namespace: ")
service_name = input("enter service name: ")
milvus_service_dns = f"{service_name}.{namespace}.svc.cluster.local"
milvus_service_port = 19530

# Connect to the Milvus instance using the service DNS name and port
connections.connect("default", host=milvus_service_dns, port=milvus_service_port)

# Define a schema for the collection
fields = [
    FieldSchema(name="id", dtype=DataType.INT64, is_primary=True, auto_id=True),
    FieldSchema(name="vector", dtype=DataType.FLOAT_VECTOR, dim=4),
]
schema = CollectionSchema(fields, "example_collection")

# Create a collection
collection = Collection(name="example_collection", schema=schema)

# Insert some random vectors
vectors = [[random.random() for _ in range(4)] for _ in range(4)]
collection.insert([vectors])

# Load the collection to memory
collection.load()

# Create an index for the collection
index_params = {"index_type": "IVF_FLAT", "metric_type": "L2", "params": {"nlist": 128}}
collection.index("vector", index_params)

# Perform a search with a random vector
query_vector = [[random.random() for _ in range(4)]]
results = collection.search(
    query_vector,
    "vector",
    params={"metric_type": "L2", "params": {"nprobe": 10}},
    limit=2,
)

# Print the results
for result in results:
    print(f"Result: id={result.id}, distance={result.distance}")
