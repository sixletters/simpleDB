from enum import Enum
class Command(Enum):
    COMMAND_GET = "get"
    COMMAND_PUT = "put"
    COMMAND_UPDATE = "update"
    COMMAND_DELETE = "delete"

class Client():

    def get_data_from_input_array(self, input_array):
        if input_array[1] == Command.COMMAND_GET.value:
            return input_array[1], input_array[2], None
        if input_array[1] == Command.COMMAND_PUT.value:
            return input_array[1], input_array[2], input_array[3]
        if input_array[1] == Command.COMMAND_UPDATE.value:
            return input_array[1], input_array[2], input_array[3]
        if input_array[1] == Command.COMMAND_DELETE.value:
            return input_array[1], input_array[2], None

    def validate_input_array(self, input_array):
        if input_array[0] != "scbctl" or len(input_array) <= 1:
            raise ValueError("Invalid input format.")
        command = input_array[1]
        if command not in (Command.COMMAND_GET.value, Command.COMMAND_PUT.value, Command.COMMAND_UPDATE.value, Command.COMMAND_DELETE.value):
            raise ValueError("Invalid Command.")
        if command == Command.COMMAND_GET.value and len(input_array) != 3:
            raise ValueError("Invalid input for GET command.")
        if command == Command.COMMAND_PUT.value and len(input_array) != 4:
            raise ValueError("Invalid input for PUT command")
        if command == Command.COMMAND_UPDATE.value and len(input_array) != 4:
            raise ValueError("Invalid input for UPDATE command.")
        if command == Command.COMMAND_DELETE.value and len(input_array) != 3:
            raise ValueError("Invalid input for DELETE command")

    def get_data_if_valid(self, input):
        input_array = input.split(" ")
        self.validate_input_array(input_array)
        return self.get_data_from_input_array(input_array)

    def run(self):
        while True:
            user_input = input("What command do you want, bro?: ")
            try:
                command, key, val = self.get_data_if_valid(user_input)
                data = {
                    "key": key,
                    "val": val
                }
                resp = KVStoreController.execute_command(command, data)
                print(resp)
            except ValueError as e:
                print(f"Invalid Input: {e}")

class KVStoreService():
    def __init__(self):
        self.KVStore = {}

    def get_value(self, data):
        if data["key"] not in self.KVStore:
            raise ValueError("No such Key Value pair.")
        return self.KVStore[data["key"]]
    
    def insert_value(self, data):
        if data["key"] in self.KVStore:
            raise ValueError("Key Value pair already exist")
        self.KVStore[data["key"]] = data["val"]
        return self.KVStore

    def update_value(self, data):
        if data["key"] not in self.KVStore:
            raise ValueError("Key Value pair does not exist")
        self.KVStore[data["key"]] = data["val"]
        return self.KVStore

    def delete_value(self, data):
        if data["key"] not in self.KVStore:
            raise ValueError("Key Value pair does not exist")
        self.KVStore.pop(data["key"])
        return self.KVStore

class KVStoreController():
    service = KVStoreService()

    commands = {
        Command.COMMAND_GET.value: service.get_value,
        Command.COMMAND_PUT.value:  service.insert_value,
        Command.COMMAND_UPDATE.value: service.update_value,
        Command.COMMAND_DELETE.value: service.delete_value,
    }

    def execute_command(command, data):
        service_function = KVStoreController.commands[command]
        return service_function(data)







if __name__ == "__main__":
    client = Client()
    try:
        print("####### Client has started #######")
        client.run()
    except (Exception, KeyboardInterrupt):
        print("####### Client has stopped #######")
