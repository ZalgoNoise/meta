package file

import (
	"bufio"
	"os"
)

// File struct serves as a file handler for browsing
// a file's content in a read-only manner
type File struct {
	Path string
	Raw [][]byte
}

// Open method will retrieve the contents of the file,
// in a slice of byte arrays
func (f *File) Open(path string) error {
	var err error
	
	f.Path = path
	
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		f.Raw = append(f.Raw, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Process method will retrieve the contents of the file,
// passing each line as a byte array to an input function
// to (externally) process each line (while still storing)
// the data in a slice of byte arrays
func (f *File) Process(path string, work func([]byte) error) error {
	var err error
	
	f.Path = path
	
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := work(scanner.Bytes()); err != nil {
			return err
		}
		
		f.Raw = append(f.Raw, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// ProcessFields method will retrieve the contents of the file,
// passing each line as a byte array to an input function
// to (externally) process each line (while still storing)
// the data in a slice of byte arrays
func (f *File) ProcessFields(path string, work func([][]byte) error) error {
	var err error
	
	f.Path = path
	
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := work(func(b []byte) [][]byte {

				var bufArr [][]byte
				var buf []byte
				counter := 0


				for _, v := range b {
					
					if v == 32 {
						if len(buf) == 0 {
							continue
						}
						bufArr = append(bufArr, buf)
						buf = []byte{}
						counter++
						continue
					}
					buf = append(buf, v)
				}
				if len(buf) > 0 {
					bufArr = append(bufArr, buf)
				}
				return bufArr
			}(scanner.Bytes()),
		); err != nil {
			return err
		}
		
		f.Raw = append(f.Raw, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// GetFields method will retrieve the contents of the file,
// passing each line as a byte array into a parser function that
// splits the fields in each line (while still storing
// the data in a slice of byte arrays)
func (f *File) GetFields(path string) ([][][]byte, error) {
	var err error
	var lines [][][]byte
	var lineFields [][]byte
	
	f.Path = path
	
	file, err := os.Open(path)
	if err != nil {
		return nil,err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineFields = getFields(scanner.Bytes())
		lines = append(lines, lineFields)
		
		
		f.Raw = append(f.Raw, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getFields(b []byte) [][]byte {
	var bufArr [][]byte
	var buf []byte
	counter := 0


	for i := 0 ; i < len(b) ; i++ {
		
		// " " == 32
		// "\t" == 9
		if b[i] == 32 || b[i] == 9 {

			if len(buf) == 0 {
				continue
			}
			bufArr = append(bufArr, buf)
			buf = []byte{}
			counter++
			continue
		}
		buf = append(buf, b[i])
	}
	if len(buf) > 0 {
		bufArr = append(bufArr, buf)
	}
	
	return bufArr
}