package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		clientOV                     *ov.OVClient
		server_profile_template_name = "Test SPT"
		enclosure_group_name         = "TESTEG"
		server_hardware_type_name    = "SY 480 Gen9 1"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		2200,
		"*")

	server_hardware_type, err := ovc.GetServerHardwareTypeByName(server_hardware_type_name)

	enc_grp, err := ovc.GetEnclosureGroupByName(enclosure_group_name)

	conn_settings := ov.ConnectionSettings{
		ManageConnections: true,
	}
	initialScopeUris := new([]utils.Nstring)
	scp, scperr := ovc.GetScopeByName("ScopeTest")
	if scperr != nil {
		*initialScopeUris = append(*initialScopeUris, scp.URI)
	}

	server_profile_template := ov.ServerProfile{
		Type:                  "ServerProfileTemplateV8",
		Name:                  server_profile_template_name,
		EnclosureGroupURI:     enc_grp.URI,
		ServerHardwareTypeURI: server_hardware_type.URI,
		ConnectionSettings:    conn_settings,
		InitialScopeUris:      *initialScopeUris,
	}

	err = ovc.CreateProfileTemplate(server_profile_template)
	if err != nil {
		fmt.Println("Server Profile Template Creation Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template Created---------------#")
	}

	sort := "name:asc"
	spt_list, err := ovc.GetProfileTemplates("", "", "", sort, "")
	if err != nil {
		fmt.Println("Server Profile Template Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template List---------------#")

		for i := 0; i < len(spt_list.Members); i++ {
			fmt.Println(spt_list.Members[i].Name)
		}
	}

	spt, err := ovc.GetProfileTemplateByName(server_profile_template_name)
	if err != nil {
		fmt.Println("Server Profile Template Retrieval By Name Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template by Name---------------#")
		fmt.Println(spt.Name)
	}

	spt.Name = "Renamed Test SPT"
	err = ovc.UpdateProfileTemplate(spt)
	if err != nil {
		fmt.Println("Server Profile Template Updation Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template Updated---------------#")
	}

	spt_list, err = ovc.GetProfileTemplates("", "", "", sort, "")
	if err != nil {
		fmt.Println("Server Profile Template Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template List---------------#")

		for i := 0; i < len(spt_list.Members); i++ {
			fmt.Println(spt_list.Members[i].Name)
		}
	}

	err = ovc.DeleteProfileTemplate(spt.Name)
	if err != nil {
		fmt.Println("Server Profile Template Delete Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template Deleted---------------#")
	}
}
