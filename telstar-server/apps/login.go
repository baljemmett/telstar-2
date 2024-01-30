package apps

import (
	"bitbucket.org/johnnewcombe/telstar-library/types"
	"bitbucket.org/johnnewcombe/telstar/config"
	"bitbucket.org/johnnewcombe/telstar/dal"
	"bitbucket.org/johnnewcombe/telstar/session"
	"errors"
)

func Login(sessionId string, settings config.Config, args []string) (bool, error) {

	var (
		user     types.User
		userId   string
		password string
		err error
	)
	if !(len(args) <= 3) { // un, password, current pageId
		return false, errors.New("Incorrect number of arguments for login")
	}
	userId = args[0]
	password = args[1]

	if user, err = dal.GetUser(settings.Database.Connection, userId); err != nil {
		return false, err
	}

	if !dal.CheckPasswordHash(password, user.Password) {
		user.Authenticated = false
		return false, errors.New("password not valid")
	}

	// we update the post action frame depending upon the success of login
	user.Authenticated = true
	session.UpdateCurrentUser(sessionId, user)

	return user.Authenticated, nil
}


