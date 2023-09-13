package SendPriorityMessageToAssets

import (
	"encoding/base64"
	"encoding/json"
	"strconv"

	"fmt"
	// "io"
	"net/http"
	"net/url"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	Status := SendPriorityMessageToAssets(input.IP, input.CustomerId, input.Username, input.Password, input.StaffIdList, input.Message)

	output := &Output{Status: Status}

	// fmt.Println("Output: ", output.Status)
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

func SendPriorityMessageToAssets(IP string, CustomerId string, username string, password string, StaffIdList string, Message string)string{

	// filtering url string's special characters
	FilteredMessage := url.QueryEscape(Message)

	// Create an HTTP client
	client := &http.Client{}

	var asset Asset
	assetCheck := json.Unmarshal([]byte(StaffIdList), &asset) // check if Asset Obj
	if (assetCheck == nil){ // if no error puting into asset struct, then obj
		StaffIdList = strconv.Itoa(asset.ID)
	}

	//Hardcoded variables
	MessagePriority := "1" //high
	MessageOptions := "00000110" // Two-Tone
	MessageId := "1"
	MessageExpirationSeconds := "0" // never

	// Create the request
	url := "http://" + IP + "/XpertRestApi/api/Device/DisplayPriorityMessageOnTagNotif?" +
		"StaffIdList=" + StaffIdList + "&" +
		"Message=" + FilteredMessage + "&" +
		"MessageId=" + MessageId + "&" +
		"MessagePriority=" + MessagePriority + "&" +
		"MessageOptions=" + MessageOptions + "&" +
		"MessageExpirationSeconds=" + MessageExpirationSeconds + "&" +
		"CustomerId=" + CustomerId
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err.Error()
	}

	req.Header.Add("Content-Type", "application/json")

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err.Error()
	}
	defer resp.Body.Close()

	returnedStatus := "false"
	if(resp.Status == "200 OK"){
		returnedStatus = "true"
	}

	return returnedStatus
}