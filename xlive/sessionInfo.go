// =================================================================================
//
//		go-xlive - https://www.github.com/hairlesshobo/go-xlive
//
//	 go-xlive is a tool for working with multitrack audio sessions recorded
//   with a Behringer X-Live addon card
//
//		Copyright (c) 2024 Steve Cross <flip@foxhollow.cc>
//
//		Licensed under the Apache License, Version 2.0 (the "License");
//		you may not use this file except in compliance with the License.
//		You may obtain a copy of the License at
//
//		     http://www.apache.org/licenses/LICENSE-2.0
//
//		Unless required by applicable law or agreed to in writing, software
//		distributed under the License is distributed on an "AS IS" BASIS,
//		WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//		See the License for the specific language governing permissions and
//		limitations under the License.
//
// =================================================================================

// package xlive provides an interface for working with multitrack audio sessions
// recorded by a Behringer X-Live addon card
package xlive

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

//
// data types
//

// SessionInfo describes the contents of a SE_LOG.BIN file after it has been
// read with ReadSessionFile
type SessionInfo struct {
	SessionDtm   time.Time
	ChannelCount uint32
	SampleRate   uint32
	DateCode     uint32
	TakeCount    uint32
	MarkerCount  uint32
	TotalLength  uint32 // samples per channel
	TakeSize     []uint32
	TakeMarkers  []uint32
}

//
// public functions
//

// ReadSessionFile opens a SE_LOG.BIN file, parses it and returns a SessionInfo
// object containing information about the X-Live recording session
func ReadSessionFile(sessionFile string) SessionInfo {
	f, err := os.Open(sessionFile)

	if err != nil {
		panic("unable to open file: " + err.Error())
	}
	defer f.Close()

	sessionDtm, _ := getSessionDtm(readUint32(f))

	sessionInfo := SessionInfo{
		SessionDtm:   sessionDtm,
		ChannelCount: readUint32(f),
		SampleRate:   readUint32(f),
		DateCode:     readUint32(f),
		TakeCount:    readUint32(f),
		MarkerCount:  readUint32(f),
		TotalLength:  readUint32(f),
	}

	for i := 0; i < int(sessionInfo.TakeCount); i++ {
		sessionInfo.TakeSize = append(sessionInfo.TakeSize, readUint32(f))
	}

	// dummy reads the rest of the not used take lengths
	for i := sessionInfo.TakeCount; i < 256; i++ { // MAX_NO_TAKE = 256
		readBytes(f, 4)
	}

	for i := 0; i < int(sessionInfo.MarkerCount); i++ {
		sessionInfo.TakeMarkers = append(sessionInfo.TakeMarkers, readUint32(f))
	}

	return sessionInfo
}

//
// private functions
//

func readBytes(f *os.File, count int) []byte {
	buffer := make([]byte, count)
	n, err := f.Read(buffer)

	if err != nil {
		panic("Failed to read: " + err.Error())
	}

	return buffer[0:n]
}

func readUint32(f *os.File) uint32 {
	return binary.LittleEndian.Uint32(readBytes(f, 4))
}

func getSessionDtm(sessionNumber uint32) (time.Time, error) {
	// uggh.. bitwise operations

	year := ((sessionNumber & uint32(0b111111<<25)) >> 25) + 1980
	month := ((sessionNumber & uint32(0b1111<<21)) >> 21)
	day := ((sessionNumber & uint32(0b11111<<16)) >> 16)
	hour := ((sessionNumber & uint32(0b11111<<11)) >> 11)
	minute := ((sessionNumber & uint32(0b111111<<5)) >> 5)
	second := ((sessionNumber & uint32(0b111111<<0)) >> 0) * 2

	timeFmt := "2006-01-02 03:04:05 MST"
	zone, _ := time.Now().Zone()
	dtm, err := time.Parse(timeFmt, fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d %s", year, month, day, hour, minute, second, zone))

	return dtm, err
}
