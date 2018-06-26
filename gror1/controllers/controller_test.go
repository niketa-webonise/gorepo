package controllers

import (
   // "encoding/hex"
    "errors"
   // "fmt"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gorilla/mux"
    "../models"
)

type CreateDockerFailImplTest struct {
}
type CreateDockerSuccessImplTest struct {
}

type GetItemIDSuccessImplTest struct {
}

type GetItemIDFailImplTest struct {
}

type GetItemReqSuccessImplTest struct {
}

type GetListSuccessImplTest struct {
}

type GetListFailImplTest struct {
}

type UpdateTest struct{
}

var testCaseCreateFail = []struct {
    Method       string
    URL          string
    Message      string
    JSONBody     string
    expectStatus int
}{
    {
        Method:       "POST",
        URL:          "/docker/config",
        Message:      "unprocessable entity",
        JSONBody:     "{\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\" }",
        expectStatus: 422,
    },
}
var testCaseCreateSuccess = []struct {
    Method       string
    URL          string
    Message      string
    JSONBody     string
    expectStatus int
}{
    {
        Method:       "POST",
        URL:          "/docker/config",
        Message:      "successfully created",
        JSONBody:     "{\"systemInfo\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"} }",
        expectStatus: 200,
    },
}

var testCaseGetItemIDSuccess = []struct {
    Method       string
    URL          string
    JSONBody     string
    expectStatus int
}{
    {
        Method:       "GET",
        URL:          "/docker/config/5b28b442a90362768113e47e",
        expectStatus: 400,
    },
}

var testCaseGetItemIDFail = []struct {
    ID           string
    Method       string
    URL          string
    JSONBody     string
    expectStatus int
}{
    {
        ID:           "1100",
        Method:       "GET",
        URL:          "/docker/config/1100",
        expectStatus: 400,
    },
}

var testCaseGetItemFullSuccess = []struct {
    ID           string
    Method       string
    URL          string
    JSONBody     string
    expectStatus int
}{
    {
        ID:           "5b28b442a90362768113e47e",
        Method:       "GET",
        URL:          "/docker/config/5b28b442a90362768113e47e",
        expectStatus: 200,
    },
}

var testCaseSuccessDockerList = []struct {
    Method       string
    URL          string
    expectStatus int
}{
    {
        Method:       "GET",
        URL:          "/docker/config",
        expectStatus: 200,
    },
}

var testCaseUpdateID = []struct {
    ID           string
    Method       string
    URL          string
    JSONBody     string
    expectStatus int
}{
    {
        ID:           "5b28b442a90362768113e47e",
        Method:       "PUT",
        URL:          "/docker/config/5b28b442a90362768113e47e",
        JSONBody:     "{\"system\": {\"grorVersion\": \"1.0.0\",\"name\": \"cocooncam\"} }",
        expectStatus: 200,
    },
}

func (s CreateDockerSuccessImplTest) InsertData(bytevalue []byte) error {
    return nil
}

func (s CreateDockerFailImplTest) InsertData(bytevalue []byte) error {
    return errors.New("The request could not be completed because of a conflict")
}

func (s GetItemIDSuccessImplTest) GetItem(bytevalue []byte) (models.Root, error) {
    var rootobject models.Root
    return rootobject, errors.New("id is found but conflict in JSON body")
}

func (s GetItemIDFailImplTest) GetItem(bytevalue []byte) (models.Root, error) {
    var rootobject models.Root
    return rootobject, errors.New("id not found")
}

func (s GetItemReqSuccessImplTest) GetItem(bytevalue []byte) (models.Root, error) {
    var rootobject models.Root
    return rootobject, nil
}

func (s UpdateTest) UpdateData(bytevalue []byte) ( error) {
    
    return  nil
}

func (s GetListSuccessImplTest) GetList() ([]string, []string, error) {
    names := []string{"cacoon-cam", "cacooncam"}
    ids := []string{"5b2cd0a9a90362508f80f71d", "5b28b442a90362768113e47e"}
    return names, ids, nil
}

func (s GetListFailImplTest) GetList() ([]string, []string, error) {
    return []string{}, []string{}, errors.New("unable list")
}

func TestCreateDockerConfig(t *testing.T) {

    s := CreateDockerControllerImpl{
        CreateDockerService: CreateDockerSuccessImplTest{},
    }
    for _, test := range testCaseCreateSuccess {

        router := mux.NewRouter()
        router.HandleFunc("/docker/config/{id}", s.CreateDockerConfig()).Methods("GET")
        ts := httptest.NewServer(router)
        defer ts.Close()

        _, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))
        if err != nil {
            t.Fatal(err)
        }

        // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
        rr := httptest.NewRecorder()
        //handler := http.HandlerFunc(s.CreateDockerConfig())

        // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
        // directly and pass in our Request and ResponseRecorder.
        //handler.ServeHTTP(rr, req)

        // Check the status code is what we expect.
        if status := rr.Code; status != test.expectStatus {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, test.expectStatus)
        }

    }

    s = CreateDockerControllerImpl{
        CreateDockerService: CreateDockerFailImplTest{},
    }

    for _, test := range testCaseCreateFail {

        router := mux.NewRouter()
        ts := httptest.NewServer(router)
        defer ts.Close()
        //fmt.Println("testCreateDockerFail JSON body", strings.NewReader(test.JSONBody))
        req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.JSONBody))
        if err != nil {
            t.Fatal(err)
        }

        // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(s.CreateDockerConfig())

        // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
        // directly and pass in our Request and ResponseRecorder.
        handler.ServeHTTP(rr, req)

        // Check the status code is what we expect.
        if status := rr.Code; status != test.expectStatus {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, test.expectStatus)
        }
    }
}



// func TestGetItem(t *testing.T) {
//     s := GetDockerItemControllerImpl{
//         GetDockerService: GetItemIDSuccessImplTest{},
//     }

//     for _, test := range testCaseGetItemIDSuccess {
//         router := mux.NewRouter()
//         //mux.Vars(req)[h.name] = h.value
//         ts := httptest.NewServer(router)
//         defer ts.Close()

//         req, err := http.NewRequest(test.Method, test.URL, nil)
//         if err != nil {
//             t.Fatal(err)
//         }
//         rr := httptest.NewRecorder()
//         handler := http.HandlerFunc(s.GetDockerConfig())

//         handler.ServeHTTP(rr, req)

//         if status := rr.Code; status != test.expectStatus {
//             t.Errorf("handler returned wrong status code: got %v want %v",
//                 status, http.StatusBadRequest)
//         }

//     }

//     // s = GetDockerItemControllerImpl{
//     //     GetDockerService: GetItemIDFailImplTest{},
//     // }

//     // for _, test := range testCaseGetItemIDFail {
//     //     router := mux.NewRouter()
//     //     ts := httptest.NewServer(router)
//     //     defer ts.Close()

//     //     req, err := http.NewRequest(test.Method, test.URL, nil)
//     //     if err != nil {
//     //         t.Fatal(err)
//     //     }
//     //     rr := httptest.NewRecorder()
//     //     handler := http.HandlerFunc(s.GetDockerConfig())

//     //     handler.ServeHTTP(rr, req)

//     //     IDval, _ := hex.DecodeString(test.ID)
//     //     if status := rr.Code; status != test.expectStatus {
//     //         t.Errorf("handler returned wrong status code: got %v want %v",
//     //             status, http.StatusBadRequest)
//     //     } else {
//     //         fmt.Printf("Invalid ID:%s", hex.EncodeToString(IDval))
//     //     }
//     // }
//     s = GetDockerItemControllerImpl{
//         GetDockerService: GetItemReqSuccessImplTest{},
//     }


    

//     for _, test := range testCaseGetItemFullSuccess {
        
//         router := mux.NewRouter()
//         router.HandleFunc("/docker/config/{id}", s.GetDockerConfig())
//             ts := httptest.NewServer(router)
//             //ts.URL = "/docker/config/{id}"
//             defer ts.Close()

//         req, err := http.NewRequest(test.Method, test.URL, nil)

//         //mux.SetURLVars(req, map[string]string{"id": "1100"})

//         //req = SetURLVars(req, map[string]string{"id": "1100"})
//         //mux.Vars(req) = map[string]string{"id": "1100"} 

//        // mux.Vars(req) = make(map[string]string, 15)
//          //   mux.Vars(req)["id"] = "1100"

//         if err != nil {
//             t.Fatal(err)
//         }
//         rr := httptest.NewRecorder()
//         //handler := http.HandlerFunc(s.GetDockerConfig())
//         router.ServeHTTP(rr, req)
//         //fmt.Println("response", rr)*/

//         //resp, err:= http.Get(test.URL)

//         //fmt.Println(resp.StatusCode)
//         //if err != nil {
//           //  panic(err)
//         //}

//         if status := rr.Code; status != test.expectStatus {
//             t.Errorf("handler returned wrong status code: got %v want %v",
//                 status, test.expectStatus)
//         }
//     }
// }

func TestUpdate(t *testing.T){
    s := UpdateDockerControllerImpl{
        UpdateDockerService: UpdateTest{},
    }


    

    for _, test := range testCaseUpdateID {
        
        router := mux.NewRouter()
        router.HandleFunc("/docker/config/{id}", s.UpdateDockerConfig())
            ts := httptest.NewServer(router)
            //ts.URL = "/docker/config/{id}"
            defer ts.Close()

        req, err := http.NewRequest(test.Method, test.URL, 
            strings.NewReader(test.JSONBody))

        //mux.SetURLVars(req, map[string]string{"id": "1100"})

        //req = SetURLVars(req, map[string]string{"id": "1100"})
        //mux.Vars(req) = map[string]string{"id": "1100"} 

       // mux.Vars(req) = make(map[string]string, 15)
         //   mux.Vars(req)["id"] = "1100"

        if err != nil {
            t.Fatal(err)
        }
        rr := httptest.NewRecorder()
        //handler := http.HandlerFunc(s.GetDockerConfig())
        router.ServeHTTP(rr, req)
        //fmt.Println("response", rr)*/

        //resp, err:= http.Get(test.URL)

        //fmt.Println(resp.StatusCode)
        //if err != nil {
          //  panic(err)
        //}

        if status := rr.Code; status != test.expectStatus {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, test.expectStatus)
        }
    }
}

/*
func TestGetItemOld(t *testing.T) {
    s := GetDockerItemControllerImpl{
        GetDockerService: GetItemIDSuccessImplTest{},
    }

    for _, test := range testCaseGetItemIDSuccess {

        handler := http.HandlerFunc(s.GetDockerConfig())
        mux.HandlerFunc("/docker/config/{id}",handler)

        router := mux.NewRouter()
        ts := httptest.NewServer(router)
        defer ts.Close()

        req, err := http.NewRequest(test.Method, test.URL, nil)
        if err != nil {
            t.Fatal(err)
        }
        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(s.GetDockerConfig())

        handler.ServeHTTP(rr, req)

        //http.Handle("/docker/config/{id}",handler)

  //log.Println("Listening...")
  //http.ListenAndServe(":3000", mux)

        if status := rr.Code; status != test.expectStatus {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, http.StatusBadRequest)
        }

    }

    s = GetDockerItemControllerImpl{
        GetDockerService: GetItemIDFailImplTest{},
    }

    for _, test := range testCaseGetItemIDFail {
        router := mux.NewRouter()
        ts := httptest.NewServer(router)
        defer ts.Close()

        req, err := http.NewRequest(test.Method, test.URL, nil)
        if err != nil {
            t.Fatal(err)
        }
        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(s.GetDockerConfig())

        handler.ServeHTTP(rr, req)

        IDval, _ := hex.DecodeString(test.ID)
        if status := rr.Code; status != test.expectStatus {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, http.StatusBadRequest)
        } else {
            fmt.Printf("Invalid ID:%s", hex.EncodeToString(IDval))
        }
    }
    s = GetDockerItemControllerImpl{
        GetDockerService: GetItemReqSuccessImplTest{},
    }

    for _, test := range testCaseGetItemFullSuccess {
        router := mux.NewRouter()
        ts := httptest.NewServer(router)
        defer ts.Close()

        req, err := http.NewRequest(test.Method, test.URL, nil)

        if err != nil {
            t.Fatal(err)
        }
        rr := httptest.NewRecorder()
        handler := http.HandlerFunc(s.GetDockerConfig())
        handler.ServeHTTP(rr, req)
        //fmt.Println("response", rr)

        if status := rr.Code; status != test.expectStatus {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, test.expectStatus)
        }
    }
}*/