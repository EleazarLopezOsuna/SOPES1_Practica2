package main

import (
	"context"
	encoding "encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"math"
	"math/rand"
	"net"
	"strconv"

	pb "github.com/JaredOsuna/SOPES1_Practica2/grpc-server/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMatchesServer
}

type Match struct {
	Game_Id   string `json:"game_id"`
	Players   string `json:"players"`
	Game_Name string `json:"game_name"`
	Winner    string `json:"winner"`
	Queue     string `json:"queue"`
}

func getRandomPlayer(players int) int {
	return rand.Intn(players-1) + 1
}

func game1(players int) (int, string) {
	winner := getRandomPlayer(players)

	if winner == 1 {
		return 2, "Game 1"
	}
	if winner%2 == 0 {
		return winner, "Game 1"
	}
	return winner + 1, "Game 1"
}

func game2(players int) (int, string) {
	winner := getRandomPlayer(players)

	if winner == 1 {
		return 1, "Game 2"
	}
	if winner%2 == 0 {
		return winner - 1, "Game 2"
	}
	return winner, "Game 2"
}

func game3() (int, string) {
	return 1, "Game 3"
}

func game4(players int) (int, string) {
	return int(math.RoundToEven(float64(players / 2))), "Game 4"
}

func game5(players int) (int, string) {
	return players, "Game 5"
}

func (s *server) AddMatch(ctx context.Context, in *pb.MatchRequest) (*pb.MatchReply, error) {
	gameId := int(in.GetGameId())
	numberPlayers := int(in.GetNumberPlayers())
	msgGrpc := " Game ID " + strconv.Itoa(int(gameId)) + " with " + strconv.Itoa(numberPlayers) + " players"
	var winner int
	var gameName string
	if gameId == 1 {
		winner, gameName = game1(numberPlayers)
	} else if gameId == 2 {
		winner, gameName = game2(numberPlayers)
	} else if gameId == 3 {
		winner, gameName = game3()
	} else if gameId == 4 {
		winner, gameName = game4(numberPlayers)
	} else if gameId == 5 {
		winner, gameName = game5(numberPlayers)
	}
	match := Match{
		strconv.Itoa(gameId),
		strconv.Itoa(numberPlayers),
		gameName,
		strconv.Itoa(winner),
		"Kafka"}

	jsonObj, err := encoding.Marshal(match)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	//randomKey := strings.Replace(uuid.New().String(), "-", "", -1)

	msgResult := " Winner: player " + strconv.Itoa(winner)

	/*_host := os.Getenv("KAFKA_HOST")
	_port := os.Getenv("KAFKA_PORT")
	_topic := os.Getenv("KAFKA_TOPIC")
	_broker := fmt.Sprintf("%v:%v", _host, _port)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{_broker},
		Topic:   _topic,
	})

	err = w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(randomKey),
		Value: jsonObj,
	})*/

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "my-cluster-kafka-bootstrap.kafka"})
	if err != nil {
		return &pb.MatchReply{Message: err.Error()}, nil
	}

	topic := "my-topic"

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          jsonObj,
	}, nil)

	defer p.Close()
	if err != nil {
		return &pb.MatchReply{Message: err.Error()}, nil
	}

	return &pb.MatchReply{Message: msgGrpc + msgResult}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMatchesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
