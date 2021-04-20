package ipc

func SplitMessages(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	if len(data) >= MessageHeaderLength {
		if length := int(BytesToUint32(data)); len(data) >= length {
			// We have a full newline-terminated line.
			return length, data[0:length], nil
		}
	}

	// Request more data.
	return 0, nil, nil
}