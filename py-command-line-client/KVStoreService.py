import SimpleDB

class KVStoreService():
    def __init__(self, server_address):
        self.KVStore = SimpleDB(server_address)

    def get_value(self, data):
        if data["key"] not in self.KVStore:
            raise ValueError("No such Key Value pair.")
        return self.KVStore.get_value(data["key"])
    
    def insert_value(self, data):
        if data["key"] in self.KVStore:
            raise ValueError("Key Value pair already exist")
        return self.KVStore.set_key_value(data["key"], data["val"])


    # To implement update and delete methods
    # def update_value(self, data):
    #     if data["key"] not in self.KVStore:
    #         raise ValueError("Key Value pair does not exist")
    #     self.KVStore[data["key"]] = data["val"]
    #     return self.KVStore

    # def delete_value(self, data):
    #     if data["key"] not in self.KVStore:
    #         raise ValueError("Key Value pair does not exist")
    #     self.KVStore.pop(data["key"])
    #     return self.KVStore