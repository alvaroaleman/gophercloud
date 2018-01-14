/*
Package hypervisors returns details about the hypervisors and shows
summary statistics for all hypervisors over all compute nodes in the OpenStack cloud.

Example of Retrieving Details of All Hypervisors

	allPages, err := hypervisors.List(computeClient).AllPages()
	if err != nil {
		panic(err)
	}

	allHypervisors, err := hypervisors.ExtractHypervisors(allPages)
	if err != nil {
		panic(err)
	}

	for _, hypervisor := range allHypervisors {
		fmt.Printf("%+v\n", hypervisor)
	}

Example of Show Hypervisor Statistics

	hypervisorsStatistics, err := hypervisors.GetStatistics(computeClient).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", hypervisorsStatistics)

*/
package hypervisors
