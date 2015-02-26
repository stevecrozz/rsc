package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/rightscale/rsc/gen"
)

// The analyzer struct holds the analysis results
type ApiAnalyzer struct {
	// Raw resources as defined in API json metadata
	RawResources map[string]map[string]interface{}
	// Raw types as defined in API json metadata
	RawTypes map[string]map[string]interface{}
	// Version being analyzed
	Version string
	// Name of golang struct to generate for API client
	ClientName string

	// Temporary data structures used by analysis

	// Descriptor being built
	descriptor *gen.ApiDescriptor
	// Map of ruby type name to go type name
	// Also used to record types that have been processed.
	typeNames map[string]string
	// Map of uids to types, used to create new types from inline structs or hashes
	typeUids map[string]*gen.ObjectDataType
}

// Factory method for API analyzer
func NewApiAnalyzer(version, clientName string, resources, types map[string]map[string]interface{}) *ApiAnalyzer {
	return &ApiAnalyzer{
		RawResources: resources,
		RawTypes:     types,
		Version:      version,
		ClientName:   clientName,
		typeNames:    make(map[string]string),
		typeUids:     make(map[string]*gen.ObjectDataType),
	}
}

// Create API descriptor from raw resources and types
func (a *ApiAnalyzer) Analyze() (*gen.ApiDescriptor, error) {
	var descriptor = gen.ApiDescriptor{
		Version:   a.Version,
		Resources: make(map[string]*gen.Resource),
		Types:     make(map[string]*gen.ObjectDataType),
	}
	a.descriptor = &descriptor

	// Sort resource names so iterations are always done in the same order
	var rawResourceNames = make([]string, len(a.RawResources))
	var idx = 0
	for name, _ := range a.RawResources {
		rawResourceNames[idx] = name
		idx += 1
	}
	sort.Strings(rawResourceNames)

	// 1. Do a first pass to collect all media types (resource and responses) so the go type
	// names are the most natural.
	for _, name := range rawResourceNames {
		var resource = a.RawResources[name]
		var err = a.AnalyzeMediaType(resource["media_type"].(string))
		if err != nil {
			return nil, err
		}
		actions, ok := resource["actions"]
		if !ok {
			continue
		}
		for _, ac := range actions.([]interface{}) {
			var action = ac.(map[string]interface{})
			responses, ok := action["responses"]
			if ok {
				for _, r := range responses.(map[string]interface{}) {
					mediaType, ok := r.(map[string]interface{})["media_type"]
					if ok {
						var m = mediaType.(map[string]interface{})
						err := a.AnalyzeMediaType(m["name"].(string))
						if err != nil {
							return nil, err
						}
					}
				}
			}
		}
	}

	// 2. Now do another pass to analyze each resource
	for _, name := range rawResourceNames {
		var err = a.AnalyzeResource(name, a.RawResources[name], &descriptor)
		if err != nil {
			return nil, err
		}
	}

	// 3. Sort and register all types
	var typeNames = make([]string, len(a.typeNames))
	idx = 0
	for _, name := range a.typeNames {
		var uniq = makeUniq(name, a.descriptor.ResourceNames)
		if uniq != name {
			var existing = descriptor.Types[name]
			existing.Name = uniq
			delete(descriptor.Types, name)
			descriptor.Types[uniq] = existing
		}
		typeNames[idx] = uniq
		idx += 1
	}
	sort.Strings(typeNames)
	descriptor.TypeNames = typeNames

	// We're done
	return &descriptor, nil
}

// Analyze media type, recurse through all types and create corresponding ObjectDataTypes and
// ActionParams.
func (a *ApiAnalyzer) AnalyzeMediaType(name string) error {
	var mtype, ok = a.RawTypes[name]
	if !ok {
		return fmt.Errorf("Unknown media type %s", name)
	}
	if a.GetType(name) != nil {
		return nil // Already analyzed
	}
	var obj = a.CreateType(name)
	var attributes = mtype["attributes"].(map[string]interface{})
	var fields = make([]*gen.ActionParam, len(attributes))
	var idx = 0
	for _, attName := range sortedKeys(attributes) {
		var att = attributes[attName]
		var param, err = a.AnalyzeAttribute(attName, attName, att.(map[string]interface{}))
		if err != nil {
			return err
		}
		fields[idx] = param
		idx += 1
	}
	obj.Fields = fields

	return nil
}

// GetType looks up a generated type by metadata name.
// Returns nil if not found.
func (a *ApiAnalyzer) GetType(metadataName string) *gen.ObjectDataType {
	if existing, ok := a.typeNames[metadataName]; ok {
		return a.descriptor.Types[existing]
	}
	return nil
}

// Creates a unique go type name, records it and returns a new type with that name.
func (a *ApiAnalyzer) CreateType(metadataName string) *gen.ObjectDataType {
	var taken = make([]string, len(a.typeNames))
	var idx = 0
	for _, n := range a.typeNames {
		taken[idx] = n
		idx += 1
	}
	var goName = makeUniq(toGoTypeName(metadataName, false), taken)

	a.typeNames[metadataName] = goName
	var obj = gen.ObjectDataType{Name: goName}
	a.descriptor.Types[goName] = &obj

	return &obj
}

// Use given unique id to look up existing type and create new one if not found
func (a *ApiAnalyzer) GetOrCreate(uid, metadataName string) *gen.ObjectDataType {
	if existing, ok := a.typeUids[uid]; ok {
		return existing
	}
	var obj = a.CreateType(metadataName)
	a.typeUids[uid] = obj
	return obj
}

// Make a unique name given a prefix and a set of names
func makeUniq(base string, taken []string) string {
	var idx = 1
	var uniq = base
	var inuse = true
	for inuse {
		inuse = false
		for _, gn := range taken {
			if gn == uniq {
				inuse = true
				break
			}
		}
		if inuse {
			idx += 1
			uniq = base + strconv.Itoa(idx)
		}
	}
	return uniq
}
