import grpc
from ..pkg.stubs import client_without_internode_pb2
from ..pkg.stubs import client_without_internode_pb2_grpc

# Simple Db SDK
class SimpleDB:
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