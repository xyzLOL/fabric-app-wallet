package task

import "testing"

func TestAccountCreateTask_Create(t *testing.T) {
	var task Task = new(AccountCreateTask)
	var accountuuid = "d769e2dc-2359-4efe-866a-38c4b588fbc5"
	taskuuid, err := task.Create(accountuuid, TASK_TYPE_CREATE_ACCOUNT, TASK_STATE_INIT)
	if err != nil {
		t.Errorf("failed to create task for accountcreate event: %v", err)
	}
	taskLogger.Debugf("successed in creating task %s for accountcreate event", taskuuid)

}

func TestAccountTransferTask_Create(t *testing.T){

}
