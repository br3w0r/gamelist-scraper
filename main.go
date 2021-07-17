package main

import (
	"log"
	"net"
	"strconv"

	"bitbucket.org/br3w0r/gamelist-scraper/entity"
	"bitbucket.org/br3w0r/gamelist-scraper/helpers"
	pb "bitbucket.org/br3w0r/gamelist-scraper/proto"
	"bitbucket.org/br3w0r/gamelist-scraper/scraper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	ADDRESS   string = helpers.GetEnvOrDefault("ADDRESS", "localhost")
	TLS       int8   = 0
	CERT_FILE string = helpers.GetEnvOrDefault("CERT_FILE", "")
	KEY_FILE  string = helpers.GetEnvOrDefault("KEY_FILE", "")
)

type gameScrapeServer struct {
	pb.UnimplementedGameScrapeServer
}

func (s *gameScrapeServer) ScrapeGames(_ *pb.Empty, stream pb.GameScrape_ScrapeGamesServer) error {
	c := make(chan entity.ScraperResp)
	go scraper.Scrape(c)
	for resp := range c {
		if resp.Err != nil {
			return resp.Err
		}

		proto := resp.Game.ConvertToProto()
		if err := stream.Send(&proto); err != nil {
			return err
		}
	}
	return nil
}

func ParseEnv() {
	env := helpers.GetEnvOrDefault("TLS", "0")
	tls, err := strconv.ParseInt(env, 0, 0)
	if err != nil {
		log.Fatalf("failed to parse TLS env key to int: %s", env)
	}
	TLS = int8(tls)
}

func main() {
	ParseEnv()

	lis, err := net.Listen("tcp", ADDRESS+":8888")
	if err != nil {
		log.Fatalf("failed to listen to IP: %s", ADDRESS+":8888")
	}
	var opts []grpc.ServerOption
	if TLS == 1 {
		creds, err := credentials.NewServerTLSFromFile(CERT_FILE, KEY_FILE)
		if err != nil {
			log.Fatalf("failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGameScrapeServer(grpcServer, &gameScrapeServer{})

	log.Printf("Starting to serve on address: %s\n", ADDRESS)
	grpcServer.Serve(lis)
}
