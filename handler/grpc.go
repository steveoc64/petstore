package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/steveoc64/petstore/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpcRun runs the RPC server
// Looks ugly, but its just common boilerplate that would normally be in a lib
func (s *PetstoreServer) grpcRun() {
	endpoint := fmt.Sprintf(":%d", s.rpcPort)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		s.log.WithError(err).Fatal("failed to listen")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPetstoreServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	s.log.WithField("endpoint", endpoint).Println("Serving gRPC")

	grpcServer.Serve(lis)
}

// rpcProxy hooks up the REST endpoints.
// Looks ugly, but its just common boilerplate that would normally be in a lib
func (s *PetstoreServer) rpcProxy(log *logrus.Logger) error {
	// Use our custom error handler
	runtime.HTTPError = CustomHTTPError
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// The UpdatePet resource uses both form-encoded data and POST, so need some custom
	// code here to handle that
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(apiKeyMatcher),
		// this gets around that nasty bit about having 2 patterns that
		// match on the GET pet/{PetID} and GET pet/fetchByStatus
		runtime.WithLastMatchWins(),
	)
	runtime.SetHTTPBodyMarshaler(mux)

	formWrapper := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// handle incoming form data - rewrite it as JSON for grpc handling
			contentType := strings.ToLower(strings.Split(r.Header.Get("Content-Type"), ";")[0])
			switch contentType {
			case "application/x-www-form-urlencoded":
				if err := r.ParseForm(); err != nil {
					log.WithFields(logrus.Fields{
						"method":  r.Method,
						"url":     r.URL,
						"browser": r.UserAgent(),
						"ip":      r.RemoteAddr,
					}).WithError(err).Error("Failed to parse form")
					http.Error(w, err.Error(), http.StatusMethodNotAllowed) // strange value, but thats whats in the swagger spec
					return
				}
				jsonMap := make(map[string]interface{}, len(r.Form))
				for k, v := range r.Form {
					if len(v) > 0 {
						jsonMap[k] = v[0]
					}
				}
				jsonBody, err := json.Marshal(jsonMap)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				r.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
				r.ContentLength = int64(len(jsonBody))
				r.Header.Set("Content-Type", "application/json")
			case "multipart/form-data":
				// 4 MB image limit should do for now
				if err := r.ParseMultipartForm(4 << 20); err != nil {
					log.WithFields(logrus.Fields{
						"method":  r.Method,
						"url":     r.URL,
						"browser": r.UserAgent(),
						"ip":      r.RemoteAddr,
					}).WithError(err).Error("Failed to parse multipart form")
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				file, handler, err := r.FormFile("file")
				if err != nil {
					log.WithFields(logrus.Fields{
						"method":  r.Method,
						"url":     r.URL,
						"browser": r.UserAgent(),
						"ip":      r.RemoteAddr,
					}).WithError(err).Error("Failed to get file data")
					return
				}
				defer file.Close()
				//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
				//fmt.Printf("File Size: %+v\n", handler.Size)
				//fmt.Printf("MIME Header: %+v\n", handler.Header)

				tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
				if err != nil {
					fmt.Println(err)
				}
				defer tempFile.Close()

				// read all of the contents of our uploaded file into a
				// byte array
				fileBytes, err := ioutil.ReadAll(file)
				if err != nil {
					fmt.Println(err)
				}
				// write this byte array to our temporary file
				tempFile.Write(fileBytes)
				// return that we have successfully uploaded our file!
				fmt.Fprintf(w, "Successfully Uploaded File\n")
			default:
				println("content type", contentType)
			}
			mux.ServeHTTP(w, r)
		})
	}

	/*
		// TODO - these elapsed timing metrics should be sent to a metrics
		// service here ... replace logging code with metrics code
		timingWrapper := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				t1 := time.Now()
				mux.ServeHTTP(w, r)
				s.log.WithField("elapsed", time.Since(t1).String()).Info("Call Duration")
			})
		}
	*/

	opts := []grpc.DialOption{grpc.WithInsecure()}
	rpcendpoint := fmt.Sprintf(":%d", s.rpcPort)
	webendpoint := fmt.Sprintf(":%d", s.restPort)
	err := pb.RegisterPetstoreServiceHandlerFromEndpoint(ctx, mux, rpcendpoint, opts)
	if err != nil {
		return err
	}

	s.log.WithField("endpoint", webendpoint).Println("Serving REST Proxy")
	return http.ListenAndServe(webendpoint, formWrapper(mux))
}

// apiKeyMatcher looks for the API_KEY in the header, and includes it in the grpc data
func apiKeyMatcher(key string) (string, bool) {
	switch key {
	case "Api_key", "api_key":
		return key, true
	default:
		return key, false
	}
}

type errorBody struct {
	Err string `json:"error,omitempty"`
}

// CustomHTTPError for stripping error contents back
func CustomHTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	contentType := marshaler.ContentType()
	if httpBodyMarshaler, ok := marshaler.(*runtime.HTTPBodyMarshaler); ok {
		pb := s.Proto()
		contentType = httpBodyMarshaler.ContentTypeFromMessage(pb)
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Del("Trailer")

	errMsg := s.Message()
	statusCode := runtime.HTTPStatusFromCode(status.Code(err))

	// Examine leader on the message, and use that for the custom error code
	if len(errMsg) > 3 && errMsg[3] == ':' {
		v, err := strconv.Atoi(errMsg[:3])
		if err == nil {
			statusCode = v
			errMsg = errMsg[4:]
		}
	}
	e := errorBody{Err: errMsg}
	w.WriteHeader(statusCode)
	jErr := json.NewEncoder(w).Encode(e)

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}
