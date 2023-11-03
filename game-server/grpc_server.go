package gameserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-abranches-dev/connect-4/common"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"log"
	"slices"
	"sync"
	"time"
)

type GameServer struct {
	UnimplementedRouteServer
}

type checkExpiredSessionsService struct {
	sessions *[]*common.Session
}

const (
	checkExpiredSessionInterval      = 2 * time.Second
	sessionLastPOLExpirationInterval = 3 * time.Second

	maxSessions = 6
)

var (
	openSessions *[]*common.Session

	openSessionsMutex sync.Mutex
)

func NewGameServer() *GameServer {
	return &GameServer{}
}

func (gs *GameServer) VerifyCompatibility(ctx context.Context,
	payload *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error) {

	gsv := version

	gcv := versions.Version{
		Tag: payload.GameClientVersion,
	}
	if !gsv.Supports(gcv) {
		return nil, utils.FormatErrors([]string{
			fmt.Sprintf("game client version %q is not compatible with the game server version %q.",
				gcv.Tag, gsv.Tag),
			fmt.Sprintf("game client versions compatible with game server version %q: %s", gsv.Tag, versions.GetGameClientsVersions(gsv.SupportedVersions)),
		})
	}

	return &VerifyCompatibilityResponse{GameServerVersion: gsv.Tag}, nil
}

func newCheckExpiredSessionsService(sessions *[]*common.Session) *checkExpiredSessionsService {
	return &checkExpiredSessionsService{
		sessions: sessions,
	}
}

func removeSession(sessions *[]*common.Session, token string) {
	*sessions = slices.DeleteFunc(*sessions, func(s *common.Session) bool {
		return s.Token.String() == token
	})
}

func (s checkExpiredSessionsService) Run(noOpenSessions chan bool) {
	if len(*openSessions) == 0 {
		noOpenSessions <- true
	}
	var sessionsToRemove []string
	for _, os := range *openSessions {
		expirationTime := os.LastPOL.Add(sessionLastPOLExpirationInterval)
		now := time.Now()
		if expirationTime.Before(now) {
			sessionsToRemove = append(sessionsToRemove, os.Token.String())
		}
	}
	if len(sessionsToRemove) > 0 {
		for _, sid := range sessionsToRemove {
			removeSession(openSessions, sid)
		}
		log.Printf(">>> number of active game clients: %d\n\n", len(*openSessions))
	}
}

func startCheckingSessionsExpired(openSessions *[]*common.Session) {
	stopTicker := make(chan bool)
	service := newCheckExpiredSessionsService(openSessions)
	common.StartTicker(checkExpiredSessionInterval, service, stopTicker)
}

func (gs *GameServer) Login(ctx context.Context, in *LoginPayload) (*LoginResponse, error) {
	if openSessions != nil && len(*openSessions) >= maxSessions {
		return nil, fmt.Errorf("there isn't any session available, try later")
	}
	if openSessions == nil {
		openSessions = new([]*common.Session)
	}
	ns := common.NewSession()
	if !slices.Contains(*openSessions, ns) {
		openSessionsMutex.Lock()
		*openSessions = append(*openSessions, ns)

		openSessionsMutex.Unlock()
		if len(*openSessions) == 1 {
			go startCheckingSessionsExpired(openSessions)
		}
	}

	log.Printf(">>> number of active game clients: %d\n\n", len(*openSessions))

	return &LoginResponse{
		Token: ns.Token.String(),
	}, nil
}

func getSessionByToken(token uuid.UUID) (*common.Session, error) {
	if openSessions == nil {
		return nil, errors.New("there aren't any game session active")
	}

	sIdx := slices.IndexFunc(*openSessions, func(s *common.Session) bool {
		return s.Token == token
	})
	if sIdx == -1 {
		return nil, errors.New("the token provided doesn't match any active session")
	}

	return (*openSessions)[sIdx], nil
}

func updateSessionLastPOLTimestamp(session *common.Session) {
	session.LastPOL = time.Now()
}

func (gs *GameServer) POL(ctx context.Context, in *POLPayload) (*POLResponse, error) {
	token, err := uuid.Parse(in.SessionToken)
	if err != nil {
		return &POLResponse{
			Err: "invalid token format",
		}, nil
	}

	session, err := getSessionByToken(token)
	if err != nil {
		return &POLResponse{
			Err: "invalid token",
		}, nil
	}

	updateSessionLastPOLTimestamp(session)

	return &POLResponse{}, nil
}
