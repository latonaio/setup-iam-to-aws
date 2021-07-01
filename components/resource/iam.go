package resource

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Iam struct {
}

func (d *Deployment) createNewNewUser(
	ctx *pulumi.Context,
	region *Region,
	user User,
) (*iam.User, error) {
	newUser, err := iam.NewUser(ctx,
		fmt.Sprintf("%s", user.UserName),
		&iam.UserArgs{
			Path: pulumi.String("/system/"),
			Tags: pulumi.StringMap{
				"mail": pulumi.String(user.MailAddress),
			},
		})

	if err != nil {
		return nil, err
	}

	// アクセスキーは管理者がGUI上で作成し、Slackなどでお伝えするのが理想なのでコメントアウト
	//newAccessKey, err := iam.NewAccessKey(ctx,
	//	fmt.Sprintf("%s%s", region.ResourceName, "-new-accessKey"),
	//	&iam.AccessKeyArgs{
	//		User: newUser.Name,
	//	})
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//// pulumi-log-upで実行したときにout.txtの中に吐き出される
	//ctx.Export("user secret", pulumi.Sprintf("%s", newAccessKey.Secret.ToStringOutput()))

	// 予め作成しておいたGroupにユーザーを紐づける

	var associationGroupList pulumi.StringArray

	for _, v := range user.AssociationGroup {
		associationGroupList = append(associationGroupList, pulumi.String(v))
	}

	_, err = iam.NewUserGroupMembership(ctx,
		fmt.Sprintf("%s%s", user.UserName, "-new-group"),
		&iam.UserGroupMembershipArgs{
			User: newUser.Name,
			Groups: associationGroupList,
		})

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

