package lib

import (
	"fmt"
	"github.com/ser163/png2j"
	"image/color"
	"io/ioutil"
	"log"
	"os"
)

// Converts a PNG file to a JPG file.
func ConvertToOrgJpg(src string, dst string, quality uint8) error {
	// Check file existence
	if !checkFileExist(src) {
		return fmt.Errorf("source file %s not found", src)
	}
	var qua = int(quality)
	var bg color.Color
	bg = color.White
	// Convert
	err := png2j.ConBase(src, dst, bg, qua)
	if err != nil {
		return err
	}
	return nil
}

// Converts a PNG file to a JPG file.
func ConvertToJpg(src string, dst string, width uint, height uint, quality uint8) error {
	// Check file existence
	if !checkFileExist(src) {
		return fmt.Errorf("source file %s not found", src)
	}

	var qua = int(quality)
	var bg color.Color
	bg = color.White

	// Resize
	tmpfile, err := ioutil.TempFile("", "png2j-19860505-")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpfile.Close()
	if err != nil {
		return err
	}

	err = png2j.ConBase(src, tmpfile.Name(), bg, qua)
	if err != nil {
		return err
	}

	err = png2j.ReSizeImage(tmpfile.Name(), width, height, dst)
	if err != nil {
		return err
	}
	err = os.Remove(tmpfile.Name())
	if err != nil {
		return err
	}

	return nil
}

// Check file existence
func checkFileExist(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
