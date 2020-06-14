package thumbnail

import "log"

func ImageFile(infile string) (string, error) {}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}
