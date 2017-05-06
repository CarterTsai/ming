package command

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/jung-kurt/gofpdf"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

// Create create command
func Create() cli.Command {

	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	return cli.Command{
		Name: "create",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "imgPath, p",
				Value: "./img",
			},
			cli.Float64Flag{
				Name:  "columnNum, c",
				Value: 2,
			},
			cli.Float64Flag{
				Name:  "margin, o",
				Value: 0,
			},
			cli.Float64Flag{
				Name:  "xInitPosition, x",
				Value: 10,
			},
			cli.Float64Flag{
				Name:  "yInitPosition, y",
				Value: 10,
			},
			cli.BoolFlag{
				Name:   "debug, d",
				Hidden: true,
			},
		},
		Action: func(c *cli.Context) error {

			var root, _ = os.Getwd()

			var imageMapPath = make(map[string][]string)

			files, e := ioutil.ReadDir(c.String("imgPath"))

			if e != nil {
				sugar.Info(e.Error())
			}

			for _, f := range files {

				err := filepath.Walk(filepath.Join(root, c.String("imgPath"), f.Name()), func(imgFilePath string, info os.FileInfo, err error) error {

					_, imgErr := imgio.Open(imgFilePath)
					if imgErr != nil {
						return nil
					}

					imageMapPath[f.Name()] = append(imageMapPath[f.Name()], imgFilePath)

					if c.Bool("debug") {
						sugar.Info(imgFilePath)
					}
					return nil
				})

				if c.Bool("debug") {
					sugar.Info(imageMapPath)
				}

				pdf := gofpdf.New("P", "mm", "A4", "")
				pdf.AddPage()

				w, h := pdf.GetPageSize()

				if c.Bool("debug") {
					sugar.Info(w)
					sugar.Info(h)
				}

				// parameter
				columnNum := c.Float64("columnNum")
				offset := c.Float64("margin")
				imageWidth := float64((w - (offset * 2) - 20) / columnNum)
				imageHeight := float64(0)
				xInitPosition := c.Float64("xInitPosition")
				yInitPosition := c.Float64("yInitPosition")
				columnCount := float64(0)

				if f.Name() == ".DS_Store" {
					continue
				}

				fmt.Printf("Create %s.pdf\n", filepath.Base(f.Name()))

				for _, v := range imageMapPath {
					imageHeight = h / float64(len(v))
					for i, v1 := range v {

						if i > 0 && math.Mod(float64(i), columnNum) == 0 {
							columnCount++
						}

						x := xInitPosition + math.Mod(float64(i), columnNum)*(imageWidth+offset)
						y := yInitPosition + imageHeight*columnCount

						pdf.Image(v1, x, y, imageWidth, imageHeight, false, "", 0, "")
						fmt.Println(v1)
						if c.Bool("debug") {
							sugar.Infof("i %s\n", i)
							sugar.Infof("x %s\n", x)
							sugar.Infof("y %s\n", y)
							sugar.Infof("width %s\n", imageWidth)
							sugar.Infof("mod %s\n", math.Mod(float64(i), columnNum))
							sugar.Infof("count %s\n", columnCount)
							sugar.Info(y)
						}
					}
				}

				imageMapPath = make(map[string][]string)
				local := filepath.Join(root, f.Name()+".pdf")

				fmt.Println(local)

				if c.Bool("debug") {
					sugar.Info(local)
				}

				erri := pdf.OutputFileAndClose(local)

				if erri != nil {
					panic(erri)
				}

				if err != nil {
					panic(err)
				}
			}

			return nil
		},
	}
}
