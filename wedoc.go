package wecomapi

// 文档 API 公开方法。生成的 exec 方法见 wedoc.md.go。

// WedocCreateDoc 新建文档
func (c *App) WedocCreateDoc(req WedocCreateDocReq) (*WedocCreateDocResp, error) {
	resp, err := c.execWedocCreateDoc(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocRenameDoc 重命名文档
func (c *App) WedocRenameDoc(req WedocRenameDocReq) error {
	_, err := c.execWedocRenameDoc(reqWedoc{Body: req})
	return err
}

// WedocDeleteDoc 删除文档
func (c *App) WedocDeleteDoc(req WedocDeleteDocReq) error {
	_, err := c.execWedocDeleteDoc(reqWedoc{Body: req})
	return err
}

// WedocGetDocBaseInfo 获取文档基础信息
func (c *App) WedocGetDocBaseInfo(req WedocGetDocBaseInfoReq) (*WedocGetDocBaseInfoResp, error) {
	resp, err := c.execWedocGetDocBaseInfo(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocShareDoc 分享文档
func (c *App) WedocShareDoc(req WedocShareDocReq) (*WedocShareDocResp, error) {
	resp, err := c.execWedocShareDoc(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocBatchUpdateDocument 编辑文档内容
func (c *App) WedocBatchUpdateDocument(req WedocBatchUpdateDocumentReq) error {
	_, err := c.execWedocBatchUpdateDocument(reqWedoc{Body: req})
	return err
}

// WedocGetDocument 获取文档数据
func (c *App) WedocGetDocument(req WedocGetDocumentReq) (*WedocGetDocumentResp, error) {
	resp, err := c.execWedocGetDocument(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocBatchUpdateSpreadsheet 编辑表格内容
func (c *App) WedocBatchUpdateSpreadsheet(req WedocBatchUpdateSpreadsheetReq) (*WedocBatchUpdateSpreadsheetResp, error) {
	resp, err := c.execWedocBatchUpdateSpreadsheet(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetSheetProperties 获取表格行列信息
func (c *App) WedocGetSheetProperties(req WedocGetSheetPropertiesReq) (*WedocGetSheetPropertiesResp, error) {
	resp, err := c.execWedocGetSheetProperties(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetSheetRangeData 获取表格数据
func (c *App) WedocGetSheetRangeData(req WedocGetSheetRangeDataReq) (*WedocGetSheetRangeDataResp, error) {
	resp, err := c.execWedocGetSheetRangeData(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddSheet 添加子表
func (c *App) WedocAddSheet(req WedocAddSheetReq) (*WedocAddSheetResp, error) {
	resp, err := c.execWedocAddSheet(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeleteSheet 删除子表
func (c *App) WedocDeleteSheet(req WedocDeleteSheetReq) error {
	_, err := c.execWedocDeleteSheet(reqWedoc{Body: req})
	return err
}

// WedocUpdateSheet 更新子表
func (c *App) WedocUpdateSheet(req WedocUpdateSheetReq) error {
	_, err := c.execWedocUpdateSheet(reqWedoc{Body: req})
	return err
}

// WedocGetSheet 查询子表
func (c *App) WedocGetSheet(req WedocGetSheetReq) (*WedocGetSheetResp, error) {
	resp, err := c.execWedocGetSheet(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddView 添加视图
func (c *App) WedocAddView(req WedocAddViewReq) (*WedocAddViewResp, error) {
	resp, err := c.execWedocAddView(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeleteViews 删除视图
func (c *App) WedocDeleteViews(req WedocDeleteViewsReq) error {
	_, err := c.execWedocDeleteViews(reqWedoc{Body: req})
	return err
}

// WedocUpdateView 更新视图
func (c *App) WedocUpdateView(req WedocUpdateViewReq) (*WedocUpdateViewResp, error) {
	resp, err := c.execWedocUpdateView(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetViews 查询视图
func (c *App) WedocGetViews(req WedocGetViewsReq) (*WedocGetViewsResp, error) {
	resp, err := c.execWedocGetViews(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddFields 添加字段
func (c *App) WedocAddFields(req WedocAddFieldsReq) (*WedocAddFieldsResp, error) {
	resp, err := c.execWedocAddFields(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeleteFields 删除字段
func (c *App) WedocDeleteFields(req WedocDeleteFieldsReq) error {
	_, err := c.execWedocDeleteFields(reqWedoc{Body: req})
	return err
}

// WedocUpdateFields 更新字段
func (c *App) WedocUpdateFields(req WedocUpdateFieldsReq) (*WedocUpdateFieldsResp, error) {
	resp, err := c.execWedocUpdateFields(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetFields 查询字段
func (c *App) WedocGetFields(req WedocGetFieldsReq) (*WedocGetFieldsResp, error) {
	resp, err := c.execWedocGetFields(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddRecords 添加记录
func (c *App) WedocAddRecords(req WedocAddRecordsReq) (*WedocAddRecordsResp, error) {
	resp, err := c.execWedocAddRecords(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeleteRecords 删除记录
func (c *App) WedocDeleteRecords(req WedocDeleteRecordsReq) error {
	_, err := c.execWedocDeleteRecords(reqWedoc{Body: req})
	return err
}

// WedocUpdateRecords 更新记录
func (c *App) WedocUpdateRecords(req WedocUpdateRecordsReq) (*WedocUpdateRecordsResp, error) {
	resp, err := c.execWedocUpdateRecords(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetRecords 查询记录
func (c *App) WedocGetRecords(req WedocGetRecordsReq) (*WedocGetRecordsResp, error) {
	resp, err := c.execWedocGetRecords(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddFieldGroup 添加编组
func (c *App) WedocAddFieldGroup(req WedocAddFieldGroupReq) (*WedocAddFieldGroupResp, error) {
	resp, err := c.execWedocAddFieldGroup(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeleteFieldGroups 删除编组
func (c *App) WedocDeleteFieldGroups(req WedocDeleteFieldGroupsReq) error {
	_, err := c.execWedocDeleteFieldGroups(reqWedoc{Body: req})
	return err
}

// WedocUpdateFieldGroup 更新编组
func (c *App) WedocUpdateFieldGroup(req WedocUpdateFieldGroupReq) (*WedocUpdateFieldGroupResp, error) {
	resp, err := c.execWedocUpdateFieldGroup(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetFieldGroups 获取编组
func (c *App) WedocGetFieldGroups(req WedocGetFieldGroupsReq) (*WedocGetFieldGroupsResp, error) {
	resp, err := c.execWedocGetFieldGroups(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddPage 添加页面
func (c *App) WedocAddPage(req WedocAddPageReq) (*WedocAddPageResp, error) {
	resp, err := c.execWedocAddPage(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocUpdatePage 更新页面
func (c *App) WedocUpdatePage(req WedocUpdatePageReq) (*WedocUpdatePageResp, error) {
	resp, err := c.execWedocUpdatePage(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeletePage 删除页面
func (c *App) WedocDeletePage(req WedocDeletePageReq) error {
	_, err := c.execWedocDeletePage(reqWedoc{Body: req})
	return err
}

// WedocGetPageHierarchy 获取页面结构
func (c *App) WedocGetPageHierarchy(req WedocGetPageHierarchyReq) (*WedocGetPageHierarchyResp, error) {
	resp, err := c.execWedocGetPageHierarchy(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddBlocks 添加内容块
func (c *App) WedocAddBlocks(req WedocAddBlocksReq) (*WedocAddBlocksResp, error) {
	resp, err := c.execWedocAddBlocks(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocUpdateBlocks 更新内容块
func (c *App) WedocUpdateBlocks(req WedocUpdateBlocksReq) (*WedocUpdateBlocksResp, error) {
	resp, err := c.execWedocUpdateBlocks(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeleteBlocks 删除内容块
func (c *App) WedocDeleteBlocks(req WedocDeleteBlocksReq) error {
	_, err := c.execWedocDeleteBlocks(reqWedoc{Body: req})
	return err
}

// WedocGetBlockList 获取内容块列表
func (c *App) WedocGetBlockList(req WedocGetBlockListReq) (*WedocGetBlockListResp, error) {
	resp, err := c.execWedocGetBlockList(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocExportTask 提交导出任务
func (c *App) WedocExportTask(req WedocExportTaskReq) (*WedocExportTaskResp, error) {
	resp, err := c.execWedocExportTask(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetExportResult 查询任务结果
func (c *App) WedocGetExportResult(req WedocGetExportResultReq) (*WedocGetExportResultResp, error) {
	resp, err := c.execWedocGetExportResult(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetSmartsheetInfo 获取数据源
func (c *App) WedocGetSmartsheetInfo(req WedocGetSmartsheetInfoReq) (*WedocGetSmartsheetInfoResp, error) {
	resp, err := c.execWedocGetSmartsheetInfo(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddSmartsheet 添加数据表
func (c *App) WedocAddSmartsheet(req WedocAddSmartsheetReq) (*WedocAddSmartsheetResp, error) {
	resp, err := c.execWedocAddSmartsheet(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeleteSmartsheet 删除数据表
func (c *App) WedocDeleteSmartsheet(req WedocDeleteSmartsheetReq) error {
	_, err := c.execWedocDeleteSmartsheet(reqWedoc{Body: req})
	return err
}

// WedocUpdateSmartsheet 更新数据表
func (c *App) WedocUpdateSmartsheet(req WedocUpdateSmartsheetReq) (*WedocUpdateSmartsheetResp, error) {
	resp, err := c.execWedocUpdateSmartsheet(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocPublishSmartDoc 发布智能文档
func (c *App) WedocPublishSmartDoc(req WedocPublishSmartDocReq) (*WedocPublishSmartDocResp, error) {
	resp, err := c.execWedocPublishSmartDoc(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocCancelPublishSmartDoc 取消发布智能文档
func (c *App) WedocCancelPublishSmartDoc(req WedocCancelPublishSmartDocReq) error {
	_, err := c.execWedocCancelPublishSmartDoc(reqWedoc{Body: req})
	return err
}

// WedocUpdatePublishSetting 修改发布页可查看范围
func (c *App) WedocUpdatePublishSetting(req WedocUpdatePublishSettingReq) error {
	_, err := c.execWedocUpdatePublishSetting(reqWedoc{Body: req})
	return err
}

// WedocGetDocAuth 获取文档权限信息
func (c *App) WedocGetDocAuth(req WedocGetDocAuthReq) (*WedocGetDocAuthResp, error) {
	resp, err := c.execWedocGetDocAuth(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocModifyDocJoinRule 修改文档加入规则
func (c *App) WedocModifyDocJoinRule(req WedocModifyDocJoinRuleReq) error {
	_, err := c.execWedocModifyDocJoinRule(reqWedoc{Body: req})
	return err
}

// WedocModifyDocMember 修改文档成员与权限
func (c *App) WedocModifyDocMember(req WedocModifyDocMemberReq) error {
	_, err := c.execWedocModifyDocMember(reqWedoc{Body: req})
	return err
}

// WedocModifyDocSafetySetting 修改文档安全设置
func (c *App) WedocModifyDocSafetySetting(req WedocModifyDocSafetySettingReq) error {
	_, err := c.execWedocModifyDocSafetySetting(reqWedoc{Body: req})
	return err
}

// WedocGetSheetPriv 查询智能表格子表权限
func (c *App) WedocGetSheetPriv(req WedocGetSheetPrivReq) (*WedocGetSheetPrivResp, error) {
	resp, err := c.execWedocGetSheetPriv(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocUpdateSheetPriv 更新智能表格子表权限
func (c *App) WedocUpdateSheetPriv(req WedocUpdateSheetPrivReq) error {
	_, err := c.execWedocUpdateSheetPriv(reqWedoc{Body: req})
	return err
}

// WedocCreateSheetPrivRule 新增智能表格指定成员额外权限
func (c *App) WedocCreateSheetPrivRule(req WedocCreateSheetPrivRuleReq) (*WedocCreateSheetPrivRuleResp, error) {
	resp, err := c.execWedocCreateSheetPrivRule(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocModifySheetPrivRuleMember 更新智能表格指定成员额外权限
func (c *App) WedocModifySheetPrivRuleMember(req WedocModifySheetPrivRuleMemberReq) error {
	_, err := c.execWedocModifySheetPrivRuleMember(reqWedoc{Body: req})
	return err
}

// WedocDeleteSheetPrivRule 删除智能表格指定成员额外权限
func (c *App) WedocDeleteSheetPrivRule(req WedocDeleteSheetPrivRuleReq) error {
	_, err := c.execWedocDeleteSheetPrivRule(reqWedoc{Body: req})
	return err
}

// WedocCreateForm 创建收集表
func (c *App) WedocCreateForm(req WedocCreateFormReq) (*WedocCreateFormResp, error) {
	resp, err := c.execWedocCreateForm(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocModifyForm 编辑收集表
func (c *App) WedocModifyForm(req WedocModifyFormReq) error {
	_, err := c.execWedocModifyForm(reqWedoc{Body: req})
	return err
}

// WedocGetFormInfo 获取收集表信息
func (c *App) WedocGetFormInfo(req WedocGetFormInfoReq) (*WedocGetFormInfoResp, error) {
	resp, err := c.execWedocGetFormInfo(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetFormStatistic 收集表的统计信息查询
func (c *App) WedocGetFormStatistic(req WedocGetFormStatisticReq) (*WedocGetFormStatisticResp, error) {
	resp, err := c.execWedocGetFormStatistic(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocGetFormAnswer 读取收集表答案
func (c *App) WedocGetFormAnswer(req WedocGetFormAnswerReq) (*WedocGetFormAnswerResp, error) {
	resp, err := c.execWedocGetFormAnswer(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocWebhookAddRecords 添加记录
func (c *App) WedocWebhookAddRecords(key string, req WedocWebhookAddRecordsReq) (*WedocWebhookAddRecordsResp, error) {
	resp, err := c.execWedocWebhookAddRecords(reqWedocWebhook{Key: key, Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocWebhookUpdateRecords 更新记录
func (c *App) WedocWebhookUpdateRecords(key string, req WedocWebhookUpdateRecordsReq) (*WedocWebhookUpdateRecordsResp, error) {
	resp, err := c.execWedocWebhookUpdateRecords(reqWedocWebhook{Key: key, Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocAddVipAccounts 分配高级功能账号
func (c *App) WedocAddVipAccounts(req WedocAddVipAccountsReq) (*WedocAddVipAccountsResp, error) {
	resp, err := c.execWedocAddVipAccounts(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocDeleteVipAccounts 取消高级功能账号
func (c *App) WedocDeleteVipAccounts(req WedocDeleteVipAccountsReq) (*WedocDeleteVipAccountsResp, error) {
	resp, err := c.execWedocDeleteVipAccounts(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocListVipAccounts 获取高级功能账号列表
func (c *App) WedocListVipAccounts(req WedocListVipAccountsReq) (*WedocListVipAccountsResp, error) {
	resp, err := c.execWedocListVipAccounts(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocUploadDocImage 上传文档图片
func (c *App) WedocUploadDocImage(req WedocUploadDocImageReq) (*WedocUploadDocImageResp, error) {
	resp, err := c.execWedocUploadDocImage(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocUploadDocFile 上传文档文件
func (c *App) WedocUploadDocFile(req WedocUploadDocFileReq) (*WedocUploadDocFileResp, error) {
	resp, err := c.execWedocUploadDocFile(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocInitFileUpload 分块上传初始化
func (c *App) WedocInitFileUpload(req WedocInitFileUploadReq) (*WedocInitFileUploadResp, error) {
	resp, err := c.execWedocInitFileUpload(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// WedocUploadFilePart 分块上传文件
func (c *App) WedocUploadFilePart(req WedocUploadFilePartReq) error {
	_, err := c.execWedocUploadFilePart(reqWedoc{Body: req})
	return err
}

// WedocFinishFileUpload 分块上传完成
func (c *App) WedocFinishFileUpload(req WedocFinishFileUploadReq) (*WedocFinishFileUploadResp, error) {
	resp, err := c.execWedocFinishFileUpload(reqWedoc{Body: req})
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
