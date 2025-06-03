import grpc
from python_gen.greet.v1 import greet_pb2
from python_gen.greet.v1 import greet_pb2_grpc

def run():
    # Create a channel
    with grpc.insecure_channel('localhost:8080') as channel:
        # Create a stub (client)
        stub = greet_pb2_grpc.GreetServiceStub(channel)
        
        # Create a request
        request = greet_pb2.GreetRequest(name='John')
        
        # Make the call
        try:
            response = stub.Greet(request)
            print(f"Response from server: {response.greeting}")
        except grpc.RpcError as e:
            print(f"RPC failed: {e}")

if __name__ == '__main__':
    run()
