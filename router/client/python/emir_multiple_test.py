from gandi import GandiClient

client = GandiClient("localhost:8080")

while True:
    print(client.get("test50", ids=[1, 2, 3, 4, 5]))