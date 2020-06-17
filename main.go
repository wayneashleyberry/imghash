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
		Use:  "imghash [file]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(args[0])
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("decode image: %w", err)
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

	fmt.Println(md5hash)
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
