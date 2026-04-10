package deptService

import (
	model "gin-web-admin/app/models"
	"gin-web-admin/internal/data"
)

type DeptTreeNode struct {
	ID       int             `json:"id"`
	ParentID int             `json:"parentId"`
	DeptName string          `json:"deptName"`
	OrderNum int             `json:"orderNum"`
	Leader   string          `json:"leader"`
	Phone    string          `json:"phone"`
	Email    string          `json:"email"`
	Status   int             `json:"status"`
	Remark   string          `json:"remark"`
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

type Service struct {
	store *data.Store
}

var defaultService *Service

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func SetDefaultService(s *Service) {
	defaultService = s
}

func serviceInstance() *Service {
	if defaultService == nil {
		panic("dept service not initialized")
	}
	return defaultService
}

func GetDeptList() ([]*model.Dept, error) {
	return serviceInstance().GetDeptList()
}

func (s *Service) GetDeptList() ([]*model.Dept, error) {
	return model.GetAllDepts(nil)
}

func GetDeptTree() ([]*DeptTreeNode, error) {
	return serviceInstance().GetDeptTree()
}

func (s *Service) GetDeptTree() ([]*DeptTreeNode, error) {
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
			roots = append(roots, n)
			continue
		}
		p.Children = append(p.Children, n)
	}

	return roots, nil
}

func CreateDept(payload CreateDeptStruct) error {
	return serviceInstance().CreateDept(payload)
}

func (s *Service) CreateDept(payload CreateDeptStruct) error {
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
	return serviceInstance().UpdateDept(id, payload)
}

func (s *Service) UpdateDept(id int, payload UpdateDeptStruct) error {
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
	return serviceInstance().DeleteDept(id)
}

func (s *Service) DeleteDept(id int) error {
	if err := model.DeleteDept(id); err != nil {
		return err
	}
	return nil
}
