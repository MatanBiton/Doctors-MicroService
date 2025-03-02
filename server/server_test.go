package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dpb "github.com/TekClinic/Doctors-MicroService/doctors_protobuf"
)

// Test GetDoctor with a valid ID
func TestGetDoctor_Success(t *testing.T) {
	// Initialize test server
	server, err := createTestDoctorsServer()
	require.NoError(t, err)
	require.NotNil(t, server) // Ensure server is not nil
	require.NotNil(t, server.db) // Ensure database is not nil

	// Insert a test doctor
	doctor := Doctor{
		ID:           1,
		Active:       true,
		Name:         "Dr. John Doe",
		Gender:       1,
		PhoneNumber:  "+1234567890",
		Specialities: []string{"Cardiology"},
		SpecialNote:  "Experienced heart surgeon",
	}
	_, err = server.db.NewInsert().Model(&doctor).Exec(context.Background())
	require.NoError(t, err)

	// Call GetDoctor with a mock token
	req := &dpb.GetDoctorRequest{
		Id:    1,
		Token: "mock_token",
	}

	resp, err := server.GetDoctor(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Dr. John Doe", resp.Doctor.Name)
}

// Test GetDoctor with a non-existing ID
func TestGetDoctor_NotFound(t *testing.T) {
	server, err := createTestDoctorsServer()
	require.NoError(t, err)

	req := &dpb.GetDoctorRequest{
		Id:    99, // Non-existent doctor
		Token: "valid_token",
	}

	resp, err := server.GetDoctor(context.Background(), req)

	assert.Nil(t, resp)
	assert.Equal(t, codes.NotFound, status.Code(err))
}

// Test CreateDoctor with valid data
func TestCreateDoctor_Success(t *testing.T) {
	server, err := createTestDoctorsServer()
	require.NoError(t, err)

	req := &dpb.CreateDoctorRequest{
		Token:        "valid_token",
		Name:         "Dr. Alice Smith",
		Gender:       2,
		PhoneNumber:  "+1987654321",
		Specialities: []string{"Neurology"},
		SpecialNote:  "Expert in brain surgery",
	}

	resp, err := server.CreateDoctor(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Id > 0) // ID should be auto-generated
}

// Test CreateDoctor with missing required fields
func TestCreateDoctor_InvalidData(t *testing.T) {
	server, err := createTestDoctorsServer()
	require.NoError(t, err)

	req := &dpb.CreateDoctorRequest{
		Token: "valid_token",
		// Missing required Name & PhoneNumber
	}

	resp, err := server.CreateDoctor(context.Background(), req)

	assert.Nil(t, resp)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}
