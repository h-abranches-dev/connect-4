package gameengine

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-abranches-dev/connect-4/common"
	"github.com/h-abranches-dev/connect-4/pkg/utils"
	"github.com/h-abranches-dev/connect-4/pkg/versions"
	"log"
	"sync"
	"time"
)

type checkExpiredGSSessionService struct {
	gsSession *common.Session
}

type GameEngine struct {
	UnimplementedRouteServer
}

const (
	delayConnectionDuration            = 3 * time.Second
	checkExpiredGSSessionInterval      = 2 * time.Second
	gsSessionLastPOLExpirationInterval = 3 * time.Second
)

var (
	gsSession *common.Session

	sessionMutex sync.Mutex
)

func NewGameEngine() *GameEngine {
	return &GameEngine{}
}

func (ge *GameEngine) VerifyCompatibility(ctx context.Context,
	payload *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error) {

	gev := version

	gsv := versions.Version{
		Tag: payload.GameServerVersion,
	}
	if !gev.Supports(gsv) {
		return nil, utils.FormatErrors([]string{
			fmt.Sprintf("game server version %q is not compatible with the game engine version %q.",
				gsv.Tag, gev.Tag),
			fmt.Sprintf("game server versions compatible with game engine version %q: %s", gev.Tag,
				versions.GetGameServersVersions(gev.SupportedVersions)),
		})
	}

	return &VerifyCompatibilityResponse{GameEngineVersion: gev.Tag}, nil
}

func delayConnection() {
	time.Sleep(delayConnectionDuration)
}

func newGSSession() *common.Session {
	ns := common.NewSession()
	sessionMutex.Lock()
	gsSession = ns
	sessionMutex.Unlock()

	return gsSession
}

func newCheckExpiredGSSessionService(gsSession *common.Session) *checkExpiredGSSessionService {
	return &checkExpiredGSSessionService{
		gsSession: gsSession,
	}
}

func (s checkExpiredGSSessionService) Run(noGSSession chan bool) {
	if gsSession == nil {
		log.Printf(">>> no game server connected\n\n")
		noGSSession <- true
	} else {
		expirationTime := gsSession.LastPOL.Add(gsSessionLastPOLExpirationInterval)
		now := time.Now()
		if expirationTime.Before(now) {
			gsSession = nil
		}
	}
}

func startCheckingGameServerSessionExpired(session *common.Session) {
	stopTicker := make(chan bool)
	service := newCheckExpiredGSSessionService(session)
	common.StartTicker(checkExpiredGSSessionInterval, service, stopTicker)
}

func (ge *GameEngine) Connect(ctx context.Context, in *ConnectPayload) (*ConnectResponse, error) {
	if gsSession != nil {
		delayConnection()
		if gsSession != nil {
			return nil, fmt.Errorf("there isn't any session available, please try later")
		}
	}

	gsSession = newGSSession()

	startCheckingGameServerSessionExpired(gsSession)

	log.Printf(">>> game server connected\n\n")

	return &ConnectResponse{
		Token: gsSession.Token.String(),
	}, nil
}

func getGSSessionByToken(token uuid.UUID) (*common.Session, error) {
	if gsSession == nil {
		return nil, errors.New("the game server hasn't any session active")
	} else if gsSession != nil && gsSession.Token != token {
		return nil, errors.New("the token provided doesn't match the game server session")
	} else if gsSession != nil && gsSession.Token == token {
		return gsSession, nil
	}
	return nil, errors.New("unexpected error")
}

func updateGSSessionLastPOLTimestamp() {
	gsSession.LastPOL = time.Now()
}

func (ge *GameEngine) POL(ctx context.Context, in *POLPayload) (*POLResponse, error) {
	token, err := uuid.Parse(in.SessionToken)
	if err != nil {
		return &POLResponse{
			Err: "invalid token format",
		}, nil
	}

	_, err = getGSSessionByToken(token)
	if err != nil {
		return &POLResponse{
			Err: "invalid token",
		}, nil
	}

	updateGSSessionLastPOLTimestamp()

	return &POLResponse{}, nil
}
