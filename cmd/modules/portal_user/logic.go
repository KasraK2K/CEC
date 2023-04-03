package portal_user

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
	md "app/cmd/models"
	"app/pkg/helper"
)

func (l *logic) List(filter md.PortalUserFilter) ([]md.PortalUser, common.Status, error) {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}

	results, status, err := Repository.List(filter, []string{"password"}...)
	if err != nil {
		return []md.PortalUser{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(portalUser md.PortalUser) (md.PortalUser, common.Status, error) {
	if len(portalUser.Email) > 0 {
		portalUser.Email = strings.ToLower(portalUser.Email)
	}

	// Hash password
	if len(portalUser.Password) > 0 {
		hash, err := helper.HashPassword(portalUser.Password)
		if err != nil {
			return md.PortalUser{}, http.StatusInternalServerError, err
		}
		portalUser.Password = hash
	}

	result, status, err := Repository.Insert(portalUser)
	if err != nil {
		return md.PortalUser{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter md.PortalUserFilter, update md.PortalUserUpdate) (md.PortalUser, common.Status, error) {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}

	if len(update.Email) > 0 {
		update.Email = strings.ToLower(update.Email)
	}

	var portalUser md.PortalUser
	err := mapstructure.Decode(update, &portalUser)
	if err != nil {
		return md.PortalUser{}, http.StatusInternalServerError, err
	}

	// Hash password
	if len(update.Password) > 0 {
		hash, err := helper.HashPassword(update.Password)
		if err != nil {
			return md.PortalUser{}, http.StatusInternalServerError, err
		}
		update.Password = hash
	}

	result, status, err := Repository.Update(filter, portalUser)
	if err != nil {
		return md.PortalUser{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter md.PortalUserFilter) (md.PortalUserFilter, common.Status, error) {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}

	result, status, err := Repository.Archive(filter)
	if err != nil {
		return md.PortalUserFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter md.PortalUserFilter) (md.PortalUserFilter, common.Status, error) {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}

	result, status, err := Repository.Restore(filter)
	if err != nil {
		return md.PortalUserFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Login(payload md.PortalUserLoginPayload) (string, common.Status, error) {
	if len(payload.Email) > 0 {
		payload.Email = strings.ToLower(payload.Email)
	}
	filter := md.PortalUserFilter{Email: payload.Email}

	results, status, err := Repository.List(filter)
	if err != nil {
		return "", status, err
	}

	if len(results) == 0 {
		return "", http.StatusNotFound, errors.New("email or password is wrong")
	}

	portalUser := results[0]
	if !helper.ComparePassword(portalUser.Password, payload.Password) {
		return "", http.StatusNotFound, errors.New("email or password is wrong")
	}

	payloadClaims := helper.PayloadClaims{
		ID:         portalUser.ID,
		Permission: portalUser.Permission,
		Platform:   payload.Platform,
		UserType:   helper.Token.UserType.Portal,
	}
	token, err := helper.Token.CreateToken(payloadClaims)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}

func (l *logic) ForgotPassword(email string) (string, common.Status, error) {
	if len(email) > 0 {
		email = strings.ToLower(email)
	}
	filter := md.PortalUserFilter{Email: email}
	password := helper.RandomString(30)
	update := md.PortalUserUpdate{Password: password}

	_, status, err := l.Update(filter, update)
	if err != nil {
		return "", status, err
	}

	body := "<html><body>Your password is changed and your new password is <h3 style=\"display:inline\">%s</strong></body></html>"

	payload := helper.EmailPayload{
		Recipients: []string{email},
		Body:       fmt.Sprintf(body, password),
		Subject:    "Change Password - CEC",
	}
	_, _, err = helper.SendEmail(payload)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("your password has been changed but we have a problem on sending email. error: %s", err.Error())
	}

	return "password successfully changed and sent to your email", status, nil
}
