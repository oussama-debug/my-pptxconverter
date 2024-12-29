/*
 Copyright 2024 Oussama Jaaouani <oussama@jaaouani.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package presentation

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/oussama-debug/pptx/internal/logger"
	models "github.com/oussama-debug/pptx/internal/presentation/models"
)

type PPTXDocument struct {
	path   string
	files  map[string]*zip.File
	medias map[string]*zip.File
	slides map[string]*zip.File
	size   uint16
	info   *models.PPTXInformation
	reader *zip.Reader
}

func NewPPTXDocument(path string) (*PPTXDocument, error) {
	logger := logger.Get()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("PPTXDocument %s does not exist", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(file, fileInfo.Size())

	pptxdoc := &PPTXDocument{
		path:   path,
		reader: reader,
		files:  make(map[string]*zip.File),
		medias: make(map[string]*zip.File),
		slides: make(map[string]*zip.File),
		size:   uint16(fileInfo.Size()),
		info:   models.NewPPTXInformation(),
	}

	logger.Info().Msgf("Found in PPTXDocument: %d file(s)", len(reader.File))

	for _, file := range reader.File {
		if strings.HasPrefix(file.Name, "ppt/media") {
			pptxdoc.medias[file.Name] = file
		} else if strings.HasPrefix(file.Name, "ppt/slides") {
			pptxdoc.slides[file.Name] = file
		} else {
			pptxdoc.files[file.Name] = file
			logger.Debug().Msgf("File in PPTXDocument Archive: %s", file.Name)
		}
	}

	// reading document information
	if err := pptxdoc.readXMLFile("docProps/app.xml", &pptxdoc.info); err != nil {
		return nil, err
	}

	return pptxdoc, err
}

func (p *PPTXDocument) String() {
	logger := logger.Get()
	logger.Info().Msgf("PPTX Document size is: %d", p.size)
	logger.Info().Msgf("PPTX Document app version is: %s", p.info.GetAppVersion())
	logger.Info().Msgf("PPTX Document company is: %s", p.info.GetCompany())
	logger.Info().Msgf("PPTX Document words count is: %d", p.info.GetWordsCount())
}

func (p *PPTXDocument) readXMLFile(filename string, v interface{}) error {
	file, ok := p.files[filename]
	if !ok {
		return fmt.Errorf("file %s not found in the archive", filename)
	}

	handle, err := file.Open()
	if err != nil {
		return err
	}

	defer handle.Close()

	decoder := xml.NewDecoder(handle)
	return decoder.Decode(v)
}

func (p *PPTXDocument) GetInformation() *models.PPTXInformation {
	return p.info
}

func (p *PPTXDocument) GetSlides() map[string]*zip.File {
	return p.slides
}

func (p *PPTXDocument) GetMedias() map[string]*zip.File {
	return p.medias
}

func (p *PPTXDocument) GetFiles() map[string]*zip.File {
	return p.files
}

func (p *PPTXDocument) GetSize() uint16 {
	return p.size
}
