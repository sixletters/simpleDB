import KVStoreService
from enum.command import Command

class KVStoreController():
    def __init__(self, server_address):
        self.service = KVStoreService(server_address)

        self.commands = {
            Command.COMMAND_GET.value: self.service.get_value,
            Command.COMMAND_PUT.value:  self.service.insert_value,
            Command.COMMAND_UPDATE.value: self.service.update_value,
            Command.COMMAND_DELETE.value: self.service.delete_value,
        }

    def execute_command(self, command, data):
        service_function = self.commands[command]
        return service_function(data)