package templates

const ApiTpl = `package api

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/golang/protobuf/ptypes/empty"
	commonPb "gitlab.ronshubao.com/grpc-insure-proto/google"
	Pb "gitlab.ronshubao.com/grpc-insure-proto/product"
	grpcUtil "gitlab.ronshubao.com/grpc-insure/grpc-util"
	"gitlab.ronshubao.com/grpc-insure/product/service/{{tableName}}_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type {{UCamelTableName}}Server struct{}

// Add 添加
func (t {{UCamelTableName}}Server) Add(ctx context.Context, r *Pb.{{UCamelTableName}}Entity) (*commonPb.Id, error) {
	valid := validation.Validation{} //实例化一个验证对象
{{ApiValidData}}
	if valid.HasErrors() {
		return &commonPb.Id{}, status.Errorf(codes.InvalidArgument, valid.Errors[0].Key+" "+valid.Errors[0].Message)
	}

	{{LCamelTableName}} := {{tableName}}_service.{{UCamelTableName}}{
{{ApiSaveData}}
	}
	Id, err := {{LCamelTableName}}.Save()

	return &commonPb.Id{
		Id: int64(Id),
	}, err
}

// Update 更新
func (t {{UCamelTableName}}Server) Update(ctx context.Context, r *Pb.{{UCamelTableName}}Entity) (*empty.Empty, error) {
	valid := validation.Validation{} //实例化一个验证对象
{{ApiValidData}}
	if valid.HasErrors() {
		return nil, status.Errorf(codes.InvalidArgument, valid.Errors[0].Key+" "+valid.Errors[0].Message)
	}

	{{LCamelTableName}} := {{tableName}}_service.{{UCamelTableName}}{
{{ApiSaveData}}
	}
	_, err := {{LCamelTableName}}.Save()

	return nil, err
}

// Delete 删除
func (t {{UCamelTableName}}Server) Delete(ctx context.Context, r *commonPb.Id) (*empty.Empty, error) {
	valid := validation.Validation{} //实例化一个验证对象
	valid.Required(r.Id, "id")
	if valid.HasErrors() {
		return &empty.Empty{}, status.Errorf(codes.InvalidArgument, valid.Errors[0].Key+" "+valid.Errors[0].Message)
	}

	{{LCamelTableName}}Service := {{tableName}}_service.{{UCamelTableName}}{Id: int(r.Id)}
	err := {{LCamelTableName}}Service.Delete()

	return &empty.Empty{}, err
}

// Search is get all
func (t {{UCamelTableName}}Server) Search(ctx context.Context, r *commonPb.SearchRequest) (*Pb.{{UCamelTableName}}SearchResponse, error) {
	fmt.Println(r.Param)
	pageNum, pageSize := grpcUtil.ParsePage(r.Param)

	queryParam := grpcUtil.URLQuery{QueryParam: r.Param}
	log.Println(queryParam.ParseSingleQueryParam("name"))
	Query := map[string]string{
		// "name": queryParam.ParseSingleQueryParam("name"),
	}

	{{LCamelTableName}}Service := {{tableName}}_service.{{UCamelTableName}}{
		PageNum:  pageNum,
		PageSize: pageSize,
		Query:    Query,
	}

	// 总数量
	total, err := {{LCamelTableName}}Service.Count()
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	{{LCamelTableName}}s, err := {{LCamelTableName}}Service.GetAll()
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	var data []*Pb.{{UCamelTableName}}Entity
	for _, v := range {{LCamelTableName}}s {
		data = append(data, &Pb.{{UCamelTableName}}Entity{
			Id:       int32(v.Id),
{{ApiAllBackData}}
			CreatedOn: v.CreatedOn.Format(grpcUtil.TimeFormat),
		})
	}

	return &Pb.{{UCamelTableName}}SearchResponse{
		PageInfo: &commonPb.SearchPageResponse{
			Page:       int32(pageNum),
			PageSize:   int32(pageSize),
			TotalPage:  int32(grpcUtil.GetTotalPage(total, pageSize)),
			TotalCount: int32(total),
		},
		Data: data,
	}, nil
}

// View is single entity
func (t {{UCamelTableName}}Server) View(ctx context.Context, r *commonPb.Id) (*Pb.{{UCamelTableName}}Entity, error) {
	valid := validation.Validation{}
	valid.Required(r.Id, "id")
	if valid.HasErrors() {
		return &Pb.{{UCamelTableName}}Entity{}, status.Errorf(codes.InvalidArgument, valid.Errors[0].Message)
	}
	{{LCamelTableName}}Service := {{tableName}}_service.{{UCamelTableName}}{Id: int(r.Id)}
	{{LCamelTableName}}, err := {{LCamelTableName}}Service.Get()
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}
	return &Pb.{{UCamelTableName}}Entity{
		Id:       int32({{LCamelTableName}}.Id),
{{ApiViewBackData}}
		CreatedOn: {{LCamelTableName}}.CreatedOn.Format(grpcUtil.TimeFormat),
	}, nil
}
`
