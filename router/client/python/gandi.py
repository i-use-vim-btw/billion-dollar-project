import requests

class GandiClient:
    def __init__(self, uri, api_key="default") -> None:
        self.__URI = uri
        self.header = {'Content-Type': 'application/json', 'Api_key': api_key}

    def create_collection(self, collection_name="default", dim=128, id="localhost:34779") -> None:

        data = {
            "collectionName": collection_name,
            "dimension": dim,
            "id": id
        }
        res = requests.post(
            'http://' + self.__URI + "/gandi/collections/create",
            headers=self.header,
            json=data
        )

        res = res.json()

        print(res)

        if res["code"] == 200:
            print("Collection succesfully created")
        else:
            print("Error in creating collection")
    
    
    # Not finished, do not use
    def create_index(self, collection_name="default", field_name="vector", params={"nlist": 16384, "index_type": "IVF_FLAT"}, id="localhost:19530"):
        data = {
            "collectionName": collection_name,
            "indexParams": [{
                "fieldName": field_name,
                "params": params
                }],
            "id": id
        }
        
        res = requests.post(
            'http://' + self.__URI + "/gandi/indexes/create",
            headers=self.header,
            json=data
        )
        
        res = res.json()
        
        print(res)
        
        if res["code"] == 200:
            print("Index succesfully created")
        else:
            print("Error in creating index")
    
    def insert(self, collection_name="default", data=[], id="localhost:19530"):
        in_data = {
            'data': data,
            'collectionName': collection_name,
            "id": id
        }

        res = requests.post(
            'http://' + self.__URI + '/gandi/entities/insert',
            headers=self.header,
            json=in_data
        )

        res = res.json()

        print(res)

        if res["code"] == 200:
            print("Insert successful")
        else:
            print("Insert failed")
    
    
    def get(self, collection_name="default", ids=[], id="localhost:19530"):
        data = {
            "collectionName": collection_name,
            'id': ids,
            "host": id
        }

        res = requests.post(
            'http://' + self.__URI + '/gandi/entities/get',
            headers=self.header,
            json=data
        )

        res = res.json()

        print(res)

        if res["code"] == 200:
            print(res["data"])
        else:
            print("Get failed")

    def upsert(self, collection_name="default", data=[], id="localhost:19530"):
        up_data = {
            'data': data,
            'collectionName': collection_name,
            "id": id
        }

        res = requests.post(
            'http://' + self.__URI + '/gandi/entities/upsert',
            headers=self.header,
            json=up_data
        )

        res = res.json()

        print(res)

        if res["code"] == 200:
            print("Upsert successful")
        else:
            print("Upsert failed")
            
    def delete(self, collection_name="default", database_name="default", partition_name="_default", filter="", id="localhost:19530"):
        delete_data = {
            'collectionName': collection_name,
            'databaseName': database_name,
            'partitionName': partition_name,
            'filter': filter,
            "id": id
        }

        res = requests.post(
            'http://' + self.__URI + '/gandi/entities/delete',
            headers=self.header,
            json=delete_data
        )

        res = res.json()

        print(res)

        if res["code"] == 200:
            print("Delete successful")
        else:
            print("Delete failed")
        

      
client = GandiClient("localhost:8080")


collName = "test60"

client.create_collection(collection_name=collName, dim=5)


client.insert(collection_name=collName, data=[
    {"id": i, "vector": [i/10]*5} for i in range(1, 6)
])


client.get(collection_name=collName, ids=[i for i in range(1, 3)])


client.upsert(collection_name=collName, data=[
    {"id": 6, "vector": [6/10]*5},
    {"id": 7, "vector": [7/10]*5}
])


client.delete(collection_name=collName, filter="id in [1,2]")
