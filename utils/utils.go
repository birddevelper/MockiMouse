package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	cfg "github.com/birddevelper/mockimouse/config"
)

func GetParamFromJson(body []byte, path string) (string, error) {
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	// Remove optional $root prefix
	path = strings.TrimPrefix(path, "$root.")
	path = strings.TrimPrefix(path, "$root")

	// Split path by dot but preserve array indices
	segments := parsePathSegments(path)

	var current interface{} = data

	for _, seg := range segments {
		// Handle map keys and array indices
		key, index, isArray := parseSegment(seg)

		if key != "" {
			m, ok := current.(map[string]interface{})
			if !ok {
				return "", errors.New("expected object for key: " + key)
			}
			current, ok = m[key]
			if !ok {
				return "", errors.New("key not found: " + key)
			}
		}

		if isArray {
			a, ok := current.([]interface{})
			if !ok {
				return "", errors.New("expected array for index in segment: " + seg)
			}
			if index < 0 || index >= len(a) {
				return "", fmt.Errorf("index out of range in segment: %s", seg)
			}
			current = a[index]
		}
	}

	return fmt.Sprintf("%v", current), nil
}

// parsePathSegments splits path while preserving [index] as part of the segment
func parsePathSegments(path string) []string {
	var segments []string
	var buf strings.Builder
	bracket := 0

	for _, ch := range path {
		switch ch {
		case '.':
			if bracket == 0 {
				segments = append(segments, buf.String())
				buf.Reset()
			} else {
				buf.WriteRune(ch)
			}
		case '[':
			bracket++
			buf.WriteRune(ch)
		case ']':
			bracket--
			buf.WriteRune(ch)
		default:
			buf.WriteRune(ch)
		}
	}
	if buf.Len() > 0 {
		segments = append(segments, buf.String())
	}
	return segments
}

// parseSegment separates key and array index
func parseSegment(segment string) (key string, index int, isArray bool) {
	re := regexp.MustCompile(`^([^\[\]]*)(\[(\d+)\])?$`)
	matches := re.FindStringSubmatch(segment)

	if len(matches) == 4 {
		key = matches[1]
		if matches[2] != "" {
			isArray = true
			index, _ = strconv.Atoi(matches[3])
		}
	}
	return
}

func PrintEndpointsInfo(endpoints []cfg.EndPoint, port int) {

	logo := `                                                                                                   
                                                                                                    
	++++.  ++++                   ++++       =@@+ ++++   ++++                                        
	@@@@#  @@@@.                  @@@@      .@@@= @@@@* .@@@@                                        
	:@@@@ *@@@%    =++       ++   :@@@            -@@@@ #@@@#    =++               -+:+.    =++      
	:@@@@ @@@@%   @@@@@-   -@@@@+ :@@@ @@@@ #@@@@ -@@@@ @@@@#   @@@@@:  @@@@=@@@@ +@@@@%   @@@@@     
	:@@@@-@@@@%  @@@ @@@   @@@:@@ :@@@  @@   @@@@ -@@@@:@@@@#  @@@ @@@  #@@@ @@@@ @@@-@@  @@@ @@@    
	:@@@@@@@@@%  @@@ @@@# =@@@+@@ :@@@ #@:   @@@@ -@@@@@@@@@# .@@@ @@@# #@@@ @@@@ @@@.@@  @@@ @@@    
	:@@@@@@@@@% %@@@ @@@@ @@@@:@@ :@@@ @@    @@@@ -@@@@@@@@@# @@@@ @@@@ #@@@ @@@@ @@@@   @@@@ @@@.   
	:@@@@@@@@@% @@@@ @@@@ @@@@    :@@@-@@-   @@@@ -@@@@@@@@@# @@@@ @@@@ #@@@ @@@@ @@@@=  @@@@ @@@-   
	:@@@@@*@@@% @@@@ @@@@ @@@@    :@@@@@@@   @@@@ -@@@@@#@@@# @@@@ @@@@ #@@@ @@@@ :@@@@  @@@@@@@@-   
	:@#@@@:@@@% @@@@ @@@@ @@@@    :@@@@@@@   @@@@ -@*@@@-@@@# @@@@ @@@@ #@@@ @@@@  @@@@= @@@@        
	:@#@@@:@@@% @@@@ @@@@ @@@@    :@@@@@@@   @@@@ -@*@@@-@@@# @@@@ @@@@ #@@@ @@@@  :@@@@ @@@@        
	:@#-@%:@@@% @@@@ @@@@ @@@@    :@@@=@@@=  @@@@ -@*=@#-@@@# @@@@ @@@@ #@@@ @@@@   @@@@ @@@@        
	:@# +.:@@@% #@@@ @@@@ @@@@ @@ :@@@ @@@@  @@@@ -@* + -@@@# %@@@ @@@@ #@@@ @@@@ ++.@@@ #@@@  %@-   
	:@#   :@@@%  @@@ @@@- =@@@ @@ :@@@ @@@@  @@@@ -@*   -@@@#  @@@ @@@: #@@@ @@@@ @@ @@@  @@@  @@.   
	=@@   %@@@@  %@@.@@@   @@@ @@ =@@@:*@@@- @@@@ +@@   %@@@@  @@@ @@@  +@@@@@@@@ #@@@@@  %@@=.@@    
	@@@.  @@@@@.  %@@@@     @@@@  @@@@@ @@@@*@@@@.@@@.  @@@@@   @@@@@    @@@%@@@@:-@@@@=   %@@@@     
																									 `
	fmt.Println(logo)

	fmt.Println("MockiMouse initialized with port ", port)
	fmt.Println("-------------------------")
	for i, endpoint := range endpoints {
		fmt.Println(i+1, "- ", endpoint.Name, " : ", endpoint.Method, " ", endpoint.Path)
	}
	fmt.Println("-------------------------")

}
