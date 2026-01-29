package deptService

import (
	model "gin-web-admin/app/models"
)

type DeptTreeNode struct {
	ID       int            `json:"id"`
	ParentID int            `json:"parentId"`
	DeptName string         `json:"deptName"`
	OrderNum int            `json:"orderNum"`
	Leader   string         `json:"leader"`
	Phone    string         `json:"phone"`
	Email    string         `json:"email"`
	Status   int            `json:"status"`
	Remark   string         `json:"remark"`
	Children []*DeptTreeNode `json:"children,omitempty"`
}

type CreateDeptStruct struct {
	ParentID int    `json:"parentId" binding:"gte=0"`
	DeptName string `json:"deptName" binding:"required"`
	OrderNum int    `json:"orderNum"`
	Leader   string `json:"leader"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	Remark   string `json:"remark"`
}

type UpdateDeptStruct struct {
	ParentID int    `json:"parentId"`
	DeptName string `json:"deptName"`
	OrderNum int    `json:"orderNum"`
	Leader   string `json:"leader"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	Remark   string `json:"remark"`
}

func GetDeptList() ([]*model.Dept, error) {
	return model.GetAllDepts(nil)
}

func GetDeptTree() ([]*DeptTreeNode, error) {
	list, err := model.GetAllDepts(nil)
	if err != nil {
		return nil, err
	}

	nodes := make([]*DeptTreeNode, 0, len(list))
	idMap := make(map[int]*DeptTreeNode, len(list))
	for _, d := range list {
		n := &DeptTreeNode{
			ID:       d.ID,
			ParentID: d.ParentID,
			DeptName: d.DeptName,
			OrderNum: d.OrderNum,
			Leader:   d.Leader,
			Phone:    d.Phone,
			Email:    d.Email,
			Status:   d.Status,
			Remark:   d.Remark,
		}
		nodes = append(nodes, n)
		idMap[n.ID] = n
	}

	roots := make([]*DeptTreeNode, 0)
	for _, n := range nodes {
		if n.ParentID == 0 {
			roots = append(roots, n)
			continue
		}
		p := idMap[n.ParentID]
		if p == nil {
			// 父节点不存在时，视作根节点
			roots = append(roots, n)
			continue
		}
		p.Children = append(p.Children, n)
	}

	return roots, nil
}

func CreateDept(payload CreateDeptStruct) error {
	dept := model.Dept{
		ParentID: payload.ParentID,
		DeptName: payload.DeptName,
		OrderNum: payload.OrderNum,
		Leader:   payload.Leader,
		Phone:    payload.Phone,
		Email:    payload.Email,
		Status:   payload.Status,
		Remark:   payload.Remark,
	}
	return model.CreateDept(dept)
}

func UpdateDept(id int, payload UpdateDeptStruct) error {
	data := map[string]any{}
	if payload.ParentID != 0 {
		data["parent_id"] = payload.ParentID
	}
	if payload.DeptName != "" {
		data["dept_name"] = payload.DeptName
	}
	data["order_num"] = payload.OrderNum
	data["leader"] = payload.Leader
	data["phone"] = payload.Phone
	data["email"] = payload.Email
	data["status"] = payload.Status
	data["remark"] = payload.Remark
	return model.UpdateDept(id, data)
}

func DeleteDept(id int) error {
	return model.DeleteDept(id)
}
