package shop

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"
	"zonst/qipai/protocolutil"
	"zonst/qipai/messages"

	"github.com/golang/protobuf/proto"
)

func TestShopPropList(t *testing.T) {
	hostInfo := "0.0.0.0:8113"
	conn, err := net.Dial("tcp", hostInfo)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		log.Println(err.Error())
		return
	}
	proMiniGameShopPropListRequest := messages.ProMiniGameShopPropListRequest{
		MessageId:  proto.Int32(660),
		GameId:     proto.Int32(9),
		PlatformId: proto.Int32(1000),
		PageType:   messages.MiniGamePageType_MINIGAME_SHOP_PAGE_HEAD.Enum(),
	}
	header := &protocolutil.ClientHeader{
		MessageType: protocolutil.ClientRequestMessageType,
		DestID:      0,
		MessageID:   uint16(proMiniGameShopPropListRequest.GetMessageId()),
	}

	packet := protocolutil.NewPacket(header)
	packet.SetContent(&proMiniGameShopPropListRequest)
	buf := packet.Serialize()
	conn.Write(buf)
	rs, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("rs:", rs)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
