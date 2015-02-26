//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated
// Feb 25, 2015 at 5:09pm (PST)
// Command:
// $ ./api15gen -metadata=../../rsapi15 -output=../../rsapi15
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

// Consists of a map of resource name to resource metadata.
var api_metadata = map[string]*metadata.Resource{
	"Account": &metadata.Resource{
		Name:        "Account",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/accounts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/accounts/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"AccountGroup": &metadata.Resource{
		Name:        "AccountGroup",
		Description: `An Account Group specifies which RightScale accounts will have access to import a shared RightScale component (e.g. ServerTemplate, RightScript, etc.) from the MultiCloud Marketplace.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists the AccountGroups owned by this Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/account_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/account_groups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single AccountGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/account_groups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/account_groups/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Alert": &metadata.Resource{
		Name:        "Alert",
		Description: `An Alert represents an AlertSpec bound to a running Instance.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "disable",
				Description: `Disables the Alert indefinitely. Idempotent.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/alerts/:id/disable",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)/disable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:server_id/alerts/:id/disable",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)/disable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:server_array_id/alerts/:id/disable",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)/disable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:deployment_id/alerts/:id/disable",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)/disable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/alerts/:id/disable",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)/disable$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "enable",
				Description: `Enables the Alert indefinitely. Idempotent.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/alerts/:id/enable",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)/enable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:server_id/alerts/:id/enable",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)/enable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:server_array_id/alerts/:id/enable",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)/enable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:deployment_id/alerts/:id/enable",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)/enable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/alerts/:id/enable",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)/enable$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists all Alerts.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/alerts",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers/:server_id/alerts",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:server_array_id/alerts",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:deployment_id/alerts",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/alerts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/alerts$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "quench",
				Description: `Suppresses the Alert from being triggered for a given time period. Idempotent.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/alerts/:id/quench",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)/quench$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:server_id/alerts/:id/quench",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)/quench$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:server_array_id/alerts/:id/quench",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)/quench$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:deployment_id/alerts/:id/quench",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)/quench$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/alerts/:id/quench",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)/quench$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "duration",
						Description: `The time period in seconds to suppress Alert from being triggered.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"duration"},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows the attributes of a specified Alert.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/alerts/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers/:server_id/alerts/:id",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:server_array_id/alerts/:id",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:deployment_id/alerts/:id",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/alerts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/alerts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alerts/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"AlertSpec": &metadata.Resource{
		Name:        "AlertSpec",
		Description: `An AlertSpec defines the conditions under which an Alert is triggered and escalated.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new AlertSpec with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:server_id/alert_specs",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:server_array_id/alert_specs",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:server_template_id/alert_specs",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/alert_specs",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/alert_specs$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "alert_spec[escalation_name]",
						Description: `Escalate to the named alert escalation when the alert is triggered. Must either escalate or vote.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[subject_href]",
						Description: `The href of the resource that this AlertSpec should be associated with. The subject can be a ServerTemplate, Server, ServerArray, or Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[description]",
						Description: `The description of the AlertSpec.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[threshold]",
						Description: `The threshold of the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[condition]",
						Description: `The condition (operator) in the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{">", ">=", "<", "<=", "==", "!="},
					},
					&metadata.ActionParam{
						Name:        "alert_spec[vote_type]",
						Description: `Vote to grow or shrink a ServerArray when the alert is triggered. Must either escalate or vote.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"grow", "shrink"},
					},
					&metadata.ActionParam{
						Name:        "alert_spec[variable]",
						Description: `The RRD variable of the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[duration]",
						Description: `The duration in minutes of the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[vote_tag]",
						Description: `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[name]",
						Description: `The name of the AlertSpec.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[file]",
						Description: `The RRD path/file_name of the condition sentence.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"alert_spec"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given AlertSpec.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/servers/:server_id/alert_specs/:id",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/server_arrays/:server_array_id/alert_specs/:id",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/server_templates/:server_template_id/alert_specs/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/alert_specs/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alert_specs/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `No description provided for index.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers/:server_id/alert_specs",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:server_array_id/alert_specs",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/alert_specs",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/alert_specs",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/alert_specs$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "with_inherited",
						Description: `Flag indicating whether or not to include AlertSpecs from the ServerTemplate in the index.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `No description provided for show.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers/:server_id/alert_specs/:id",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:server_array_id/alert_specs/:id",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/alert_specs/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/alert_specs/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alert_specs/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates an AlertSpec with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/servers/:server_id/alert_specs/:id",
						Variables:  []string{"server_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/server_arrays/:server_array_id/alert_specs/:id",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/server_templates/:server_template_id/alert_specs/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/alert_specs/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/alert_specs/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/alert_specs/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "alert_spec[escalation_name]",
						Description: `Escalate to the named alert escalation when the alert is triggered.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[description]",
						Description: `The description of the AlertSpec.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[threshold]",
						Description: `The threshold of the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[vote_type]",
						Description: `Vote to grow or shrink a ServerArray when the alert is triggered.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"grow", "shrink"},
					},
					&metadata.ActionParam{
						Name:        "alert_spec[condition]",
						Description: `The condition (operator) in the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{">", ">=", "<", "<=", "==", "!="},
					},
					&metadata.ActionParam{
						Name:        "alert_spec[variable]",
						Description: `The RRD variable of the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[duration]",
						Description: `The duration in minutes of the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[vote_tag]",
						Description: `Should correspond to a vote tag on a ServerArray if vote to grow or shrink.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[file]",
						Description: `The RRD path/file_name of the condition sentence.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "alert_spec[name]",
						Description: `The name of the AlertSpec.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"alert_spec"},
			},
		},
	},
	"AuditEntry": &metadata.Resource{
		Name:        "AuditEntry",
		Description: `An Audit Entry can be used to track various activities of a resource.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "append",
				Description: `Updates the summary and appends more details to a given AuditEntry`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/audit_entries/:id/append",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/audit_entries/([^/]+)/append$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "summary",
						Description: `The updated summary for the audit entry, maximum length is 255 characters.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "detail",
						Description: `The details to be appended to the audit entry record.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "offset",
						Description: `The offset where the new details should be appended to in the audit entry's existing details section. Also used in ordering of summary updates. Defaults to end.`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"detail", "notify", "offset", "summary"},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new AuditEntry with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/audit_entries",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/audit_entries$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "audit_entry[auditee_href]",
						Description: `The href of the resource that this audit entry should be associated with (e.g. an instance's href).`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_entry[summary]",
						Description: `The summary of the audit entry to be created, maximum length is 255 characters.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_entry[detail]",
						Description: `The initial details of the audit entry to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "user_email",
						Description: `The email of the user (who created/triggered the audit entry). Only usable with instance role.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"audit_entry", "notify", "user_email"},
			},

			&metadata.Action{
				Name:        "detail",
				Description: `shows the details of a given AuditEntry.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/audit_entries/:id/detail",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/audit_entries/([^/]+)/detail$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists AuditEntries of the account`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/audit_entries",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/audit_entries$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "start_date",
						Description: `The start date for retrieving audit entries, the format must be YYYY/MM/DD HH:MM:SS [+/-]ZZZZ e.g., 2011/06/25 00:00:00 +0000`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "end_date",
						Description: `The end date for retrieving audit entries (the format must be the same as start date). The time period between start and end date must be less than 3 months (93 days).`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "limit",
						Description: `Limit the audit entries to this number. The limit should >= 1 and <= 1000`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Lists the attributes of a given audit entry.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/audit_entries/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/audit_entries/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates the summary of a given AuditEntry.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/audit_entries/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/audit_entries/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "audit_entry[summary]",
						Description: `The updated summary for the audit entry, maximum length is 255 characters.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "audit_entry[offset]",
						Description: `The offset where the next details will be appended. Used in ordering of summary updates.`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "notify",
						Description: `The event notification category. Defaults to 'None'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"audit_entry", "notify"},
			},
		},
	},
	"Backup": &metadata.Resource{
		Name:        "Backup",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "cleanup",
				Description: `Deletes old backups that meet the given criteria`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/backups/cleanup",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/backups/cleanup$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cloud_href",
						Description: `Backups belonging to only this cloud are considered for cleanup. Otherwise, all backups in the account with the same lineage will be considered.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "keep_last",
						Description: `The number of backups that should be kept.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "monthlies",
						Description: `The number of monthly backups(the latest one in each month) that should be kept.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "weeklies",
						Description: `The number of weekly backups(the latest one in each week) that should be kept.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "yearlies",
						Description: `The number of yearly backups(the latest one in each year) that should be kept.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "dailies",
						Description: `The number of daily backups(the latest one in each day) that should be kept.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "lineage",
						Description: `The lineage of the backups that are to be cleaned-up.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"cloud_href", "dailies", "keep_last", "lineage", "monthlies", "weeklies", "yearlies"},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Takes in an array of volumeattachmenthrefs and takes a snapshot of each.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/backups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/backups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup[volume_attachment_hrefs][]",
						Description: `List of volume attachment hrefs that are to be backed-up.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[description]",
						Description: `The description to be set on each of the volume snapshots`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[from_master]",
						Description: `Setting this to 'true' will create a tag 'rs_backup:from_master=true' on the snapshots so that one can filter them later.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "backup[lineage]",
						Description: `A unique value to create backups belonging to a particular system.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[name]",
						Description: `The name to be set on each of the volume snapshots.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"backup"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given backup by deleting all of its snapshots, this call will succeed even if the backup has not completed.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/backups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/backups/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists all of the backups with the given lineage tag`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/backups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/backups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "lineage",
						Description: `Backups belonging to this lineage.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "restore",
				Description: `Restores the given Backup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/backups/:id/restore",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/backups/([^/]+)/restore$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup[volume_type_href]",
						Description: `The href of the volume type. Each volume is created with this volume type instead of the default volume type for the cloud. A Name, Resource UID and optional Size is associated with a volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[description]",
						Description: `Each volume is created with this description instead of the volume snapshot's description`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance_href",
						Description: `The instance href that the backup will be restored to.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[iops]",
						Description: `The number of IOPS (I/O Operations Per Second) each volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[name]",
						Description: `Each volume is created with this name instead of the volume snapshot's name`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "backup[size]",
						Description: `Each volume is created with this size in gigabytes (GB) instead of the volume snapshot's size (must be equal or larger). Some volume types have predefined sizes and do not allow selecting a custom size on volume creation.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"backup", "instance_href"},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Lists the attributes of a given backup`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/backups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/backups/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates the committed tag for all of the VolumeSnapshots in the given Backup to the given value.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/backups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/backups/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "backup[committed]",
						Description: `Setting this to 'true' will update the 'rs_backup:committed=false' tag to 'rs_backup:committed=true' on all the snapshots.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"backup"},
			},
		},
	},
	"ChildAccount": &metadata.Resource{
		Name:        "ChildAccount",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create an enterprise ChildAccount for this Account`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/child_accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/child_accounts$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "child_account[cluster_href]",
						Description: `The href of the cluster in which to create the account. If not specified, will default to the cluster of the parent account.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "child_account[name]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"child_account"},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the enterprise ChildAccounts available for this Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/child_accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/child_accounts$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update an enterprise ChildAccount for this Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/accounts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/accounts/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/child_accounts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/child_accounts/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "child_account[name]",
						Description: `The updated name for the account.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"child_account"},
			},
		},
	},
	"Cloud": &metadata.Resource{
		Name:        "Cloud",
		Description: `Represents a Cloud (within the context of the account in the session).`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists the clouds available to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/clouds$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single cloud`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"CloudAccount": &metadata.Resource{
		Name:        "CloudAccount",
		Description: `Represents a Cloud Account (An association between the account and a cloud).`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a CloudAccount by passing in the respective credentials for each cloud.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/cloud_accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cloud_accounts$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cloud_account[cloud_href]",
						Description: `The href of the cloud if it is known. For valid values see support portal link above.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "cloud_account[creds]",
						Description: ``,
						Type:        "map",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "cloud_account[token]",
						Description: `The cloud token to identify a private cloud`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"cloud_account"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete a CloudAccount.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/cloud_accounts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cloud_accounts/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the CloudAccounts (non-aws) available to this Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/cloud_accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cloud_accounts$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/cloud_accounts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cloud_accounts/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"Cookbook": &metadata.Resource{
		Name:        "Cookbook",
		Description: `Represents a given instance of a single cookbook.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "destroy",
				Description: `Destroys a Cookbook. Only available for cookbooks that have no Cookbook Attachments.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/cookbooks/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "follow",
				Description: `Follows (or unfollows) a Cookbook. Only available for cookbooks that are in the Alternate namespace.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/cookbooks/:id/follow",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/follow$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should follow (true) or unfollow (false) a Cookbook.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"value"},
			},

			&metadata.Action{
				Name:        "freeze",
				Description: `Freezes (or unfreezes) a Cookbook. Only available for cookbooks that are in the Primary namespace.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/cookbooks/:id/freeze",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/freeze$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should freeze (true) or unfreeze (false) a Cookbook.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"value"},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the Cookbooks available to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/cookbooks",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbooks$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "extended_designer"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "obsolete",
				Description: `Marks a Cookbook as obsolete (or un-obsolete).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/cookbooks/:id/obsolete",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/obsolete$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "value",
						Description: `Indicates if this action should obsolete (true) or un-obsolete (false) a Cookbook.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"value"},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single Cookbook.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/cookbooks/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "extended_designer"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"CookbookAttachment": &metadata.Resource{
		Name:        "CookbookAttachment",
		Description: `Cookbook Attachment is used to associate a particular cookbook with a ServerTemplate. A Cookbook Attachment must be in place before a recipe can be bound to a runlist using RunnableBinding.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Attach a cookbook to a given resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/cookbooks/:cookbook_id/cookbook_attachments",
						Variables:  []string{"cookbook_id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/cookbook_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:server_template_id/cookbook_attachments",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/cookbook_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachment[server_template_href]",
						Description: `The href of the server template to attach the cookbook to.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "cookbook_attachment[cookbook_href]",
						Description: `The href of the cookbook to attach.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"cookbook_attachment"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Detach a cookbook from a given resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/cookbooks/:cookbook_id/cookbook_attachments/:id",
						Variables:  []string{"cookbook_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/cookbook_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/server_templates/:server_template_id/cookbook_attachments/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/cookbook_attachments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists Cookbook Attachments.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/cookbooks/:cookbook_id/cookbook_attachments",
						Variables:  []string{"cookbook_id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/cookbook_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/cookbook_attachments",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/cookbook_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "multi_attach",
				Description: `Attach multiple cookbooks to a given resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:server_template_id/cookbook_attachments/multi_attach",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments/multi_attach$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/cookbook_attachments/multi_attach",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments/multi_attach$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachments[server_template_href]",
						Description: `The href of the server template to attach the cookbooks to.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "cookbook_attachments[cookbook_hrefs][]",
						Description: `The hrefs of the cookbooks to be attached`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"cookbook_attachments"},
			},

			&metadata.Action{
				Name:        "multi_detach",
				Description: `Detach multiple cookbooks from a given resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:server_template_id/cookbook_attachments/multi_detach",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments/multi_detach$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/cookbook_attachments/multi_detach",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments/multi_detach$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "cookbook_attachments[cookbook_attachment_hrefs][]",
						Description: `The hrefs of the cookbook attachments to be detached`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"cookbook_attachments"},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single cookbook attachment to a ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/cookbooks/:cookbook_id/cookbook_attachments/:id",
						Variables:  []string{"cookbook_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/cookbooks/([^/]+)/cookbook_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/cookbook_attachments/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/cookbook_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/cookbook_attachments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/cookbook_attachments/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Credential": &metadata.Resource{
		Name:        "Credential",
		Description: `A Credential provides an abstraction for sensitive textual information, such as passphrases or cloud credentials, whose value should be encrypted when stored in the database and not generally shown in the UI or through the API...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new Credential with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/credentials",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/credentials$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credential[description]",
						Description: `The description of the Credential to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "credential[value]",
						Description: `The value of the Credential to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "credential[name]",
						Description: `The name of the Credential to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"credential"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a Credential.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/credentials/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/credentials/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the Credentials available to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/credentials",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/credentials$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single Credential. NOTE: Credential values may be updated through the API, but values cannot be retrieved via the API.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/credentials/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/credentials/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates attributes of a Credential.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/credentials/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/credentials/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "credential[description]",
						Description: `The updated description of the Credential.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "credential[value]",
						Description: `The updated value of the Credential.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "credential[name]",
						Description: `The updated name of the Credential.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"credential"},
			},
		},
	},
	"Datacenter": &metadata.Resource{
		Name:        "Datacenter",
		Description: `Datacenters represent isolated facilities within a cloud`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Datacenters for a particular cloud.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/datacenters",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/datacenters$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single Datacenter.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/datacenters/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/datacenters/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Deployment": &metadata.Resource{
		Name:        "Deployment",
		Description: `Deployments represent logical groupings of related assets such as servers, server arrays, default configuration settings...etc.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "clone",
				Description: `Clones a given deployment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:id/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/clone$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the cloned deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&metadata.ActionParam{
						Name:        "deployment[description]",
						Description: `The description for the cloned deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "deployment[name]",
						Description: `The name for the cloned deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"deployment"},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new deployment with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/deployments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&metadata.ActionParam{
						Name:        "deployment[description]",
						Description: `The description of the deployment to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "deployment[name]",
						Description: `The name of the deployment to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"deployment"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given deployment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/deployments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists deployments of the account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/deployments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "lock",
				Description: `Locks a given deployment. Idempotent.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:id/lock",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/lock$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "servers",
				Description: `Lists the servers belonging to this deployment`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:id/servers",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Lists the attributes of a given deployment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "unlock",
				Description: `Unlocks a given deployment. Idempotent.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:id/unlock",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/unlock$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates attributes of a given deployment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/deployments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "deployment[server_tag_scope]",
						Description: `The routing scope for tags for servers in the deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
						ValidValues: []string{"deployment", "account"},
					},
					&metadata.ActionParam{
						Name:        "deployment[description]",
						Description: `The updated description for the deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "deployment[name]",
						Description: `The updated name for the deployment.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"deployment"},
			},
		},
	},
	"HealthCheck": &metadata.Resource{
		Name:        "HealthCheck",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Check health of RightApi controllers`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/health-check/",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/health-check/$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"IdentityProvider": &metadata.Resource{
		Name:        "IdentityProvider",
		Description: `An Identity Provider represents a SAML identity provider (IdP) that is linked to your RightScale Enterprise account, and is trusted by the RightScale dashboard to authenticate your enterprise's end users...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists the identity providers associated with this enterprise account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/identity_providers",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/identity_providers$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show the specified identity provider, if associated with this enterprise account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/identity_providers/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/identity_providers/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Image": &metadata.Resource{
		Name:        "Image",
		Description: `Images represent base VM image existing in a cloud`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists all Images for the given Cloud.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/images",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/images$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows information about a single Image.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/images/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/images/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Input": &metadata.Resource{
		Name:        "Input",
		Description: `Inputs help extract dynamic information, usually specified at runtime, from repeatable configuration operations that can be codified.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Retrieves the full list of existing inputs of the specified resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/inputs",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/inputs$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:deployment_id/inputs",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/inputs$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/inputs",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/inputs$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs_2_0"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "multi_update",
				Description: `Performs a bulk update of inputs on the specified resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/inputs/multi_update",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/inputs/multi_update$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/deployments/:deployment_id/inputs/multi_update",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/inputs/multi_update$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/server_templates/:server_template_id/inputs/multi_update",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/inputs/multi_update$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value to be updated with. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The name of the input to be updated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"inputs"},
			},
		},
	},
	"Instance": &metadata.Resource{
		Name:        "Instance",
		Description: `Instances represent an entity that is runnable in the cloud.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates and launches a raw instance using the provided parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][ebs_optimized]",
						Description: `Whether the instance is able to connect to IOPS-enabled volumes.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[placement_group_href]",
						Description: `The placement group to launch the instance in. Not supported by all clouds & instance types.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[security_group_hrefs][]",
						Description: `The hrefs of the security groups.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[instance_type_href]",
						Description: `The href of the instance type.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ramdisk_image_href]",
						Description: `The href of the ramdisk image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[kernel_image_href]",
						Description: `The href of the kernel image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[deployment_href]",
						Description: `The href of the deployment to which the Instance will be added.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ssh_key_href]",
						Description: `The href of the SSH key to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[name]",
						Description: `The name of the instance.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"instance"},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists instances of a given cloud, server array.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:server_array_id/current_instances",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/current_instances$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "launch",
				Description: `Launches an instance using the parameters that this instance has been configured with.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/launch",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/launch$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:server_id/launch",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/launch$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:server_array_id/launch",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/launch$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of that input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The input name. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "api_behavior",
						Description: `When set to 'async', an instance resource will be returned immediately and processing will be handled in the background. Errors will not be returned and must be checked through the instance's audit entries. Default value is 'sync'`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"async", "sync"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"api_behavior", "inputs"},
			},

			&metadata.Action{
				Name:        "lock",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/lock",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/lock$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "multi_run_executable",
				Description: `Runs a script or a recipe in the running instances.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/multi_run_executable",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/multi_run_executable$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:server_array_id/multi_run_executable",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/multi_run_executable$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of these inputs. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The name of inputs needed. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock(s) on the Instance(s).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "recipe_name",
						Description: `The name of the recipe to be run.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{"ignore_lock", "inputs", "recipe_name", "right_script_href"},
			},

			&metadata.Action{
				Name:        "multi_terminate",
				Description: `Terminates running instances.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/multi_terminate",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/multi_terminate$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:server_array_id/multi_terminate",
						Variables:  []string{"server_array_id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/multi_terminate$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "terminate_all",
						Description: `Specifies the ability to terminate all instances.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{"terminate_all"},
			},

			&metadata.Action{
				Name:        "reboot",
				Description: `Reboot a running instance.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/reboot",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/reboot$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:server_id/reboot",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/reboot$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "run_executable",
				Description: `Runs a script or a recipe in the running instance.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/run_executable",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/run_executable$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script_href",
						Description: `The href of the RightScript to run. Should be of the form '/api/right_scripts/:id'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][value]",
						Description: `The value of these inputs. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "inputs[][name]",
						Description: `The name of inputs needed. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ignore_lock",
						Description: `Specifies the ability to ignore the lock on the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "recipe_name",
						Description: `The name of the recipe to run.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"ignore_lock", "inputs", "recipe_name", "right_script_href"},
			},

			&metadata.Action{
				Name:        "set_custom_lodgement",
				Description: `This method is deprecated.  Please use InstanceCustomLodgement.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/set_custom_lodgement",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/set_custom_lodgement$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "quantity[][value]",
						Description: `The value of the quantity. Should be a positive integer.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "quantity[][name]",
						Description: `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "timeframe",
						Description: `The timeframe (either a month or a single day) for which the quantity value is valid (currently for the PDT timezone only).`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"quantity", "timeframe"},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows attributes of a single instance.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended", "full", "full_inputs_2_0", "tiny"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "start",
				Description: `Starts an instance that has been stopped, resuming it to its previously saved volume state.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/start",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/start$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "stop",
				Description: `Stores the instance's current volume state to resume later using the 'start' action.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/stop",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/stop$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "terminate",
				Description: `Terminates a running instance.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/terminate",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/terminate$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:server_id/terminate",
						Variables:  []string{"server_id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/terminate$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "unlock",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:id/unlock",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/unlock$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates attributes of a single instance.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/clouds/:cloud_id/instances/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "instance[cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[multi_cloud_image_href]",
						Description: `The href of the updated MultiCloudImage for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "instance[server_template_href]",
						Description: `The href of the updated ServerTemplate for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[security_group_hrefs][]",
						Description: `The hrefs of the updated security groups.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ramdisk_image_href]",
						Description: `The href of the updated ramdisk image for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[instance_type_href]",
						Description: `The href of the updated Instance Type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[kernel_image_href]",
						Description: `The href of the updated kernel image for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[deployment_href]",
						Description: `The href of the updated Deployment for the Instance. This is only supported for Instances that are not associated with a Server or ServerArray.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[datacenter_href]",
						Description: `The href of the updated Datacenter / Zone for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[ssh_key_href]",
						Description: `The href of the updated SSH key for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[image_href]",
						Description: `The href of the updated Image for the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "instance[name]",
						Description: `The updated name to give the Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"instance"},
			},
		},
	},
	"InstanceCustomLodgement": &metadata.Resource{
		Name:        "InstanceCustomLodgement",
		Description: `An InstanceCustomLodgement represents a way to create custom reports about a specific instance with a user defined quantity.  Replaces the legacy Instances#setcustomlodgement interface.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a lodgement with the quantity and timeframe specified.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/instance_custom_lodgements$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "quantity[][value]",
						Description: `The value of the quantity. Should be a positive integer.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "quantity[][name]",
						Description: `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "timeframe",
						Description: `The time-frame (either a month "YYYY_MM" or a single day "YYYY_MM_DD") for which the quantity value is valid (currently for the PDT timezone only).`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"quantity", "timeframe"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Destroy the specified lodgement.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/instance_custom_lodgements/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `List InstanceCustomLodgements of a given cloud and instance.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/instance_custom_lodgements$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show the specified lodgement.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/instance_custom_lodgements/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update a lodgement with the quantity specified.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/instance_custom_lodgements/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/instance_custom_lodgements/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "quantity[][value]",
						Description: `The value of the quantity. Should be a positive integer.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "quantity[][name]",
						Description: `The name of the quantity. A customer-specific string, e.g. "MB/s" or "GB/Month".`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"quantity"},
			},
		},
	},
	"InstanceType": &metadata.Resource{
		Name:        "InstanceType",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists instance types.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instance_types",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instance_types$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single Instance type.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instance_types/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instance_types/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"IpAddress": &metadata.Resource{
		Name:        "IpAddress",
		Description: `An IpAddress provides an abstraction for IPv4 addresses bindable to Instance resources running in a Cloud.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new IpAddress with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address[deployment_href]",
						Description: `The href of the Deployment that owns this IpAddress.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address[network_href]",
						Description: `(OpenStack Only) The href of the Network that the IpAddress will be associated to. This parameter is required for OpenStack with Neutron clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address[domain]",
						Description: `(Amazon Only) Pass vpc to create this IP for EC2-VPC only environments. Pass ec2_classic to create this IP for EC2-Classic environments. Defaults to ec2_classic.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"ec2_classic", "vpc"},
					},
					&metadata.ActionParam{
						Name:        "ip_address[name]",
						Description: `The name of the IpAddress to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"ip_address"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given IpAddress.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the IpAddresses available to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single IpAddress.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates attributes of a given IpAddress.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address[deployment_href]",
						Description: `The href of the Deployment that owns this IpAddress.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address[name]",
						Description: `The updated name of the IpAddress.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"ip_address"},
			},
		},
	},
	"IpAddressBinding": &metadata.Resource{
		Name:        "IpAddressBinding",
		Description: `An IpAddressBinding represents an abstraction for binding an IpAddress to an instance.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates an ip address binding which attaches a specified IpAddress resource to a specified instance, and also allows for configuration of port forwarding rules...`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings",
						Variables:  []string{"cloud_id", "ip_address_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)/ip_address_bindings$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/ip_address_bindings",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_address_bindings$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ip_address_binding[public_ip_address_href]",
						Description: `The IpAddress to bind to the specified instance. Required unless port forwarding rule params are passed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[instance_href]",
						Description: `The Instance to which this IpAddress should be bound.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[private_port]",
						Description: `Incoming network traffic will get forwarded to this port number on the specified Instance. If not specified, will use public port. Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[public_port]",
						Description: `The incoming port for port forwarding. Incoming network traffic on this port will get forwarded (to the IP:Private Port of the specified Instance). Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "ip_address_binding[protocol]",
						Description: `Transport layer protocol of traffic that may be forwarded from public port to private port on the Instance. Required unless public_ip_address_href is passed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"UDP", "TCP"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"ip_address_binding"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `No description provided for destroy.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id",
						Variables:  []string{"cloud_id", "ip_address_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)/ip_address_bindings/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/ip_address_bindings/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_address_bindings/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the ip address bindings available to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings",
						Variables:  []string{"cloud_id", "ip_address_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)/ip_address_bindings$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_address_bindings",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_address_bindings$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single ip address binding.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_addresses/:ip_address_id/ip_address_bindings/:id",
						Variables:  []string{"cloud_id", "ip_address_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_addresses/([^/]+)/ip_address_bindings/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ip_address_bindings/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ip_address_bindings/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"MonitoringMetric": &metadata.Resource{
		Name:        "MonitoringMetric",
		Description: `A monitoring metric is a stream of data that is captured in an instance. Metrics can be monitored, graphed and can be used as the basis for triggering alerts.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "data",
				Description: `Gives the raw monitoring data for a particular metric`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id/data",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/monitoring_metrics/([^/]+)/data$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "start",
						Description: `An integer number of seconds from current time. e.g. -300`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "end",
						Description: `An integer number of seconds from current time. e.g. -150 or 0 `,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the monitoring metrics available for the instance and their corresponding graph hrefs.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/monitoring_metrics$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "period",
						Description: `The time scale for which the graph is generated. Default is 'day'`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"now", "day", "yday", "week", "lweek", "month", "quarter", "year"},
					},
					&metadata.ActionParam{
						Name:        "title",
						Description: `The title of the graph.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "size",
						Description: `The size of the graph to be generated. Default is 'small'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"thumb", "tiny", "small", "large", "xlarge"},
					},
					&metadata.ActionParam{
						Name:        "tz",
						Description: `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows attributes of a single monitoring metric.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/monitoring_metrics/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/monitoring_metrics/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "period",
						Description: `The time scale for which the graph is generated. Default is 'day'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"now", "day", "yday", "week", "lweek", "month", "quarter", "year"},
					},
					&metadata.ActionParam{
						Name:        "title",
						Description: `The title of the graph.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "size",
						Description: `The size of the graph to be generated. Default is 'small'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"thumb", "tiny", "small", "large", "xlarge"},
					},
					&metadata.ActionParam{
						Name:        "tz",
						Description: `The time zone in which the graph will be displayed. Default will be 'America/Los_Angeles'. For more zones, see User Settings -> Preferences. `,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"MultiCloudImage": &metadata.Resource{
		Name:        "MultiCloudImage",
		Description: `A MultiCloudImage is a RightScale component that functions as a pointer to machine images in specific clouds (e...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "clone",
				Description: `Clones a given MultiCloudImage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/multi_cloud_images/:id/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/clone$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image[description]",
						Description: `The description for the cloned MultiCloudImage.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image[name]",
						Description: `The name for the cloned MultiCloudImage.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"multi_cloud_image"},
			},

			&metadata.Action{
				Name:        "commit",
				Description: `Commits a given MultiCloudImage. Only HEAD revisions can be committed.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/multi_cloud_images/:id/commit",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/commit$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "commit_message",
						Description: `The message associated with the commit.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"commit_message"},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new MultiCloudImage with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:server_template_id/multi_cloud_images",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image[description]",
						Description: `The description of the MultiCloudImage to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image[name]",
						Description: `The name of the MultiCloudImage to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"multi_cloud_image"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given MultiCloudImage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/server_templates/:server_template_id/multi_cloud_images/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/multi_cloud_images/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the MultiCloudImages available to this account. HEAD revisions have a revision of 0.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/multi_cloud_images",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single MultiCloudImage. HEAD revisions have a revision of 0.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/multi_cloud_images/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/multi_cloud_images/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates attributes of a given MultiCloudImage. Only HEAD revisions can be updated (revision 0).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/server_templates/:server_template_id/multi_cloud_images/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/multi_cloud_images/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/multi_cloud_images/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image[description]",
						Description: `The updated description for the MultiCloudImage.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image[name]",
						Description: `The updated name for the MultiCloudImage.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"multi_cloud_image"},
			},
		},
	},
	"MultiCloudImageSetting": &metadata.Resource{
		Name:        "MultiCloudImageSetting",
		Description: `A MultiCloudImageSetting defines which settings should be used when a server is launched in a cloud...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new setting for an existing MultiCloudImage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/multi_cloud_images/:multi_cloud_image_id/settings",
						Variables:  []string{"multi_cloud_image_id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[instance_type_href]",
						Description: `The href of the instance type. Mandatory if specifying cloud_href.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[ramdisk_image_href]",
						Description: `The href of the ramdisk image. Optional if specifying cloud_href.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[kernel_image_href]",
						Description: `The href of the kernel image. Optional if specifying cloud_href.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[cloud_href]",
						Description: `The href of the Cloud to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[image_href]",
						Description: `The href of the Image to use. Mandatory if specifying cloud_href.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"multi_cloud_image_setting"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a MultiCloudImage setting.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/multi_cloud_images/:multi_cloud_image_id/settings/:id",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the MultiCloudImage settings.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/multi_cloud_images/:multi_cloud_image_id/settings",
						Variables:  []string{"multi_cloud_image_id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single MultiCloudImage setting.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/multi_cloud_images/:multi_cloud_image_id/settings/:id",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a settings for a MultiCLoudImage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/multi_cloud_images/:multi_cloud_image_id/settings/:id",
						Variables:  []string{"multi_cloud_image_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/multi_cloud_images/([^/]+)/settings/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[instance_type_href]",
						Description: `The href of the instance type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[ramdisk_image_href]",
						Description: `The href of the ramdisk image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[kernel_image_href]",
						Description: `The href of the kernel image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[cloud_href]",
						Description: `The href of the Cloud to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "multi_cloud_image_setting[user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"multi_cloud_image_setting"},
			},
		},
	},
	"Network": &metadata.Resource{
		Name:        "Network",
		Description: `A Network is a logical grouping of network devices.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new network.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/networks",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/networks$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network[instance_tenancy]",
						Description: `The launch policy for AWS instances in the Network. Specify 'default' to allow instances to decide their own launch policy. Specify 'dedicated' to force all instances to be launched as 'dedicated'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[description]",
						Description: `The description for the Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[cidr_block]",
						Description: `The range of IP addresses for the Network. This parameter is required for Amazon clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[cloud_href]",
						Description: `The Cloud to create the Network in`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[name]",
						Description: `The name for the Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"network"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes the given network(s).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/networks/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/networks/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists networks in this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/networks",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/networks$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows attributes of a single network.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/networks/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/networks/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates the given network.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/networks/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/networks/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network[route_table_href]",
						Description: `Sets the default RouteTable for this Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[description]",
						Description: `The updated description for the Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network[name]",
						Description: `The updated name for the Network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"network"},
			},
		},
	},
	"NetworkGateway": &metadata.Resource{
		Name:        "NetworkGateway",
		Description: `A NetworkGateway is an interface that allows traffic to be routed between networks.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a new NetworkGateway.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/network_gateways",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_gateways$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_gateway[description]",
						Description: `The description to be set on the NetworkGateway.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[cloud_href]",
						Description: `The cloud to create the NetworkGateway in.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[name]",
						Description: `The name to be set on the NetworkGateway.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[type]",
						Description: `The type of the NetworkGateway.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"vpn", "internet"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"network_gateway"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing NetworkGateway.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/network_gateways/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_gateways/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the NetworkGateways available to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_gateways",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_gateways$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single NetworkGateway.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_gateways/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_gateways/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update an existing NetworkGateway.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/network_gateways/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_gateways/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_gateway[network_href]",
						Description: `Pass a blank string to detach from the specified Network, or pass a valid Network href to attach to the specified network.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[description]",
						Description: `The description to be set on the NetworkGateway.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_gateway[name]",
						Description: `The name to be set on the NetworkGateway.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"network_gateway"},
			},
		},
	},
	"NetworkOptionGroup": &metadata.Resource{
		Name:        "NetworkOptionGroup",
		Description: `A key/value pair hash containing options for configuring a Network.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a new NetworkOptionGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/network_option_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group[description]",
						Description: `Description of this NetworkOptionGroup`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[cloud_href]",
						Description: `The Cloud to create this NetworkOptionGroup in`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[options]",
						Description: ``,
						Type:        "map",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[name]",
						Description: `Name of this NetworkOptionGroup`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[type]",
						Description: `Type of this NetworkOptionGroup`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"network_option_group"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing NetworkOptionGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/network_option_groups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `List NetworkOptionGroups available in this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_option_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single NetworkOptionGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_option_groups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update an existing NetworkOptionGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/network_option_groups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_groups/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group[description]",
						Description: `Update the description`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[options]",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "network_option_group[name]",
						Description: `Update the name`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"network_option_group"},
			},
		},
	},
	"NetworkOptionGroupAttachment": &metadata.Resource{
		Name:        "NetworkOptionGroupAttachment",
		Description: `Resource for attaching NetworkOptionGroups to Networks.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a new NetworkOptionGroupAttachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/network_option_group_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group_attachment[network_option_group_href]",
						Description: `The NetworkOptionGroup to attach to the specified resource.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "network_option_group_attachment[network_href]",
						Description: `The Network to attach the specified NetworkOptionGroup to.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"network_option_group_attachment"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing NetworkOptionGroupAttachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/network_option_group_attachments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `List NetworkOptionGroupAttachments in this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_option_group_attachments",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single NetworkOptionGroupAttachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/network_option_group_attachments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update an existing NetworkOptionGroupAttachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/network_option_group_attachments/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/network_option_group_attachments/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "network_option_group_attachment[network_option_group_href]",
						Description: `The NetworkOptionGroup to attach to the specified resource.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"network_option_group_attachment"},
			},
		},
	},
	"Oauth2": &metadata.Resource{
		Name:        "Oauth2",
		Description: `Note that all API calls irrespective of the resource it is acting on, should pass a header "X-Api-Version" with the value "1...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Perform an OAuth 2`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/oauth2/",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/oauth2/$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_link_version",
						Description: `The RightLink gem version the client conforms to (only needed for instance agent clients).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "client_secret",
						Description: `The client secret (only needed for confidential clients).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "refresh_token",
						Description: `The refresh token obtained from OAuth grant.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "r_s_version",
						Description: `The RightAgent protocol version the client conforms to (only needed for instance agent clients).`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "account_id",
						Description: `The client's account ID (only needed for instance agent clients).`,
						Type:        "int",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "grant_type",
						Description: `Type of grant.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"refresh_token"},
					},
					&metadata.ActionParam{
						Name:        "client_id",
						Description: `The client ID (only needed for confidential clients).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"account_id", "client_id", "client_secret", "grant_type", "r_s_version", "refresh_token", "right_link_version"},
			},
		},
	},
	"Permission": &metadata.Resource{
		Name:        "Permission",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a permission, thereby granting some user a particular role with respect to the current account...`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/permissions",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/permissions$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "permission[role_title]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "permission[user_href]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"permission"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Destroy a permission, thereby revoking a user's role with respect to the current account...`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/permissions/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/permissions/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `List all permissions for all users of the current acount.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/permissions",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/permissions$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single permission.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/permissions/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/permissions/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"PlacementGroup": &metadata.Resource{
		Name:        "PlacementGroup",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a PlacementGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/placement_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/placement_groups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "placement_group[description]",
						Description: `The description of the Placement Group to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "placement_group[cloud_href]",
						Description: `The Href of the Cloud in which the PlacementGroup should be created. Note: This feature is not supported for all clouds.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "placement_group[name]",
						Description: `The name of the Placement Group to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"placement_group"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Destroys a PlacementGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/placement_groups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/placement_groups/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists all PlacementGroups in an account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/placement_groups",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/placement_groups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows information about a single PlacementGroup.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/placement_groups/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/placement_groups/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Preference": &metadata.Resource{
		Name:        "Preference",
		Description: `A Preference is a user and account-specific setting. Preferences are used in many part of the RightScale platform and can be used for custom purposes if desired.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes the given preference.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/preferences/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/preferences/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists all preferences.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/preferences",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/preferences$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a single preference.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/preferences/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/preferences/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `If 'id' is known, updates preference with given contents.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/preferences/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/preferences/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "preference[contents]",
						Description: `The updated contents for the Preference.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"preference"},
			},
		},
	},
	"Publication": &metadata.Resource{
		Name:        "Publication",
		Description: `A Publication is a revisioned component shared with a set of Account Groups.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "import",
				Description: `Imports the given publication and its subordinates to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/publications/:id/import",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/publications/([^/]+)/import$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the publications available to this account. Only non-HEAD revisions are possible.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/publications",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/publications$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single publication. Only non-HEAD revisions are possible.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/publications/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/publications/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"PublicationLineage": &metadata.Resource{
		Name:        "PublicationLineage",
		Description: `A Publication Lineage contains lineage information for a Publication in the MultiCloudMarketplace.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single publication lineage. Only non-HEAD revisions are possible.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/publication_lineages/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/publication_lineages/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"RecurringVolumeAttachment": &metadata.Resource{
		Name:        "RecurringVolumeAttachment",
		Description: `A RecurringVolumeAttachment specifies a Volume/VolumeSnapshot to attach to a Server/ServerArray the next time an instance is launched.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new recurring volume attachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/recurring_volume_attachments",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/recurring_volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/recurring_volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments",
						Variables:  []string{"cloud_id", "volume_snapshot_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/recurring_volume_attachments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment[runnable_href]",
						Description: `The href of the server or server array to attach to.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment[storage_href]",
						Description: `The href of the volume or volume snapshot to be attached on launch of a next instance.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "recurring_volume_attachment[device]",
						Description: `The device location where the volume or volume snapshot will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"recurring_volume_attachment"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given recurring volume attachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/recurring_volume_attachments/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments/:id",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments/:id",
						Variables:  []string{"cloud_id", "volume_snapshot_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists all recurring volume attachments.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/recurring_volume_attachments",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/recurring_volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/recurring_volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments",
						Variables:  []string{"cloud_id", "volume_snapshot_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/recurring_volume_attachments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single recurring volume attachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/recurring_volume_attachments/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/recurring_volume_attachments/:id",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volume_snapshots/:volume_snapshot_id/recurring_volume_attachments/:id",
						Variables:  []string{"cloud_id", "volume_snapshot_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)/recurring_volume_attachments/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Repository": &metadata.Resource{
		Name:        "Repository",
		Description: `A Repository is a location from which you can download and import design objects such as Chef cookbooks. Using this resource you can add and modify repository information and import assets discovered in the repository.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "cookbook_import",
				Description: `Performs a Cookbook import, which allows you to use the specified cookbooks in your design objects.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/repositories/:id/cookbook_import",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/cookbook_import$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "repository_commit_reference",
						Description: `Optional commit reference indicating last succeeded commit. Must match the Repository's fetch_status.succeeded_commit attribute or the import will not be performed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "with_dependencies",
						Description: `A flag indicating whether dependencies should automatically be imported.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "asset_hrefs[]",
						Description: `Hrefs of the assets that should be imported.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "namespace",
						Description: `The namespace to import into.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"primary", "alternate"},
					},
					&metadata.ActionParam{
						Name:        "follow",
						Description: `A flag indicating whether imported cookbooks should be followed.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"asset_hrefs", "follow", "namespace", "repository_commit_reference", "with_dependencies"},
			},

			&metadata.Action{
				Name:        "cookbook_import_preview",
				Description: `Retrieves a preview of the effects of a Cookbook import.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/repositories/:id/cookbook_import_preview",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/cookbook_import_preview$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "asset_hrefs[]",
						Description: `Hrefs of the assets that should be imported.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "namespace",
						Description: `The namespace to import into.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"primary", "alternate"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"asset_hrefs", "namespace"},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a Repository.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/repositories",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/repositories$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "repository[asset_paths][cookbooks][]",
						Description: `The cookbook paths for the repository`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][password]",
						Description: `The password, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][username]",
						Description: `The user name, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][ssh_key]",
						Description: `The SSH key, or credential, for the repository (only valid for git repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[commit_reference]",
						Description: `The revision for the repository`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[source_type]",
						Description: `The source type for the repository.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"git", "svn", "download"},
					},
					&metadata.ActionParam{
						Name:        "repository[auto_import]",
						Description: `Whether cookbooks should automatically be imported upon repository creation.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "repository[description]",
						Description: `The description for the repository.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[source]",
						Description: `The URL for the repository.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[name]",
						Description: `The repository name.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"repository"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes the specified Repositories.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/repositories/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists all Repositories for this Account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/repositories",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/repositories$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "refetch",
				Description: `Refetches all RepositoryAssets associated with the Repository.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/repositories/:id/refetch",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/refetch$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "auto_import",
						Description: `Whether cookbooks should automatically be imported after repositories are fetched.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"auto_import"},
			},

			&metadata.Action{
				Name:        "resolve",
				Description: `Show a list of repositories that have imported cookbooks with the given names.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/repositories/resolve",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/repositories/resolve$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "imported_cookbook_name[]",
						Description: `A list of cookbook names that were imported by the repository.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"imported_cookbook_name"},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows a specified Repository.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/repositories/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates a specified Repository.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/repositories/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "repository[asset_paths][cookbooks][]",
						Description: `The updated cookbook paths for the repository`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][password]",
						Description: `The updated password, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][username]",
						Description: `The updated user name, or credential, for the repository (only valid for svn or download repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[credentials][ssh_key]",
						Description: `The updated SSH key for the repository (only valid for git repositories).`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[commit_reference]",
						Description: `The updated commit reference (tag, branch, revision...) for the repository`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[source_type]",
						Description: `The updated source type for the repository.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"git", "svn", "download"},
					},
					&metadata.ActionParam{
						Name:        "repository[description]",
						Description: `The updated description for the repository.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "repository[source]",
						Description: `The updated URL for the repository.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "repository[name]",
						Description: `The updated repository name.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"repository"},
			},
		},
	},
	"RepositoryAsset": &metadata.Resource{
		Name:        "RepositoryAsset",
		Description: `A RepositoryAsset represents an item discovered in a Repository`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `List a repository's current assets.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/repositories/:repository_id/repository_assets",
						Variables:  []string{"repository_id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/repository_assets$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single asset.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/repositories/:repository_id/repository_assets/:id",
						Variables:  []string{"repository_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/repositories/([^/]+)/repository_assets/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"RightScript": &metadata.Resource{
		Name:        "RightScript",
		Description: `A RightScript is an executable piece of code that can be run on a server during the boot, operational, or decommission phases...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "commit",
				Description: `Commits the given RightScript. Only HEAD revisions (revision 0) can be committed.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/right_scripts/:id/commit",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/commit$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script[commit_message]",
						Description: `The message to be included with the requested commit`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"right_script"},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists RightScripts.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/right_scripts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/right_scripts$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "latest_only",
						Description: `Whether or not to return only the latest version for each lineage.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single RightScript.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/right_scripts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show_source",
				Description: `Returns the script source for a RightScript`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/right_scripts/:id/source",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/source$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates RightScript name/description`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/right_scripts/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "right_script[description]",
						Description: `The new description for the RightScript`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "right_script[name]",
						Description: `The new name for the RightScript`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"right_script"},
			},

			&metadata.Action{
				Name:        "update_source",
				Description: `Updates the source of the given RightScript`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/right_scripts/:id/source",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/right_scripts/([^/]+)/source$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"Route": &metadata.Resource{
		Name:        "Route",
		Description: `A Route defines how networking traffic should be routed from one destination to another...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a new Route.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/routes",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/routes$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/route_tables/:route_table_id/routes",
						Variables:  []string{"route_table_id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route[destination_cidr_block]",
						Description: `The destination (CIDR IP address) for the Route.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[route_table_href]",
						Description: `The RouteTable to create the Route in.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_href]",
						Description: `The href of the Route's next hop.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_type]",
						Description: `The Route's next hop type.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"instance", "network_interface", "network_gateway", "ip_string"},
					},
					&metadata.ActionParam{
						Name:        "route[description]",
						Description: `The description to be set on the Route.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_ip]",
						Description: `The IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"route"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing Route.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/routes/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/routes/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/route_tables/:route_table_id/routes/:id",
						Variables:  []string{"route_table_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `List Routes available in this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/routes",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/routes$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/route_tables/:route_table_id/routes",
						Variables:  []string{"route_table_id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single Route.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/routes/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/routes/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/route_tables/:route_table_id/routes/:id",
						Variables:  []string{"route_table_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update an existing Route.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/routes/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/routes/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/route_tables/:route_table_id/routes/:id",
						Variables:  []string{"route_table_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)/routes/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route[destination_cidr_block]",
						Description: `The updated destination (CIDR IP address) for the Route.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_href]",
						Description: `The updated href of the Route's next hop. Required if route[next_hop_type] is 'instance', 'network_interface', or 'network_gateway'. Not allowed otherwise.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_type]",
						Description: `The updated Route's next hop type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"instance", "network_interface", "network_gateway", "ip_string"},
					},
					&metadata.ActionParam{
						Name:        "route[description]",
						Description: `The updated description of the Route.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "route[next_hop_ip]",
						Description: `The updated IP Address of the Route's next hop. Required if route[next_hop_type] is 'ip_string'. Not allowed otherwise.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"route"},
			},
		},
	},
	"RouteTable": &metadata.Resource{
		Name:        "RouteTable",
		Description: `Grouped listing of Routes`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a new RouteTable.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/route_tables",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/route_tables$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route_table[network_href]",
						Description: `The Network to create the RouteTable in.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route_table[description]",
						Description: `The description to be set on the RouteTable.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "route_table[cloud_href]",
						Description: `The cloud to create the RouteTable in.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "route_table[name]",
						Description: `The name to be set on the RouteTable.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"route_table"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete an existing RouteTable.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/route_tables/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `List RouteTables available in this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/route_tables",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/route_tables$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single RouteTable.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/route_tables/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update an existing RouteTable.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/route_tables/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/route_tables/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "route_table[description]",
						Description: `The description to be set on the RouteTable.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "route_table[name]",
						Description: `The name to be set on the RouteTable.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"route_table"},
			},
		},
	},
	"RunnableBinding": &metadata.Resource{
		Name:        "RunnableBinding",
		Description: `A RunnableBinding represents an item in a runlist of a ServerTemplate`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Bind an executable to the given ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:server_template_id/runnable_bindings",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_binding[right_script_href]",
						Description: `The RightScript href. Note: recipe cannot be specified when this param is given.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_binding[position]",
						Description: `The position of the executable in the execution order. If not specified, will be added to the end. If specified, will be inserted in that location and cause all others to move down.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_binding[sequence]",
						Description: `The sequence at which this executable should be run. Default is 'operational'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"boot", "decommission", "operational"},
					},
					&metadata.ActionParam{
						Name:        "runnable_binding[recipe]",
						Description: `The Chef recipe name. Note: right_script_href cannot be specified when this param is given.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"runnable_binding"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Unbind an executable from the given resource.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/server_templates/:server_template_id/runnable_bindings/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the executables bound to the ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/runnable_bindings",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "multi_update",
				Description: `Update attributes for multiple bound executables.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/server_templates/:server_template_id/runnable_bindings/multi_update",
						Variables:  []string{"server_template_id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings/multi_update$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "runnable_bindings[][right_script_href]",
						Description: `The updated RightScript href. Note: recipe cannot be specified when this param is given.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_bindings[][position]",
						Description: `The updated position of the RunnableBinding in the execution order. If specified, will be inserted in that location and cause all others to move down.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_bindings[][sequence]",
						Description: `The sequence at which this executable should be run.  Default is 'operational'.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"boot", "decommission", "operational"},
					},
					&metadata.ActionParam{
						Name:        "runnable_bindings[][recipe]",
						Description: `The updated Chef recipe name. Note: right_script_href cannot be specified when this param is given.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "runnable_bindings[][id]",
						Description: `The ID of the RunnableBinding to update.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"runnable_bindings"},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single executable binding.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:server_template_id/runnable_bindings/:id",
						Variables:  []string{"server_template_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/runnable_bindings/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"SecurityGroup": &metadata.Resource{
		Name:        "SecurityGroup",
		Description: `Security Groups represent network security profiles that contain lists of firewall rules for different ports and source IP addresses, as well as trust relationships amongst different security groups...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a security group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/security_groups",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group[network_href]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group[description]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group[name]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"security_group"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete security group(s)`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/security_groups/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists Security Groups.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/security_groups",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "tiny"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single Security Group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/security_groups/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "tiny"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"SecurityGroupRule": &metadata.Resource{
		Name:        "SecurityGroupRule",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a security group rule for a security group.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/security_group_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules",
						Variables:  []string{"cloud_id", "security_group_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol_details][start_port]",
						Description: `Start of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol_details][icmp_type]",
						Description: `ICMP type. Required if protocol is 'icmp'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol_details][icmp_code]",
						Description: `ICMP code. Required if protocol is 'icmp'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol_details][end_port]",
						Description: `End of port range (inclusive). Required if protocol is 'tcp' or 'udp'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[security_group_href]",
						Description: `Security Group to add rule to.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[source_type]",
						Description: `Source type. May be a CIDR block or another Security Group.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"cidr_ips", "group"},
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[group_owner]",
						Description: `Owner of source Security Group. Required if source_type is 'group'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[group_name]",
						Description: `Name of source Security Group. Required if source_type is 'group'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[direction]",
						Description: `Direction of traffic.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"ingress", "egress"},
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[protocol]",
						Description: `Protocol to filter on.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"tcp", "udp", "icmp", "all"},
					},
					&metadata.ActionParam{
						Name:        "security_group_rule[cidr_ips]",
						Description: `An IP address range in CIDR notation. Required if source_type is 'cidr_ips'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"security_group_rule"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Delete security group rule(s)`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/security_group_rules/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id",
						Variables:  []string{"cloud_id", "security_group_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists SecurityGroupRules.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/security_group_rules",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules",
						Variables:  []string{"cloud_id", "security_group_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single SecurityGroupRule.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/security_group_rules/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id",
						Variables:  []string{"cloud_id", "security_group_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: ``,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/security_group_rules/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/security_group_rules/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/clouds/:cloud_id/security_groups/:security_group_id/security_group_rules/:id",
						Variables:  []string{"cloud_id", "security_group_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/security_groups/([^/]+)/security_group_rules/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "security_group_rule[description]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"security_group_rule"},
			},
		},
	},
	"Server": &metadata.Resource{
		Name:        "Server",
		Description: `Servers represent the notion of a server/machine from the RightScale's perspective`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "clone",
				Description: `Clones a given server.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:id/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/clone$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new server, and configures its corresponding "next" instance with the received parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/servers$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:deployment_id/servers",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][multi_cloud_image_href]",
						Description: `The href of the Multi Cloud Image to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[instance][placement_group_href]",
						Description: `The href of the Placement Group.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][security_group_hrefs][]",
						Description: `The hrefs of the security groups.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][server_template_href]",
						Description: `The href of the Server Template.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][instance_type_href]",
						Description: `The href of the Instance Type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][ramdisk_image_href]",
						Description: `The href of the Ramdisk Image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][kernel_image_href]",
						Description: `The href of the Kernel Image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][inputs][][value]",
						Description: `The value of that Input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][inputs]",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][inputs][][name]",
						Description: `The Input name. This format is used for passing legacy 1.0-style Inputs. Will eventually be deprecated.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][ssh_key_href]",
						Description: `The href of the SSH key to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][subnet_hrefs][]",
						Description: `The hrefs of the updated subnets.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][cloud_href]",
						Description: `The href of the cloud that the Server should be added to.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][image_href]",
						Description: `The href of the Image to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[deployment_href]",
						Description: `The href of the deployment to which the Server will be added.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[description]",
						Description: `The Server description.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[optimized]",
						Description: `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[name]",
						Description: `The name of the Server.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given server.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/servers/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/deployments/:deployment_id/servers/:id",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists servers.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/servers$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:deployment_id/servers",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "launch",
				Description: `Launches the "next" instance of this server`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:id/launch",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/launch$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows the information of a single server.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/servers/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:deployment_id/servers/:id",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "terminate",
				Description: `Terminates the current instance of this server`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/:id/teminate",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)/teminate$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates attributes of a single server.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/servers/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/servers/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/deployments/:deployment_id/servers/:id",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server[automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[description]",
						Description: `The updated description for the server.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[optimized]",
						Description: `A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server[name]",
						Description: `The updated server name.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server"},
			},

			&metadata.Action{
				Name:        "wrap_instance",
				Description: `Wrap an existing instance and set current instance for new server`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/servers/wrap_instance",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/servers/wrap_instance$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:deployment_id/servers/wrap_instance",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/servers/wrap_instance$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server[instance][multi_cloud_image_href]",
						Description: `The href of the Multi Cloud Image to use.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][server_template_href]",
						Description: `The href of the Server Template.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][inputs]",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[deployment_href]",
						Description: `The href of the deployment to which the Server will be added.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[instance][href]",
						Description: `The href of the Instance around which the server should be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[description]",
						Description: `The Server description.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server[name]",
						Description: `The name of the Server.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server"},
			},
		},
	},
	"ServerArray": &metadata.Resource{
		Name:        "ServerArray",
		Description: `A server array represents a logical group of instances and allows to resize(grow/shrink) that group based on certain elasticity parameters.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "clone",
				Description: `Clones a given server array.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:id/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/clone$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new server array, and configures its corresponding "next" instance with the received parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_arrays$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/deployments/:deployment_id/server_arrays",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]",
						Description: `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][automatic_instance_store_mapping]",
						Description: `A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][collect_audit_entries]",
						Description: `The audit SQS queue that will store audit entries.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][alert_specific_params][voters_tag_predicate]",
						Description: `The Voters Tag that RightScale will use in order to determine when to scale up/down.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][algorithm]",
						Description: `The algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"max_10", "avg_10"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_performance]",
						Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][alert_specific_params][decision_threshold]",
						Description: `The percentage of servers that must agree in order to trigger an alert before an action is taken.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][max_age]",
						Description: `The threshold (in seconds) before a resize action occurs on the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][regexp]",
						Description: `The regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_type_uid]",
						Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][iam_instance_profile]",
						Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_specific_attributes][root_volume_size]",
						Description: `The size for root disk. Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_calm_time]",
						Description: `The time (in minutes) on how long you want to wait before you repeat another action.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_down_by]",
						Description: `The number of servers to scale down by.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][max_count]",
						Description: `The maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][min_count]",
						Description: `The minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_up_by]",
						Description: `The number of servers to scale up by.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][associate_public_ip_address]",
						Description: `Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][bounds][max_count]",
						Description: `The maximum number of servers that can be operational at the same time in the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][bounds][min_count]",
						Description: `The minimum number of servers that must be operational at all times in the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][time]",
						Description: `Specifies the time when an alert-based array resizes.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][day]",
						Description: `Specifies the day when an alert-based array resizes.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][multi_cloud_image_href]",
						Description: `The href of the MultiCloudImage to be used.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][ip_forwarding_enabled]",
						Description: `Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][security_group_hrefs][]",
						Description: `The hrefs of the Security Groups.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][server_template_href]",
						Description: `The ServerTemplate that will be used to create the worker instances in the server array.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][placement_group_href]",
						Description: `The href of the Placement Group.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][ramdisk_image_href]",
						Description: `The href of the Ramdisk Image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][instance_type_href]",
						Description: `The href of the Instance Type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][weight]",
						Description: `Instance allocation (should total 100%).`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][kernel_image_href]",
						Description: `The href of the Kernel Image.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][inputs][][value]",
						Description: `The value of that Input. Should be of the form 'text:my_value' or 'cred:MY_CRED' etc.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][inputs]",
						Description: ``,
						Type:        "map",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][datacenter_href]",
						Description: `The href of the Datacenter / Zone. For multiple Datacenters, use 'datacenter_policy' instead.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][max]",
						Description: `Max instances (0 for unlimited).`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][inputs][][name]",
						Description: `The Input name.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][ssh_key_href]",
						Description: `The href of the SSH Key to be used.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][subnet_hrefs][]",
						Description: `The hrefs of the updated Subnets.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][cloud_href]",
						Description: `The href of the Cloud that the array will be associated with.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][image_href]",
						Description: `The href of the Image to be used.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[instance][user_data]",
						Description: `User data that RightScale automatically passes to your instance at boot time.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[deployment_href]",
						Description: `The href of the deployment for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[description]",
						Description: `The description for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[array_type]",
						Description: `The array type for the Server Array.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"alert", "queue"},
					},
					&metadata.ActionParam{
						Name:        "server_array[optimized]",
						Description: `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[state]",
						Description: `The status of the server array. If active, the server array is enabled for scaling actions.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"enabled", "disabled"},
					},
					&metadata.ActionParam{
						Name:        "server_array[name]",
						Description: `The name for the Server Array.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server_array"},
			},

			&metadata.Action{
				Name:        "current_instances",
				Description: `List the running instances belonging to the server array. See Instances#index for details.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:id/current_instances",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/current_instances$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given server array.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/server_arrays/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/deployments/:deployment_id/server_arrays/:id",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists server arrays.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_arrays$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:deployment_id/server_arrays",
						Variables:  []string{"deployment_id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "launch",
				Description: `Launches a new instance in the server array with the configuration defined in the 'next_instance'`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:id/launch",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/launch$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "multi_run_executable",
				Description: `Run an executable on all instances of this array`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:id/multi_run_executable",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/multi_run_executable$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "multi_terminate",
				Description: `Terminate all instances of this array`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_arrays/:id/multi_terminate",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/multi_terminate$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows the information of a single server array.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/deployments/:deployment_id/server_arrays/:id",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "instance_detail"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates attributes of a single server array.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/server_arrays/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/deployments/:deployment_id/server_arrays/:id",
						Variables:  []string{"deployment_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/deployments/([^/]+)/server_arrays/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]",
						Description: `Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and "Items per instance" is set to 10, the server array will resize to 5 worker instances (50/10).  Default = 10`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][collect_audit_entries]",
						Description: `The updated audit SQS queue that will store audit entries.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][alert_specific_params][voters_tag_predicate]",
						Description: `The updated Voters Tag that RightScale will use in order to determine when to scale up/down.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][algorithm]",
						Description: `The updated algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"max_10", "avg_10"},
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][alert_specific_params][decision_threshold]",
						Description: `The updated percentage of servers that must agree in order to trigger an alert before an action is taken.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][max_age]",
						Description: `The updated threshold (in seconds) before a resize action occurs on the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][queue_specific_params][item_age][regexp]",
						Description: `The updated regexp that helps the system determine an item's "age" in the queue. Example: created_at: (\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d UTC)`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_calm_time]",
						Description: `The updated time (in minutes) on how long you want to wait before you repeat another action.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_down_by]",
						Description: `The updated number of servers to scale down by.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][min_count]",
						Description: `The updated minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][max_count]",
						Description: `The updated maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][pacing][resize_up_by]",
						Description: `The updated number of servers to scale up by.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][bounds][min_count]",
						Description: `The updated minimum number of servers that must be operational at all times in the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][bounds][max_count]",
						Description: `The updated maximum number of servers that can be operational at the same time in the server array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][datacenter_href]",
						Description: `The href of the Datacenter / Zone.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][time]",
						Description: `The updated time when an alert-based array resizes.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[elasticity_params][schedule][][day]",
						Description: `The updated day when an alert-based array resizes.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][weight]",
						Description: `Instance allocation (should total 100%).`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[datacenter_policy][][max]",
						Description: `Max instances (0 for unlimited).`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[deployment_href]",
						Description: `The updated href of the deployment for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[description]",
						Description: `The updated description for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_array[array_type]",
						Description: `The updated array type for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"alert", "queue"},
					},
					&metadata.ActionParam{
						Name:        "server_array[optimized]",
						Description: `A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "server_array[state]",
						Description: `The updated status of the server array. If active, the server array is enabled for scaling actions.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"enabled", "disabled"},
					},
					&metadata.ActionParam{
						Name:        "server_array[name]",
						Description: `The updated name for the Server Array.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server_array"},
			},
		},
	},
	"ServerTemplate": &metadata.Resource{
		Name:        "ServerTemplate",
		Description: `ServerTemplates allow you to pre-configure servers by starting from a base image and adding scripts that run during the boot, operational, and shutdown phases...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "clone",
				Description: `Clones a given ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:id/clone",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/clone$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template[description]",
						Description: `The description for the cloned ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_template[name]",
						Description: `The name for the cloned ServerTemplate.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server_template"},
			},

			&metadata.Action{
				Name:        "commit",
				Description: `Commits a given ServerTemplate. Only HEAD revisions (revision 0) that are owned by the account can be committed.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:id/commit",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/commit$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "commit_head_dependencies",
						Description: `Commit all HEAD revisions (if any) of the associated MultiCloud Images, RightScripts and Chef repo sequences.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "freeze_repositories",
						Description: `Freeze the repositories.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "commit_message",
						Description: `The message associated with the commit.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"commit_head_dependencies", "commit_message", "freeze_repositories"},
			},

			&metadata.Action{
				Name:        "create",
				Description: `Creates a new ServerTemplate with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_templates$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template[description]",
						Description: `The description of the ServerTemplate to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_template[name]",
						Description: `The name of the ServerTemplate to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server_template"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/server_templates/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "detect_changes_in_head",
				Description: `Identifies RightScripts attached to the resource that differ from their HEAD.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:id/detect_changes_in_head",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/detect_changes_in_head$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the ServerTemplates available to this account. HEAD revisions have a revision of 0.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_templates$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "publish",
				Description: `Publishes a given ServerTemplate and its subordinates.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:id/publish",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/publish$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "account_group_hrefs[]",
						Description: `List of hrefs of account groups to publish to.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "descriptions[short]",
						Description: `Short Description.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "descriptions[notes]",
						Description: `New Revision Notes.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "descriptions[long]",
						Description: `Long Description.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "allow_comments",
						Description: `Allow users to leave comments on this ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "email_comments",
						Description: `Email me when a user comments on this ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "categories[]",
						Description: `List of Categories.`,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"account_group_hrefs", "allow_comments", "categories", "descriptions", "email_comments"},
			},

			&metadata.Action{
				Name:        "resolve",
				Description: `Enumerates all attached cookbooks, missing dependencies and bound executables.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:id/resolve",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/resolve$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single ServerTemplate. HEAD revisions have a revision of 0.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_templates/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "inputs", "inputs_2_0"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "swap_repository",
				Description: `In-place replacement of attached cookbooks from a given repository.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_templates/:id/swap_repository",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)/swap_repository$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "source_repository_href",
						Description: `The repository whose cookbook attachments are to be replaced.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "target_repository_href",
						Description: `The repository whose cookbook attachments are to be utilized.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"source_repository_href", "target_repository_href"},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates attributes of a given ServerTemplate. Only HEAD revisions can be updated (revision 0).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/server_templates/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_templates/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template[description]",
						Description: `The updated description for the ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "server_template[name]",
						Description: `The updated name for the ServerTemplate.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server_template"},
			},
		},
	},
	"ServerTemplateMultiCloudImage": &metadata.Resource{
		Name:        "ServerTemplateMultiCloudImage",
		Description: `This resource represents links between ServerTemplates and MultiCloud Images and enables you to effectively add/delete MultiCloudImages to ServerTemplates and make them the default one...`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new ServerTemplateMultiCloudImage with the given parameters.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_template_multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "server_template_multi_cloud_image[multi_cloud_image_href]",
						Description: `The href of the MultiCloud Image to be used.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "server_template_multi_cloud_image[server_template_href]",
						Description: `The href of the ServerTemplate to be used.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"server_template_multi_cloud_image"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given ServerTemplateMultiCloudImage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/server_template_multi_cloud_images/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists the ServerTemplateMultiCloudImages available to this account.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_template_multi_cloud_images",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "make_default",
				Description: `Makes a given ServerTemplateMultiCloudImage the default for the ServerTemplate.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/server_template_multi_cloud_images/:id/make_default",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images/([^/]+)/make_default$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single ServerTemplateMultiCloudImage which represents an association between a ServerTemplate and a MultiCloudImage.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_template_multi_cloud_images/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/server_template_multi_cloud_images/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Session": &metadata.Resource{
		Name:        "Session",
		Description: `The sessions resource is in charge of creating API sessions that are bound to a given account`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "accounts",
				Description: `List all the accounts that a user has access to.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/sessions/accounts",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/sessions/accounts$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "password",
						Description: `The corresponding password.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "email",
						Description: `The email to login with if not using existing session.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: `Extended view shows account permissions and products`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Returns a list of root resources so an authenticated session can use them as a starting point or a way to know what features are available within its privileges...`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/sessions",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/sessions$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index_instance_session",
				Description: `Shows the full attributes of the instance (that has the token used to log-in).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/sessions/instance",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/sessions/instance$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"SshKey": &metadata.Resource{
		Name:        "SshKey",
		Description: `Ssh Keys represent a created SSH Key that exists in the cloud.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new ssh key.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/ssh_keys",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ssh_keys$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "ssh_key[name]",
						Description: `The name for the SSH key to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"ssh_key"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given ssh key.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/ssh_keys/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ssh_keys/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists ssh keys.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ssh_keys",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ssh_keys$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single ssh key.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/ssh_keys/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/ssh_keys/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "sensitive"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"Subnet": &metadata.Resource{
		Name:        "Subnet",
		Description: `A Subnet is a logical grouping of network devices`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new subnet.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/subnets",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/subnets",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "subnet[datacenter_href]",
						Description: `The associated Datacenter.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[network_href]",
						Description: `The associated Network.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[description]",
						Description: `The description for the Subnet.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[cidr_block]",
						Description: `The range of IP addresses for the Subnet.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "subnet[name]",
						Description: `The name for the Subnet.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"subnet"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes the given subnet(s).`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/subnets/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/subnets/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists subnets of a given cloud.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/subnets",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/subnets",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Shows attributes of a single subnet.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/subnets/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/subnets/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Updates name and description for the current subnet.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/subnets/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/subnets/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/clouds/:cloud_id/subnets/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/subnets/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "subnet[route_table_href]",
						Description: `The RouteTable to associate/dissociate with this Subnet. Pass this parameter with an empty string to reset this Subnet to use the default RouteTable.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "subnet[description]",
						Description: `The updated description for the Subnet.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
					&metadata.ActionParam{
						Name:        "subnet[name]",
						Description: `The updated name for the Subnet.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    false,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"subnet"},
			},
		},
	},
	"Tag": &metadata.Resource{
		Name:        "Tag",
		Description: `A tag or machine tag is a useful way of attaching useful metadata to an object/resource.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "by_resource",
				Description: `Get tags for a list of resource hrefs.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/tags/by_resource",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/tags/by_resource$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which tags are to be returned.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"resource_hrefs"},
			},

			&metadata.Action{
				Name:        "by_tag",
				Description: `Search for resources having a list of tags in a specific resource_type.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/tags/by_tag",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/tags/by_tag$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "include_tags_with_prefix",
						Description: `If included, all tags matching this prefix will be returned. If not included, no tags will be returned.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "resource_type",
						Description: `Search among a single resource type.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
						ValidValues: []string{"servers", "instances", "volumes", "volume_snapshots", "deployments", "server_templates", "multi_cloud_images", "images", "server_arrays", "accounts"},
					},
					&metadata.ActionParam{
						Name:        "with_deleted",
						Description: `If set to 'true', tags for deleted resources will also be returned. Default value is 'false'.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "match_all",
						Description: `If set to 'true', resources having all the tags specified in the 'tags' parameter are returned. Otherwise, resources having any of the tags are returned.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "tags[]",
						Description: `The tags which must be present on the resource.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"include_tags_with_prefix", "match_all", "resource_type", "tags", "with_deleted"},
			},

			&metadata.Action{
				Name:        "multi_add",
				Description: `Add a list of tags to a list of hrefs. The tags must be either plain_tags or machine_tags.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/tags/multi_add",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/tags/multi_add$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which the tags are to be added.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "tags[]",
						Description: `Tags to be added.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"resource_hrefs", "tags"},
			},

			&metadata.Action{
				Name:        "multi_delete",
				Description: `Delete a list of tags on a list of hrefs. The tags must be either plain_tags or machine_tags.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/tags/multi_delete",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/tags/multi_delete$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "resource_hrefs[]",
						Description: `Hrefs of the resources for which tags are to be deleted.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "tags[]",
						Description: `Tags to be deleted.`,
						Type:        "[]string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"resource_hrefs", "tags"},
			},
		},
	},
	"Task": &metadata.Resource{
		Name:        "Task",
		Description: `Tasks represent processes that happen (or have happened) asynchronously within the context of an Instance.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `Displays information of a given task within the context of an instance.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/live/tasks/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/live/tasks/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/server_arrays/:server_array_id/live/tasks/:id",
						Variables:  []string{"server_array_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/server_arrays/([^/]+)/live/tasks/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"User": &metadata.Resource{
		Name:        "User",
		Description: `A User represents an individual RightScale user`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Create a user. If a user already exists with the same email, that user will be returned.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/users",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/users$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user[identity_provider_href]",
						Description: `The RightScale API href of the Identity Provider through which this user will login to RightScale. Required to create an SSO-authenticated user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[principal_uid]",
						Description: `The principal identifier (SAML NameID or OpenID identity URL) of this user. Required to create an SSO-authenticated user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[timezone_name]",
						Description: `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[first_name]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[last_name]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[password]",
						Description: `The password of this user. Required to create a password-authenticated user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[company]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[email]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[phone]",
						Description: ``,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"user"},
			},

			&metadata.Action{
				Name:        "index",
				Description: `List the users available to the account the user is logged in to`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/users",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/users$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{"filter"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Show information about a single user.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/users/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/users/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "update",
				Description: `Update a user's contact information, change her password, or update SSO her settings`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "PUT",
						Pattern:    "/api/users/:id",
						Variables:  []string{"id"},
						Regexp:     regexp.MustCompile(`^/api/users/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "user[identity_provider_href]",
						Description: `The updated RightScale API href of the associated Identity Provider.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[current_password]",
						Description: `The current password for the user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[principal_uid]",
						Description: `The updated principal identifier (SAML NameID or OpenID identity URL) of this user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[current_email]",
						Description: `The existing email of this user.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[timezone_name]",
						Description: `This can be in the form of country/region or timezone name. For example 'America/Los_Angeles' or 'GB' or 'UTC'. A complete list of acceptable values is available in the Settings > User Settings > Preferences page.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[new_password]",
						Description: `The new password for this user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[first_name]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[last_name]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[new_email]",
						Description: `The updated email of this user.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[company]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "user[phone]",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"user"},
			},
		},
	},
	"UserData": &metadata.Resource{
		Name:        "UserData",
		Description: ``,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "show",
				Description: `No description provided for show.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/user_data/",
						Variables:  []string{},
						Regexp:     regexp.MustCompile(`^/api/user_data/$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},
		},
	},
	"Volume": &metadata.Resource{
		Name:        "Volume",
		Description: `A Volume provides a highly reliable, efficient and persistent storage solution that can be mounted to a cloud instance (in the same datacenter / zone).`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new volume.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/volumes",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume[parent_volume_snapshot_href]",
						Description: `The href of the snapshot from which the volume will be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[placement_group_href]",
						Description: `The href of the Placement Group.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[volume_type_href]",
						Description: `The href of the volume type. A Name, Resource UID and optional Size is associated with a Volume Type.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[deployment_href]",
						Description: `The href of the Deployment that owns this Volume.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[datacenter_href]",
						Description: `The href of the Datacenter / Zone that the Volume will be in. This parameter is required for non-OpenStack clouds.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[description]",
						Description: `The description of the Volume to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[encrypted]",
						Description: `A flag indicating whether Volume should be encrypted. Only available on clouds supporting volume encryption.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
					&metadata.ActionParam{
						Name:        "volume[iops]",
						Description: `The number of IOPS (I/O Operations Per Second) this Volume should support. Only available on clouds supporting performance provisioning.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[name]",
						Description: `The name for the Volume to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume[size]",
						Description: `The size of a Volume to be created in gigabytes (GB). Some Volume Types have predefined sizes and do not allow selecting a custom size on Volume creation.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"volume"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given volume.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/volumes/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists volumes.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volumes",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single volume.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volumes/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default", "extended"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"VolumeAttachment": &metadata.Resource{
		Name:        "VolumeAttachment",
		Description: `A VolumeAttachment represents a relationship between a volume and an instance.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new volume attachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/volume_attachments",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/volume_attachments",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_attachment",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachment$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume_attachment[instance_href]",
						Description: `The href of the instance to which the volume will be attached to.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_attachment[volume_href]",
						Description: `The href of the volume to be attached.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_attachment[device]",
						Description: `The device location where the volume will be mounted. Value must be of format /dev/xvd[bcefghij]. This is not reliable and will be deprecated.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"volume_attachment"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given volume attachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/volume_attachments/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_attachment",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachment$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "force",
						Description: `Specifies whether to force the detachment of a volume.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"true", "false"},
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists all volume attachments.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/volume_attachments",
						Variables:  []string{"cloud_id", "instance_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volume_attachments",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_attachments$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single volume attachment.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/instances/:instance_id/volume_attachments/:id",
						Variables:  []string{"cloud_id", "instance_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/instances/([^/]+)/volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volume_attachments/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_attachments/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_attachments",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachments$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_attachment",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_attachment$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"VolumeSnapshot": &metadata.Resource{
		Name:        "VolumeSnapshot",
		Description: `A VolumeSnapshot represents a Cloud storage volume at a particular point in time`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "create",
				Description: `Creates a new volume_snapshot.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots$`),
					},
					&metadata.PathPattern{
						HttpMethod: "POST",
						Pattern:    "/api/clouds/:cloud_id/volume_snapshots",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "volume_snapshot[parent_volume_href]",
						Description: `The href of the Volume from which the Volume Snapshot will be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_snapshot[deployment_href]",
						Description: `The href of the Deployment that owns this Volume Snapshot.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_snapshot[description]",
						Description: `The description for the Volume Snapshot to be created.`,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "volume_snapshot[name]",
						Description: `The name for the Volume Snapshot to be created.`,
						Type:        "string",
						Mandatory:   true,
						NonBlank:    true,
					},
				},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{"volume_snapshot"},
			},

			&metadata.Action{
				Name:        "destroy",
				Description: `Deletes a given volume_snapshot.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "DELETE",
						Pattern:    "/api/clouds/:cloud_id/volume_snapshots/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)$`),
					},
				},
				Params:            []*metadata.ActionParam{},
				QueryParamNames:   []string{},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "index",
				Description: `Lists all volume_snapshots.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots",
						Variables:  []string{"cloud_id", "volume_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volume_snapshots",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single volume_snapshot.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volumes/:volume_id/volume_snapshots/:id",
						Variables:  []string{"cloud_id", "volume_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volumes/([^/]+)/volume_snapshots/([^/]+)$`),
					},
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volume_snapshots/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_snapshots/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
	"VolumeType": &metadata.Resource{
		Name:        "VolumeType",
		Description: `A VolumeType describes the type of volume, particularly the size.`,
		Actions: []*metadata.Action{
			&metadata.Action{
				Name:        "index",
				Description: `Lists Volume Types.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volume_types",
						Variables:  []string{"cloud_id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_types$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "filter[]",
						Description: ``,
						Type:        "[]string",
						Mandatory:   false,
						NonBlank:    true,
					},
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"filter, view"},
				PayloadParamNames: []string{},
			},

			&metadata.Action{
				Name:        "show",
				Description: `Displays information about a single Volume Type.`,
				PathPatterns: []*metadata.PathPattern{
					&metadata.PathPattern{
						HttpMethod: "GET",
						Pattern:    "/api/clouds/:cloud_id/volume_types/:id",
						Variables:  []string{"cloud_id", "id"},
						Regexp:     regexp.MustCompile(`^/api/clouds/([^/]+)/volume_types/([^/]+)$`),
					},
				},
				Params: []*metadata.ActionParam{
					&metadata.ActionParam{
						Name:        "view",
						Description: ``,
						Type:        "string",
						Mandatory:   false,
						NonBlank:    true,
						ValidValues: []string{"default"},
					},
				},
				QueryParamNames:   []string{"view"},
				PayloadParamNames: []string{},
			},
		},
	},
}
