package convertor

import (
	"github.com/lookandhate/microservice-courese/auth/internal/model"
	"github.com/lookandhate/microservice-courese/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateUserFromProto converts data from protobuf generated struct to service CreateUserModel.
func CreateUserFromProto(user *auth_v1.CreateRequest) *model.CreateUserModel {
	if user == nil {
		return nil
	}

	return &model.CreateUserModel{
		Name:            user.GetName(),
		Email:           user.GetEmail(),
		Password:        user.GetPassword(),
		PasswordConfirm: user.GetPasswordConfirm(),
		Role:            model.UserRole(user.GetRole()),
	}
}

// CreateUserModelToRepo converts model from service layer to repo layer.
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

// UserModelToGetResponseProto converts from UserModel to proto response.
func UserModelToGetResponseProto(user *model.UserModel) *auth_v1.GetResponse {
	if user == nil {
		return nil
	}

	return &auth_v1.GetResponse{
		Id:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

// UserUpdateFromProto converts from proto UpdateRequest to UpdateUserModel.
func UserUpdateFromProto(user *auth_v1.UpdateRequest) *model.UpdateUserModel {
	if user == nil {
		return nil
	}

	return &model.UpdateUserModel{
		Name:     user.GetName().GetValue(),
		Email:    user.GetEmail().GetValue(),
		Role:     int(user.GetRole()),
		Password: user.GetPassword().GetValue(),
		ID:       int(user.GetId()),
	}
}
