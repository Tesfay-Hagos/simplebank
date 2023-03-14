package test

/*
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"tesfayprep/user_authentication/controller"
	"tesfayprep/user_authentication/model"
	"testing"
	"time"
)

func TestUserregisterandlogin(t *testing.T) {
	server := controller.Newserver()
	t.Run("Test Registering users", func(t *testing.T) {
		Newuser := model.UserInfo{Username: "Tskadkan", Password: "tskadkaney2121", Email: "tsadkan2121@gmail.com", CreatedAt: time.Now()}
		//create request and response writer
		buff := convtobuff(Newuser)
		req := httptest.NewRequest(http.MethodPost, "/register", &buff)
		resp := httptest.NewRecorder()
		//initiate server
		server.Handler.ServeHTTP(resp, req)
		response := model.JsonResponse{}
		json.NewDecoder(resp.Body).Decode(&response)
		//check response
		if response.Type != "success" {
			t.Errorf("Test Failed")
		}
	})
}
func TestLoginandGet(t *testing.T) {
	server := controller.Newserver()
	t.Run("Test login user and generate token", func(t *testing.T) {
		Newuser := model.UserInfo{Password: "tshomehu2121", Email: "teshhagos@gmail.com"}
		buff := convtobuff(Newuser)
		req := httptest.NewRequest(http.MethodPost, "/login", &buff)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		//t.Errorf("Teoken of login password:%s", resp.Body)
		//t.Logf("\nToken:%s", resp.Body.String())
		if resp.Result().StatusCode != 200 {
			t.Errorf("Test Failed wit statuscode:%d", resp.Result().StatusCode)
		}

		Newuserr := model.ResetPasswordform{Token: resp.Body.String()}
		buffr := convtobuff(Newuserr)
		reqr := httptest.NewRequest(http.MethodPost, "/getalluser", &buffr)
		respr := httptest.NewRecorder()
		server.Handler.ServeHTTP(respr, reqr)
		//t.Errorf("Teoken of login password:%s", respr.Body.String())
		if respr.Result().StatusCode != 200 {
			t.Errorf("Test Failed with statuscode:%d", respr.Result().StatusCode)
		}

		//Test Change password
		Newuserc := model.UserInfo{Password: "tshomehu2121", Email: "teshhagos@gmail.com"}
		buffchangec := convtobuff(Newuserc)
		reqc := httptest.NewRequest(http.MethodPut, "/changepassword", &buffchangec)
		reqc.Header.Add("token", resp.Body.String())
		respc := httptest.NewRecorder()
		server.Handler.ServeHTTP(respc, reqc)
		if respc.Result().StatusCode != 200 {
			t.Errorf("Test Failed with statuscode:%d", respr.Result().StatusCode)
		}

	})
}
func TestPasswordReset(t *testing.T) {
	server := controller.Newserver()
	Newuser := model.ResetPassword{Email: "teshhagos@gmail.com"}
	buff := convtobuff(Newuser)
	req := httptest.NewRequest(http.MethodPost, "/passwordresetrequest", &buff)
	resp := httptest.NewRecorder()
	server.Handler.ServeHTTP(resp, req)
	if resp.Body.String() == "" {
		t.Errorf("Test Failed")

	}
	//t.Errorf("OTP code of resetpass:%s", resp.Body.String())
	otpcode := resp.Body.String()
	Newuserr := model.ResetPasswordform{Email: "shamthagos@gmail.com", Token: otpcode, NewPassword: "Hello123", ConfirmNewPassword: "Hello123"}
	buffr := convtobuff(Newuserr)
	reqr := httptest.NewRequest(http.MethodPut, "/passwordresetwithtoken", &buffr)
	respr := httptest.NewRecorder()
	server.Handler.ServeHTTP(respr, reqr)
	//t.Errorf("Boddy of resetpass:%s", respr.Body.String())
	if respr.Code != http.StatusOK {
		t.Errorf("Test Failed with statuscode:%d", respr.Result().StatusCode)
	}
}

func convtobuff(user interface{}) bytes.Buffer {
	body, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	buff := bytes.NewBuffer(body)
	return *buff

}
*/
