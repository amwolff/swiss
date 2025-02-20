// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package swb

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type InstantDeliveryServersInput struct {
	Name_in     []string `json:"name_in"`
	Location_in []string `json:"location_in"`
	Region_in   []string `json:"region_in"`
}

// GetName_in returns InstantDeliveryServersInput.Name_in, and is useful for accessing the field via an interface.
func (v *InstantDeliveryServersInput) GetName_in() []string { return v.Name_in }

// GetLocation_in returns InstantDeliveryServersInput.Location_in, and is useful for accessing the field via an interface.
func (v *InstantDeliveryServersInput) GetLocation_in() []string { return v.Location_in }

// GetRegion_in returns InstantDeliveryServersInput.Region_in, and is useful for accessing the field via an interface.
func (v *InstantDeliveryServersInput) GetRegion_in() []string { return v.Region_in }

type PaginatedInstantDeliveryServersInput struct {
	// First page has pageIndex value `0`.
	PageIndex int `json:"pageIndex"`
	// The maximum number of results per page is `50`.
	PageSize int                         `json:"pageSize"`
	Filter   InstantDeliveryServersInput `json:"filter"`
}

// GetPageIndex returns PaginatedInstantDeliveryServersInput.PageIndex, and is useful for accessing the field via an interface.
func (v *PaginatedInstantDeliveryServersInput) GetPageIndex() int { return v.PageIndex }

// GetPageSize returns PaginatedInstantDeliveryServersInput.PageSize, and is useful for accessing the field via an interface.
func (v *PaginatedInstantDeliveryServersInput) GetPageSize() int { return v.PageSize }

// GetFilter returns PaginatedInstantDeliveryServersInput.Filter, and is useful for accessing the field via an interface.
func (v *PaginatedInstantDeliveryServersInput) GetFilter() InstantDeliveryServersInput {
	return v.Filter
}

// __availableMachinesInput is used internally by genqlient
type __availableMachinesInput struct {
	Input PaginatedInstantDeliveryServersInput `json:"input"`
}

// GetInput returns __availableMachinesInput.Input, and is useful for accessing the field via an interface.
func (v *__availableMachinesInput) GetInput() PaginatedInstantDeliveryServersInput { return v.Input }

// availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse includes the requested fields of the GraphQL type PaginatedInstantDeliveryServerResponse.
type availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse struct {
	// Total number of items in the full result set.
	EntriesTotalCount int `json:"entriesTotalCount"`
	// Total number of pages which constitute the full result set.
	PageCount int `json:"pageCount"`
	// Current page index which was returned, the first index is 0.
	CurrentPageIndex int `json:"currentPageIndex"`
	// Number of items per page.
	PageSize int `json:"pageSize"`
	// Index of the next page, when none is available `null`.
	NextPageIndex int `json:"nextPageIndex"`
	// Index of the previous page, first index is 0, if currently on first page, the value will be `null`.
	PreviousPageIndex int `json:"previousPageIndex"`
	// Indicates whether this is the last page.
	IsLastPage bool `json:"isLastPage"`
	// Indicates whether this is the first page.
	IsFirstPage bool `json:"isFirstPage"`
	// Resulting paginated items.
	Entries []availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer `json:"entries"`
}

// GetEntriesTotalCount returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.EntriesTotalCount, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetEntriesTotalCount() int {
	return v.EntriesTotalCount
}

// GetPageCount returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.PageCount, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetPageCount() int {
	return v.PageCount
}

// GetCurrentPageIndex returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.CurrentPageIndex, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetCurrentPageIndex() int {
	return v.CurrentPageIndex
}

// GetPageSize returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.PageSize, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetPageSize() int {
	return v.PageSize
}

// GetNextPageIndex returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.NextPageIndex, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetNextPageIndex() int {
	return v.NextPageIndex
}

// GetPreviousPageIndex returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.PreviousPageIndex, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetPreviousPageIndex() int {
	return v.PreviousPageIndex
}

// GetIsLastPage returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.IsLastPage, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetIsLastPage() bool {
	return v.IsLastPage
}

// GetIsFirstPage returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.IsFirstPage, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetIsFirstPage() bool {
	return v.IsFirstPage
}

// GetEntries returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse.Entries, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse) GetEntries() []availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer {
	return v.Entries
}

// availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer includes the requested fields of the GraphQL type InstantDeliveryServer.
type availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer struct {
	// The unique identifier of the server.
	Name string `json:"name"`
	// The location of the server.
	Location availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation `json:"location"`
	// Overall server uplink capacity in Gbps.
	UplinkCapacity int `json:"uplinkCapacity"`
	// Current hardware configuration of the server.
	Hardware availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardware `json:"hardware"`
}

// GetName returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer.Name, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer) GetName() string {
	return v.Name
}

// GetLocation returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer.Location, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer) GetLocation() availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation {
	return v.Location
}

// GetUplinkCapacity returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer.UplinkCapacity, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer) GetUplinkCapacity() int {
	return v.UplinkCapacity
}

// GetHardware returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer.Hardware, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServer) GetHardware() availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardware {
	return v.Hardware
}

// availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardware includes the requested fields of the GraphQL type Hardware.
type availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardware struct {
	Cpus []availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareCpusCpu `json:"cpus"`
	Rams []availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareRamsRam `json:"rams"`
}

// GetCpus returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardware.Cpus, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardware) GetCpus() []availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareCpusCpu {
	return v.Cpus
}

// GetRams returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardware.Rams, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardware) GetRams() []availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareRamsRam {
	return v.Rams
}

// availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareCpusCpu includes the requested fields of the GraphQL type Cpu.
type availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareCpusCpu struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// GetName returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareCpusCpu.Name, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareCpusCpu) GetName() string {
	return v.Name
}

// GetCount returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareCpusCpu.Count, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareCpusCpu) GetCount() int {
	return v.Count
}

// availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareRamsRam includes the requested fields of the GraphQL type Ram.
type availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareRamsRam struct {
	Count int `json:"count"`
	// RAM size in gigabytes (GB).
	Size int `json:"size"`
}

// GetCount returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareRamsRam.Count, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareRamsRam) GetCount() int {
	return v.Count
}

// GetSize returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareRamsRam.Size, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerHardwareRamsRam) GetSize() int {
	return v.Size
}

// availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation includes the requested fields of the GraphQL type Location.
type availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation struct {
	// Location name.
	Name string `json:"name"`
	// Location region.
	Region string `json:"region"`
	// Location short name.
	Short string `json:"short"`
}

// GetName returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation.Name, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation) GetName() string {
	return v.Name
}

// GetRegion returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation.Region, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation) GetRegion() string {
	return v.Region
}

// GetShort returns availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation.Short, and is useful for accessing the field via an interface.
func (v *availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponseEntriesInstantDeliveryServerLocation) GetShort() string {
	return v.Short
}

// availableMachinesResponse is returned by availableMachines on success.
type availableMachinesResponse struct {
	// List of all in-stock servers available for instant delivery.<br>You can look up servers by location, region, or hardware configuration.
	InstantDeliveryServers availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse `json:"instantDeliveryServers"`
}

// GetInstantDeliveryServers returns availableMachinesResponse.InstantDeliveryServers, and is useful for accessing the field via an interface.
func (v *availableMachinesResponse) GetInstantDeliveryServers() availableMachinesInstantDeliveryServersPaginatedInstantDeliveryServerResponse {
	return v.InstantDeliveryServers
}

// The query or mutation executed by availableMachines.
const availableMachines_Operation = `
query availableMachines ($input: PaginatedInstantDeliveryServersInput) {
	instantDeliveryServers(input: $input) {
		entriesTotalCount
		pageCount
		currentPageIndex
		pageSize
		nextPageIndex
		previousPageIndex
		isLastPage
		isFirstPage
		entries {
			name
			location {
				name
				region
				short
			}
			uplinkCapacity
			hardware {
				cpus {
					name
					count
				}
				rams {
					count
					size
				}
			}
		}
	}
}
`

func availableMachines(
	ctx_ context.Context,
	client_ graphql.Client,
	input PaginatedInstantDeliveryServersInput,
) (*availableMachinesResponse, error) {
	req_ := &graphql.Request{
		OpName: "availableMachines",
		Query:  availableMachines_Operation,
		Variables: &__availableMachinesInput{
			Input: input,
		},
	}
	var err_ error

	var data_ availableMachinesResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
