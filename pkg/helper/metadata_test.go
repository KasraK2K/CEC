package helper

import "testing"

func TestAddMetaData(t *testing.T) {
	mockData := "Test data"

	// Test case 1 - Success
	expectedSuccess := true
	expectedResult := mockData
	md := AddMetaData(mockData, false)
	if md.SUCCESS != expectedSuccess {
		t.Errorf("AddMetaData() success flag = %v; expected %v", md.SUCCESS, expectedSuccess)
	}
	if md.RESULT != expectedResult {
		t.Errorf("AddMetaData() result = %v; expected %v", md.RESULT, expectedResult)
	}

	// Test case 2 - Error
	expectedSuccess = false
	expectedError := mockData
	md = AddMetaData(mockData, true)
	if md.SUCCESS != expectedSuccess {
		t.Errorf("AddMetaData() success flag = %v; expected %v", md.SUCCESS, expectedSuccess)
	}
	if md.ERRORS != expectedError {
		t.Errorf("AddMetaData() error = %v; expected %v", md.ERRORS, expectedError)
	}
}
