package callback

import tim "github.com/xjcloudy/tencent-im-restapi-go"

const (
	// group callback command
	Group_CallbackBeforeCreateGroup     = "Group.CallbackBeforeCreateGroup"
	Group_CallbackAfterCreateGrou       = "Group.CallbackAfterCreateGroup"
	Group_CallbackBeforeApplyJoinGroup  = "Group.CallbackBeforeApplyJoinGroup"
	Group_CallbackBeforeInviteJoinGroup = "Group.CallbackBeforeInviteJoinGroup"
	Group_CallbackAfterNewMemberJoin    = "Group.CallbackAfterNewMemberJoin"
	Group_CallbackAfterMemberExit       = "Group.CallbackAfterMemberExit"
	Group_CallbackBeforeSendMsg         = "Group.CallbackBeforeSendMsg"
	Group_CallbackAfterSendMsg          = "Group.CallbackAfterSendMsg"
	Group_CallbackAfterGroupFull        = "Group.CallbackAfterGroupFull"
	Group_CallbackAfterGroupDestroyed   = "Group.CallbackAfterGroupDestroyed"
	Group_CallbackAfterGroupInfoChanged = "Group.CallbackAfterGroupInfoChanged"
)

type GroupCallbackBase struct {
	CallbackCommand string
	GroupID         string        `json:"GroupId"` // 群组 ID
	Type            tim.GroupType // 群组类型
}
type GroupCallbackAfterSendMsg struct {
	GroupCallbackBase
	FromAccount string           `json:"From_Account"` // 发送者
	MsgSeq      int64            // 消息的序列号
	MsgTime     int64            // 消息的时间
	MsgBody     []tim.MsgElement // 消息体
}
type GroupCallbackBeforeSendMsg struct {
	GroupCallbackBase
	FromAccount     string           `json:"From_Account"`     // 发送者
	OperatorAccount string           `json:"Operator_Account"` // 请求的发起者
	Random          int32            // 随机数
	MsgBody         []tim.MsgElement // 消息体
}
type GroupBeforeCreateGroup struct {
	OperatorAccount string                   `json:"Operator_Account"` // 操作者
	OwnerAccount    string                   `json:"Owner_Account"`    // 群主
	Name            string                   // 群组名称
	CreatedNum      int32                    //该用户已创建的同类的群组个数
	MemberList      []tim.GroupMemberAccount // 初始成员列表

}
type GroupAfterCreateGroup struct {
	GroupCallbackBase
	OperatorAccount     string                   `json:"Operator_Account"` // 操作者
	OwnerAccount        string                   `json:"Owner_Account"`    // 群主
	Name                string                   // 群组名称
	MemberList          []tim.GroupMemberAccount // 初始成员列表
	UserDefinedDataList []tim.KV                 //用户建群时的自定义字段
}
