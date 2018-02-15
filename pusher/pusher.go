package pusher

import (
	"io"
	"os"
	"time"

	"github.com/micahwedemeyer/gphoto2go"

	"github.com/kochman/hotshots/config"
	"github.com/kochman/hotshots/log"
)

type Pusher struct {
	cfg config.Config
}

func New(cfg *config.Config) (*Pusher, error) {
	p := &Pusher{
		cfg: *cfg,
	}

	return p, nil
}

func (p *Pusher) Run() {
	ticker := time.Tick(time.Second)
	for range ticker {
		camera := &gphoto2go.Camera{}
		err := camera.Init()
		if err < 0 {
			log.Errorf("Unable to initialize camera: %s", gphoto2go.CameraResultToString(err))
		}

		folders := camera.RListFolders("/")
		for _, folder := range folders {
			files, _ := camera.ListFiles(folder)
			for _, fileName := range files {
				destFileName := "tmp/" + fileName
				if _, err := os.Stat(destFileName); err == nil {
					// file exists, don't copy
					continue
				}
				fileWriter, err := os.Create(destFileName)
				if err != nil {
					log.WithError(err).Error("Unable to create file")
					continue
				}
				cameraFileReader := camera.FileReader(folder, fileName)
				defer cameraFileReader.Close()
				_, err = io.Copy(fileWriter, cameraFileReader)
				if err != nil {
					log.WithError(err).Error("Unable to copy photo")
					continue
				}
				log.Infof("Copied photo to %s", destFileName)

				// Important, since there is memory used in the transfer that needs to be freed up
			}
		}
		err = camera.Exit()
		if err < 0 {
			log.Errorf("Unable to exit camera: %s", gphoto2go.CameraResultToString(err))
		}
	}
}
