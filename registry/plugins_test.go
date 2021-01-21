package registry

import (
	"fmt"
	"testing"
)

func TestListPlugins(t *testing.T) {
	plugins, err := ListPlugins()
	if err != nil {
		t.Errorf("ListPlugins failed with error: %s", err)
	}

	for _, p := range plugins {
		fmt.Printf("\nName: %s", p.Name)
		fmt.Printf("\nCreateTime: %s", p.CreateTime)
		fmt.Printf("\nUpdateTime: %s", p.UpdateTime)

		versions, err := p.Versions()
		if err != nil {
			t.Errorf("ListPlugins - p.Versions failed with error: %s", err)
		}
		for _, v := range versions {
			fmt.Printf("\n\tImageName: %s", v.ImageName)
			fmt.Printf("\n\tDigest: %s", v.Digest)
			fmt.Printf("\n\tCreateTime: %s", v.CreateTime)
			fmt.Printf("\n\tUpdateTime: %s", v.UpdateTime)
			fmt.Printf("\n\tTags:")
			for _, t := range v.Tags {
				fmt.Printf("\n\t\t%+s", t)

			}
		}
	}

}
