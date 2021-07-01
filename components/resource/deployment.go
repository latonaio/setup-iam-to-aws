package resource

type Region struct {
	ResourceName string `json:"ResourceName"`
	KeyPairName  string `json:"keyPairName"`
	Location     string `json:"Location"`
	Iam          *Iam   `json:"iam"`
}

type Deployment struct {
}

func newRegion(region Region) *Region {
	return &Region{
		ResourceName: region.ResourceName,
		KeyPairName:  region.KeyPairName,
		Location:     region.Location,
		Iam:          region.Iam,
	}
}
