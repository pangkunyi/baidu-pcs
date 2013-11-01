package pcs

const (
	LIST_ORDER_BY_TIME = "time" //修改时间
	LIST_ORDER_BY_NAME = "name" //文件名
	LIST_ORDER_BY_SIZE = "size" //大小，注意目录无大小
	LIST_ORDER_ASC     = "asc"  //升序
	LIST_ORDER_DESC    = "desc" //降序
)

type ListReq struct {
	BasicFileReq
	By    string `pcs:"by"`    //optional, 排序字段，缺省根据文件类型排序：
	Order string `pcs:"order"` //optional, 缺省采用降序排序。
	Limit string `pcs:"limit"` //optional, 返回条目控制，参数格式为：n1-n2。 返回结果集的[n1, n2)之间的条目，缺省返回所有条目；n1从0开始。
}

type ListResp struct {
	Files     []PcsFileInfo `json:"list"`
	RequestId uint64        `json:"request_id"`
}

type PcsFileInfo struct {
	FsId  uint64 `json:"fs_id"` // optional, 文件或目录在PCS的临时唯一标识id。
	Path  string `json:"path"`  // optional, 文件或目录的绝对路径。
	CTime uint   `json:"ctime"` // optional, 文件或目录的创建时间。
	MTime uint   `json:"mtime"` // optional, 文件或目录的最后修改时间。
	Md5   string `json:"md5"`   // optional, 文件的md5值。
	Size  uint64 `json:"size"`  // optional, 文件大小（byte）。
	IsDir uint   `json:"isdir"` // optional, 是否是目录的标识符： “0”为文件,  “1”为目录
}

func NewListReq() *ListReq {
	req := &ListReq{}
	req.Method = "list"
	return req
}

func List(req *ListReq) (resp *ListResp, err error) {
	data := parseUrlValues(req)
	resp = new(ListResp)
	err = getJson(PCS_URL, data, resp)
	return
}
