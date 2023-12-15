from enum import Enum
from enum.command import Command
import KVStoreController


class Client():
    def __init__(self, server_address):
        self.KVStoreController = KVStoreController(server_address)

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
                resp = self.KVStoreController.execute_command(command, data)
                print(resp)
            except ValueError as e:
                print(f"Invalid Input: {e}")









if __name__ == "__main__":
    client = Client('localhost:8080')
    try:
        print("####### Client has started #######")
        client.run()
    except (Exception, KeyboardInterrupt):
        print("####### Client has stopped #######")