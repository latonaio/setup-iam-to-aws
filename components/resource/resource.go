package resource

import (
	"encoding/json"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"io/ioutil"
)

type UserList []User

type User struct {
	UserName    string `json:"userName"`
	MailAddress string `json:"mailAddress"`
	AssociationGroup []string `json:"associationGroup"`
}

func readUserJson(ctx *pulumi.Context) (UserList, error) {
	user, err := ioutil.ReadFile("./user.json")

	if err != nil {
		ctx.Export("Unmarshal error", pulumi.Printf("%v", err))
		return nil, err
	}

	var u UserList

	err = json.Unmarshal(user, &u)

	if err != nil {
		ctx.Export("Unmarshal error in readUserJson", pulumi.Printf("%v", err))
		return nil, err
	}

	return u, err
}

func Setup(ctx *pulumi.Context) error {
	config, _ := ioutil.ReadFile("./config.json")

	var r Region

	err := json.Unmarshal(config, &r)

	if err != nil {
		ctx.Export("Unmarshal error", pulumi.Printf("%v", err))
		return err
	}

	region := newRegion(r)

	deployment := new(Deployment)

	userList, err := readUserJson(ctx)

	if err != nil {
		ctx.Export("readUserJson error", pulumi.Printf("%v", err))
		return err
	}

	for _, user := range userList {
		_, err = deployment.createNewNewUser(ctx, region, user)

		if err != nil {
			ctx.Export("createNewNewUser error", pulumi.Printf("%v", err))
			ctx.Export("user name: ", pulumi.Printf("%v", user.UserName))
			return err
		}
	}

	return nil
}

