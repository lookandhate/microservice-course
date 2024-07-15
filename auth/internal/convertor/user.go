package convertor

import (
	"github.com/lookandhate/microservice-courese/auth/internal/model"
	"github.com/lookandhate/microservice-courese/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateUserFromProto converts data from protobuf generated struct to service CreateUserModel
func CreateUserFromProto(user *auth_v1.CreateRequest) *model.CreateUserModel {
	if user == nil {
		return nil
	}

	return &model.CreateUserModel{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            model.UserRole(user.Role),
	}
}

// CreateUserModelToRepo converts model from service layer to repo layer
func CreateUserModelToRepo(user *model.CreateUserModel) *model.CreateUserRepositoryModel {
	if user == nil {
		return nil
	}

	return &model.CreateUserRepositoryModel{
		Name:  user.Name,
		Email: user.Email,
		// TODO hashing
		Password: user.Password,
		Role:     user.Role,
	}
}

func UserModelToGetResponseProto(user *model.UserModel) *auth_v1.GetResponse {
	return &auth_v1.GetResponse{
		Id:        int64(user.Id),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
