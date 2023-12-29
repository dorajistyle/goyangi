package aws

import (
	"github.com/goamz/goamz/aws"
	"github.com/spf13/viper"
)

// Auth return the aws authentication.
func Auth() aws.Auth {
	return aws.Auth{
		AccessKey: viper.GetString("aws.accessKeyID"),
		SecretKey: viper.GetString("aws.secretAccessKey"),
	}
}

// Region return the aws region from string.
func Region(regionName string) aws.Region {
	switch regionName {
	case "APNortheast":
		return aws.APNortheast
	case "APSoutheast":
		return aws.APSoutheast
	case "APSoutheast2":
		return aws.APSoutheast2
	case "EUCentral":
		return aws.EUCentral
	case "EUWest":
		return aws.EUWest
	case "USEast":
		return aws.USEast
	case "USWest":
		return aws.USWest
	case "USWest2":
		return aws.USWest2
	case "USGovWest":
		return aws.USGovWest
	case "SAEast":
		return aws.SAEast
		// case "CNNorth":
		// return aws.CNNorth
	}
	return aws.Region{}
}
