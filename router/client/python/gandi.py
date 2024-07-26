import requests


# Don't forget timeout and error handling
class GandiClient:
    def __init__(self, uri, api_key="default") -> None:
        self.__URI = uri
        self.header = {"Content-Type": "application/json", "Api_key": api_key}

    def create_collection(
        self, collection_name="default", dim=128, id="localhost:8080"
    ) -> None:

        data = {"collectionName": collection_name, "dimension": dim, "id": id}
        res = requests.post(
            "http://" + self.__URI + "/gandi/collections/create",
            headers=self.header,
            json=data,
        )

        res = res.json()

        print(res)

    # Not finished, do not use
    def create_index(
        self,
        collection_name="default",
        field_name="vector",
        params={"nlist": 16384, "index_type": "IVF_FLAT"},
        id="localhost:19530",
    ):
        data = {
            "collectionName": collection_name,
            "indexParams": [{"fieldName": field_name, "params": params}],
            "id": id,
        }

        res = requests.post(
            "http://" + self.__URI + "/gandi/indexes/create",
            headers=self.header,
            json=data,
        )

        res = res.json()

        print(res)

    def insert(self, collection_name="default", data=[], id="localhost:19530"):
        in_data = {"data": data, "collectionName": collection_name, "id": id}

        res = requests.post(
            "http://" + self.__URI + "/gandi/vectors/insert",
            headers=self.header,
            json=in_data,
        )

        res = res.json()

        print(res)

    def get(self, collection_name="default", ids=[], id="localhost:19530"):
        data = {"collectionName": collection_name, "id": ids, "host": id}

        res = requests.post(
            "http://" + self.__URI + "/gandi/vectors/get",
            headers=self.header,
            json=data,
        )

        res = res.json()

        print(res)

    def upsert(self, collection_name="default", data=[], id="localhost:19530"):
        up_data = {"data": data, "collectionName": collection_name, "id": id}

        res = requests.post(
            "http://" + self.__URI + "/gandi/vectors/upsert",
            headers=self.header,
            json=up_data,
        )

        res = res.json()

        print(res)

    def delete(
        self,
        collection_name="default",
        database_name="default",
        partition_name="_default",
        filter="",
        id="localhost:19530",
    ):
        delete_data = {
            "collectionName": collection_name,
            "databaseName": database_name,
            "partitionName": partition_name,
            "filter": filter,
            "id": id,
        }

        res = requests.post(
            "http://" + self.__URI + "/gandi/vectors/delete",
            headers=self.header,
            json=delete_data,
        )

        res = res.json()

        print(res)

    def get_stats(self, collection_name="default", id="localhost:19530"):
        data = {"collectionName": collection_name, "id": id}

        res = requests.post(
            "http://" + self.__URI + "/gandi/collections/get_load_state",
            headers=self.header,
            json=data,
        )

        res = res.json()

        print(res)


client = GandiClient("localhost:8080")


collName = "test68"

client.get_stats(collection_name=collName)

client.create_collection(collection_name=collName, dim=5)


client.insert(
    collection_name=collName,
    data=[{"id": i, "vector": [i / 10] * 5} for i in range(1, 250)],
)


client.get(collection_name=collName, ids=[i for i in range(150, 153)])

client.get_stats(collection_name=collName)

client.upsert(
    collection_name=collName,
    data=[{"id": 6, "vector": [6 / 10] * 5}, {"id": 7, "vector": [7 / 10] * 5}],
)


client.delete(collection_name=collName, filter="id in [1,2]")
