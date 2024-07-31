import sys
import os

# Add the src directory to PYTHONPATH
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..', 'src')))

from gandi.models import (
    CollectionRequest,
    CreateVectorsInsertRequest,
    CreateVectorsGetRequest,
    CreateVectorsDeleteRequest,
)
from gandi import Gandi
coll_name = "test107"
host = "127.0.0.1:44179"

sdk = Gandi(access_token="YOUR_ACCESS_TOKEN", base_url="http://localhost:8080")

request_body = CollectionRequest(
    collection_name=coll_name, dimension=8, host=host)

result = sdk.collections.create_collections_create(request_body=request_body)

"""
print(result)

result = sdk.collections.create_collections_describe(
    CreateCollectionsDescribeRequest(collection_name=coll_name, host=host)
)
"""

print(result)

result = sdk.vectors.create_vectors_insert(
    CreateVectorsInsertRequest(
        collection_name=coll_name,
        data=[{"vector": [i / 10] * 8, "id": i} for i in range(1, 10)],
        host=host,
    )
)

print(result)

result = sdk.vectors.create_vectors_get(
    CreateVectorsGetRequest(collection_name=coll_name,
                            id_=[1, 3, 5], host=host)
)

print(result)

result = sdk.vectors.create_vectors_delete(
    CreateVectorsDeleteRequest(
        collection_name=coll_name, filter="id in [2, 4]", host=host
    )
)

print(result)


result = sdk.vectors.create_vectors_get(
    CreateVectorsGetRequest(collection_name=coll_name,
                            id_=[1, 2, 3, 4], host=host)
)

print(result)
