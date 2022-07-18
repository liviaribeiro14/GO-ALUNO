import "google.golang.org/grpc"

func main() {
	connection, err := net.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect to grpc server: %w", err)
	}

	defer connection.Close()

	client : = pb.NewUserServiceClient(connection)
	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id: "0",
		Name: "Joao",
		Email: "joao@gmail.com"
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatal("Could not make grpc server: %w", err)
	}

	fmt.Println(res)
}