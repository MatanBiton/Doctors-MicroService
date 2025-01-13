package testing

import (
	"context"
	//"fmt"
	//"net/http"
	//"strings"
	"testing"

	//"github.com/TekClinic/API-Gateway/middlewares"
	//"github.com/TekClinic/API-Gateway/schemas"
	doctors "github.com/TekClinic/Doctors-MicroService/doctors_protobuf"
	//"github.com/gin-gonic/gin"
)


const resourceNameDoctor = "doctor"

type DoctorsParams struct {
	Skip   int32  `form:"skip,default=0"`
	Limit  int32  `form:"limit,default=20"`
	Search string `form:"search" binding:"omitempty,min=1,max=100"`
}

// create a test that checks if the getDoctors function works as expected with hard-codded paramters and no gin integration
func TestGetDoctor(t *testing.T) {
	// create a new instance of the doctor service client
	client := doctors.NewDoctorsServiceClient(nil)

	// create a new context
	ctx := context.Background()

	// create a new request
	req := &doctors.GetDoctorRequest{
		Token: "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwbDdyUUphVlc2Wk5seVNiR01FeU1ZOUc2NzREbVRnazVnZUZvTXJRdFdBIn0.eyJleHAiOjE3MzY0OTY4NjcsImlhdCI6MTczNjQ5NjU2NywiYXV0aF90aW1lIjoxNzM2NDk2NTY3LCJqdGkiOiI4MzE5MDM4NS02NGJkLTRhMTktYWZiMi0xMTk0MDAxYmNiYjkiLCJpc3MiOiJodHRwOi8vYXV0aC50ZWtjbGluaWMub3JnOjgwL3JlYWxtcy90ZWtjbGluaWMiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZTZiYWRiOGMtMjNlZi00NmU3LWI5YWMtZDc4MjZmMThhZWVmIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2ViLWFwcCIsInNlc3Npb25fc3RhdGUiOiIxN2NjZDRjZi1lNDUzLTRmN2UtYjYxZC1jYzM3NjE0M2RkYTUiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy10ZWtjbGluaWMiLCJvZmZsaW5lX2FjY2VzcyIsImFkbWluIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6Im9wZW5pZCBlbWFpbCBwcm9maWxlIiwic2lkIjoiMTdjY2Q0Y2YtZTQ1My00ZjdlLWI2MWQtY2MzNzYxNDNkZGE1IiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3RfdXNlciJ9.WYW5zrKMsOupeMMoBqJIcdV07Ur-5SGEK5F19radkMB4F1onn6Sx1Hs14jPqwJPzPYtxu4wfh6nxn3Qw61FY-7V1kjZ_qGNWGWTW9QEw1oPIPa3WTiNSRs6ojWE9Jx8-wfJThH0WeI_t4fT_N8D-vWvPQLM7sktbQubffJoqp5IJ0NpVlgn2VNnokOiMMn64xKx2l-fu1gJV_euOZD7KVjZSspf8GuWQ4CEcFtaIONf9QWO5ViCK4qoSLW6r4MijCu8BmxWBM0HEd4qJe9XPvxBhi8Ra3sMSZ-p5wUhvP3pXHvDtfIC4VbdWcLWAVxAB90J5TktqnPvEY4quhj60gw",
		Id: 1,
	}

	// call the getDoctor function
	res, err := client.GetDoctor(ctx, req)

	// check if the response is nil
	if res == nil {
		t.Errorf("Expected response to be non-nil")
	}

	// check if the error is nil
	if err != nil {
		t.Errorf("Expected error to be nil")
	}
}


// func getDoctors(service doctors.DoctorsServiceClient) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// fetch params from the query
// 		var params DoctorsParams
// 		err := ctx.ShouldBindQuery(&params)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, schemas.ErrorResponse{
// 				Message: err.Error(),
// 			})
// 			return
// 		}
		
// 		// call doctor microservice
// 		response, err := service.GetDoctorsIDs(ctx, &doctors.GetDoctorsIDsRequest{
// 			Token:  ctx.GetString(middlewares.TokenKey),
// 			Limit:  params.Limit,
// 			Offset: params.Skip,
// 			Search: params.Search,
// 		})
// 		if err != nil {
// 			HandleGRPCError(err, ctx)
// 			return
// 		}
// 		fmt.Println(response.GetResults())
// 		ctx.JSON(http.StatusOK,
// 			CreateNamedAPIResourceList(ctx, resourceNameDoctor,
// 				params.Skip, params.Limit, response.GetCount(), response.GetResults()))
// 	}
// }

// type DoctorParams struct {
// 	ID int32 `uri:"id" binding:"required"`
// }

// func getDoctor(service doctors.DoctorsServiceClient) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// fetch params from the path
// 		var uriParams DoctorParams
// 		err := ctx.ShouldBindUri(&uriParams)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, schemas.ErrorResponse{
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		// call doctor microservice
// 		response, err := service.GetDoctor(ctx, &doctors.GetDoctorRequest{
// 			Token: ctx.GetString(middlewares.TokenKey),
// 			Id:    uriParams.ID,
// 		})
// 		if err != nil {
// 			HandleGRPCError(err, ctx)
// 			return
// 		}

// 		if response.GetDoctor() == nil {
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, schemas.ErrorResponse{
// 				Message: "Invalid response from the server.",
// 			})
// 			return
// 		}

// 		doctor := response.GetDoctor()

// 		specialities := doctor.GetSpecialities()
// 		if specialities == nil {
// 			specialities = []string{}
// 		}

// 		ctx.JSON(http.StatusOK, schemas.Doctor{
// 			DoctorBase: schemas.DoctorBase{
// 				Name:         doctor.GetName(),
// 				Gender:       strings.ToLower(doctor.GetGender().String()),
// 				PhoneNumber:  doctor.GetPhoneNumber(),
// 				Specialities: specialities,
// 				SpecialNote:  doctor.GetSpecialNote(),
// 			},
// 			ID:     doctor.GetId(),
// 			Active: doctor.GetActive(),
// 		})
// 	}
// }

// func createDoctor(service doctors.DoctorsServiceClient) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// fetch params from the body
// 		var bodyParams schemas.DoctorBase
// 		err := ctx.ShouldBindJSON(&bodyParams)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, schemas.ErrorResponse{
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		// call doctor microservice
// 		response, err := service.CreateDoctor(ctx, &doctors.CreateDoctorRequest{
// 			Token:        ctx.GetString(middlewares.TokenKey),
// 			Name:         bodyParams.Name,
// 			Gender:       doctors.Doctor_Gender(doctors.Doctor_Gender_value[strings.ToUpper(bodyParams.Gender)]),
// 			PhoneNumber:  bodyParams.PhoneNumber,
// 			Specialities: bodyParams.Specialities,
// 			SpecialNote:  bodyParams.SpecialNote,
// 		})
// 		if err != nil {
// 			HandleGRPCError(err, ctx)
// 			return
// 		}

// 		ctx.JSON(http.StatusCreated, schemas.IDHolder{
// 			ID: response.GetId(),
// 		})
// 	}
// }

// func deleteDoctor(service doctors.DoctorsServiceClient) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// fetch params from the path
// 		var uriParams DoctorParams
// 		err := ctx.ShouldBindUri(&uriParams)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, schemas.ErrorResponse{
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		// call doctor microservice
// 		_, err = service.DeleteDoctor(ctx, &doctors.DeleteDoctorRequest{
// 			Token: ctx.GetString(middlewares.TokenKey),
// 			Id:    uriParams.ID,
// 		})
// 		if err != nil {
// 			HandleGRPCError(err, ctx)
// 			return
// 		}

// 		ctx.Status(http.StatusOK)
// 	}
// }

// type UpdateDoctorParams struct {
// 	ID int32 `uri:"id" binding:"required"`
// }

// func updateDoctor(service doctors.DoctorsServiceClient) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var uriParams UpdateDoctorParams
// 		err := ctx.ShouldBindUri(&uriParams)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, schemas.ErrorResponse{
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		var bodyParams schemas.DoctorUpdate
// 		err = ctx.ShouldBindJSON(&bodyParams)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, schemas.ErrorResponse{
// 				Message: err.Error(),
// 			})
// 			return
// 		}

// 		// call doctor microservice
// 		response, err := service.UpdateDoctor(ctx, &doctors.UpdateDoctorRequest{
// 			Token: ctx.GetString(middlewares.TokenKey),
// 			Doctor: &doctors.Doctor{
// 				Id:           uriParams.ID,
// 				Active:       bodyParams.Active,
// 				Name:         bodyParams.Name,
// 				Gender:       doctors.Doctor_Gender(doctors.Doctor_Gender_value[strings.ToUpper(bodyParams.Gender)]),
// 				PhoneNumber:  bodyParams.PhoneNumber,
// 				Specialities: bodyParams.Specialities,
// 				SpecialNote:  bodyParams.SpecialNote,
// 			},
// 		})
// 		if err != nil {
// 			HandleGRPCError(err, ctx)
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, schemas.IDHolder{
// 			ID: response.GetId(),
// 		})
// 	}
// }

// func RegisterDoctorRoutes(router *gin.Engine) {
// 	client := InitiateClient(resourceNameDoctor, doctors.NewDoctorsServiceClient)

// 	// deprecated
// 	router.GET("/doctor", getDoctors(client))
// 	router.POST("/doctor", createDoctor(client))
// 	router.GET("/doctor/:id", getDoctor(client))
// 	router.DELETE("/doctor/:id", deleteDoctor(client))
// 	// end deprecated

// 	router.GET("/doctors", getDoctors(client))
// 	router.POST("/doctors", createDoctor(client))
// 	router.GET("/doctors/:id", getDoctor(client))
// 	router.PUT("/doctors/:id", updateDoctor(client))
// 	router.DELETE("/doctors/:id", deleteDoctor(client))
// }
