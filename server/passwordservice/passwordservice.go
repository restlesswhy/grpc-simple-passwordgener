package passwordservice

import (
	"context"
	"crypto/sha512"
	"fmt"

	pb "github.com/restlesswhy/grps-passwordgener/proto"
)

const (
	salt = "als;fkam'fl"
)

type PasswordServer struct {
	pb.UnimplementedPasswordGeneratorServiceServer
}

func (s *PasswordServer) Generate(ctx context.Context, req *pb.PasswordRequest) (*pb.PasswordResponse, error) {
	// res := new(pb.PasswordResponse)
	// passwordHash := Hasher(req.GetSample())
	fmt.Println(req.GetSample())
	fmt.Println(Hasher(req.GetSample()))
	return &pb.PasswordResponse{Password: Hasher(req.GetSample())}, nil
}

func Hasher(password string) []byte {
	hash := sha512.New()
	hash.Write([]byte(password))
	return hash.Sum([]byte(salt))
}