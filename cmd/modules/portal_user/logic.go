package portal_user

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
	"app/pkg/helper"
)

func (l *logic) List(filter PortalUserFilter) ([]PortalUser, common.Status, error) {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}

	results, status, err := Repository.List(filter, []string{"password"}...)
	if err != nil {
		return []PortalUser{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(portalUser PortalUser) (PortalUser, common.Status, error) {
	// Hash password
	if len(portalUser.Password) > 0 {
		hash, err := helper.HashPassword(portalUser.Password)
		if err != nil {
			return PortalUser{}, http.StatusInternalServerError, err
		}
		portalUser.Password = hash
	}

	result, status, err := Repository.Insert(portalUser)
	if err != nil {
		return PortalUser{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter PortalUserFilter, portalUser PortalUser) (PortalUser, common.Status, error) {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}

	var portalUserUpdate PortalUserUpdate
	err := mapstructure.Decode(portalUser, &portalUserUpdate)
	if err != nil {
		return PortalUser{}, http.StatusInternalServerError, err
	}

	// Hash password
	if len(portalUser.Password) > 0 {
		hash, err := helper.HashPassword(portalUser.Password)
		if err != nil {
			return PortalUser{}, http.StatusInternalServerError, err
		}
		portalUser.Password = hash
	}

	result, status, err := Repository.Update(filter, portalUser)
	if err != nil {
		return PortalUser{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter PortalUserFilter) (PortalUserFilter, common.Status, error) {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}

	result, status, err := Repository.Archive(filter)
	if err != nil {
		return PortalUserFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter PortalUserFilter) (PortalUserFilter, common.Status, error) {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}

	result, status, err := Repository.Restore(filter)
	if err != nil {
		return PortalUserFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Login(payload PortalUserLoginPayload) (string, common.Status, error) {
	if len(payload.Email) > 0 {
		payload.Email = strings.ToLower(payload.Email)
	}
	filter := PortalUserFilter{Email: payload.Email}

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
		ID:       portalUser.ID,
		RoleID:   1,
		Platform: uint8(payload.Platform),
		UserType: helper.Token.UserType.Portal,
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
	filter := PortalUserFilter{Email: email}
	password := helper.RandomString(30)
	update := PortalUser{Password: password}

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
