class Client():
    def __init__(self, host='127.0.0.1', port=4002):
        if isinstance(host, list) and not isinstance(host, str):
            self.host_list = host
        else:
            self.host_list = [[host, port]]

    
    def put(key, val):
        # TODO
        return

    def get(key, val):
        # TODO
        return