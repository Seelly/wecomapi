# imports

```go
import "encoding/json"
```

# 文档 API 模型

## Models

### `WedocAddBlocksReq` 添加内容块请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`PageID`|`page_id`|`string`| 
`Blocks`|`blocks`|`[]WedocBlock`| 

### `WedocAddBlocksResp` 添加内容块响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Blocks`|`blocks`|`[]WedocBlock`| 

### `WedocAddField` AddField

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldTitle`|`field_title`|`string`| 
`FieldType`|`field_type`|`string`| 
`PropertyNumber`|`property_number,omitempty`|`*WedocNumberFieldProperty`| 
`PropertyCheckbox`|`property_checkbox,omitempty`|`*WedocCheckboxFieldProperty`| 
`PropertyDateTime`|`property_date_time,omitempty`|`*WedocDateTimeFieldProperty`| 
`PropertyAttachment`|`property_attachment,omitempty`|`*WedocAttachmentFieldProperty`| 
`PropertyUser`|`property_user,omitempty`|`*WedocUserFieldProperty`| 
`PropertyURL`|`property_url,omitempty`|`*WedocUrlFieldProperty`| 
`PropertySelect`|`property_select,omitempty`|`*WedocSelectFieldProperty`| 
`PropertyCreatedTime`|`property_created_time,omitempty`|`*WedocCreatedTimeFieldProperty`| 
`PropertyModifiedTime`|`property_modified_time,omitempty`|`*WedocModifiedTimeFieldProperty`| 
`PropertyProgress`|`property_progress,omitempty`|`*WedocProgressFieldProperty`| 
`PropertySingleSelect`|`property_single_select,omitempty`|`*WedocSingleSelectFieldProperty`| 
`PropertyReference`|`property_reference,omitempty`|`*WedocReferenceFieldProperty`| 
`PropertyLocation`|`property_location,omitempty`|`*WedocLocationFieldProperty`| 
`PropertyAutoNumber`|`property_auto_number,omitempty`|`*WedocAutoNumberFieldProperty`| 
`PropertyCurrency`|`property_currency,omitempty`|`*WedocCurrencyFieldProperty`| 
`PropertyWwGroup`|`property_ww_group,omitempty`|`*WedocWwGroupFieldProperty`| 
`PropertyPercentage`|`property_percentage,omitempty`|`*WedocPercentageFieldProperty`| 
`PropertyBarcode`|`property_barcode,omitempty`|`*WedocBarcodeFieldProperty`| 

### `WedocAddFieldGroupReq` 添加编组

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`Name`|`name`|`string`| 
`Children`|`children,omitempty`|`[]WedocFieldGroupChild`| 

### `WedocAddFieldGroupResp` 添加编组

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldGroup`|`field_group`|`WedocFieldGroup`| 

### `WedocAddFieldsReq` 添加字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`Fields`|`fields`|`[]WedocAddField`| 

### `WedocAddFieldsResp` 添加字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`Fields`|`fields`|`[]WedocField`| 

### `WedocAddPageReq` 添加页面请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Info`|`info`|`WedocPageInfo`| 

### `WedocAddPageResp` 添加页面响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Info`|`info`|`WedocPageInfo`| 

### `WedocAddRecord` AddRecord

Name|JSON|Type|Doc
:---|:---|:---|:--
`Values`|`values`|`map[string]any`| 

### `WedocAddRecordsReq` 添加记录

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`KeyType`|`key_type,omitempty`|`string`| 
`Records`|`records`|`[]WedocAddRecord`| 

### `WedocAddRecordsResp` 添加记录

Name|JSON|Type|Doc
:---|:---|:---|:--
`Records`|`records`|`[]WedocCommonRecord`| 

### `WedocAddSheetReq` 添加子表

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Properties`|`properties,omitempty`|`*WedocSheetProperties`| 

### `WedocAddSheetResp` 添加子表

Name|JSON|Type|Doc
:---|:---|:---|:--
`Properties`|`properties`|`WedocSheetProperties`| 

### `WedocAddSheetResponse` AddSheetResponse

Name|JSON|Type|Doc
:---|:---|:---|:--
`Properties`|`properties`|`WedocProperties`| 

### `WedocAddSmartsheetReq` 添加数据表请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Info`|`info`|`WedocSmartDocSmartsheetInfo`| 

### `WedocAddSmartsheetResp` 添加数据表响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Info`|`info`|`WedocSmartDocSmartsheetInfo`| 

### `WedocAddViewReq` 添加视图

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`ViewTitle`|`view_title`|`string`| 
`ViewType`|`view_type`|`string`| 
`PropertyGantt`|`property_gantt,omitempty`|`*WedocGanttViewProperty`| 
`PropertyCalendar`|`property_calendar,omitempty`|`*WedocCalendarViewProperty`| 

### `WedocAddViewResp` 添加视图

Name|JSON|Type|Doc
:---|:---|:---|:--
`View`|`view`|`WedocView`| 

### `WedocAddVipAccountsReq` 分配高级功能账号请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserIDList`|`userid_list`|`[]string`| 

### `WedocAddVipAccountsResp` 分配高级功能账号响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`SuccUserIDList`|`succ_userid_list`|`[]string`| 
`FailUserIDList`|`fail_userid_list`|`[]string`| 

### `WedocAttachmentFieldProperty` AttachmentFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`DisplayMode`|`display_mode`|`string`| 

### `WedocAutoNumberFieldProperty` AutoNumberFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`| 
`Rules`|`rules`|`[]WedocNumberRule`| 
`ReformatExistingRecord`|`reformat_existing_record`|`bool`| 

### `WedocBarcodeFieldProperty` BarcodeFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`MobileScanOnly`|`mobile_scan_only`|`bool`| 

### `WedocBatchUpdateDocumentReq` 编辑文档内容请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Version`|`version,omitempty`|`uint32`| 
`Requests`|`requests`|`[]WedocDocumentUpdateRequest`| 

### `WedocBatchUpdateSpreadsheetReq` 编辑表格内容请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Requests`|`requests`|`[]WedocSpreadsheetUpdateRequest`| 

### `WedocBatchUpdateSpreadsheetResp` 编辑表格内容响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Data`|`data`|`WedocSpreadsheetUpdateData`| 

### `WedocBlock` 智能文档内容块

Name|JSON|Type|Doc
:---|:---|:---|:--
`ID`|`id,omitempty`|`string`| 
`Type`|`type,omitempty`|`string`| 
`Title`|`title,omitempty`|`string`| 
`Content`|`content,omitempty`|`string`| 
`ParentID`|`parent_id,omitempty`|`string`| 
`AfterID`|`after_id,omitempty`|`string`| 
`Children`|`children,omitempty`|`[]string`| 
`Props`|`props,omitempty`|`*WedocBlockProps`| 

### `WedocBlockProps` 内容块属性，不同类型的 Block 只使用其中一个字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`TextProps`|`text_props,omitempty`|`*WedocTextBlockProps`| 
`ImageProps`|`image_props,omitempty`|`*WedocImageBlockProps`| 
`FileProps`|`file_props,omitempty`|`*WedocFileBlockProps`| 
`LinkProps`|`link_props,omitempty`|`*WedocLinkBlockProps`| 
`TableProps`|`table_props,omitempty`|`*WedocTableBlockProps`| 
`ViewProps`|`view_props,omitempty`|`*WedocSmartSheetViewBlockProps`| 
`ColumnProps`|`column_props,omitempty`|`*WedocColumnListProps`| 
`HighlightProps`|`highlight_props,omitempty`|`*WedocHighlightBlockProps`| 
`CodeProps`|`code_props,omitempty`|`*WedocCodeBlockProps`| 
`TodoProps`|`todo_props,omitempty`|`*WedocTodoBlockProps`| 

### `WedocBoolValue` 复选框条件值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Value`|`value`|`bool`| 

### `WedocCalendarViewProperty` CalendarViewProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`StartDateFieldID`|`start_date_field_id`|`string`| 
`EndDateFieldID`|`end_date_field_id`|`string`| 

### `WedocCancelPublishSmartDocReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 

### `WedocCellAttachmentValue` CellAttachmentValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`Name`|`name`|`string`| 
`FileBase64`|`file_base64`|`string`| 

### `WedocCellAutoNumberValue` CellAutoNumberValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`Seq`|`seq`|`string`| 
`Text`|`text`|`string`| 

### `WedocCellData` 单元格数据

Name|JSON|Type|Doc
:---|:---|:---|:--
`CellValue`|`cell_value,omitempty`|`*WedocCellValue`| 
`CellFormat`|`cell_format,omitempty`|`*WedocCellFormat`| 

### `WedocCellFormat` 单元格样式

Name|JSON|Type|Doc
:---|:---|:---|:--
`TextFormat`|`text_format,omitempty`|`*WedocTextFormat`| 

### `WedocCellImageValue` CellImageValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`Title`|`title`|`string`| 
`ImageBase64`|`image_base64`|`string`| 

### `WedocCellLink` 单元格超链接

Name|JSON|Type|Doc
:---|:---|:---|:--
`URL`|`url`|`string`| 
`Text`|`text`|`string`| 

### `WedocCellLocationValue` CellLocationValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`SourceType`|`source_type`|`uint32`| 
`ID`|`id`|`string`| 
`Latitude`|`latitude`|`string`| 
`Longitude`|`longitude`|`string`| 
`Title`|`title`|`string`| 

### `WedocCellReferenceValue` CellReferenceValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`RecordID`|`record_id`|`string`| 

### `WedocCellTextValue` CellTextValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`| 
`Text`|`text`|`string`| 
`Link`|`link`|`string`| 

### `WedocCellUrlValue` CellUrlValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`Text`|`text`|`string`| 
`Link`|`link`|`string`| 

### `WedocCellUserValue` CellUserValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`user_id`|`string`| 

### `WedocCellValue` 单元格内容，text/link 二选一

Name|JSON|Type|Doc
:---|:---|:---|:--
`Text`|`text,omitempty`|`string`| 
`Link`|`link,omitempty`|`*WedocCellLink`| 

### `WedocCheckboxFieldProperty` CheckboxFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`Checked`|`checked`|`bool`| 

### `WedocCodeBlockProps` 代码块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`CodeLanguage`|`code_language,omitempty`|`string`| 
`CodeWrap`|`code_wrap,omitempty`|`bool`| 

### `WedocColor` RGBA 颜色

Name|JSON|Type|Doc
:---|:---|:---|:--
`Red`|`red`|`uint32`| 
`Green`|`green`|`uint32`| 
`Blue`|`blue`|`uint32`| 
`Alpha`|`alpha`|`uint32`| 

### `WedocColumnListProps` 分栏块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`ColumnNum`|`column_num,omitempty`|`uint32`| 

### `WedocCommonRecord` CommonRecord

Name|JSON|Type|Doc
:---|:---|:---|:--
`RecordID`|`record_id`|`string`| 
`Values`|`values`|`map[string]any`| 

### `WedocCondition` Condition

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`| 
`FieldType`|`field_type`|`string`| 
`Operator`|`operator`|`string`| 
`StringValue`|`string_value,omitempty`|`*WedocStringValue`| StringValue 文本、网址、电话、邮箱、地理位置、单选、多选等列类型使用；选项列为选项 ID
`NumberValue`|`number_value,omitempty`|`*WedocNumberValue`| NumberValue 数字、进度列类型使用
`BoolValue`|`bool_value,omitempty`|`*WedocBoolValue`| BoolValue 复选框列类型使用
`UserValue`|`user_value,omitempty`|`*WedocUserValue`| UserValue 人员、创建人、最后编辑人列类型使用
`DateTimeValue`|`date_time_value,omitempty`|`*WedocFilterDataTimeValue`| DateTimeValue 日期、创建时间、最后编辑时间列类型使用

### `WedocCreateDocReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`SpaceID`|`spaceid,omitempty`|`string`| 
`FatherID`|`fatherid,omitempty`|`string`| 
`DocType`|`doc_type`|`uint32`| 
`DocName`|`doc_name`|`string`| 
`AdminUsers`|`admin_users,omitempty`|`[]string`| 

### `WedocCreateDocResp` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`URL`|`url`|`string`| 
`DocID`|`docid`|`string`| 

### `WedocCreateFormReq` 创建收集表请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`SpaceID`|`spaceid,omitempty`|`string`| 
`FatherID`|`fatherid,omitempty`|`string`| 
`FormInfo`|`form_info`|`WedocFormInfo`| 

### `WedocCreateFormResp` 创建收集表响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`FormID`|`formid`|`string`| 

### `WedocCreateSheetPrivRuleReq` 新增智能表格指定成员额外权限请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Name`|`name`|`string`| 

### `WedocCreateSheetPrivRuleResp` 新增智能表格指定成员额外权限响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`RuleID`|`rule_id`|`uint32`| 

### `WedocCreatedTimeFieldProperty` CreatedTimeFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`Format`|`format`|`string`| 

### `WedocCurrencyFieldProperty` CurrencyFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`CurrencyType`|`currency_type`|`string`| 
`DecimalPlaces`|`decimal_places`|`int64`| 
`UseSeparate`|`use_separate`|`bool`| 

### `WedocDateTimeFieldProperty` DateTimeFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`Format`|`format`|`string`| 
`AutoFill`|`auto_fill`|`bool`| 

### `WedocDeleteBlocksReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`PageID`|`page_id`|`string`| 
`IDs`|`ids`|`[]string`| 

### `WedocDeleteContent` 删除内容

Name|JSON|Type|Doc
:---|:---|:---|:--
`Range`|`range`|`*WedocRange`| 

### `WedocDeleteDimensionResponse` DeleteDimensionResponse

Name|JSON|Type|Doc
:---|:---|:---|:--
`Deleted`|`deleted`|`int64`| 

### `WedocDeleteDocReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid,omitempty`|`string`| 
`FormID`|`formid,omitempty`|`string`| 

### `WedocDeleteFieldGroupsReq` 删除编组

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`FieldGroupIDs`|`field_group_ids`|`[]string`| 

### `WedocDeleteFieldsReq` 删除字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`FieldIDs`|`field_ids`|`[]string`| 

### `WedocDeletePageReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`PageID`|`page_id`|`string`| 

### `WedocDeleteRecordsReq` 删除记录

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`RecordIDs`|`record_ids`|`[]string`| 

### `WedocDeleteSheetPrivRuleReq` 删除智能表格指定成员额外权限请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`RuleIDList`|`rule_id_list`|`[]uint32`| 

### `WedocDeleteSheetReq` 删除子表

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 

### `WedocDeleteSheetResponse` DeleteSheetResponse

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id`|`string`| 

### `WedocDeleteSmartsheetReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`BlockID`|`block_id`|`string`| 

### `WedocDeleteViewsReq` 删除视图

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`ViewIDs`|`view_ids`|`[]string`| 

### `WedocDeleteVipAccountsReq` 取消高级功能账号请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserIDList`|`userid_list`|`[]string`| 

### `WedocDeleteVipAccountsResp` 取消高级功能账号响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`SuccUserIDList`|`succ_userid_list`|`[]string`| 
`FailUserIDList`|`fail_userid_list`|`[]string`| 

### `WedocDocAccessRule` 文档查看规则

Name|JSON|Type|Doc
:---|:---|:---|:--
`EnableCorpInternal`|`enable_corp_internal`|`bool`| 
`CorpInternalAuth`|`corp_internal_auth`|`uint32`| 
`EnableCorpExternal`|`enable_corp_external`|`bool`| 
`CorpExternalAuth`|`corp_external_auth`|`uint32`| 
`CorpInternalApproveOnlyByAdmin`|`corp_internal_approve_only_by_admin`|`bool`| 
`CorpExternalApproveOnlyByAdmin`|`corp_external_approve_only_by_admin`|`bool`| 
`BanShareExternal`|`ban_share_external`|`bool`| 

### `WedocDocBaseInfo` 文档基础信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`DocName`|`doc_name`|`string`| 
`CreateTime`|`create_time`|`uint64`| 
`ModifyTime`|`modify_time`|`uint64`| 
`DocType`|`doc_type`|`uint32`| 

### `WedocDocCoAuth` 文档查看权限特定部门

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`uint32`| 
`DepartmentID`|`departmentid`|`uint64`| 
`Auth`|`auth`|`uint32`| 

### `WedocDocMember` 文档成员/通知范围成员

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`uint32`| 
`UserID`|`userid,omitempty`|`string`| 
`DepartmentID`|`departmentid,omitempty`|`uint64`| 
`TmpExternalUserID`|`tmp_external_userid,omitempty`|`string`| 
`Auth`|`auth,omitempty`|`uint32`| 

### `WedocDocMemberDelete` 删除文档成员

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`uint32`| 
`UserID`|`userid,omitempty`|`string`| 
`TmpExternalUserID`|`tmp_external_userid,omitempty`|`string`| 

### `WedocDocSecureSetting` 文档安全设置

Name|JSON|Type|Doc
:---|:---|:---|:--
`EnableReadonlyCopy`|`enable_readonly_copy`|`bool`| 
`Watermark`|`watermark,omitempty`|`*WedocDocWatermark`| 
`EnableReadonlyComment`|`enable_readonly_comment`|`bool`| 

### `WedocDocWatermark` 文档水印设置

Name|JSON|Type|Doc
:---|:---|:---|:--
`MarginType`|`margin_type,omitempty`|`uint32`| 
`ShowVisitorName`|`show_visitor_name,omitempty`|`bool`| 
`ShowText`|`show_text,omitempty`|`bool`| 
`Text`|`text,omitempty`|`string`| 

### `WedocDocumentNode` 文档内容节点

Name|JSON|Type|Doc
:---|:---|:---|:--
`Begin`|`begin`|`uint32`| 
`End`|`end`|`uint32`| 
`Property`|`property,omitempty`|`json.RawMessage`| 
`Type`|`type`|`string`| 
`Children`|`children,omitempty`|`[]WedocDocumentNode`| 
`Text`|`text,omitempty`|`string`| 

### `WedocDocumentUpdateRequest` 文档批量更新操作，一次只能填一个字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`ReplaceText`|`replace_text,omitempty`|`*WedocReplaceText`| 
`InsertText`|`insert_text,omitempty`|`*WedocInsertText`| 
`DeleteContent`|`delete_content,omitempty`|`*WedocDeleteContent`| 
`InsertImage`|`insert_image,omitempty`|`*WedocInsertImage`| 
`InsertPageBreak`|`insert_page_break,omitempty`|`*WedocInsertPageBreak`| 
`InsertTable`|`insert_table,omitempty`|`*WedocInsertTable`| 
`InsertParagraph`|`insert_paragraph,omitempty`|`*WedocInsertParagraph`| 
`UpdateTextProperty`|`update_text_property,omitempty`|`*WedocUpdateTextProperty`| 

### `WedocExportTaskReq` 提交导出任务

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`ContentType`|`content_type`|`uint32`| 
`PageID`|`page_id,omitempty`|`string`| 

### `WedocExportTaskResp` 提交导出任务

Name|JSON|Type|Doc
:---|:---|:---|:--
`TaskID`|`task_id`|`string`| 

### `WedocField` Field

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`| 
`FieldTitle`|`field_title`|`string`| 
`FieldType`|`field_type`|`string`| 
`PropertyNumber`|`property_number`|`WedocNumberFieldProperty`| 
`PropertyCheckbox`|`property_checkbox`|`WedocCheckboxFieldProperty`| 
`PropertyDateTime`|`property_date_time`|`WedocDateTimeFieldProperty`| 
`PropertyAttachment`|`property_attachment`|`WedocAttachmentFieldProperty`| 
`PropertyUser`|`property_user`|`WedocUserFieldProperty`| 
`PropertyURL`|`property_url`|`WedocUrlFieldProperty`| 
`PropertySelect`|`property_select`|`WedocSelectFieldProperty`| 
`PropertyCreatedTime`|`property_created_time`|`WedocCreatedTimeFieldProperty`| 
`PropertyModifiedTime`|`property_modified_time`|`WedocModifiedTimeFieldProperty`| 
`PropertyProgress`|`property_progress`|`WedocProgressFieldProperty`| 
`PropertySingleSelect`|`property_single_select`|`WedocSingleSelectFieldProperty`| 
`PropertyReference`|`property_reference`|`WedocReferenceFieldProperty`| 
`PropertyLocation`|`property_location`|`WedocLocationFieldProperty`| 
`PropertyAutoNumber`|`property_auto_number`|`WedocAutoNumberFieldProperty`| 
`PropertyCurrency`|`property_currency`|`WedocCurrencyFieldProperty`| 
`PropertyWwGroup`|`property_ww_group`|`WedocWwGroupFieldProperty`| 
`PropertyPercentage`|`property_percentage`|`WedocPercentageFieldProperty`| 
`PropertyBarcode`|`property_barcode`|`WedocBarcodeFieldProperty`| 

### `WedocFieldGroup` 智能表格编组

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldGroupID`|`field_group_id,omitempty`|`string`| 
`Name`|`name,omitempty`|`string`| 
`Children`|`children,omitempty`|`[]WedocFieldGroupChild`| 

### `WedocFieldGroupChild` 编组内的字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`| 

### `WedocFileBlockProps` 文件块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`FileID`|`file_id,omitempty`|`string`| 

### `WedocFilterDataTimeValue` FilterDataTimeValue

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`| 
`Value`|`value`|`[]string`| 

### `WedocFilterSpec` FilterSpec

Name|JSON|Type|Doc
:---|:---|:---|:--
`Conjunction`|`conjunction`|`string`| 
`Conditions`|`conditions`|`[]WedocCondition`| 

### `WedocFinishFileUploadReq` 分块上传完成

Name|JSON|Type|Doc
:---|:---|:---|:--
`UploadKey`|`upload_key`|`string`| 

### `WedocFinishFileUploadResp` 分块上传完成

Name|JSON|Type|Doc
:---|:---|:---|:--
`FileURL`|`file_url`|`string`| 

### `WedocFormAnswer` 收集表答案

Name|JSON|Type|Doc
:---|:---|:---|:--
`AnswerID`|`answer_id`|`uint64`| 
`UserName`|`user_name`|`string`| 
`CTime`|`ctime`|`uint64`| 
`MTime`|`mtime`|`uint64`| 
`Reply`|`reply`|`WedocFormReply`| 
`AnswerStatus`|`answer_status`|`uint32`| 
`UserID`|`userid,omitempty`|`string`| 
`TmpExternalUserID`|`tmp_external_userid,omitempty`|`string`| 

### `WedocFormDepartmentReply` 部门题答案

Name|JSON|Type|Doc
:---|:---|:---|:--
`List`|`list`|`[]WedocFormDepartmentReplyItem`| 

### `WedocFormDepartmentReplyItem` 部门题答案项

Name|JSON|Type|Doc
:---|:---|:---|:--
`DepartmentID`|`department_id`|`uint64`| 

### `WedocFormDurationReply` 时长题答案

Name|JSON|Type|Doc
:---|:---|:---|:--
`BeginTime`|`begin_time`|`uint32`| 
`EndTime`|`end_time`|`uint32`| 
`TimeScale`|`time_scale`|`uint32`| 
`DayRange`|`day_range`|`uint32`| 
`Days`|`days`|`float64`| 
`Hours`|`hours`|`float64`| 

### `WedocFormFileExtendReply` 文件题答案

Name|JSON|Type|Doc
:---|:---|:---|:--
`Name`|`name`|`string`| 
`FileID`|`fileid`|`string`| 

### `WedocFormFillInRange` 可填写范围

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserIDs`|`userids,omitempty`|`[]string`| 
`DepartmentIDs`|`departmentids,omitempty`|`[]uint64`| 

### `WedocFormInfo` 收集表信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`FormID`|`formid,omitempty`|`string`| 
`FormTitle`|`form_title`|`string`| 
`FormDesc`|`form_desc,omitempty`|`string`| 
`FormHeader`|`form_header,omitempty`|`string`| 
`FormQuestion`|`form_question`|`WedocFormQuestion`| 
`FormSetting`|`form_setting,omitempty`|`*WedocFormSetting`| 
`RepeatedID`|`repeated_id,omitempty`|`[]string`| 

### `WedocFormItem` 单个问题

Name|JSON|Type|Doc
:---|:---|:---|:--
`QuestionID`|`question_id`|`uint32`| 
`Title`|`title`|`string`| 
`Pos`|`pos`|`uint32`| 
`Status`|`status`|`uint32`| 
`ReplyType`|`reply_type`|`uint32`| 
`MustReply`|`must_reply`|`bool`| 
`Note`|`note,omitempty`|`string`| 
`Placeholder`|`placeholder,omitempty`|`string`| 
`OptionItem`|`option_item,omitempty`|`[]WedocFormOptionItem`| 
`QuestionExtendSetting`|`question_extend_setting,omitempty`|`json.RawMessage`| 

### `WedocFormMemberReply` 成员题答案

Name|JSON|Type|Doc
:---|:---|:---|:--
`List`|`list`|`[]WedocFormMemberReplyItem`| 

### `WedocFormMemberReplyItem` 成员题答案项

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 

### `WedocFormOptionExtendReply` 选择题「其他」答案

Name|JSON|Type|Doc
:---|:---|:---|:--
`OptionReply`|`option_reply`|`uint32`| 
`ExtendText`|`extend_text`|`string`| 

### `WedocFormOptionItem` 单选题选项

Name|JSON|Type|Doc
:---|:---|:---|:--
`Key`|`key`|`uint32`| 
`Value`|`value`|`string`| 
`Status`|`status`|`uint32`| 

### `WedocFormQuestion` 收集表问题列表

Name|JSON|Type|Doc
:---|:---|:---|:--
`Items`|`items`|`[]WedocFormItem`| 

### `WedocFormReply` 答案明细

Name|JSON|Type|Doc
:---|:---|:---|:--
`Items`|`items`|`[]WedocFormReplyItem`| 

### `WedocFormReplyItem` 单个问题的答案

Name|JSON|Type|Doc
:---|:---|:---|:--
`QuestionID`|`question_id`|`uint64`| 
`TextReply`|`text_reply,omitempty`|`string`| 
`OptionReply`|`option_reply,omitempty`|`[]uint32`| 
`OptionExtendReply`|`option_extend_reply,omitempty`|`[]WedocFormOptionExtendReply`| 
`FileExtendReply`|`file_extend_reply,omitempty`|`[]WedocFormFileExtendReply`| 
`DepartmentReply`|`department_reply,omitempty`|`*WedocFormDepartmentReply`| 
`MemberReply`|`member_reply,omitempty`|`*WedocFormMemberReply`| 
`DurationReply`|`duration_reply,omitempty`|`*WedocFormDurationReply`| 

### `WedocFormSetting` 收集表设置

Name|JSON|Type|Doc
:---|:---|:---|:--
`FillOutAuth`|`fill_out_auth,omitempty`|`uint32`| 
`FillInRange`|`fill_in_range,omitempty`|`*WedocFormFillInRange`| 
`SettingManagerRange`|`setting_manager_range,omitempty`|`*WedocFormSettingManagerRange`| 
`TimedRepeatInfo`|`timed_repeat_info,omitempty`|`*WedocFormTimedRepeatInfo`| 
`AllowMultiFill`|`allow_multi_fill,omitempty`|`bool`| 
`TimedFinish`|`timed_finish,omitempty`|`uint32`| 
`CanAnonymous`|`can_anonymous,omitempty`|`bool`| 
`CanNotifySubmit`|`can_notify_submit,omitempty`|`bool`| 

### `WedocFormSettingManagerRange` 收集表管理员

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserIDs`|`userids,omitempty`|`[]string`| 

### `WedocFormSubmitUser` 已填写人

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid,omitempty`|`string`| 
`TmpExternalUserID`|`tmp_external_userid,omitempty`|`string`| 
`SubmitTime`|`submit_time`|`uint64`| 
`AnswerID`|`answer_id`|`uint64`| 
`UserName`|`user_name`|`string`| 

### `WedocFormTimedRepeatInfo` 定时重复设置

Name|JSON|Type|Doc
:---|:---|:---|:--
`Enable`|`enable,omitempty`|`bool`| 
`RemindTime`|`remind_time,omitempty`|`uint32`| 
`RuleCTime`|`rule_ctime,omitempty`|`uint32`| 
`RuleMTime`|`rule_mtime,omitempty`|`uint32`| 
`RepeatType`|`repeat_type,omitempty`|`uint32`| 
`WeekFlag`|`week_flag,omitempty`|`uint32`| 
`SkipHoliday`|`skip_holiday,omitempty`|`bool`| 
`DayOfMonth`|`day_of_month,omitempty`|`uint32`| 
`ForkFinishType`|`fork_finish_type,omitempty`|`uint32`| 

### `WedocFormUnfillUser` 未填写人

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid,omitempty`|`string`| 
`UserName`|`user_name`|`string`| 

### `WedocGanttViewProperty` GanttViewProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`StartDateFieldID`|`start_date_field_id`|`string`| 
`EndDateFieldID`|`end_date_field_id`|`string`| 

### `WedocGetBlockListReq` 获取内容块列表请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`PageID`|`page_id`|`string`| 
`IDs`|`ids,omitempty`|`[]string`| 
`Start`|`start,omitempty`|`int32`| 
`Limit`|`limit,omitempty`|`int32`| 

### `WedocGetBlockListResp` 获取内容块列表响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Blocks`|`blocks`|`[]WedocBlock`| 
`HasMore`|`has_more`|`string`| 
`NextStart`|`next_start`|`int32`| 

### `WedocGetDocAuthReq` 获取文档权限信息请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 

### `WedocGetDocAuthResp` 获取文档权限信息响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`AccessRule`|`access_rule`|`WedocDocAccessRule`| 
`SecureSetting`|`secure_setting`|`WedocDocSecureSetting`| 
`DocMemberList`|`doc_member_list`|`[]WedocDocMember`| 
`CoAuthList`|`co_auth_list`|`[]WedocDocCoAuth`| 

### `WedocGetDocBaseInfoReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 

### `WedocGetDocBaseInfoResp` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocBaseInfo`|`doc_base_info,omitempty`|`*WedocDocBaseInfo`| 
`DocID`|`docid`|`string`| 
`DocName`|`doc_name`|`string`| 
`CreateTime`|`create_time`|`uint64`| 
`ModifyTime`|`modify_time`|`uint64`| 
`DocType`|`doc_type`|`uint32`| 

### `WedocGetDocumentReq` 获取文档数据请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 

### `WedocGetDocumentResp` 获取文档数据响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Version`|`version`|`uint32`| 
`Document`|`document`|`WedocDocumentNode`| 

### `WedocGetExportResultReq` 查询任务结果

Name|JSON|Type|Doc
:---|:---|:---|:--
`TaskID`|`task_id`|`string`| 

### `WedocGetExportResultResp` 查询任务结果

Name|JSON|Type|Doc
:---|:---|:---|:--
`TaskDone`|`task_done`|`bool`| 
`Content`|`content`|`string`| 

### `WedocGetFieldGroupsReq` 获取编组

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`Offset`|`offset,omitempty`|`uint32`| 
`Limit`|`limit,omitempty`|`uint32`| 

### `WedocGetFieldGroupsResp` 获取编组

Name|JSON|Type|Doc
:---|:---|:---|:--
`Total`|`total`|`uint32`| 
`HasMore`|`has_more`|`bool`| 
`Next`|`next`|`uint32`| 
`FieldGroups`|`field_groups`|`[]WedocFieldGroup`| 

### `WedocGetFieldsReq` 查询字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`ViewID`|`view_id,omitempty`|`string`| 
`FieldIDs`|`field_ids,omitempty`|`[]string`| 
`FieldTitles`|`field_titles,omitempty`|`[]string`| 
`Offset`|`offset,omitempty`|`int64`| 
`Limit`|`limit,omitempty`|`int64`| 

### `WedocGetFieldsResp` 查询字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`Total`|`total`|`int64`| 
`Fields`|`fields`|`[]WedocField`| 

### `WedocGetFormAnswerReq` 读取收集表答案请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`RepeatedID`|`repeated_id`|`string`| 
`AnswerIDs`|`answer_ids`|`[]uint64`| 

### `WedocGetFormAnswerResp` 读取收集表答案响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Answer`|`answer,omitempty`|`*WedocFormAnswer`| 
`AnswerList`|`answer_list,omitempty`|`[]WedocFormAnswer`| 

### `WedocGetFormInfoReq` 获取收集表信息请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`FormID`|`formid`|`string`| 

### `WedocGetFormInfoResp` 获取收集表信息响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`FormInfo`|`form_info`|`WedocFormInfo`| 

### `WedocGetFormStatisticReq` 收集表统计信息查询请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`RepeatedID`|`repeated_id`|`string`| 
`ReqType`|`req_type`|`uint32`| 
`StartTime`|`start_time,omitempty`|`*uint64`| 
`EndTime`|`end_time,omitempty`|`*uint64`| 
`Limit`|`limit,omitempty`|`*uint64`| 
`Cursor`|`cursor,omitempty`|`*uint64`| 

### `WedocGetFormStatisticResp` 收集表统计信息查询响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`FillCnt`|`fill_cnt`|`uint64`| 
`FillUserCnt`|`fill_user_cnt`|`uint64`| 
`UnfillUserCnt`|`unfill_user_cnt`|`uint64`| 
`SubmitUsers`|`submit_users,omitempty`|`[]WedocFormSubmitUser`| 
`UnfillUsers`|`unfill_users,omitempty`|`[]WedocFormUnfillUser`| 
`HasMore`|`has_more`|`bool`| 
`Cursor`|`cursor`|`uint64`| 

### `WedocGetPageHierarchyReq` 获取页面结构请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 

### `WedocGetPageHierarchyResp` 获取页面结构响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Pages`|`pages`|`[]WedocPage`| 

### `WedocGetRecordsReq` 查询记录

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`ViewID`|`view_id,omitempty`|`string`| 
`RecordIDs`|`record_ids,omitempty`|`[]string`| 
`KeyType`|`key_type,omitempty`|`string`| 
`FieldTitles`|`field_titles,omitempty`|`[]string`| 
`FieldIDs`|`field_ids,omitempty`|`[]string`| 
`Sort`|`sort,omitempty`|`[]WedocSort`| 
`Offset`|`offset,omitempty`|`uint32`| 
`Limit`|`limit,omitempty`|`uint32`| 
`Ver`|`ver,omitempty`|`uint32`| 
`FilterSpec`|`filter_spec,omitempty`|`*WedocFilterSpec`| 
`NeedFormulaResult`|`need_formula_result,omitempty`|`bool`| 

### `WedocGetRecordsResp` 查询记录

Name|JSON|Type|Doc
:---|:---|:---|:--
`Total`|`total`|`uint32`| 
`HasMore`|`has_more`|`bool`| 
`Next`|`next`|`uint32`| 
`Records`|`records`|`[]WedocRecord`| 
`Ver`|`ver`|`uint32`| 

### `WedocGetSheetPrivReq` 查询智能表格子表权限请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Type`|`type`|`uint32`| 
`RuleIDList`|`rule_id_list,omitempty`|`[]uint32`| 

### `WedocGetSheetPrivResp` 查询智能表格子表权限响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`RuleList`|`rule_list`|`[]WedocSheetPrivRule`| 

### `WedocGetSheetPropertiesReq` 获取表格行列信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 

### `WedocGetSheetPropertiesResp` 获取表格行列信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`Properties`|`properties`|`[]WedocProperties`| 

### `WedocGetSheetRangeDataReq` 获取表格数据请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`Range`|`range`|`string`| 

### `WedocGetSheetRangeDataResp` 获取表格数据响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Data`|`data`|`WedocSheetRangeData`| 

### `WedocGetSheetReq` 查询子表

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id,omitempty`|`string`| 
`NeedAllTypeSheet`|`need_all_type_sheet,omitempty`|`bool`| 

### `WedocGetSheetResp` 查询子表

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetList`|`sheet_list`|`[]WedocSheetInfo`| 

### `WedocGetSmartsheetInfoReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 

### `WedocGetSmartsheetInfoResp` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`SSDocID`|`ss_docid`|`string`| 

### `WedocGetViewsReq` 查询视图

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`ViewIDs`|`view_ids,omitempty`|`[]string`| 
`Offset`|`offset,omitempty`|`uint32`| 
`Limit`|`limit,omitempty`|`uint32`| 

### `WedocGetViewsResp` 查询视图

Name|JSON|Type|Doc
:---|:---|:---|:--
`Total`|`total`|`uint32`| 
`HasMore`|`has_more`|`bool`| 
`Next`|`next`|`uint32`| 
`Views`|`views`|`[]WedocView`| 

### `WedocGridData` 表格数据

Name|JSON|Type|Doc
:---|:---|:---|:--
`StartRow`|`start_row`|`uint32`| 
`StartColumn`|`start_column`|`uint32`| 
`Rows`|`rows`|`[]WedocRowData`| 

### `WedocGroup` 分组设置中的字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`| 
`Desc`|`desc,omitempty`|`bool`| 

### `WedocGroupSpec` GroupSpec

Name|JSON|Type|Doc
:---|:---|:---|:--
`Groups`|`groups`|`[]WedocGroup`| 

### `WedocHighlightBlockProps` 背景块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`BackgroundImageURL`|`background_image_url,omitempty`|`string`| 
`ExtendBackground`|`extend_background,omitempty`|`bool`| 
`BlockColor`|`block_color,omitempty`|`string`| 
`BorderColor`|`border_color,omitempty`|`string`| 

### `WedocImageBlockProps` 图片块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`URL`|`url,omitempty`|`string`| 
`Width`|`width,omitempty`|`float64`| 
`Height`|`height,omitempty`|`float64`| 
`Stroke`|`stroke,omitempty`|`bool`| 

### `WedocInitFileUploadReq` 分块上传初始化

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`FileName`|`file_name`|`string`| 
`Size`|`size`|`uint64`| 
`BlockSHA`|`block_sha`|`[]string`| 

### `WedocInitFileUploadResp` 分块上传初始化

Name|JSON|Type|Doc
:---|:---|:---|:--
`HitExist`|`hit_exist`|`bool`| 
`UploadKey`|`upload_key`|`string`| 
`FileURL`|`file_url`|`string`| 

### `WedocInsertImage` 插入图片

Name|JSON|Type|Doc
:---|:---|:---|:--
`ImageID`|`image_id`|`string`| 
`Location`|`location`|`*WedocLocation`| 
`Width`|`width,omitempty`|`uint32`| 
`Height`|`height,omitempty`|`uint32`| 

### `WedocInsertPageBreak` 插入分页符

Name|JSON|Type|Doc
:---|:---|:---|:--
`Location`|`location`|`*WedocLocation`| 

### `WedocInsertParagraph` 插入段落

Name|JSON|Type|Doc
:---|:---|:---|:--
`Location`|`location`|`*WedocLocation`| 

### `WedocInsertTable` 插入表格

Name|JSON|Type|Doc
:---|:---|:---|:--
`Rows`|`rows`|`uint32`| 
`Cols`|`cols`|`uint32`| 
`Location`|`location`|`*WedocLocation`| 

### `WedocInsertText` 插入文本

Name|JSON|Type|Doc
:---|:---|:---|:--
`Text`|`text`|`string`| 
`Location`|`location`|`*WedocLocation`| 

### `WedocLinkBlockProps` 链接块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`LinkURL`|`link_url,omitempty`|`string`| 
`LinkName`|`link_name,omitempty`|`string`| 
`LinkDescription`|`link_description,omitempty`|`string`| 

### `WedocListVipAccountsReq` 获取高级功能账号列表请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`Cursor`|`cursor,omitempty`|`string`| 
`Limit`|`limit,omitempty`|`int32`| 

### `WedocListVipAccountsResp` 获取高级功能账号列表响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`HasMore`|`has_more`|`bool`| 
`NextCursor`|`next_cursor`|`string`| 
`UserIDList`|`userid_list`|`[]string`| 

### `WedocLocation` 文档位置

Name|JSON|Type|Doc
:---|:---|:---|:--
`Index`|`index`|`uint32`| 

### `WedocLocationFieldProperty` LocationFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`InputType`|`input_type`|`string`| 

### `WedocModifiedTimeFieldProperty` ModifiedTimeFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`Format`|`format`|`string`| 

### `WedocModifyDocJoinRuleReq` 修改文档加入规则请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`EnableCorpInternal`|`enable_corp_internal,omitempty`|`*bool`| 
`CorpInternalAuth`|`corp_internal_auth,omitempty`|`*uint32`| 
`EnableCorpExternal`|`enable_corp_external,omitempty`|`*bool`| 
`CorpExternalAuth`|`corp_external_auth,omitempty`|`*uint32`| 
`CorpInternalApproveOnlyByAdmin`|`corp_internal_approve_only_by_admin,omitempty`|`*bool`| 
`CorpExternalApproveOnlyByAdmin`|`corp_external_approve_only_by_admin,omitempty`|`*bool`| 
`BanShareExternal`|`ban_share_external,omitempty`|`*bool`| 
`UpdateCoAuthList`|`update_co_auth_list,omitempty`|`*bool`| 
`CoAuthList`|`co_auth_list,omitempty`|`[]WedocDocCoAuth`| 

### `WedocModifyDocMemberReq` 修改文档成员与权限请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`UpdateFileMemberList`|`update_file_member_list,omitempty`|`[]WedocDocMember`| 
`DelFileMemberList`|`del_file_member_list,omitempty`|`[]WedocDocMemberDelete`| 

### `WedocModifyDocSafetySettingReq` 修改文档安全设置请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`EnableReadonlyCopy`|`enable_readonly_copy,omitempty`|`*bool`| 
`Watermark`|`watermark,omitempty`|`*WedocDocWatermark`| 

### `WedocModifyFormReq` 编辑收集表请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`Oper`|`oper`|`uint32`| 
`FormID`|`formid`|`string`| 
`FormInfo`|`form_info,omitempty`|`*WedocFormInfo`| 

### `WedocModifySheetPrivRuleMemberReq` 更新智能表格指定成员额外权限请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`RuleID`|`rule_id`|`uint32`| 
`AddMemberRange`|`add_member_range,omitempty`|`*WedocSheetPrivMemberRange`| 
`DelMemberRange`|`del_member_range,omitempty`|`*WedocSheetPrivMemberRange`| 

### `WedocNumberFieldProperty` NumberFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`DecimalPlaces`|`decimal_places`|`int64`| 
`UseSeparate`|`use_separate`|`bool`| 

### `WedocNumberRule` 自动编号规则

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`| 
`Value`|`value`|`string`| 

### `WedocNumberValue` 数字类条件值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Value`|`value`|`float64`| 

### `WedocOption` Option

Name|JSON|Type|Doc
:---|:---|:---|:--
`Style`|`style`|`int64`| 
`Text`|`text`|`string`| 

### `WedocPage` 智能文档页面（页面层级）

Name|JSON|Type|Doc
:---|:---|:---|:--
`PageID`|`page_id`|`string`| 
`ParentID`|`parent_id`|`string`| 
`Title`|`title`|`string`| 

### `WedocPageInfo` 智能文档页面信息（添加/更新页面）

Name|JSON|Type|Doc
:---|:---|:---|:--
`PageID`|`page_id,omitempty`|`string`| 
`Title`|`title,omitempty`|`string`| 
`ParentID`|`parent_id,omitempty`|`string`| 
`AfterID`|`after_id,omitempty`|`string`| 
`LayoutMode`|`layout_mode,omitempty`|`uint32`| 

### `WedocPercentageFieldProperty` PercentageFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`DecimalPlaces`|`decimal_places`|`int64`| 
`UseSeparate`|`use_separate`|`bool`| 

### `WedocProgressFieldProperty` ProgressFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`DecimalPlaces`|`decimal_places`|`int64`| 

### `WedocProperties` 在线表格工作表属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id`|`string`| 
`Title`|`title`|`string`| 
`RowCount`|`row_count`|`uint32`| 
`ColumnCount`|`column_count`|`uint32`| 

### `WedocPublishSmartDocReq` 发布智能文档请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`PublishRange`|`publish_range,omitempty`|`*uint32`| 
`AuthList`|`auth_list,omitempty`|`[]WedocDocMember`| 

### `WedocPublishSmartDocResp` 发布智能文档响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`ShareCode`|`share_code`|`string`| 
`PublishURL`|`publish_url`|`string`| 
`Version`|`version`|`uint64`| 
`PublishTime`|`publish_time`|`uint64`| 
`PublishDocTitle`|`publish_doc_title`|`string`| 

### `WedocRange` 文档范围

Name|JSON|Type|Doc
:---|:---|:---|:--
`StartIndex`|`start_index`|`uint32`| 
`Length`|`length`|`uint32`| 

### `WedocRecord` Record

Name|JSON|Type|Doc
:---|:---|:---|:--
`RecordID`|`record_id`|`string`| 
`CreateTime`|`create_time`|`string`| 
`UpdateTime`|`update_time`|`string`| 
`Values`|`values`|`map[string]any`| 
`CreatorName`|`creator_name`|`string`| 
`UpdaterName`|`updater_name`|`string`| 

### `WedocReferenceFieldProperty` ReferenceFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`SubID`|`sub_id`|`string`| 
`FieldID`|`field_id`|`string`| 
`IsMultiple`|`is_multiple`|`bool`| 
`ViewID`|`view_id`|`string`| 

### `WedocRenameDocReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid,omitempty`|`string`| 
`FormID`|`formid,omitempty`|`string`| 
`NewName`|`new_name`|`string`| 

### `WedocReplaceText` 替换文本

Name|JSON|Type|Doc
:---|:---|:---|:--
`Text`|`text`|`string`| 
`Ranges`|`ranges`|`[]WedocRange`| 

### `WedocRowData` 行数据

Name|JSON|Type|Doc
:---|:---|:---|:--
`Values`|`values`|`[]WedocCellData`| 

### `WedocSelectFieldProperty` SelectFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`IsQuickAdd`|`is_quick_add`|`bool`| 
`Options`|`options`|`[]WedocOption`| 

### `WedocShareDocReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid,omitempty`|`string`| 
`FormID`|`formid,omitempty`|`string`| 

### `WedocShareDocResp` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`ShareURL`|`share_url`|`string`| 

### `WedocSheetFieldPriv` 字段权限

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldRangeType`|`field_range_type,omitempty`|`uint32`| 
`FieldRuleList`|`field_rule_list,omitempty`|`[]WedocSheetFieldRule`| 
`FieldDefaultRule`|`field_default_rule,omitempty`|`*WedocSheetFieldRule`| 

### `WedocSheetFieldRule` 单个字段权限

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id,omitempty`|`string`| 
`FieldType`|`field_type,omitempty`|`string`| 
`CanEdit`|`can_edit`|`bool`| 
`CanInsert`|`can_insert`|`bool`| 
`CanView`|`can_view`|`bool`| 

### `WedocSheetInfo` 智能表格子表信息（查询子表返回）

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id`|`string`| 
`Title`|`title`|`string`| 
`IsVisible`|`is_visible`|`bool`| 
`Type`|`type`|`string`| 

### `WedocSheetPrivItem` 子表权限

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id`|`string`| 
`Priv`|`priv,omitempty`|`string`| 
`CanInsertRecord`|`can_insert_record,omitempty`|`bool`| 
`CanDeleteRecord`|`can_delete_record,omitempty`|`bool`| 
`CanCreateModifyDeleteView`|`can_create_modify_delete_view,omitempty`|`bool`| 
`FieldPriv`|`field_priv,omitempty`|`*WedocSheetFieldPriv`| 
`RecordPriv`|`record_priv,omitempty`|`*WedocSheetRecordPriv`| 
`Clear`|`clear,omitempty`|`bool`| 

### `WedocSheetPrivMemberRange` 额外权限的成员范围

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserIDList`|`userid_list,omitempty`|`[]string`| 

### `WedocSheetPrivRule` 一条智能表格内容权限规则

Name|JSON|Type|Doc
:---|:---|:---|:--
`RuleID`|`rule_id,omitempty`|`uint32`| 
`Type`|`type`|`uint32`| 
`Name`|`name,omitempty`|`string`| 
`PrivList`|`priv_list,omitempty`|`[]WedocSheetPrivItem`| 

### `WedocSheetProperties` 智能表格子表属性（添加/更新子表使用）

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id,omitempty`|`string`| 
`Title`|`title,omitempty`|`string`| 
`Index`|`index,omitempty`|`int32`| 

### `WedocSheetRecordPriv` 记录权限

Name|JSON|Type|Doc
:---|:---|:---|:--
`RecordRangeType`|`record_range_type,omitempty`|`uint32`| 
`RecordRuleList`|`record_rule_list,omitempty`|`[]WedocSheetRecordRule`| 
`OtherPriv`|`other_priv,omitempty`|`uint32`| 

### `WedocSheetRecordRule` 记录权限条件

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`| 
`FieldType`|`field_type,omitempty`|`string`| 
`OperType`|`oper_type`|`uint32`| 
`Value`|`value,omitempty`|`[]string`| 

### `WedocSingleSelectFieldProperty` SingleSelectFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`IsQuickAdd`|`is_quick_add`|`bool`| 
`Options`|`options`|`[]WedocOption`| 

### `WedocSmartDocSmartsheetInfo` 智能文档中的数据表信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`BlockID`|`block_id,omitempty`|`string`| 
`Title`|`title,omitempty`|`string`| 
`SheetID`|`sheet_id,omitempty`|`string`| 
`SSDocID`|`ss_docid,omitempty`|`string`| 
`AfterID`|`after_id,omitempty`|`string`| 

### `WedocSmartSheetViewBlockProps` 数据表视图块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`doc_id,omitempty`|`string`| 
`TableID`|`table_id,omitempty`|`string`| 
`ViewID`|`view_id,omitempty`|`string`| 
`TableBlockID`|`table_block_id,omitempty`|`string`| 
`EnableTableTitle`|`enable_table_title,omitempty`|`bool`| 
`EnableAddRow`|`enable_add_row,omitempty`|`bool`| 

### `WedocSort` Sort

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldTitle`|`field_title`|`string`| 
`Desc`|`desc,omitempty`|`bool`| 

### `WedocSortInfo` 排序设置中的字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`| 
`Desc`|`desc,omitempty`|`bool`| 

### `WedocSortSpec` SortSpec

Name|JSON|Type|Doc
:---|:---|:---|:--
`SortInfos`|`sort_infos`|`[]WedocSortInfo`| 

### `WedocSpreadsheetAddSheetRequest` 新增工作表

Name|JSON|Type|Doc
:---|:---|:---|:--
`Title`|`title`|`string`| 
`RowCount`|`row_count`|`uint32`| 
`ColumnCount`|`column_count`|`uint32`| 

### `WedocSpreadsheetDeleteDimensionRequest` 删除行列

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id`|`string`| 
`Dimension`|`dimension`|`string`| 
`StartIndex`|`start_index`|`uint32`| 
`EndIndex`|`end_index`|`uint32`| 

### `WedocSpreadsheetDeleteSheetRequest` 删除工作表

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id`|`string`| 

### `WedocSpreadsheetUpdateRangeRequest` 更新单元格范围

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id`|`string`| 
`GridData`|`grid_data`|`*WedocGridData`| 

### `WedocSpreadsheetUpdateRequest` 表格批量更新操作，一次只能填一个字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`AddSheetRequest`|`add_sheet_request,omitempty`|`*WedocSpreadsheetAddSheetRequest`| 
`DeleteSheetRequest`|`delete_sheet_request,omitempty`|`*WedocSpreadsheetDeleteSheetRequest`| 
`UpdateRangeRequest`|`update_range_request,omitempty`|`*WedocSpreadsheetUpdateRangeRequest`| 
`DeleteDimensionRequest`|`delete_dimension_request,omitempty`|`*WedocSpreadsheetDeleteDimensionRequest`| 

### `WedocSpreadsheetUpdateResponse` 表格批量更新结果，一次只返回一个字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`AddSheetResponse`|`add_sheet_response,omitempty`|`*WedocAddSheetResponse`| 
`DeleteSheetResponse`|`delete_sheet_response,omitempty`|`*WedocDeleteSheetResponse`| 
`UpdateRangeResponse`|`update_range_response,omitempty`|`*WedocUpdateRangeResponse`| 
`DeleteDimensionResponse`|`delete_dimension_response,omitempty`|`*WedocDeleteDimensionResponse`| 

### `WedocStringValue` 文本类条件值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Value`|`value`|`[]string`| 

### `WedocTableBlockProps` 表格块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`CellPositions`|`cell_positions,omitempty`|`[]WedocTableCellMerge`| 
`EnableRowHeader`|`enable_row_header,omitempty`|`bool`| 
`EnableColumnHeader`|`enable_column_header,omitempty`|`bool`| 

### `WedocTableCellMerge` 表格单元格合并信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`StartCell`|`start_cell,omitempty`|`*WedocTableCellPosition`| 
`EndCell`|`end_cell,omitempty`|`*WedocTableCellPosition`| 

### `WedocTableCellPosition` 表格单元格坐标

Name|JSON|Type|Doc
:---|:---|:---|:--
`Row`|`row`|`uint32`| 
`Column`|`column`|`uint32`| 

### `WedocTextBlockProps` 文本块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Align`|`align,omitempty`|`uint32`| 
`TextAlign`|`text_align,omitempty`|`uint32`| 
`BlockColor`|`block_color,omitempty`|`string`| 

### `WedocTextFormat` 文本样式

Name|JSON|Type|Doc
:---|:---|:---|:--
`Font`|`font,omitempty`|`string`| 
`FontSize`|`font_size,omitempty`|`uint32`| 
`Bold`|`bold,omitempty`|`bool`| 
`Italic`|`italic,omitempty`|`bool`| 
`Strikethrough`|`strikethrough,omitempty`|`bool`| 
`Underline`|`underline,omitempty`|`bool`| 
`Color`|`color,omitempty`|`*WedocColor`| 

### `WedocTextProperty` 文本属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Bold`|`bold,omitempty`|`bool`| 
`Color`|`color,omitempty`|`string`| 
`BackgroundColor`|`background_color,omitempty`|`string`| 

### `WedocTodoBlockProps` 待办块属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Checked`|`checked,omitempty`|`bool`| 

### `WedocUpdateBlocksReq` 更新内容块请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`PageID`|`page_id`|`string`| 
`Blocks`|`blocks`|`[]WedocBlock`| 

### `WedocUpdateBlocksResp` 更新内容块响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Blocks`|`blocks`|`[]WedocBlock`| 

### `WedocUpdateField` UpdateField

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`| 
`FieldTitle`|`field_title,omitempty`|`string`| 
`FieldType`|`field_type`|`string`| 
`PropertyText`|`property_text,omitempty`|`json.RawMessage`| 
`PropertyNumber`|`property_number,omitempty`|`*WedocNumberFieldProperty`| 
`PropertyCheckbox`|`property_checkbox,omitempty`|`*WedocCheckboxFieldProperty`| 
`PropertyDateTime`|`property_date_time,omitempty`|`*WedocDateTimeFieldProperty`| 
`PropertyAttachment`|`property_attachment,omitempty`|`*WedocAttachmentFieldProperty`| 
`PropertyUser`|`property_user,omitempty`|`*WedocUserFieldProperty`| 
`PropertyURL`|`property_url,omitempty`|`*WedocUrlFieldProperty`| 
`PropertySelect`|`property_select,omitempty`|`*WedocSelectFieldProperty`| 
`PropertyCreatedUser`|`property_created_user,omitempty`|`json.RawMessage`| 
`PropertyModifiedUser`|`property_modified_user,omitempty`|`json.RawMessage`| 
`PropertyCreatedTime`|`property_created_time,omitempty`|`*WedocCreatedTimeFieldProperty`| 
`PropertyModifiedTime`|`property_modified_time,omitempty`|`*WedocModifiedTimeFieldProperty`| 
`PropertyProgress`|`property_progress,omitempty`|`*WedocProgressFieldProperty`| 
`PropertySingleSelect`|`property_single_select,omitempty`|`*WedocSingleSelectFieldProperty`| 
`PropertyReference`|`property_reference,omitempty`|`*WedocReferenceFieldProperty`| 
`PropertyLocation`|`property_location,omitempty`|`*WedocLocationFieldProperty`| 
`PropertyAutoNumber`|`property_auto_number,omitempty`|`*WedocAutoNumberFieldProperty`| 
`PropertyCurrency`|`property_currency,omitempty`|`*WedocCurrencyFieldProperty`| 
`PropertyWwGroup`|`property_ww_group,omitempty`|`*WedocWwGroupFieldProperty`| 
`PropertyPercentage`|`property_percentage,omitempty`|`*WedocPercentageFieldProperty`| 
`PropertyBarcode`|`property_barcode,omitempty`|`*WedocBarcodeFieldProperty`| 

### `WedocUpdateFieldGroupReq` 更新编组

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`FieldGroupID`|`field_group_id`|`string`| 
`Name`|`name,omitempty`|`string`| 
`Children`|`children,omitempty`|`[]WedocFieldGroupChild`| 

### `WedocUpdateFieldGroupResp` 更新编组

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldGroupID`|`field_group_id`|`string`| 
`Name`|`name`|`string`| 
`Children`|`children`|`[]WedocFieldGroupChild`| 

### `WedocUpdateFieldsReq` 更新字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`Fields`|`fields`|`[]WedocUpdateField`| 

### `WedocUpdateFieldsResp` 更新字段

Name|JSON|Type|Doc
:---|:---|:---|:--
`Fields`|`fields`|`[]WedocField`| 

### `WedocUpdatePageReq` 更新页面请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Info`|`info`|`WedocPageInfo`| 

### `WedocUpdatePageResp` 更新页面响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Info`|`info`|`WedocPageInfo`| 

### `WedocUpdatePublishSettingReq` 修改发布页可查看范围请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`PublishRange`|`publish_range,omitempty`|`*uint32`| 
`AuthList`|`auth_list,omitempty`|`[]WedocDocMember`| 

### `WedocUpdateRangeResponse` UpdateRangeResponse

Name|JSON|Type|Doc
:---|:---|:---|:--
`UpdatedCells`|`updated_cells`|`uint32`| 

### `WedocUpdateRecord` UpdateRecord

Name|JSON|Type|Doc
:---|:---|:---|:--
`RecordID`|`record_id`|`string`| 
`Values`|`values`|`map[string]any`| 

### `WedocUpdateRecordsReq` 更新记录

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`KeyType`|`key_type,omitempty`|`string`| 
`Records`|`records`|`[]WedocUpdateRecord`| 

### `WedocUpdateRecordsResp` 更新记录

Name|JSON|Type|Doc
:---|:---|:---|:--
`Records`|`records`|`[]WedocCommonRecord`| 

### `WedocUpdateSheetPrivReq` 更新智能表格子表权限请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Type`|`type`|`uint32`| 
`RuleID`|`rule_id,omitempty`|`uint32`| 
`Name`|`name,omitempty`|`string`| 
`PrivList`|`priv_list,omitempty`|`[]WedocSheetPrivItem`| 

### `WedocUpdateSheetReq` 更新子表

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Properties`|`properties`|`WedocSheetProperties`| 

### `WedocUpdateSmartsheetReq` 更新数据表请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Info`|`info`|`WedocSmartDocSmartsheetInfo`| 

### `WedocUpdateSmartsheetResp` 更新数据表响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`Info`|`info`|`WedocSmartDocSmartsheetInfo`| 

### `WedocUpdateTextProperty` 更新文本属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`TextProperty`|`text_property`|`*WedocTextProperty`| 
`Ranges`|`ranges`|`[]WedocRange`| 

### `WedocUpdateViewReq` 更新视图

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`SheetID`|`sheet_id`|`string`| 
`ViewID`|`view_id`|`string`| 
`ViewTitle`|`view_title,omitempty`|`string`| 
`Property`|`property,omitempty`|`*WedocViewProperty`| 

### `WedocUpdateViewResp` 更新视图

Name|JSON|Type|Doc
:---|:---|:---|:--
`View`|`view`|`WedocView`| 

### `WedocUploadDocImageReq` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`Base64Content`|`base64_content`|`string`| 

### `WedocUploadDocImageResp` (top)

Name|JSON|Type|Doc
:---|:---|:---|:--
`URL`|`url`|`string`| 
`Height`|`height`|`string`| 
`Width`|`width`|`string`| 
`Size`|`size`|`string`| 

### `WedocUploadFilePartReq` 分块上传文件

Name|JSON|Type|Doc
:---|:---|:---|:--
`UploadKey`|`upload_key`|`string`| 
`Index`|`index`|`int32`| 
`FileBase64Content`|`file_base64_content`|`string`| 

### `WedocUrlFieldProperty` UrlFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`| 

### `WedocUserFieldProperty` UserFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`IsMultiple`|`is_multiple`|`bool`| 
`IsNotified`|`is_notified`|`bool`| 

### `WedocUserValue` 人员类条件值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Value`|`value`|`[]string`| 

### `WedocView` View

Name|JSON|Type|Doc
:---|:---|:---|:--
`ViewID`|`view_id`|`string`| 
`ViewTitle`|`view_title`|`string`| 
`ViewType`|`view_type`|`string`| 
`Property`|`property,omitempty`|`*WedocViewProperty`| 

### `WedocViewColorCondition` ViewColorCondition

Name|JSON|Type|Doc
:---|:---|:---|:--
`ID`|`id,omitempty`|`string`| 
`Type`|`type`|`string`| 
`Color`|`color`|`string`| 
`Conditions`|`conditions`|`[]WedocCondition`| 

### `WedocViewColorConfig` ViewColorConfig

Name|JSON|Type|Doc
:---|:---|:---|:--
`Conditions`|`conditions`|`[]WedocViewColorCondition`| 

### `WedocViewProperty` ViewProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`AutoSort`|`auto_sort`|`bool`| 
`SortSpec`|`sort_spec,omitempty`|`*WedocSortSpec`| 
`GroupSpec`|`group_spec,omitempty`|`*WedocGroupSpec`| 
`FilterSpec`|`filter_spec,omitempty`|`*WedocFilterSpec`| 
`IsFieldStatEnabled`|`is_field_stat_enabled`|`bool`| 
`FieldVisibility`|`field_visibility`|`map[string]bool`| 
`FrozenFieldCount`|`frozen_field_count`|`int32`| 
`ColorConfig`|`color_config,omitempty`|`*WedocViewColorConfig`| 

### `WedocWebhookAddRecordsReq` 智能表格 Webhook 添加记录请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`AddRecords`|`add_records`|`[]WedocAddRecord`| 

### `WedocWebhookAddRecordsResp` 智能表格 Webhook 添加记录响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`AddRecords`|`add_records`|`[]WedocCommonRecord`| 

### `WedocWebhookUpdateRecordsReq` 智能表格 Webhook 更新记录请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`UpdateRecords`|`update_records`|`[]WedocUpdateRecord`| 

### `WedocWebhookUpdateRecordsResp` 智能表格 Webhook 更新记录响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`UpdateRecords`|`update_records`|`[]WedocCommonRecord`| 

### `WedocWwGroupFieldProperty` WwGroupFieldProperty

Name|JSON|Type|Doc
:---|:---|:---|:--
`AllowMultiple`|`allow_multiple`|`bool`| 

### `WedocSpreadsheetUpdateData` 编辑表格内容响应的数据

Name|JSON|Type|Doc
:---|:---|:---|:--
`Responses`|`responses`|`[]WedocSpreadsheetUpdateResponse`| 

### `WedocSheetRangeData` 获取表格数据响应的数据

Name|JSON|Type|Doc
:---|:---|:---|:--
`Result`|`result`|`WedocGridData`| 

### `WedocUploadDocFileReq` 上传文档文件请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`| 
`FileName`|`file_name`|`string`| 
`FileBase64Content`|`file_base64_content`|`string`| 

### `WedocUploadDocFileResp` 上传文档文件响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`FileURL`|`file_url`|`string`| 

# 文档 API

## API calls

Name|Request type|Response type|URL|Access token|Doc
:---|:---|:---|:---|:---|:--
`execWedocCreateDoc`|`reqWedoc`|`respWedocCreateDoc`|`POST /cgi-bin/wedoc/create_doc`|yes|新建文档
`execWedocRenameDoc`|`reqWedoc`|`respWedocRenameDoc`|`POST /cgi-bin/wedoc/rename_doc`|yes|重命名文档
`execWedocDeleteDoc`|`reqWedoc`|`respWedocDeleteDoc`|`POST /cgi-bin/wedoc/del_doc`|yes|删除文档
`execWedocGetDocBaseInfo`|`reqWedoc`|`respWedocGetDocBaseInfo`|`POST /cgi-bin/wedoc/get_doc_base_info`|yes|获取文档基础信息
`execWedocShareDoc`|`reqWedoc`|`respWedocShareDoc`|`POST /cgi-bin/wedoc/doc_share`|yes|分享文档
`execWedocBatchUpdateDocument`|`reqWedoc`|`respWedocBatchUpdateDocument`|`POST /cgi-bin/wedoc/document/batch_update`|yes|编辑文档内容
`execWedocGetDocument`|`reqWedoc`|`respWedocGetDocument`|`POST /cgi-bin/wedoc/document/get`|yes|获取文档数据
`execWedocBatchUpdateSpreadsheet`|`reqWedoc`|`respWedocBatchUpdateSpreadsheet`|`POST /cgi-bin/wedoc/spreadsheet/batch_update`|yes|编辑表格内容
`execWedocGetSheetProperties`|`reqWedoc`|`respWedocGetSheetProperties`|`POST /cgi-bin/wedoc/spreadsheet/get_sheet_properties`|yes|获取表格行列信息
`execWedocGetSheetRangeData`|`reqWedoc`|`respWedocGetSheetRangeData`|`POST /cgi-bin/wedoc/spreadsheet/get_sheet_range_data`|yes|获取表格数据
`execWedocAddSheet`|`reqWedoc`|`respWedocAddSheet`|`POST /cgi-bin/wedoc/smartsheet/add_sheet`|yes|添加子表
`execWedocDeleteSheet`|`reqWedoc`|`respWedocDeleteSheet`|`POST /cgi-bin/wedoc/smartsheet/delete_sheet`|yes|删除子表
`execWedocUpdateSheet`|`reqWedoc`|`respWedocUpdateSheet`|`POST /cgi-bin/wedoc/smartsheet/update_sheet`|yes|更新子表
`execWedocGetSheet`|`reqWedoc`|`respWedocGetSheet`|`POST /cgi-bin/wedoc/smartsheet/get_sheet`|yes|查询子表
`execWedocAddView`|`reqWedoc`|`respWedocAddView`|`POST /cgi-bin/wedoc/smartsheet/add_view`|yes|添加视图
`execWedocDeleteViews`|`reqWedoc`|`respWedocDeleteViews`|`POST /cgi-bin/wedoc/smartsheet/delete_views`|yes|删除视图
`execWedocUpdateView`|`reqWedoc`|`respWedocUpdateView`|`POST /cgi-bin/wedoc/smartsheet/update_view`|yes|更新视图
`execWedocGetViews`|`reqWedoc`|`respWedocGetViews`|`POST /cgi-bin/wedoc/smartsheet/get_views`|yes|查询视图
`execWedocAddFields`|`reqWedoc`|`respWedocAddFields`|`POST /cgi-bin/wedoc/smartsheet/add_fields`|yes|添加字段
`execWedocDeleteFields`|`reqWedoc`|`respWedocDeleteFields`|`POST /cgi-bin/wedoc/smartsheet/delete_fields`|yes|删除字段
`execWedocUpdateFields`|`reqWedoc`|`respWedocUpdateFields`|`POST /cgi-bin/wedoc/smartsheet/update_fields`|yes|更新字段
`execWedocGetFields`|`reqWedoc`|`respWedocGetFields`|`POST /cgi-bin/wedoc/smartsheet/get_fields`|yes|查询字段
`execWedocAddRecords`|`reqWedoc`|`respWedocAddRecords`|`POST /cgi-bin/wedoc/smartsheet/add_records`|yes|添加记录
`execWedocDeleteRecords`|`reqWedoc`|`respWedocDeleteRecords`|`POST /cgi-bin/wedoc/smartsheet/delete_records`|yes|删除记录
`execWedocUpdateRecords`|`reqWedoc`|`respWedocUpdateRecords`|`POST /cgi-bin/wedoc/smartsheet/update_records`|yes|更新记录
`execWedocGetRecords`|`reqWedoc`|`respWedocGetRecords`|`POST /cgi-bin/wedoc/smartsheet/get_records`|yes|查询记录
`execWedocAddFieldGroup`|`reqWedoc`|`respWedocAddFieldGroup`|`POST /cgi-bin/wedoc/smartsheet/add_field_group`|yes|添加编组
`execWedocDeleteFieldGroups`|`reqWedoc`|`respWedocDeleteFieldGroups`|`POST /cgi-bin/wedoc/smartsheet/delete_field_groups`|yes|删除编组
`execWedocUpdateFieldGroup`|`reqWedoc`|`respWedocUpdateFieldGroup`|`POST /cgi-bin/wedoc/smartsheet/update_field_group`|yes|更新编组
`execWedocGetFieldGroups`|`reqWedoc`|`respWedocGetFieldGroups`|`POST /cgi-bin/wedoc/smartsheet/get_field_groups`|yes|获取编组
`execWedocAddPage`|`reqWedoc`|`respWedocAddPage`|`POST /cgi-bin/wedoc/smartdoc/add_page`|yes|添加页面
`execWedocUpdatePage`|`reqWedoc`|`respWedocUpdatePage`|`POST /cgi-bin/wedoc/smartdoc/update_page`|yes|更新页面
`execWedocDeletePage`|`reqWedoc`|`respWedocDeletePage`|`POST /cgi-bin/wedoc/smartdoc/delete_page`|yes|删除页面
`execWedocGetPageHierarchy`|`reqWedoc`|`respWedocGetPageHierarchy`|`POST /cgi-bin/wedoc/smartdoc/get_page_hierarchy`|yes|获取页面结构
`execWedocAddBlocks`|`reqWedoc`|`respWedocAddBlocks`|`POST /cgi-bin/wedoc/smartdoc/add_blocks`|yes|添加内容块
`execWedocUpdateBlocks`|`reqWedoc`|`respWedocUpdateBlocks`|`POST /cgi-bin/wedoc/smartdoc/update_blocks`|yes|更新内容块
`execWedocDeleteBlocks`|`reqWedoc`|`respWedocDeleteBlocks`|`POST /cgi-bin/wedoc/smartdoc/delete_blocks`|yes|删除内容块
`execWedocGetBlockList`|`reqWedoc`|`respWedocGetBlockList`|`POST /cgi-bin/wedoc/smartdoc/get_block_list`|yes|获取内容块列表
`execWedocExportTask`|`reqWedoc`|`respWedocExportTask`|`POST /cgi-bin/wedoc/smartdoc/export_task`|yes|提交导出任务
`execWedocGetExportResult`|`reqWedoc`|`respWedocGetExportResult`|`POST /cgi-bin/wedoc/smartdoc/get_export_result`|yes|查询任务结果
`execWedocGetSmartsheetInfo`|`reqWedoc`|`respWedocGetSmartsheetInfo`|`POST /cgi-bin/wedoc/smartdoc/get_smartsheet_info`|yes|获取数据源
`execWedocAddSmartsheet`|`reqWedoc`|`respWedocAddSmartsheet`|`POST /cgi-bin/wedoc/smartdoc/add_smartsheet`|yes|添加数据表
`execWedocDeleteSmartsheet`|`reqWedoc`|`respWedocDeleteSmartsheet`|`POST /cgi-bin/wedoc/smartdoc/delete_smartsheet`|yes|删除数据表
`execWedocUpdateSmartsheet`|`reqWedoc`|`respWedocUpdateSmartsheet`|`POST /cgi-bin/wedoc/smartdoc/update_smartsheet`|yes|更新数据表
`execWedocPublishSmartDoc`|`reqWedoc`|`respWedocPublishSmartDoc`|`POST /cgi-bin/wedoc/smartdoc/publish`|yes|发布智能文档
`execWedocCancelPublishSmartDoc`|`reqWedoc`|`respWedocCancelPublishSmartDoc`|`POST /cgi-bin/wedoc/smartdoc/cancel_publish`|yes|取消发布智能文档
`execWedocUpdatePublishSetting`|`reqWedoc`|`respWedocUpdatePublishSetting`|`POST /cgi-bin/wedoc/smartdoc/publish_setting`|yes|修改发布页可查看范围
`execWedocGetDocAuth`|`reqWedoc`|`respWedocGetDocAuth`|`POST /cgi-bin/wedoc/doc_get_auth`|yes|获取文档权限信息
`execWedocModifyDocJoinRule`|`reqWedoc`|`respWedocModifyDocJoinRule`|`POST /cgi-bin/wedoc/mod_doc_join_rule`|yes|修改文档加入规则
`execWedocModifyDocMember`|`reqWedoc`|`respWedocModifyDocMember`|`POST /cgi-bin/wedoc/mod_doc_member`|yes|修改文档成员与权限
`execWedocModifyDocSafetySetting`|`reqWedoc`|`respWedocModifyDocSafetySetting`|`POST /cgi-bin/wedoc/mod_doc_safty_setting`|yes|修改文档安全设置
`execWedocGetSheetPriv`|`reqWedoc`|`respWedocGetSheetPriv`|`POST /cgi-bin/wedoc/smartsheet/content_priv/get_sheet_priv`|yes|查询智能表格子表权限
`execWedocUpdateSheetPriv`|`reqWedoc`|`respWedocUpdateSheetPriv`|`POST /cgi-bin/wedoc/smartsheet/content_priv/update_sheet_priv`|yes|更新智能表格子表权限
`execWedocCreateSheetPrivRule`|`reqWedoc`|`respWedocCreateSheetPrivRule`|`POST /cgi-bin/wedoc/smartsheet/content_priv/create_rule`|yes|新增智能表格指定成员额外权限
`execWedocModifySheetPrivRuleMember`|`reqWedoc`|`respWedocModifySheetPrivRuleMember`|`POST /cgi-bin/wedoc/smartsheet/content_priv/mod_rule_member`|yes|更新智能表格指定成员额外权限
`execWedocDeleteSheetPrivRule`|`reqWedoc`|`respWedocDeleteSheetPrivRule`|`POST /cgi-bin/wedoc/smartsheet/content_priv/delete_rule`|yes|删除智能表格指定成员额外权限
`execWedocCreateForm`|`reqWedoc`|`respWedocCreateForm`|`POST /cgi-bin/wedoc/create_form`|yes|创建收集表
`execWedocModifyForm`|`reqWedoc`|`respWedocModifyForm`|`POST /cgi-bin/wedoc/modify_form`|yes|编辑收集表
`execWedocGetFormInfo`|`reqWedoc`|`respWedocGetFormInfo`|`POST /cgi-bin/wedoc/get_form_info`|yes|获取收集表信息
`execWedocGetFormStatistic`|`reqWedoc`|`respWedocGetFormStatistic`|`POST /cgi-bin/wedoc/get_form_statistic`|yes|收集表的统计信息查询
`execWedocGetFormAnswer`|`reqWedoc`|`respWedocGetFormAnswer`|`POST /cgi-bin/wedoc/get_form_answer`|yes|读取收集表答案
`execWedocWebhookAddRecords`|`reqWedocWebhook`|`respWedocWebhookAddRecords`|`POST /cgi-bin/wedoc/smartsheet/webhook`|no|添加记录
`execWedocWebhookUpdateRecords`|`reqWedocWebhook`|`respWedocWebhookUpdateRecords`|`POST /cgi-bin/wedoc/smartsheet/webhook`|no|更新记录
`execWedocAddVipAccounts`|`reqWedoc`|`respWedocAddVipAccounts`|`POST /cgi-bin/wedoc/vip/batch_add`|yes|分配高级功能账号
`execWedocDeleteVipAccounts`|`reqWedoc`|`respWedocDeleteVipAccounts`|`POST /cgi-bin/wedoc/vip/batch_del`|yes|取消高级功能账号
`execWedocListVipAccounts`|`reqWedoc`|`respWedocListVipAccounts`|`POST /cgi-bin/wedoc/vip/list`|yes|获取高级功能账号列表
`execWedocUploadDocImage`|`reqWedoc`|`respWedocUploadDocImage`|`POST /cgi-bin/wedoc/image_upload`|yes|上传文档图片
`execWedocUploadDocFile`|`reqWedoc`|`respWedocUploadDocFile`|`POST /cgi-bin/wedoc/file_upload`|yes|上传文档文件
`execWedocInitFileUpload`|`reqWedoc`|`respWedocInitFileUpload`|`POST /cgi-bin/wedoc/file_upload_init`|yes|分块上传初始化
`execWedocUploadFilePart`|`reqWedoc`|`respWedocUploadFilePart`|`POST /cgi-bin/wedoc/file_upload_part`|yes|分块上传文件
`execWedocFinishFileUpload`|`reqWedoc`|`respWedocFinishFileUpload`|`POST /cgi-bin/wedoc/file_upload_finish`|yes|分块上传完成

