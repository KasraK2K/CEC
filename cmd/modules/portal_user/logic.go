package portal_user

import (
	"errors"
	"net/http"

	"github.com/mitchellh/mapstructure"

	"app/cmd/common"
	"app/pkg/helper"
)

func (l *logic) List(filter PortalUserFilter) ([]PortalUser, common.Status, error) {
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
	result, status, err := Repository.Archive(filter)
	if err != nil {
		return PortalUserFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter PortalUserFilter) (PortalUserFilter, common.Status, error) {
	result, status, err := Repository.Restore(filter)
	if err != nil {
		return PortalUserFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Login(payload PortalUserLoginPayload) (string, common.Status, error) {
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
