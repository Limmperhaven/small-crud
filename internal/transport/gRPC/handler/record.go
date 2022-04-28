package grpcHandler

import (
	"gitlab.digital-spirit.ru/study/artem_crud/internal/models"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/service"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/transport/gRPC/proto"
	"golang.org/x/net/context"
)

type RecordHandler struct {
	proto.UnimplementedRecordsServer
	services *service.Service
}

func NewRecordHandler(services *service.Service) *RecordHandler {
	return &RecordHandler{
		services: services,
	}
}

func (r *RecordHandler) Create(ctx context.Context, req *proto.RecordInput) (*proto.Uuid, error) {
	var input models.Record

	if req.FirstName != "" {
		input.FirstName = req.FirstName
	}
	if req.LastName != "" {
		input.LastName = req.LastName
	}
	if req.MobilePhone != "" {
		input.MobilePhone = req.MobilePhone
	}
	if req.HomePhone != "" {
		input.HomePhone = req.HomePhone
	}

	uuid, err := r.services.Create(input)

	return &proto.Uuid{Uuid: uuid}, err
}

func (r *RecordHandler) GetByUuid(ctx context.Context, req *proto.Uuid) (*proto.Record, error) {
	rec, err := r.services.GetById(req.Uuid)

	return &proto.Record{
		Uuid:        rec.Uuid,
		FirstName:   rec.FirstName,
		LastName:    rec.LastName,
		MobilePhone: rec.MobilePhone,
		HomePhone:   rec.HomePhone,
	}, err
}

func (r *RecordHandler) GetByFilter(ctx context.Context, req *proto.RecordInput) (*proto.RecordList, error) {
	recs, err := r.services.GetByFilter(models.RecordInput{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		MobilePhone: req.MobilePhone,
		HomePhone:   req.HomePhone,
	})

	var recSlice []*proto.Record

	for _, rec := range recs {
		recSlice = append(recSlice, &proto.Record{
			Uuid:        rec.Uuid,
			FirstName:   rec.FirstName,
			LastName:    rec.LastName,
			MobilePhone: rec.MobilePhone,
			HomePhone:   rec.HomePhone,
		})
	}

	return &proto.RecordList{Records: recSlice}, err
}

func (r *RecordHandler) Update(ctx context.Context, req *proto.Record) (*proto.Empty, error) {
	return new(proto.Empty), r.services.Update(req.Uuid, models.RecordInput{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		MobilePhone: req.MobilePhone,
		HomePhone:   req.HomePhone,
	})
}

func (r *RecordHandler) Delete(ctx context.Context, req *proto.Uuid) (*proto.Empty, error) {
	return new(proto.Empty), r.services.Delete(req.Uuid)
}
