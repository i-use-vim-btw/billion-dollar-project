from pymilvus import MilvusClient
import json

# 1. Set up a Milvus client
client = MilvusClient(
    uri="http://127.0.0.1:34779"
)

collection_name = "quick_setup"

res = client.search(
    collection_name=collection_name,
    data=[[1.0, 2.0, 2.0, 1.0, 0.0]],
    limit=1,
    search_params={"metric_type": "IP", "params": {}}

)
result = json.dumps(res, indent=4)
print(result)
