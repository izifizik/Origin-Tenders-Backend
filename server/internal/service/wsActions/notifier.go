package wsActions

import (
	"origin-tender-backend/server/internal/domain/wsModels"
)

//func NotifyAllNotifiers(data string) {
//	for _, conns := range wsModels.UserConnections {
//		conns.Notification.WriteMessage(1, []byte(data))
//	}
//}

func NotifyUser(data string, userId string) {
	var conn = wsModels.UserConnections[userId]
	if conn.Notification != nil {
		wsModels.UserConnections[userId].Notification.WriteMessage(1, []byte(data))
	}

}

func NotifyAllSession(data string) {
	for _, conns := range wsModels.UserConnections {
		conns.Session.WriteMessage(1, []byte(data))
	}
}

func NotifyAllBet(data string) {
	for _, conns := range wsModels.UserConnections {
		conns.Bets.WriteMessage(1, []byte(data))
	}
}
