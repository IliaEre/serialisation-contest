package mapper

import (
	pb "proto-docs-service/grpc/docs"
	"proto-docs-service/pkg/model"
)

type ProtoMapper struct {
}

func New() *ProtoMapper {
	return &ProtoMapper{}
}

func (pm *ProtoMapper) ProtoToModelDocument(protoDoc *pb.Document) *model.Document {
	var goods []model.Goods
	for _, v := range protoDoc.GetGoods() {
		product := model.Goods{
			Name:   v.GetName(),
			Amount: int(v.GetAmount()),
			Code:   v.GetCode(),
		}
		goods = append(goods, product)
	}

	return &model.Document{
		Docs: model.Docs{
			Name: protoDoc.GetName(),
			Department: model.Department{
				Code: protoDoc.Department.GetCode(),
				Time: protoDoc.Department.GetTime(),
				Employee: model.Employee{
					Name:    protoDoc.Department.Employee.GetName(),
					Surname: protoDoc.Department.Employee.GetSurname(),
					Code:    protoDoc.Department.Employee.GetCode(),
				},
			},
			Price: model.Price{
				CategoryA: protoDoc.GetPrice().GetCategoryA(),
				CategoryB: protoDoc.GetPrice().GetCategoryB(),
				CategoryC: protoDoc.GetPrice().GetCategoryC(),
			},
			Owner: model.Owner{
				UUID:   protoDoc.GetOwner().GetUuid(),
				Secret: protoDoc.GetOwner().GetSecret(),
			},
			Data: model.Data{
				Transaction: model.Transaction{
					Type:      protoDoc.GetData().GetTransaction().GetType(),
					UUID:      protoDoc.GetData().GetTransaction().GetUuid(),
					PointCode: protoDoc.GetData().GetTransaction().GetPointCode(),
				},
			},
			Delivery: model.Delivery{
				Company: protoDoc.GetDelivery().Company,
				Address: model.Address{
					Code:      protoDoc.GetDelivery().GetAddress().GetCode(),
					Country:   protoDoc.GetDelivery().GetAddress().GetCountry(),
					Street:    protoDoc.GetDelivery().GetAddress().GetStreet(),
					Apartment: protoDoc.GetDelivery().GetAddress().GetApartment(),
				},
			},
			Goods: goods,
		},
	}
}

func (pm *ProtoMapper) ModelToProtoDocument(modelDoc *model.Document) *pb.Document {
	var protoGoods []*pb.Goods
	for _, v := range modelDoc.Docs.Goods {
		product := &pb.Goods{
			Name:   v.Name,
			Amount: int32(v.Amount),
			Code:   v.Code,
		}
		protoGoods = append(protoGoods, product)
	}

	return &pb.Document{
		Name: modelDoc.Docs.Name,
		Department: &pb.Department{
			Code: modelDoc.Docs.Department.Code,
			Time: modelDoc.Docs.Department.Time,
			Employee: &pb.Employee{
				Name:    modelDoc.Docs.Department.Employee.Name,
				Surname: modelDoc.Docs.Department.Employee.Surname,
				Code:    modelDoc.Docs.Department.Employee.Code,
			},
		},
		Price: &pb.Price{
			CategoryA: modelDoc.Docs.Price.CategoryA,
			CategoryB: modelDoc.Docs.Price.CategoryB,
			CategoryC: modelDoc.Docs.Price.CategoryC,
		},
		Owner: &pb.Owner{
			Uuid:   modelDoc.Docs.Owner.UUID,
			Secret: modelDoc.Docs.Owner.Secret,
		},
		Data: &pb.Data{
			Transaction: &pb.Transaction{
				Type:      modelDoc.Docs.Data.Transaction.Type,
				Uuid:      modelDoc.Docs.Data.Transaction.UUID,
				PointCode: modelDoc.Docs.Data.Transaction.PointCode,
			},
		},
		Delivery: &pb.Delivery{
			Company: modelDoc.Docs.Delivery.Company,
			Address: &pb.Address{
				Code:      modelDoc.Docs.Delivery.Address.Code,
				Country:   modelDoc.Docs.Delivery.Address.Country,
				Street:    modelDoc.Docs.Delivery.Address.Street,
				Apartment: modelDoc.Docs.Delivery.Address.Apartment,
			},
		},
		Goods: protoGoods,
	}
}
