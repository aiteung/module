package module

import (
	"fmt"
	"strings"

	"github.com/aiteung/atdb"
	"github.com/aiteung/musik"
	"go.mongodb.org/mongo-driver/mongo"
)

func IteungModuleCall(WAIface IteungWhatsMeowConfig, DBIface IteungDBConfig) (Modulename string, Pesan IteungMessage) {
	Pesan = Whatsmeow2Struct(WAIface)
	NormalizeAndTypoCorrection(&Pesan.Message, DBIface.MongoConn, DBIface.TypoCollection)
	if IsIteungCall(Pesan) {
		Modulename = GetModuleName(Pesan, DBIface.MongoConn, DBIface.ModuleCollection)
	}
	return
}

func Whatsmeow2Struct(WAIface IteungWhatsMeowConfig) (im IteungMessage) {
	im.Phone_number = WAIface.Info.Sender.User
	im.Chat_server = WAIface.Info.Chat.Server
	im.Group_name = ""
	im.Alias_name = WAIface.Info.PushName
	m := WAIface.Message.GetConversation()
	im.Message = m
	im.Is_group = "false"
	im.Filename = ""
	im.Filedata = ""
	im.Latitude = 0.0
	im.Longitude = 0.0
	if WAIface.Info.Chat.Server == "g.us" {
		groupInfo, err := WAIface.Waclient.GetGroupInfo(WAIface.Info.Chat)
		fmt.Println("cek err : ", err)
		if groupInfo != nil {
			im.Group = groupInfo.GroupName.Name + "@" + WAIface.Info.Chat.User
			im.Group_name = groupInfo.GroupName.Name
			im.Group_id = WAIface.Info.Chat.User
		} else {
			fmt.Println("groupInfo : ", groupInfo)
		}
		im.Is_group = "true"
	}
	return
}

func IsIteungCall(im IteungMessage) bool {
	if (strings.Contains(im.Message, "teung") && im.Chat_server == "g.us") || (im.Chat_server == "s.whatsapp.net") {
		return true
	} else {
		return false
	}
}

func GetModuleName(im IteungMessage, MongoConn *mongo.Database, ModuleCollection string) (modulename string) {
	modules := atdb.GetAllDoc[[]Module](MongoConn, ModuleCollection)
	for _, mod := range modules {
		complete, _ := musik.IsMatch(im.Message, mod.Keyword...)
		if complete {
			modulename = mod.Name
		}
	}
	return
}
