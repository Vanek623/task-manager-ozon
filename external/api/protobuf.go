package api

import (
	"context"

	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
)

type implementation struct {
	pb.UnimplementedStorageServer
	s  iTaskStorage
	cs *counters.Counters
}

const protobufGroupName = "protobufAPI"

func newProtobuf(s iTaskStorage) *implementation {
	return &implementation{
		s:  s,
		cs: counters.New(protobufGroupName),
	}
}

func (i *implementation) TaskAdd(ctx context.Context, in *pb.TaskAddRequest) (*pb.TaskAddResponse, error) {
	i.cs.Inc(counters.Incoming)
	log.WithFields(log.Fields{
		"id":          in.GetID(),
		"title":       in.GetTitle(),
		"description": in.GetDescription(),
	}).Info("Incoming add request")

	task := models.Task{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			i.cs.Inc(counters.Fail)
		} else {
			i.cs.Inc(counters.Success)
		}
	}()

	task.ID, err = uuid.FromBytes(in.GetID())
	if err != nil {
		return nil, err
	}

	if err = i.s.Add(ctx, &task); err != nil {
		return nil, err
	}

	return &pb.TaskAddResponse{}, nil
}

func (i *implementation) TaskList(ctx context.Context, in *pb.TaskListRequest) (*pb.TaskListResponse, error) {
	i.cs.Inc(counters.Incoming)

	log.WithFields(log.Fields{
		"offset": in.GetOffset(),
		"limit":  in.GetLimit(),
	}).Info("Incoming list request")

	tasks, err := i.s.List(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		i.cs.Inc(counters.Fail)
		log.Error(err)
		return nil, err
	}

	result := make([]*pb.TaskListResponse_Task, len(tasks))
	for i, t := range tasks {
		result[i] = &pb.TaskListResponse_Task{
			ID:    uuidToBytes(&t.ID),
			Title: t.Title,
		}
	}

	i.cs.Inc(counters.Success)
	return &pb.TaskListResponse{Tasks: result}, nil
}

func (i *implementation) TaskUpdate(ctx context.Context, in *pb.TaskUpdateRequest) (*pb.TaskUpdateResponse, error) {
	i.cs.Inc(counters.Incoming)
	log.WithFields(log.Fields{
		"id":          in.GetID(),
		"title":       in.GetTitle(),
		"description": in.GetDescription(),
	}).Info("Incoming update request")

	task := models.Task{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	var err error
	defer func() {
		if err != nil {
			log.Error(err)
			i.cs.Inc(counters.Fail)
		} else {
			i.cs.Inc(counters.Success)
		}
	}()

	task.ID, err = uuid.FromBytes(in.GetID())
	if err != nil {
		return nil, err
	}

	if err = i.s.Update(ctx, &task); err != nil {
		return nil, err
	}

	return &pb.TaskUpdateResponse{}, nil
}

func (i *implementation) TaskDelete(ctx context.Context, in *pb.TaskDeleteRequest) (*pb.TaskDeleteResponse, error) {
	i.cs.Inc(counters.Incoming)
	log.WithFields(log.Fields{
		"id": in.GetID(),
	}).Info("Incoming delete request")

	id, err := uuid.FromBytes(in.GetID())
	defer func() {
		if err != nil {
			log.Error(err)
			i.cs.Inc(counters.Fail)
		} else {
			i.cs.Inc(counters.Success)
		}
	}()

	if err != nil {
		return nil, err
	}

	if err = i.s.Delete(ctx, &id); err != nil {
		return nil, err
	}

	return &pb.TaskDeleteResponse{}, nil
}

func (i *implementation) TaskGet(ctx context.Context, in *pb.TaskGetRequest) (*pb.TaskGetResponse, error) {
	i.cs.Inc(counters.Incoming)
	log.WithFields(log.Fields{
		"id": in.GetID(),
	}).Info("Incoming get request")

	id, err := uuid.FromBytes(in.GetID())
	if err != nil {
		log.Error(err)
		i.cs.Inc(counters.Fail)
		return nil, err
	}

	task, err := i.s.Get(ctx, &id)
	if err != nil {
		log.Error(err)
		i.cs.Inc(counters.Fail)
		return nil, err
	}

	res := &pb.TaskGetResponse{
		Title:   task.Title,
		Edited:  task.Edited.Unix(),
		Created: task.Created.Unix(),
	}

	if task.Description != "" {
		res.Description = &task.Description
	}

	i.cs.Inc(counters.Success)
	return res, nil
}

func uuidToBytes(ID *uuid.UUID) []byte {
	bytes := make([]byte, 16)
	copy(bytes, ID[0:])

	return bytes
}
