package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"github.com/corona10/goimagehash"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:  "imghash [file] [file]",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				return run(args[0])
			}

			return compare(args[0], args[1])
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func compare(path1, path2 string) error {
	img1, err := read(path1)
	if err != nil {
		return err
	}

	img2, err := read(path2)
	if err != nil {
		return err
	}

	avg1, err := goimagehash.AverageHash(img1)
	if err != nil {
		return fmt.Errorf("average hash: %w", err)
	}

	avg2, err := goimagehash.AverageHash(img2)
	if err != nil {
		return fmt.Errorf("average hash: %w", err)
	}

	diff1, err := goimagehash.DifferenceHash(img1)
	if err != nil {
		return fmt.Errorf("average hash: %w", err)
	}

	diff2, err := goimagehash.DifferenceHash(img2)
	if err != nil {
		return fmt.Errorf("average hash: %w", err)
	}

	perception1, err := goimagehash.PerceptionHash(img1)
	if err != nil {
		return fmt.Errorf("average hash: %w", err)
	}

	perception2, err := goimagehash.PerceptionHash(img2)
	if err != nil {
		return fmt.Errorf("average hash: %w", err)
	}

	avgDistance, err := avg1.Distance(avg2)
	if err != nil {
		return err
	}

	diffDistance, err := diff1.Distance(diff2)
	if err != nil {
		return err
	}

	perceptionDistance, err := perception1.Distance(perception2)
	if err != nil {
		return err
	}

	fmt.Printf("a:%d\n", avgDistance)
	fmt.Printf("d:%d\n", diffDistance)
	fmt.Printf("p:%d\n", perceptionDistance)

	return nil
}

func read(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	return img, nil
}

func run(path string) error {
	img, err := read(path)
	if err != nil {
		return err
	}

	avg, err := goimagehash.AverageHash(img)
	if err != nil {
		return fmt.Errorf("average hash: %w", err)
	}

	diff, err := goimagehash.DifferenceHash(img)
	if err != nil {
		return fmt.Errorf("difference hash: %w", err)
	}

	perception, err := goimagehash.PerceptionHash(img)
	if err != nil {
		return fmt.Errorf("perception hash: %w", err)
	}

	md5hash, err := getImageHash(path)
	if err != nil {
		return fmt.Errorf("md5 hash: %w", err)
	}

	fmt.Println("m:" + md5hash)
	fmt.Println(avg.ToString())
	fmt.Println(diff.ToString())
	fmt.Println(perception.ToString())

	return nil
}

func getImageHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
