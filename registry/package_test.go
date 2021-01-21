package registry

import (
	"fmt"
	"testing"
)

func TestListPackages(t *testing.T) {

	packages, err := ListPackages("steampipe", "us", "plugins")

	if err != nil {
		t.Errorf("ListPackages failed with error: %s", err)
	}

	for _, pkg := range packages {
		fmt.Printf("\nImageName: %s", pkg.ImageName)
		fmt.Printf("\nName: %s", pkg.Name)
		fmt.Printf("\nCreateTime: %s", pkg.CreateTime)
		fmt.Printf("\nUpdateTime: %s\n", pkg.UpdateTime)
	}

}

func TestVersions(t *testing.T) {

	packages, err := ListPackages("steampipe", "us", "plugins")

	if err != nil {
		t.Errorf("ListPackages failed with error: %s", err)
	}

	for _, pkg := range packages {
		fmt.Printf("\nName: %s", pkg.Name)
		versions, err := pkg.Versions()
		if err != nil {
			t.Errorf("Versions failed with error: %s", err)
		}

		fmt.Printf("\n\tVersions:")

		for _, v := range versions {
			fmt.Printf("\n\t\tDigest: %s", v.Digest)
			fmt.Printf("\n\t\tName: %s", v.Name)
			fmt.Printf("\n\t\tCreateTime: %s", v.CreateTime)
			fmt.Printf("\n\t\tUpdateTime: %s", v.UpdateTime)
		}
	}

}

func TestTags(t *testing.T) {

	packages, err := ListPackages("steampipe", "us", "plugins")

	if err != nil {
		t.Errorf("ListPackages failed with error: %s", err)
	}

	for _, pkg := range packages {
		fmt.Printf("\nName: %s", pkg.Name)

		fmt.Printf("\n\tTags:")
		tags, err := pkg.Tags()
		if err != nil {
			t.Errorf("Tag() failed with error: %s", err)
		}

		for _, t := range tags {
			fmt.Printf("\n\t\tTag: %s", t.Tag)
			fmt.Printf("\n\t\tName: %s", t.Name)
			fmt.Printf("\n\t\tversion: %s", t.Version)
		}
	}

}
