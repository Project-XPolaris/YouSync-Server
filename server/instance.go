package server

import (
	context "context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"path/filepath"
	"yousync/database"
	"yousync/pb"
	"yousync/utils"
)

var Default *USyncService = &USyncService{}

type USyncService struct {
	server Server
}

type Server struct {
	pb.UnimplementedFileSyncServer
}

func (s Server) CheckChunk(ctx context.Context, in *pb.ChunkInfo) (*pb.CheckResult, error) {
	var folder database.SyncFolder
	err := database.Instance.First(&folder, in.FolderId).Error
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(folder.Path, in.Path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return &pb.CheckResult{Success: false}, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	// oversize
	if stat.Size() < int64(in.Size+in.Offset) {
		return &pb.CheckResult{Success: false}, nil
	}
	buf := make([]byte, in.Size)
	_, err = file.ReadAt(buf, int64(in.Offset))
	if err != nil {
		return nil, err
	}
	checkSum := utils.SHA256Checksum(buf)
	return &pb.CheckResult{
		Success: checkSum == in.CheckSum,
	}, nil
}
func (s Server) SyncFileChunk(ctx context.Context, in *pb.Chunk) (*pb.SyncChunkResult, error) {
	var folder database.SyncFolder
	err := database.Instance.First(&folder, in.FolderId).Error
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(folder.Path, in.Path)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	_, err = file.WriteAt(in.Data, int64(in.Offset))
	if err != nil {
		return nil, err
	}
	file.Close()
	return &pb.SyncChunkResult{Success: true}, nil
}
func (s *USyncService) Run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcServer := grpc.NewServer()
	s.server = Server{}
	pb.RegisterFileSyncServer(rpcServer, &s.server)
	log.Printf("server listening at %v", lis.Addr())
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
