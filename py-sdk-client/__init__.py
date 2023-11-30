# class Client():
#     def __init__(self, host='127.0.0.1', port=4002):
#         if isinstance(host, list) and not isinstance(host, str):
#             self.host_list = host
#         else:
#             self.host_list = [[host, port]]

    
#     def put(key, val):
#         # TODO
#         return

#     def get(key, val):
#         # TODO
#         return
    


import grpc
import key_value_pb2
import key_value_pb2_grpc

class Client:
    def __init__(self, server_address):
        self.channel = grpc.insecure_channel(server_address)
        self.stub = key_value_pb2_grpc.KeyValueServiceStub(self.channel)

    def set_key_value(self, key, value):
        request = key_value_pb2.KeyValue(key=key, value=value)
        response = self.stub.SetKeyValue(request)
        return response

    def get_value(self, key):
        request = key_value_pb2.KeyValue(key=key)
        response = self.stub.GetValue(request)
        return response.value

# Usage
if __name__ == '__main__':
    server_address = 'localhost:50051'  # Replace with your server's address
    client = Client(server_address)

    key = 'example_key'
    value = 'example_value'

    # Set a key-value pair on the server
    response = client.set_key_value(key, value)
    print(f'Success: Key "{key}" set to value "{value}"')

    # Get a value from the server
    retrieved_value = client.get_value(key)
    print(f'Value for key "{key}" is: {retrieved_value}')
