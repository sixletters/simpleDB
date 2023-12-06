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
from ..pkg.stubs import client_without_internode_pb2
from ..pkg.stubs import client_without_internode_pb2_grpc

class Client:
    def __init__(self, server_address):
        self.channel = grpc.insecure_channel(server_address)
        self.stub = client_without_internode_pb2_grpc.KeyValueServiceStub(self.channel)

    def set_key_value(self, key, value):
        request = client_without_internode_pb2.PutRequest(key=key, value=value)
        response = self.stub.Put(request)
        return response

    def get_value(self, key):
        request = client_without_internode_pb2.GetRequest(key=key)
        response = self.stub.Get(request)
        return response.value

# Usage
if __name__ == '__main__':
    server_address = 'localhost:8080'  # Replace with your server's address
    client = Client(server_address)

    key = 'dion'
    value = 'test_value'

    # Set a key-value pair on the server
    response = client.set_key_value(key, value)
    print(f'Success: Key "{key}" set to value "{value}"')

    # Get a value from the server
    retrieved_value = client.get_value(key)
    print(f'Value for key "{key}" is: {retrieved_value}')
