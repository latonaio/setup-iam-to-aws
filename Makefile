# 発行されたiamのリストを確認
get-iam-user-list:
	aws iam list-users --query 'Users[].UserName'

get-iam-group-list:
	aws iam list-groups \
    --query 'Groups[].GroupName'

# 指定したawsのuserがどこのグループに所属しているか確認
# get-belongs-group userName=name
get-belongs-group:
	aws iam list-groups-for-user \
    --user-name ${userName} \
    --query 'Groups[].GroupName'

# 実行前にどのresourceがあたるかを確認することができる
pulumi-preview:
	pulumi preview

# debug logを吐き出す
# secretはout.txtに吐き出される
pulumi-log-up:
	pulumi up --logtostderr -v=9 2> export/log.txt

# log.txtに書き出されたsecretキーを確認する
show-secret:
	cat export/log.txt | grep "user secret"
