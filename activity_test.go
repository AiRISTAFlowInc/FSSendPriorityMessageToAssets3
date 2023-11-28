package SendPriorityMessageToAssets

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	input := &Input{IP: "3.212.201.170:802", CustomerId: "2047", Username: "afadmin", Password: "admin", StaffIdList: "{\"ItemId\":400001909}", Message: "Test Prio Msg OBJ"}
	// StaffIdList "9064" OR "9064,37685" OR "{\"ItemId\":400001909}"
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{} 
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)

	assert.Equal(t, "true", output.Status)
}
