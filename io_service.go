package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"path"
	"strconv"
)

func splitFrames(tempFolder string, inputFileName string, fps int) ([]string, error) {
	cmd := exec.Command(
		"ffmpeg",
		"-i", inputFileName,
		"-vf",
		"fps="+strconv.Itoa(fps),
		path.Join(tempFolder, "%04d.png")) // 4 digits for the frame number

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("in splitFrames: %w", err)
	}

	files, err := os.ReadDir(tempFolder)
	if err != nil {
		return nil, fmt.Errorf("in splitFrames: %w", err)
	}

	filesPath := make([]string, len(files))
	for i := range files {
		filesPath[i] = path.Join(tempFolder, files[i].Name())
	}

	return filesPath, nil
}

func decodeFrame(imagePath string) (*image.RGBA, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("in decodeFrame: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("in decodeFrame: %w", err)
	}

	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		return nil, fmt.Errorf("in decodeFrame: %w", err)
	}

	return rgbaImg, nil
}

// saveImage saves an RGBA image to a file
func saveFrame(imagePath string, img *image.RGBA) error {
	outFile, err := os.Create(imagePath)
	if err != nil {
		return fmt.Errorf("in saveFrame: %w", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		return fmt.Errorf("in saveFrame: %w", err)
	}
	return nil
}

func joinFrames(tempFolder string, outputFileName string, fps int) error {
	cmd := exec.Command(
		"ffmpeg",
		"-framerate", strconv.Itoa(fps),
		"-y", // Replace output file without asking
		"-i", path.Join(tempFolder, "%04d.png"),
		"-c:v", "libx264",
		"-r", strconv.Itoa(fps),
		"-pix_fmt", "yuv420p",
		outputFileName,
	)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("in joinFrame: %w", err)
	}

	fmt.Println("Video created successfully: output.mp4")

	return nil
}

func createTemporaryFolder() (string, error) {
	folder, err := os.MkdirTemp("", "temp_img_folder")
	if err != nil {
		return "", fmt.Errorf("in createTemporaryFolder: %w", err)
	}
	return folder, nil
}

func deleteTemporaryFolder(tempFolder string) error {
	err := os.RemoveAll(tempFolder)
	if err != nil {
		return fmt.Errorf("in deleteTemporaryFolder: %w", err)
	}
	return nil
}
