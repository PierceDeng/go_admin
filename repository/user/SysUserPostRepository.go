package user

import (
	"context"
	"go_admin/config"
	"go_admin/middleware/db"
	"go_admin/model/entity"
)

type SysUserPostRepository struct {
}

var NewSysUserPostRepository = func() *SysUserPostRepository {
	return &SysUserPostRepository{}
}

func (tis SysUserPostRepository) DelUserPost(ctx context.Context, userId uint64) error {
	return db.GetDB(ctx, config.DB).Where("user_id = ?", userId).Delete(&entity.SysUserPost{}).Error
}

func (tis SysUserPostRepository) DelUserPostBatch(ctx context.Context, userIds []uint64) error {
	return db.GetDB(ctx, config.DB).Where("user_id in ?", userIds).Delete(&entity.SysUserPost{}).Error
}

func (tis SysUserPostRepository) AddUserPost(ctx context.Context, userId uint64, postIds []int64) error {

	if len(postIds) == 0 {
		return nil
	}

	var userPostList []*entity.SysUserPost

	for _, postId := range postIds {

		userPostList = append(userPostList, &entity.SysUserPost{
			UserId: userId,
			PostId: postId,
		})

	}

	return db.GetDB(ctx, config.DB).CreateInBatches(userPostList, len(userPostList)).Error
}
