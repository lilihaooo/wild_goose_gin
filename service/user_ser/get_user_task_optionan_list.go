package user_ser

import (
	"errors"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
)

func (s UserService) GetUserTaskOptionalList(taskID uint) (users []models.User, newErr error) {
	// 根据任务ID获取任务
	var taskModel models.Task
	taskModel.ID = taskID
	task, err := taskModel.TakeOneRecordByID()
	if err != nil {
		newErr = errors.New("任务不存在")
		global.Logrus.Error(err)
		return nil, newErr
	}
	// 手册信息
	manual := task.Component.Manual
	// 证书信息
	certificates := task.Certificates
	if len(*certificates) == 0 {
		// 返回具有该手册的所有users
		users, err = getHasManualAllUserList(manual.ID)
		if err != nil {
			global.Logrus.Error(err)
		}
		return
	} else {
		// 获取具有该手册同时具备该授权的用户
		var certificateIDs []uint
		for _, one := range *certificates {
			certificateIDs = append(certificateIDs, one.ID)
		}
		users, err = getHasManualAndCertificatesUserList(manual.ID, certificateIDs)
		return
	}
}

func getHasManualAllUserList(manualID uint) (users []models.User, err error) {
	var umModel models.UserManual
	umModel.ManualID = manualID
	ums, err := umModel.GetRecordsByManualID()
	if err != nil {
		return
	}
	if len(ums) > 0 {
		for _, one := range ums {
			users = append(users, one.User)
		}
	}
	return
}

func getHasManualAndCertificatesUserList(manualID uint, certificateIDs []uint) (users []models.User, err error) {
	certificateMap := make(map[uint]struct{})
	for _, id := range certificateIDs {
		certificateMap[id] = struct{}{}
	}

	var umModel models.UserManual
	umModel.ManualID = manualID
	ums, err := umModel.GetRecordsByManualIDWithCertificate()
	if err != nil {
		return
	}
	// 过滤数据
	if len(ums) > 0 {
		// 输出JSON格式的字符串到控制台
		for _, um := range ums {
			if hasAllCertificates(*um.UserManualCertificates, certificateMap) {
				users = append(users, um.User)
			}
		}
	}
	return
}

func hasAllCertificates(umcs []models.UserManualCertificate, certificateMap map[uint]struct{}) bool {
	// 复制目标集合
	newMap := make(map[uint]struct{})
	for k, v := range certificateMap {
		newMap[k] = v
	}

	for _, umc := range umcs {
		if _, exists := certificateMap[umc.CertificateID]; exists && umc.State.String() {
			delete(newMap, umc.CertificateID)
		}
	}
	return len(newMap) == 0
}
