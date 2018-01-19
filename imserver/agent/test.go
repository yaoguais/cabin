package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/VolantMQ/volantmq/configuration"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var (
	logger       *zap.Logger
	httpEndPoint = "127.0.0.1:50050"
)

func init() {
	logger = configuration.GetLogger().Named("agent")
}

func main() {
	testBroker()
	testUser()
	testGroup()
	testPublish()
	fmt.Println("PASS")
}

func request(path string, body []byte) []byte {
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s%s", httpEndPoint, path), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		logger.Fatal("Request", zap.Error(err))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Fatal("Response", zap.Error(err))
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal("Response body", zap.Error(err))
	}
	fmt.Printf("Reqeust %s %s response %s\n", path, string(body), string(respBody))
	return respBody
}

func testBroker() {
	// Update
	params := map[string]interface{}{
		"external": "127.0.0.1:8883",
		"internal": "127.0.0.1:1883",
		"conn":     0,
		"max_conn": -1,
	}
	body, _ := jsoniter.Marshal(params)
	data := request("/v1/broker/update", body)
	if !assert.ObjectsAreEqual("127.0.0.1:8883", jsoniter.Get(data, "external").ToString()) {
		logger.Fatal("Assert")
	}
	// List
	params = map[string]interface{}{}
	body, _ = jsoniter.Marshal(params)
	data = request("/v1/broker/list", body)
	if !assert.ObjectsAreEqual(1, jsoniter.Get(data, "brokers").Size()) {
		logger.Fatal("Assert")
	}
}

func testUser() {
	// Add
	id := "1"
	username := "user1"
	password := "testpassword"
	params := map[string]interface{}{
		"id":       id,
		"username": username,
		"password": password,
	}
	body, _ := jsoniter.Marshal(params)
	data := request("/v1/user/add", body)
	if !assert.ObjectsAreEqual(id, jsoniter.Get(data, "id").ToString()) {
		logger.Fatal("Assert")
	}
	// Get
	params = map[string]interface{}{
		"username": username,
	}
	body, _ = jsoniter.Marshal(params)
	data = request("/v1/user/get", body)
	if !assert.ObjectsAreEqual(id, jsoniter.Get(data, "id").ToString()) ||
		!assert.ObjectsAreEqual(username, jsoniter.Get(data, "username").ToString()) ||
		!assert.ObjectsAreEqual(password, jsoniter.Get(data, "password").ToString()) {
		logger.Fatal("Assert")
	}
	// Del
	params = map[string]interface{}{
		"username": username,
	}
	body, _ = jsoniter.Marshal(params)
	data = request("/v1/user/del", body)
	if !assert.ObjectsAreEqual(username, jsoniter.Get(data, "username").ToString()) {
		logger.Fatal("Assert")
	}
}

func testGroup() {
	// Add
	groupID := "1"
	members := []string{"1", "2", "3", "4", "5", "6", "7"}
	params := map[string]interface{}{
		"group_id": groupID,
		"members":  members,
	}
	body, _ := jsoniter.Marshal(params)
	data := request("/v1/group/member/add", body)
	if !assert.ObjectsAreEqual(true, jsoniter.Get(data, "ok").ToBool()) {
		logger.Fatal("Assert")
	}
	// List
	params = map[string]interface{}{
		"group_id": groupID,
	}
	body, _ = jsoniter.Marshal(params)
	data = request("/v1/group/member/list", body)
	if !assert.ObjectsAreEqual(len(members), jsoniter.Get(data, "members").Size()) {
		logger.Fatal("Assert")
	}
	// Del
	params = map[string]interface{}{
		"group_id": groupID,
		"members":  members,
	}
	body, _ = jsoniter.Marshal(params)
	data = request("/v1/group/member/del", body)
	if !assert.ObjectsAreEqual(true, jsoniter.Get(data, "ok").ToBool()) {
		logger.Fatal("Assert")
	}
}

func testPublish() {
	for i := 1; i < 10; i++ {
		params := map[string]interface{}{
			"topic":    "u/1",
			"payload":  "testpayload",
			"qos":      1,
			"retained": false,
		}
		body, _ := jsoniter.Marshal(params)
		data := request("/v1/publish", body)
		if !assert.ObjectsAreEqual(true, jsoniter.Get(data, "ok").ToBool()) {
			logger.Fatal("Assert")
		}
	}
}
