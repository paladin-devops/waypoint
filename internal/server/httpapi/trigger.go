package httpapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/waypoint/internal/clicontext"
	pb "github.com/hashicorp/waypoint/internal/server/gen"
	"github.com/hashicorp/waypoint/internal/serverclient"
	"github.com/hashicorp/waypoint/internal/serverconfig"
)

// TODO(briancain): write tests
// HandleTrigger will execute a run trigger, if the requested id exists
// This works by connecting back to our own local gRPC server.
func HandleTrigger(addr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := hclog.FromContext(ctx)
		log.SetLevel(hclog.Debug)

		// Authless trigger URLs should be able to make a request
		// without a token.
		token := r.URL.Query().Get("token")
		requireAuth := true
		if token == "" {
			log.Trace("no token provided, will attempt to run authless trigger")
			requireAuth = false
			//http.Error(w, "no token provided", 403)
		}

		// Connect back to our own gRPC service.
		grpcConn, err := serverclient.Connect(ctx,
			serverclient.Logger(log),
			serverclient.FromContextConfig(&clicontext.Config{
				Server: serverconfig.Client{
					Address:     addr,
					RequireAuth: requireAuth,
					AuthToken:   token,

					// Our gRPC server should always be listening on TLS.
					// We ignore it because its coming out of our own process.
					Tls:           true,
					TlsSkipVerify: true,
				},
			}),
		)
		if err != nil {
			log.Error("trigger connection back to gRPC failed", "err", err)
			return
		}
		defer grpcConn.Close()

		// Our API client
		client := pb.NewWaypointClient(grpcConn)

		requestVars := mux.Vars(r)
		runTriggerId := requestVars["id"]
		//triggerOverrideVars := r.URL.Query().Get("variables")

		var resp *pb.RunTriggerResponse
		runTriggerReq := &pb.RunTriggerRequest{
			Ref: &pb.Ref_Trigger{
				Id: runTriggerId,
			},
			VariableOverrides: nil, // TODO fix me
		}

		if requireAuth {
			// attempt to make a grpc request to run trigger by id
			resp, err = client.RunTrigger(ctx, runTriggerReq)
		} else {
			// attempt to make a grpc request to run trigger by id
			resp, err = client.AuthlessRunTrigger(ctx, runTriggerReq)
		}
		if err != nil {
			log.Error("server failed to run trigger", "id", runTriggerId, "err", err)

			if status.Code(err) == codes.PermissionDenied {
				http.Error(w, fmt.Sprintf("request not authorized to run trigger: %s", err), 401)
			} else {
				// improve http error code, which is more applicable for general queue failures?
				http.Error(w, fmt.Sprintf("server failed to run trigger: %s", err), 412)
			}

			return
		}
		if resp == nil {
			http.Error(w, fmt.Sprintf("server returned no job ids from run trigger %q", runTriggerId), 500)
			return
		}
		jobIds := resp.JobIds

		// TODO(briancain): attempt to stream output back, on request.
		for _, jId := range jobIds {
			_, err := client.GetJobStream(ctx, &pb.GetJobStreamRequest{
				JobId: jId,
			})
			if err != nil {
				log.Error("server failed to get job stream output for trigger", "job_id", jId, "err", err)
			}
		}
	}
}
