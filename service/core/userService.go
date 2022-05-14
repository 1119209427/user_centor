package core

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"user_centor/pb"
	"user_centor/service/dao"
	"user_centor/service/model"
)


type UserServices struct {
	pb.UnimplementedUserServiceServer
}
func(us *UserServices)UserChangePassword(ctx context.Context,req *pb.UserRequestChangePassword)(*pb.UserDetailResponse,error){
	var user model.User
	var resp pb.UserDetailResponse
	var count int64
	resp.Code=200
	if req.NewPassword!=req.NewPasswordConfirm{
		resp.Code=500
		err:=errors.New("两次输入密码不一致")
		return &resp, err

	}
	//检查用户是否存在
	if err:=dao.DB.Where("user_name",req.Username).First(&user).Count(&count).Error;err!=nil{
		if err.Error()=="record not found"{
			resp.Code=500
			err=errors.New("请注册后修改密码")
			return &resp,nil
		}
		resp.Code=400
		return &resp, err
	}
	if count>0{//用户存在
		//验证密码
		flag:=user.CheckPassword(req.OldPassword)
		if flag==false{
			resp.Code=500
			err:=errors.New("密码错误，请重新输入")
			return &resp, err
		}
		//加密新密码
		err:=user.SetPassword(req.NewPassword)
		if err!=nil{
			resp.Code=500
			return &resp, err
		}
		//更新密码
		if err=dao.DB.Model(&user).Where("user_name",req.Username).Update("password",req.NewPassword).Error;err!=nil{
			resp.Code=500
			return &resp, err
		}
		resp.Model=BuildUser(user)
		return &resp,nil




	}else{
		resp.Code=500
		err:=errors.New("请注册后修改密码")
		return &resp,err
	}



}




func(us *UserServices)UserLogin(ctx context.Context, req *pb.UserRequest )(*pb.UserDetailResponse,error){
	//检查用户名是否存在

	var user model.User
	var resp pb.UserDetailResponse
	resp.Code=200
	if err:=dao.DB.Where("user_name",req.Username).First(&user);err!=nil {
		if err.Error == gorm.ErrRecordNotFound {
			resp.Code = 400
			return &resp, err.Error
		}
		resp.Code = 500
		return &resp, err.Error
	}
	flag:=user.CheckPassword(req.Password)
	if flag==true{
		resp.Model=BuildUser(user)
		return &resp,nil


	}else{
		resp.Code = 500
		return &resp,nil
	}

}
func(us *UserServices)UserRegister(ctx context.Context,req *pb.UserRequest)( *pb.UserDetailResponse,error){
	var resp pb.UserDetailResponse
	resp.Code=200
	if req.Password!=req.PasswordConfirm{
		resp.Code=500
		err:=errors.New("两次输入密码不一致")
		return &resp, err
	}
	var count int64

	if err:=dao.DB.Model(&model.User{}).Where("user_name",req.Username).Count(&count).Error;err!=nil{
		if err.Error()=="record not found"{
			user:=model.User{
				UserName: req.Username,
			}
			//加密密码
			if err:=user.SetPassword(req.Password);err!=nil{
				resp.Code=500
				return &resp, err
			}
			if err:=dao.DB.Create(&user).Error;err!=nil{
				resp.Code=500
				return &resp, err
			}
			resp.Model=BuildUser(user)
			return &resp,nil


		}
		resp.Code=400
		return &resp, err

	}



	if count>0{
		resp.Code=500
		err:=errors.New("用户名已存在")
		return &resp, err
	}
	user:=model.User{
		UserName: req.Username,
	}
	//加密密码
	if err:=user.SetPassword(req.Password);err!=nil{
		resp.Code=500
		return &resp, err
	}
	if err:=dao.DB.Create(&user).Error;err!=nil{
		resp.Code=500
		return &resp, err
	}
	resp.Model=BuildUser(user)
	return &resp,nil



}
func BuildUser(item model.User)*pb.UserModel{
	var model pb.UserModel
	model.ID= uint32(item.ID)
	model.Username=item.UserName
	model.CreateTime=item.CreatedAt.Unix()
	model.UpdateTime=item.UpdatedAt.Unix()
	return &model
}
