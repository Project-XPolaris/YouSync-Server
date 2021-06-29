package server

import (
	context "context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"google.golang.org/grpc"
	"io/fs"
	"log"
	"net"
	"os"
	"path/filepath"
	"yousync/database"
	"yousync/pb"
	"yousync/service"
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
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return &pb.CheckResult{Success: false}, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
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
	if in.LastChunk {
		err = file.Truncate(int64(in.Offset + in.Size))
		if err != nil {
			return nil, err
		}
	}
	file.Close()
	return &pb.SyncChunkResult{Success: true}, nil
}
func (s Server) ReadFolderFiles(ctx context.Context, in *pb.RemoteFilesMessage) (*pb.RemoteFilesResult, error) {
	var folder database.SyncFolder
	err := database.Instance.First(&folder, in.FolderId).Error
	if err != nil {
		return nil, err
	}
	result := &pb.RemoteFilesResult{Files: []*pb.RemoteFiles{}}
	afero.Walk(service.AppFs, folder.Path, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		itemPath, err := filepath.Rel(folder.Path, path)
		if err != nil {
			logrus.Error(err)
			return nil
		}
		result.Files = append(result.Files, &pb.RemoteFiles{Path: itemPath, FolderId: in.FolderId, Size: uint64(info.Size())})
		return nil
	})
	return result, nil
}
func (s Server) GetRemoteFileChunkInfo(ctx context.Context, in *pb.GetRemoteChunkInfoMessage) (*pb.RemoteChunkInfo, error) {
	var folder database.SyncFolder
	err := database.Instance.First(&folder, in.FolderId).Error
	if err != nil {
		return nil, err
	}
	targetPath := filepath.Join(folder.Path, filepath.Clean(in.Path))
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		return nil, err
	}
	file, err := os.Open(targetPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	// oversize
	if stat.Size() < int64(in.Offset) {
		return nil, errors.New("chunk overflow")
	}
	size := in.Size
	if int64(in.Offset+size) > stat.Size() {
		size = uint64(stat.Size()) - in.Offset
	}
	buf := make([]byte, size)
	_, err = file.ReadAt(buf, int64(in.Offset))
	if err != nil {
		return nil, err
	}
	checkSum := utils.SHA256Checksum(buf)
	return &pb.RemoteChunkInfo{
		LastChunk: int64(in.Offset+in.Size) > stat.Size(),
		Checksum:  checkSum,
	}, nil
}
func (s Server) GetRemoteFileChunk(ctx context.Context, in *pb.GetRemoteChunkMessage) (*pb.RemoteChunk, error) {
	var folder database.SyncFolder
	err := database.Instance.First(&folder, in.FolderId).Error
	if err != nil {
		return nil, err
	}
	targetPath := filepath.Join(folder.Path, filepath.Clean(in.Path))
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		return nil, err
	}
	file, err := os.Open(targetPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	// oversize
	if stat.Size() < int64(in.Offset) {
		return nil, errors.New("chunk overflow")
	}
	size := in.Size
	if int64(in.Offset+size) > stat.Size() {
		size = uint64(stat.Size()) - in.Offset
	}
	buf := make([]byte, size)
	_, err = file.ReadAt(buf, int64(in.Offset))
	if err != nil {
		return nil, err
	}
	return &pb.RemoteChunk{
		Data: buf,
	}, nil
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
