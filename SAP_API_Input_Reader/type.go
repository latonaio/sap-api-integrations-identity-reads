package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey               string `json:"connection_key"`
	Result                      bool   `json:"result"`
	RedisKey                    string `json:"redis_key"`
	Filepath                    string `json:"filepath"`
	IdentityGetLoggedInUserInfo struct {
		ObjectID                  string `json:"ObjectID"`
		UserID                    string `json:"UserID"`
		UserName                  string `json:"UserName"`
		UserAccountID             string `json:"UserAccountID"`
		UUID                      string `json:"UUID"`
		EmployeeID                string `json:"EmployeeID"`
		BusinessPartnerID         string `json:"BusinessPartnerID"`
		EmployeeUUID              string `json:"EmployeeUUID"`
		Email                     string `json:"Email"`
		DateFormatCode            string `json:"DateFormatCode"`
		DateFormatCodeText        string `json:"DateFormatCodeText"`
		DecimalFormatCode         string `json:"DecimalFormatCode"`
		DecimalFormatCodeText     string `json:"DecimalFormatCodeText"`
		LogonLanguageCode         string `json:"LogonLanguageCode"`
		LogonLanguageCodeText     string `json:"LogonLanguageCodeText"`
		TimeFormatCode            string `json:"TimeFormatCode"`
		TimeFormatCodeText        string `json:"TimeFormatCodeText"`
		TimeZoneCode              string `json:"TimeZoneCode"`
		TimeZoneCodeText          string `json:"TimeZoneCodeText"`
		TechnicalUserIndicator    bool   `json:"TechnicalUserIndicator"`
		KeyUserIndicator          bool   `json:"KeyUserIndicator"`
		InactiveIndicator         bool   `json:"InactiveIndicator"`
		PasswordInactiveIndicator bool   `json:"PasswordInactiveIndicator"`
		PasswordLockedIndicator   bool   `json:"PasswordLockedIndicator"`
		PasswordPolicyCode        string `json:"PasswordPolicyCode"`
		PasswordPolicyCodeText    string `json:"PasswordPolicyCodeText"`
		UserAccountTypeCode       string `json:"UserAccountTypeCode"`
		UserAccountTypeCodeText   string `json:"UserAccountTypeCodeText"`
		StatusCode                string `json:"StatusCode"`
		StatusCodeText            string `json:"StatusCodeText"`
		CreatedOn                 string `json:"CreatedOn"`
		CreatedBy                 string `json:"CreatedBy"`
		ChangedOn                 string `json:"ChangedOn"`
		ChangedBy                 string `json:"ChangedBy"`
		CreatedByUUID             string `json:"CreatedByUUID"`
		ChangedByUUID             string `json:"ChangedByUUID"`
		EntityLastChangedOn       string `json:"EntityLastChangedOn"`
	} `json:"IdentityGetLoggedInUserInfo"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	IdentityCode string   `json:"identity_code"`
	Deleted      bool     `json:"deleted"`
}
