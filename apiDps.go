package dps

// Define Policy RestAPI to communicate with policy service
import (
	"encoding/json"
	"fmt"
	logger "github.com/ShaoWenAcerLu/BCIR_main/utilities"
	"github.com/ShaoWenAcerLu/BCIR_planBuilder"
	"github.com/ShaoWenAcerLu/BCIR_policyManager"
	"github.com/gorilla/mux"
	"net/http"
)

var BCIRPolicy policyManager.PolicyJSON

// CreatePolicy Create policy and save in DB
func CreatePolicy(w http.ResponseWriter, r *http.Request) {
	policyName := mux.Vars(r)["policy_name"]
	if !policyManager.CreatePolicy(w, r) {
		logger.LogError("Fail to create policy \"%s\"", policyName)
		return
	}

	if policy, ok := policyManager.LoadPolicy(policyName); ok {
		if policy.Status == "Draft" {
			return
		}
	}

	msgs, err := planBuilder.BuildPlan(policyName)
	if err != nil {
		msg := fmt.Sprintf("Fail to build protection plan for policy %s because %s", policyName, err)
		logger.LogError(msg)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message:    msg,
			StatusCode: 406,
		})
	} else {
		json.NewEncoder(w).Encode(SuccessResponse{
			Messages: msgs,
			Status:   "Success",
		})
	}
}

func UpdatePolicy(w http.ResponseWriter, r *http.Request) {
	policyManager.UpdatePolicy(w, r)
}

func GetPolicy(w http.ResponseWriter, r *http.Request) {
	policyManager.GetPolicy(w, r)
}

func GetPolicies(w http.ResponseWriter, r *http.Request) {
	policyManager.GetPolicies(w, r)
}

func DeletePolicy(w http.ResponseWriter, r *http.Request) {
	policyManager.DeletePolicy(w, r)
}
